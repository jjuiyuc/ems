import { connect } from "react-redux"
import { Button, Stack } from "@mui/material"
import ReportProblemIcon from "@mui/icons-material/ReportProblem"
import { Fragment as Frag, useEffect, useState } from "react"
import moment from "moment"
import { useTranslation } from "react-multi-lang"

import { apiCall } from "../utils/api"
import variables from "../configs/variables"

import AlertBox from "../components/AlertBox"
import BatteryStatusCard from "../components/BatteryStatusCard"
import Clock from "../components/Clock"
import EnergyCard from "../components/EnergyCard"
import LineChart from "../components/LineChart"
import Spinner from "../components/Spinner"

import { ReactComponent as EditIcon } from "../assets/icons/edit.svg"

const { colors } = variables

const MOCK_MODE = true

const mockTouInfo = {
    energySources: {
        offPeak: {
            allProducedLifetimeEnergyACDiff: 44.207,
            batteryProducedEnergyPercentAC: 2.04,
            batteryProducedLifetimeEnergyACDiff: 0.901,
            gridProducedEnergyPercentAC: 72.62,
            gridProducedLifetimeEnergyACDiff: 32.102,
            pvProducedEnergyPercentAC: 25.34,
            pvProducedLifetimeEnergyACDiff: 11.204
        },
        onPeak: {
            allProducedLifetimeEnergyACDiff: 66.896,
            batteryProducedEnergyPercentAC: 20.18,
            batteryProducedLifetimeEnergyACDiff: 13.498,
            gridProducedEnergyPercentAC: 5.98,
            gridProducedLifetimeEnergyACDiff: 4,
            pvProducedEnergyPercentAC: 73.84,
            pvProducedLifetimeEnergyACDiff: 49.398
        }
    },
    timeOfUse: {
        currentPeakType: "onPeak",
        midPeak: null,
        offPeak: [
            { end: "06:00:00", start: "00:00:00", touRate: 2.15 },
            { end: "14:00:00", start: "11:00:00", touRate: 2.15 }
        ],
        onPeak: [
            { end: "11:00:00", start: "06:00:00", touRate: 5.39 },
            { end: "24:00:00", start: "14:00:00", touRate: 5.39 }
        ],
        timezone: "+0800"
    }
}

const mockBatteryStatus = {
    batterySoC: 70,
    batteryProducedAveragePowerAC: 3.369,
    batteryConsumedAveragePowerAC: 0,
    batteryChargingFrom: "",
    batteryDischargingTo: "Load"
}

const mockSolarUsage = {
    timestamps: [
        1742749150, 1742752753, 1742756353, 1742759950, 1742763555,
        1742767155, 1742770749, 1742774355, 1742777949, 1742781551,
        1742785149, 1742788758, 1742792349, 1742795954, 1742799550,
        1742803165, 1742806749, 1742810347, 1742811969
    ],
    loadPvConsumedEnergyPercentACs: [
        0, 0, 0, 0, 0, 100, 100, 22.15, 18.3, 17.22,
        100, 100, 100, 100, 100, 33.77, 100, 100, 100
    ]
}

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

export default connect(mapState)(function TimeOfUse(props) {
    const showFullSections = parseInt(import.meta.env.VITE_APP_API_TOU_SHOW_FULL_SECTIONS)
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        errorT = string => t("error." + string),
        formT = string => t("form." + string),
        pageT = (string, params) => t("timeOfUse." + string, params)

    const energyCardTitle = source =>
        <>
            <span className="inline-block mr-1">
                {pageT("source", { type: pageT(source) })}
            </span>
            {/* <span className="inline-block">
            ({pageT("totalUntilNow")})
        </span> */}
        </>

    const chartSolarUsageSet = ({ data, labels }) => ({
        datasets: [{
            backgroundColor: colors.primary.main,
            borderColor: colors.primary.main,
            data: data || [],
            fill: {
                above: colors.primary["main-opacity-10"],
                target: "origin"
            },
            pointBorderColor: colors.primary["main-opacity-20"],
            hoverRadius: 0,
            pointHoverBorderWidth: 0,
            radius: 0
        }],
        labels,
        tickCallback: val => parseFloat(val.toFixed(2)) + "%",
        tooltipLabel: item => `${item.parsed.y}%`,
        x: { grid: { lineWidth: 0 } },
        y: { min: 0 }

    })
    const
        [tab, setTab] = useState("today"),
        [infoError, setInfoError] = useState(""),
        [infoLoading, setInfoLoading] = useState(false),
        [preInfoError, setPreInfoError] = useState(""),
        [preInfoLoading, setPreInfoLoading] = useState(false),
        [clockDataset, setClockDataset] = useState({
            data: [], backgroundColor: []
        }),
        [currentPeriod, setCurrentPeriod] = useState(""),
        [currentTime, setCurrentTime] = useState(""),
        [batteryStatusLoading, setBatteryStatusLoading] = useState(false),
        [batteryStatusError, setBatteryStatusError] = useState(false),
        [batteryStatus, setBatteryStatus] = useState({
            direction: "",
            target: "",
            power: 0,
            state: 0
        }),
        [lineChartUsage, setLineChartUsage] = useState(null),
        [lineChartUsageError, setLineChartUsageError] = useState(""),
        [lineChartUsageLoading, setLineChartUsageLoading] = useState(false),
        [lineChartUsageRes] = useState("hour"),
        [onPeak, setOnPeak] = useState({
            types: [
                { kwh: 0, percentage: 0, type: "grid" },
                { kwh: 0, percentage: 0, type: "solar" },
                { kwh: 0, percentage: 0, type: "battery" },
            ],
            kwh: 0,
            color: "text-negative-main"
        }),
        [offPeak, setOffPeak] = useState({
            types: [
                { kwh: 0, percentage: 0, type: "grid" },
                { kwh: 0, percentage: 0, type: "solar" },
                { kwh: 0, percentage: 0, type: "battery" },
            ],
            kwh: 0,
            color: "text-green-main"
        }),
        [midPeak, setMidPeak] = useState({
            types: [
                { kwh: 0, percentage: 0, type: "grid" },
                { kwh: 0, percentage: 0, type: "solar" },
                { kwh: 0, percentage: 0, type: "battery" },
            ],
            kwh: 0,
            color: "text-yellow-main"
        }),
        [superOffPeak, setSuperOffPeak] = useState({
            types: [
                { kwh: 0, percentage: 0, type: "grid" },
                { kwh: 0, percentage: 0, type: "solar" },
                { kwh: 0, percentage: 0, type: "battery" },
            ],
            kwh: 0
        }),
        [preOnPeak, setPreOnPeak] = useState({
            types: [
                { kwh: 0, percentage: 0, type: "grid" },
                { kwh: 0, percentage: 0, type: "solar" },
                { kwh: 0, percentage: 0, type: "battery" },
            ],
            kwh: 0,
            color: "text-negative-main"
        }),
        [preOffPeak, setPreOffPeak] = useState({
            types: [
                { kwh: 0, percentage: 0, type: "grid" },
                { kwh: 0, percentage: 0, type: "solar" },
                { kwh: 0, percentage: 0, type: "battery" },
            ],
            kwh: 0,
            color: "text-green-main"
        }),
        [preMidPeak, setPreMidPeak] = useState({
            types: [
                { kwh: 0, percentage: 0, type: "grid" },
                { kwh: 0, percentage: 0, type: "solar" },
                { kwh: 0, percentage: 0, type: "battery" },
            ],
            kwh: 0,
            color: "text-yellow-main"
        }),
        [preSuperOffPeak, setPreSuperOffPeak] = useState({
            types: [
                { kwh: 0, percentage: 0, type: "grid" },
                { kwh: 0, percentage: 0, type: "solar" },
                { kwh: 0, percentage: 0, type: "battery" },
            ],
            kwh: 0
        }),
        [prices, setPrices] = useState({}),
        [timeOfUse, setTimeOfUse] = useState([])

    const getMoment = string => {
        const [hour, minute] = string.split(":")
        return moment().hour(parseInt(hour)).minute(parseInt(minute)).second(0)
    }

    useEffect(() => {
        const
            dataset = { data: [], backgroundColor: [] },
            prices = {}

        let currentPeriod = ""

        if (timeOfUse.length === 0) return
        timeOfUse.forEach(item => {
            const
                { end, start } = item,
                endTime = getMoment(end),
                startTime = getMoment(start),
                duration = moment.duration(endTime.diff(startTime)).as("hours")

            dataset.data.push(duration)
            dataset.backgroundColor.push(colors[item.name])
            prices[item.name] = item.touRate

            if (currentTime >= startTime) {
                currentPeriod = item.name
            }
        })
        setClockDataset(dataset)
        setPrices(prices)

    }, [timeOfUse])

    const urlPrefix = `/api/${props.gatewayID}/devices`
    const callTodayCards = () => {
        if (MOCK_MODE) {
            setInfoLoading(true)
            setTimeout(() => {
                const data = mockTouInfo
                const { onPeak, offPeak } = data.energySources
                const { timeOfUse } = data

                setOnPeak(r => ({
                    ...r,
                    types: [
                        { kwh: onPeak.gridProducedLifetimeEnergyACDiff, percentage: onPeak.gridProducedEnergyPercentAC, type: "grid" },
                        { kwh: onPeak.pvProducedLifetimeEnergyACDiff, percentage: onPeak.pvProducedEnergyPercentAC, type: "solar" },
                        { kwh: onPeak.batteryProducedLifetimeEnergyACDiff, percentage: onPeak.batteryProducedEnergyPercentAC, type: "battery" }
                    ],
                    kwh: onPeak.allProducedLifetimeEnergyACDiff
                }))

                setOffPeak(r => ({
                    ...r,
                    types: [
                        { kwh: offPeak.gridProducedLifetimeEnergyACDiff, percentage: offPeak.gridProducedEnergyPercentAC, type: "grid" },
                        { kwh: offPeak.pvProducedLifetimeEnergyACDiff, percentage: offPeak.pvProducedEnergyPercentAC, type: "solar" },
                        { kwh: offPeak.batteryProducedLifetimeEnergyACDiff, percentage: offPeak.batteryProducedEnergyPercentAC, type: "battery" }
                    ],
                    kwh: offPeak.allProducedLifetimeEnergyACDiff
                }))

                const periods = []
                Object.entries(timeOfUse).forEach(([key, value]) => {
                    if (Array.isArray(value)) {
                        value.forEach(item => periods.push({ name: key, ...item }))
                    }
                })
                periods.sort((a, b) => getMoment(a.start) - getMoment(b.start))

                setCurrentTime(moment().format("HH:mm"))
                setCurrentPeriod(timeOfUse.currentPeakType)
                setTimeOfUse(periods)
                setInfoLoading(false)
            }, 300)
        }
    }
        ,
        callYesterdayCards = (preStartTime) => {
            if (MOCK_MODE) {
                setInfoLoading(true)
                setTimeout(() => {
                    const data = mockTouInfo
                    const { onPeak, offPeak } = data.energySources
                    const { timeOfUse } = data

                    setOnPeak(r => ({
                        ...r,
                        types: [
                            { kwh: onPeak.gridProducedLifetimeEnergyACDiff, percentage: onPeak.gridProducedEnergyPercentAC, type: "grid" },
                            { kwh: onPeak.pvProducedLifetimeEnergyACDiff, percentage: onPeak.pvProducedEnergyPercentAC, type: "solar" },
                            { kwh: onPeak.batteryProducedLifetimeEnergyACDiff, percentage: onPeak.batteryProducedEnergyPercentAC, type: "battery" }
                        ],
                        kwh: onPeak.allProducedLifetimeEnergyACDiff
                    }))

                    setOffPeak(r => ({
                        ...r,
                        types: [
                            { kwh: offPeak.gridProducedLifetimeEnergyACDiff, percentage: offPeak.gridProducedEnergyPercentAC, type: "grid" },
                            { kwh: offPeak.pvProducedLifetimeEnergyACDiff, percentage: offPeak.pvProducedEnergyPercentAC, type: "solar" },
                            { kwh: offPeak.batteryProducedLifetimeEnergyACDiff, percentage: offPeak.batteryProducedEnergyPercentAC, type: "battery" }
                        ],
                        kwh: offPeak.allProducedLifetimeEnergyACDiff
                    }))

                    const periods = []
                    Object.entries(timeOfUse).forEach(([key, value]) => {
                        if (Array.isArray(value)) {
                            value.forEach(item => periods.push({ name: key, ...item }))
                        }
                    })
                    periods.sort((a, b) => getMoment(a.start) - getMoment(b.start))

                    setCurrentTime(moment().format("HH:mm"))
                    setCurrentPeriod(timeOfUse.currentPeakType)
                    setTimeOfUse(periods)
                    setInfoLoading(false)
                }, 300)
            }
        },
        callLineChartUsage = () => {
            if (MOCK_MODE) {
                setLineChartUsageLoading(true)
                setTimeout(() => {
                    const { timestamps, loadPvConsumedEnergyPercentACs } = mockSolarUsage
                    setLineChartUsage({
                        data: loadPvConsumedEnergyPercentACs,
                        labels: timestamps.map(t => t * 1000)
                    })
                    setLineChartUsageLoading(false)
                }, 300)
            }
        }
    useEffect(() => {
        if (!props.gatewayID) return

        let startTime = "", endTime = ""
        let preStartTime = "", preEndTime = ""
        if (tab === "today") {
            startTime = moment().startOf("day").toISOString()
            endTime = moment().toISOString()

            callTodayCards(startTime)
            callLineChartUsage(startTime, endTime)

        } else if (tab === "yesterday") {
            preStartTime = moment().subtract(1, "day").startOf("day").toISOString()
            preEndTime = moment().subtract(1, "day").endOf("day").toISOString()

            callYesterdayCards(preStartTime)
        }
    }, [props.gatewayID, tab])

    useEffect(() => {
        if (!props.gatewayID) return

        if (MOCK_MODE) {
            setBatteryStatusLoading(true)
            setTimeout(() => {
                const data = mockBatteryStatus
                setBatteryStatus({
                    direction: data.batteryChargingFrom ? "chargingFrom" : "dischargingTo",
                    target: data.batteryChargingFrom || data.batteryDischargingTo,
                    power: data.batteryProducedAveragePowerAC + data.batteryConsumedAveragePowerAC,
                    state: data.batterySoC
                })
                setBatteryStatusLoading(false)
            }, 300)
        }
    }, [props.gatewayID])


    return <>
        <div className="page-header">
            <h1>{pageT("timeOfUse")}</h1>
            <Stack direction="row" justifyContent="flex-end" spacing={1.5}>
                <Button
                    onClick={() => setTab("today")}
                    filter={tab === "today" ? "selected" : ""}
                    radius="pill"
                    variant="contained">
                    {commonT("today")}
                </Button>
                <Button
                    filter={tab === "yesterday" ? "selected" : ""}
                    onClick={() => setTab("yesterday")}
                    radius="pill"
                    variant="contained">
                    {pageT("yesterday")}
                </Button>
            </Stack>
        </div>
        {tab === "today"
            ? <>
                <div className="gap-8 grid md:grid-cols-2 items-start">
                    <EnergyCard data={onPeak} title={energyCardTitle("onPeak")} />
                    <EnergyCard data={offPeak} title={energyCardTitle("offPeak")} />
                    {showFullSections
                        ? <>
                            <EnergyCard data={midPeak} title={energyCardTitle("midPeak")} />
                            <div className="card energyCard">
                                <div className="flex flex-wrap items-baseline mb-8">
                                    <h2 className="mr-2 whitespace-nowrap">{superOffPeak.kwh} {commonT("kwh")}</h2>
                                    <h5 className="font-bold">
                                        <span className="inline-block mr-1">
                                            {pageT("superOffPeak")} {commonT("sources")}
                                        </span>
                                    </h5>
                                </div>
                                <div className="h-2 bg-gray-600 w-full rounded-full" />
                                <div className="mx-2.5 mb-12 mt-4 lg:h-5 w-3 mr-2 sm:h-4" />
                                <div className="grid grid-cols-3 column-separator gap-x-5 sm:gap-x-10">
                                    {superOffPeak.types.map((t, i) =>
                                        <div key={"detail-" + i}
                                            className="">
                                            <h6 className="font-bold text-white">{commonT(t.type)}</h6>
                                            <h3 className="my-1">-</h3>
                                            {/* <p className="lg:test text-13px text-white">
                                                {t.kwh} {commonT("kwh")}
                                            </p> */}
                                            <div className="md:h-6 lg:h-4 w-4"></div>
                                        </div>)}
                                </div>
                            </div>
                        </> : null}
                    {infoLoading
                        ? <div className="absolute bg-black-main-opacity-95 grid inset-0
                                place-items-center rounded-3xl">
                            <Spinner />
                        </div>
                        : null}
                    <div className="card">
                        <div className="header -mr-4">
                            <h4>{pageT("timeOfUse")}</h4>
                            {/* <Button
                        color="brand"
                        radius="pill"
                        size="small"
                        variant="text">
                        <EditIcon className="h-4 mr-1 w-4" />
                        {pageT("editTimeOfUse")}
                    </Button> */}
                        </div>
                        <div className="flex flex-wrap items-center justify-around">
                            <div className="flex flex-wrap items-center justify-center">
                                <Clock dataset={clockDataset} id="touClock" />
                                <div className="grid grid-cols-3-auto gap-y-2
                                        items-center mx-8 my-4 text-white">
                                    {Object.keys(prices).map((key, i) =>
                                        <Frag key={"t-p-" + i}>
                                            <div
                                                className="h-2 rounded-full mr-3 w-2"
                                                style={{ background: colors[key] }} />
                                            <div className="text-11px">
                                                {pageT(key)}
                                            </div>
                                            <div className="font-bold ml-2 text-base">
                                                ${prices[key]} /{commonT("kwh")}
                                            </div>
                                        </Frag>)}
                                </div>
                            </div>
                            <div className="my-6 subCard w-56">
                                <h6 className="font-bold text-11px text-gray-300">
                                    {pageT("current")}
                                </h6>
                                <div className="flex flex-col items-center justify-center">
                                    <div
                                        className="font-bold mb-3 mt-2 px-5 py-2 rounded-full text-gray-900 text-sm"
                                        style={{ background: colors[currentPeriod] }}>
                                        {pageT(currentPeriod)}
                                    </div>
                                    <div className="font-bold text-2xl">
                                        {currentTime}
                                    </div>
                                </div>
                            </div>
                        </div>
                        {infoLoading
                            ? <div className="absolute bg-black-main-opacity-95 grid inset-0
                                place-items-center rounded-3xl">
                                <Spinner />
                            </div>
                            : null}
                    </div>
                    <BatteryStatusCard data={batteryStatus} />
                </div>
                <div className="mt-20 page-header">
                    <h1>{pageT("directSolarUsage")}</h1>
                </div>
                <div className="card chart max-h-80vh h-160 relative w-full">
                    <LineChart data={chartSolarUsageSet({
                        ...lineChartUsage
                    })} id="touLineChart" />
                    <ErrorBox
                        error={lineChartUsageError}
                        message={pageT("chartError")} />
                    <LoadingBox loading={lineChartUsageLoading} />
                </div>
            </>
            : null}
        {tab == "yesterday"
            ? <>
                <div className="gap-8 grid md:grid-cols-2 items-start">
                    <EnergyCard data={preOnPeak} title={energyCardTitle("onPeak")} />
                    <EnergyCard data={preOffPeak} title={energyCardTitle("offPeak")} />
                    {showFullSections
                        ? <>
                            <EnergyCard data={preMidPeak} title={energyCardTitle("midPeak")} />
                            {/* <EnergyCard data={superOffPeak} title={energyCardTitle("superOffPeak")} /> */}

                            <div className="card energyCard">
                                <div className="flex flex-wrap items-baseline mb-8">
                                    <h2 className="mr-2 whitespace-nowrap">{superOffPeak.kwh} {commonT("kwh")}</h2>
                                    <h5 className="font-bold">
                                        <span className="inline-block mr-1">
                                            {pageT("superOffPeak")} {commonT("sources")}
                                        </span>
                                    </h5>
                                </div>
                                <div className="h-2 bg-gray-600 w-full rounded-full" />
                                <div className="mx-2.5 mb-12 mt-4 lg:h-5 w-3 mr-2 sm:h-4" />
                                <div className="grid grid-cols-3 column-separator gap-x-5 sm:gap-x-10">
                                    {superOffPeak.types.map((t, i) =>
                                        <div key={"detail-" + i}
                                            className="">
                                            <h6 className="font-bold text-white">{commonT(t.type)}</h6>
                                            <h3 className="my-1">-</h3>
                                            {/* <p className="lg:test text-13px text-white">
                                                {t.kwh} {commonT("kwh")}
                                            </p> */}
                                            <div className="md:h-6 lg:h-4 w-4"></div>
                                        </div>)}
                                </div>
                            </div>
                        </>
                        : null}
                    {infoLoading
                        ? <div className="absolute bg-black-main-opacity-95 grid inset-0
                                place-items-center rounded-3xl">
                            <Spinner />
                        </div>
                        : null}
                    <div className="card">
                        <div className="header -mr-4">
                            <h4>{pageT("timeOfUse")}</h4>
                            {/* <Button
                                color="brand"
                                radius="pill"
                                size="small"
                                variant="text">
                                <EditIcon className="h-4 mr-1 w-4" />
                                {pageT("editTimeOfUse")}
                            </Button> */}
                        </div>
                        <div className="flex flex-wrap items-center justify-around">
                            <div className="flex flex-wrap items-center justify-center">
                                <Clock dataset={clockDataset} id="touClockYesterday" />
                                <div className="grid grid-cols-3-auto gap-y-2
                                        items-center mx-8 my-4 text-white">
                                    {Object.keys(prices).map((key, i) =>
                                        <Frag key={"t-p-" + i}>
                                            <div
                                                className="h-2 rounded-full mr-3 w-2"
                                                style={{ background: colors[key] }} />
                                            <div className="text-11px">
                                                {pageT(key)}
                                            </div>
                                            <div className="font-bold ml-2 text-base">
                                                ${prices[key]} /{commonT("kwh")}
                                            </div>
                                        </Frag>)}
                                </div>
                            </div>
                        </div>
                        {infoLoading
                            ? <div className="absolute bg-black-main-opacity-95 grid inset-0
                                place-items-center rounded-3xl">
                                <Spinner />
                            </div>
                            : null}
                    </div>
                </div>
            </>
            : null
        }
    </>
})
