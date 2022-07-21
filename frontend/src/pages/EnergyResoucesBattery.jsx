import moment from "moment"
import { startTransition, useEffect, useState } from "react"
import { useTranslation } from "react-multi-lang"

import EnergyResoucesCard from "../components/EnergyResoucesCard"
import EnergyResoucesTabs from "../components/EnergyResoucesTabs"
import LineChart from "../components/LineChart"
import variables from "../configs/variables"

import { ReactComponent as ChargedIcon }
    from "../assets/icons/battery_charged.svg"
import { ReactComponent as ChargingIcon }
    from "../assets/icons/battery_charging.svg"
import { ReactComponent as CycleIcon } from "../assets/icons/battery_cycle.svg"
import { ReactComponent as DischargeIcon }
    from "../assets/icons/battery_discharge.svg"

const {colors} = variables

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
const chartChargeVoltageSet = ({data, highPeak, labels, unit}) => ({
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
            pointBorderColor: colors.blue["main-opacity-20"]
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
            pointBorderColor: colors.primary["main-opacity-20"]
        }
    ],
    labels,
    tooltipLabel: item => item.parsed.y + unit[item.dataset.id],
    y: {max: 100, min: 0}
})
const chartPowerSet = ({data, highPeak, labels, unit}) => ({
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
    y: {max: 15, min: -15}
})

export default function EnergyResoucesBattery () {
    const
        [batteryPower, setBatteryPower] = useState(0),
        [capacity, setCapacity] = useState(0),
        [chargedLifetime, setChargedLifetime] = useState(0),
        [chargedToday, setChargedToday] = useState(0),
        [chargingState, setChargingState] = useState(0),
        [chargeVoltage, setChargeVoltage] = useState(null),
        [cyclesLifetime, setCyclesLifetime] = useState(0),
        [cyclesToday, setCyclesToday] = useState(0),
        [dischargedLifetime, setDischargedLifetime] = useState(0),
        [dischargedToday, setDischargedToday] = useState(0),
        [modal, setModal] = useState(""),
        [power, setPower] = useState(null),
        [powerSources, setPowerSources] = useState(""),
        [voltage, setVoltage] = useState(0)

    useEffect(() => {
        setBatteryPower(10)
        setCapacity(10)
        setChargedLifetime(500)
        setChargedToday(50)
        setChargingState(80)
        setCyclesLifetime(16)
        setCyclesToday(2)
        setDischargedLifetime(500)
        setDischargedToday(50)
        setModal("Battery F1")
        setPowerSources("Solar + Grid")
        setVoltage(40)

        // Chart Data (Fake)
        const
            currentHour = new Date().getHours(),
            hours = Array.from(new Array(currentHour + 1).keys()),
            labels = Array.from(new Array(25).keys()).map(n => {
                const time = moment().hour(n).minute(0).second(0)

                return time.format("hh A")
            }),
            chargeArrays = hours
                .map(() => Math.round(Math.random() * (75 - 30) + 30)),
            powerArrays
                = hours.map(() => Math.round(Math.random() * (5 - (-5)) + (-5))),
            voltageArrays
                = hours.map(() => Math.round(Math.random() * (40 - 39) + 39))

        setChargeVoltage({
            data: {charge: chargeArrays, voltage: voltageArrays},
            highPeak: {start: 19, end: 23},
            labels
        })
        setPower({data: powerArrays, highPeak: {start: 19, end: 23}, labels})
    }, [])

    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = string => t("energyResources.battery." + string)

    const batteryInfoData = [
        {title: pageT("modal"), value: modal},
        {title: pageT("capacity"), value: capacity},
        {title: pageT("powerSources"), value: powerSources},
        {title: pageT("batteryPower"), value: batteryPower},
        {title: pageT("voltage"), value: voltage},
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
            {title: commonT("today"), value: cyclesToday},
            {title: pageT("lifetime"),value: cyclesLifetime}
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
        ? <LineChart
            data={chartChargeVoltageSet({
                ...chargeVoltage,
                unit: {charge: "%", voltage: " " + commonT("kw")}
            })}
            id="erbChargeVoltage" />
        : null

    const powerChart = power
        ? <LineChart
            data={chartPowerSet({...power, unit: commonT("kw")})}
            id="erbPower" />
        : null

    const batteryInfoCards = batteryInfoData.map((item, i) =>
        <div className="card" key={"erb-bic-" + i}>
            <h5 className="mb-4">{item.title}</h5>
            <h2>{item.value}</h2>
        </div>)

    return <>
        <h1 className="mb-9">{t("navigator.energyResources")}</h1>
        <EnergyResoucesTabs current="battery" />
        <div className="font-bold gap-8 grid lg:grid-cols-2">
            <EnergyResoucesCard
                data={cardsData.cycles}
                icon={CycleIcon}
                title={pageT("batteryOperationCycles")} />
            <EnergyResoucesCard
                data={cardsData.chargingState}
                icon={ChargingIcon}
                title={pageT("stateOfCharge")} />
            <EnergyResoucesCard
                data={cardsData.discharged}
                icon={DischargeIcon}
                title={pageT("discharged")} />
            <EnergyResoucesCard
                data={cardsData.charged}
                icon={ChargedIcon}
                title={pageT("charged")} />
        </div>
        <div className="card chart mt-8">
            <h4 className="mb-9">{pageT("chargingDischargingPower")}</h4>
            <div className="max-h-80vh h-160 relative w-full">{powerChart}</div>
        </div>
        <div className="card chart mt-8">
            <h4 className="mb-9">{pageT("stateOfChargeVoltage")}</h4>
            <div className="max-h-80vh h-160 relative w-full">
                {chargeVoltageChart}
            </div>
        </div>
        <h1 className="mb-7 mt-20">{pageT("batteryInformation")}</h1>
        <div className="font-bold gap-5 grid md:grid-cols-2 lg:grid-cols-3">
            {batteryInfoCards}
        </div>
    </>
}