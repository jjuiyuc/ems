import { connect } from "react-redux"
import moment from "moment"
import ReportProblemIcon from "@mui/icons-material/ReportProblem"
import { useEffect, useState } from "react"
import { useTranslation } from "react-multi-lang"

import { apiCall } from "../utils/api"
import { ConvertTimeToNumber } from "../utils/utils"
import variables from "../configs/variables"

import AlertBox from "../components/AlertBox"
import EnergyResourcesCard from "../components/EnergyResourcesCard"
import EnergyResourcesTabs from "../components/EnergyResourcesTabs"
import LineChart from "../components/LineChart"
import Spinner from "../components/Spinner"

import { ReactComponent as ChargedIcon }
    from "../assets/icons/battery_charged.svg"
import { ReactComponent as ChargingIcon }
    from "../assets/icons/battery_charging.svg"
import { ReactComponent as CycleIcon } from "../assets/icons/battery_cycle.svg"
import { ReactComponent as DischargeIcon }
    from "../assets/icons/battery_discharge.svg"

const { colors } = variables

const drawHighPeak = (startHour, endHour) => chart => {
    if (chart.scales.x._gridLineItems && endHour && startHour) {
        const
            ctx = chart.ctx,
            xLines = chart.scales.x._gridLineItems,
            xLineFirst = xLines[0],
            yFirstLine = chart.scales.y._gridLineItems[0],
            xLeft = yFirstLine.x1,
            xFullWidth = yFirstLine.x2 - xLeft,
            xWidth = (endHour - startHour) / 24 * xFullWidth,
            xStart = startHour / 24 * xFullWidth + xLeft,
            yTop = xLineFirst.y1,
            yFullHeight = xLineFirst.y2 - yTop

        ctx.beginPath()
        ctx.fillStyle = "#ffffff10"
        ctx.strokeStyle = colors.gray[400]
        ctx.rect(xStart, yTop, xWidth, yFullHeight)
        ctx.fill()
        ctx.stroke()
    }
}
const chartPowerSet = ({ data, highPeak, labels, unit }) => ({
    beforeDraw: drawHighPeak(highPeak.start, highPeak.end),
    datasets: [{
        backgroundColor: colors.blue.main,
        borderColor: colors.blue.main,
        data,
        fill: {
            above: colors.blue["main-opacity-10"],
            below: colors.blue["main-opacity-10"],
            target: "origin"
        },
        pointBorderColor: colors.blue["main-opacity-20"]
    }],
    labels,
    tickCallback: (val, index) => val + " " + unit,
    tooltipLabel: item => `${item.parsed.y} ${unit}`,
    y: { max: 15, min: -15 },
    x: {
        max: moment().add(1, "day").startOf("day"),
        min: moment().startOf("day")
    }
})
const ErrorBox = ({ error, margin = "", message }) => error
    ? <AlertBox
        boxClass={`${margin} negative`}
        content={<>
            <span className="font-mono ml-2">{error}</span>
            <span className="ml-2">{message}</span>
        </>}
        icon={ReportProblemIcon}
        iconColor="negative-main" />
    : null
const LoadingBox = ({ loading }) => loading
    ? <div className="grid h-24 place-items-center"><Spinner /></div>
    : null

const mapState = state => ({ gatewayID: state.gateways.active.gatewayID })

export default connect(mapState)(function EnergyResoucesBattery(props) {
    const
        [batteryPower, setBatteryPower] = useState(0),
        [capacity, setCapacity] = useState(0),
        [chargedLifetime, setChargedLifetime] = useState(0),
        [chargedToday, setChargedToday] = useState(0),
        [chargingState, setChargingState] = useState(0),
        [chargeVoltage, setChargeVoltage] = useState(null),
        [chargeVoltageError, setChargeVoltageError] = useState(""),
        [chargeVoltageLoading, setChargeVoltageLoading] = useState(false),
        [chargeVoltageRes] = useState("hour"),
        [cyclesLifetime, setCyclesLifetime] = useState(0),
        [cyclesToday, setCyclesToday] = useState(0),
        [dischargedLifetime, setDischargedLifetime] = useState(0),
        [dischargedToday, setDischargedToday] = useState(0),
        [infoError, setInfoError] = useState(""),
        [infoLoading, setInfoLoading] = useState(false),
        [modal, setModal] = useState(""),
        [power, setPower] = useState(null),
        [powerError, setPowerError] = useState(""),
        [powerLoading, setPowerLoading] = useState(false),
        [powerRes] = useState("hour"),
        [powerSources, setPowerSources] = useState(""),
        [voltage, setVoltage] = useState(0)

    const chartChargeVoltageSet = ({ data, highPeak, labels, unit }) => ({
        beforeDraw: drawHighPeak(highPeak.start, highPeak.end),
        datasets: [
            {
                backgroundColor: colors.blue.main,
                borderColor: colors.blue.main,
                data: data.charge,
                fill: {
                    above: colors.blue["main-opacity-10"],
                    target: "origin"
                },
                id: "charge",
                pointBorderColor: colors.blue["main-opacity-20"],
                label: pageT("soc")
            },
            {
                backgroundColor: colors.primary.main,
                borderColor: colors.primary.main,
                data: data.voltage,
                fill: {
                    above: colors.primary["main-opacity-10"],
                    target: "origin"
                },
                id: "voltage",
                pointBorderColor: colors.primary["main-opacity-20"],
                label: pageT("voltage"),
                yAxisID: "y1"
            },
        ],
        labels,
        tickCallback: val => val + " " + unit.charge,
        tooltipLabel: item => `${item.dataset.label} ${item.parsed.y}`
            + unit[item.dataset.id],
        y: { max: 100, min: 0 },
        y1: { max: 100, min: 0 },
        y1TickCallback: val => val + " " + unit.voltage,
        x: {
            max: moment().add(1, "day").startOf("day"),
            min: moment().startOf("day")
        }
    })
    useEffect(() => {
        if (!props.gatewayID) return

        const
            startTime = moment().startOf("day").toISOString(),
            chartParams = resolution => new URLSearchParams({
                startTime,
                endTime: moment().toISOString(),
                resolution
            }).toString(),
            urlPrefix = `/api/${props.gatewayID}/devices/battery`

        apiCall({
            onComplete: () => setInfoLoading(false),
            onError: error => setInfoError(error),
            onStart: () => setInfoLoading(true),
            onSuccess: rawData => {
                if (!rawData || !rawData.data) return

                const { data } = rawData

                setBatteryPower(data.batteryPower || 0)
                setCapacity(data.capcity || 0)
                setChargedLifetime(data.batteryConsumedLifetimeEnergyAC || 0)
                setChargedToday(data.batteryConsumedLifetimeEnergyACDiff || 0)
                setChargingState(data.batterySoC || 0)
                setDischargedLifetime(data.batteryProducedLifetimeEnergyAC || 0)
                setDischargedToday(data.batteryProducedLifetimeEnergyACDiff || 0)
                setCyclesToday(data.batteryLifetimeOperationCyclesDiff || 0)
                setCyclesLifetime(data.batteryLifetimeOperationCycles || 0)
                setModal(data.model || "-")
                setPowerSources(data.powerSources || "-")
                setVoltage(data.voltage || 0)
            },
            url: `${urlPrefix}/energy-info?startTime=${startTime}`
        })

        const oClocks = Array.from(new Array(25).keys()).map(n =>
            parseInt(moment().hour(n).startOf("h").format("x")))

        apiCall({
            onComplete: () => setPowerLoading(false),
            onError: error => setPowerError(error),
            onStart: () => setPowerLoading(true),
            onSuccess: rawData => {
                if (!rawData || !rawData.data) return

                const
                    { data } = rawData,
                    { onPeakTime, timestamps } = data,
                    { end, start, timezone } = onPeakTime,
                    labels = [
                        ...timestamps.map(t => t * 1000),
                        ...oClocks.slice(timestamps.length)
                    ],
                    peakStart = ConvertTimeToNumber(start, timezone),
                    peakEnd = ConvertTimeToNumber(end, timezone)

                setPower({
                    data: data.batteryAveragePowerACs,
                    highPeak: { start: peakStart, end: peakEnd },
                    labels
                })
            },
            url: `${urlPrefix}/power-state?${chartParams(powerRes)}`
        })

        const chargeVoltageUrl = `${urlPrefix}/charge-voltage-state?`
            + chartParams(chargeVoltageRes)

        apiCall({
            onComplete: () => setChargeVoltageLoading(false),
            onError: error => setChargeVoltageError(error),
            onStart: () => setChargeVoltageLoading(true),
            onSuccess: rawData => {
                if (!rawData || !rawData.data) return

                const
                    { data } = rawData,
                    { onPeakTime, timestamps } = data,
                    { end, start, timezone } = onPeakTime,
                    labels = [
                        ...timestamps.map(t => t * 1000),
                        ...oClocks.slice(timestamps.length)
                    ],
                    peakStart = ConvertTimeToNumber(start, timezone),
                    peakEnd = ConvertTimeToNumber(end, timezone)

                setChargeVoltage({
                    data: {
                        charge: data.batterySoCs,
                        voltage: data.batteryVoltages
                    },
                    highPeak: { start: peakStart, end: peakEnd },
                    labels
                })
            },
            url: chargeVoltageUrl
        })
    }, [props.gatewayID])

    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = string => t("energyResources.battery." + string)

    const batteryInfoData = [
        { title: pageT("modal"), value: modal },
        { title: pageT("capacity"), value: capacity },
        { title: pageT("powerSources"), value: powerSources },
        { title: pageT("batteryPower"), value: batteryPower },
        { title: pageT("voltage"), value: voltage },
    ]

    const cardsData = {
        charged: [
            {
                title: commonT("today"),
                value: `${chargedToday} ${commonT("kwh")}`
            },
            {
                title: pageT("lifetime"),
                value: `${chargedLifetime} ${commonT("kwh")}`
            }
        ],
        chargingState: [{
            title: pageT("percentage"),
            value: `${chargingState} %`
        }],
        cycles: [
            { title: commonT("today"), value: cyclesToday },
            { title: pageT("lifetime"), value: cyclesLifetime }
        ],
        discharged: [
            {
                title: commonT("today"),
                value: `${dischargedToday} ${commonT("kwh")}`
            },
            {
                title: pageT("lifetime"),
                value: `${dischargedLifetime} ${commonT("kwh")}`
            }
        ]
    }

    const chargeVoltageChart = chargeVoltage
        ? <div className="max-h-80vh h-160 relative w-full">
            <LineChart
                data={chartChargeVoltageSet({
                    ...chargeVoltage,
                    unit: { charge: "%", voltage: " " + commonT("v") },
                })}
                id="erbChargeVoltage" />
        </div>
        : null

    const powerChart = power
        ? <div className="max-h-80vh h-160 relative w-full">
            <LineChart
                data={chartPowerSet({ ...power, unit: commonT("kw") })}
                id="erbPower" />
        </div>
        : null

    const batteryInfoCards = batteryInfoData.map((item, i) =>
        <div className="card" key={"erb-bic-" + i}>
            <h5 className="mb-4">{item.title}</h5>
            <h2>{item.value || "-"}</h2>
        </div>)

    const infoErrorBox = <ErrorBox
        error={infoError}
        margin="mb-8"
        message={pageT("infoError")} />

    return <>
        <h1 className="mb-9">{t("navigator.energyResources")}</h1>
        <EnergyResourcesTabs current="battery" />
        {infoErrorBox}
        <div className="font-bold gap-8 grid lg:grid-cols-2 relative">
            <EnergyResourcesCard
                data={cardsData.cycles}
                icon={CycleIcon}
                title={pageT("batteryOperationCycles")} />
            <EnergyResourcesCard
                data={cardsData.chargingState}
                icon={ChargingIcon}
                title={pageT("stateOfChargeSOC")} />
            <EnergyResourcesCard
                data={cardsData.discharged}
                icon={DischargeIcon}
                title={pageT("discharged")} />
            <EnergyResourcesCard
                data={cardsData.charged}
                icon={ChargedIcon}
                title={pageT("charged")} />
            {infoLoading
                ? <div className="absolute bg-black-main-opacity-95 grid inset-0
                                place-items-center rounded-3xl">
                    <Spinner />
                </div>
                : null}
        </div>
        <div className="card chart mt-8">
            <h4 className="mb-9">{pageT("chargingDischargingPower")}</h4>
            <ErrorBox error={powerError} message={pageT("chartError")} />
            <LoadingBox loading={powerLoading} />
            {powerChart}
        </div>
        <div className="card chart mt-8">
            <h4 className="mb-9">{pageT("stateOfChargeVoltage")}</h4>
            <ErrorBox
                error={chargeVoltageError}
                message={pageT("chartError")} />
            {chargeVoltageChart}
            <LoadingBox loading={chargeVoltageLoading} />
        </div>
        <h1 className="mb-7 mt-20">{pageT("batteryInformation")}</h1>
        {infoErrorBox}
        <div className="font-bold gap-5 grid md:grid-cols-2 lg:grid-cols-3">
            {batteryInfoCards}
        </div>
    </>
})