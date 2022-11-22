import { connect } from "react-redux"
import { Button, Stack } from "@mui/material"
import { Fragment as Frag, useEffect, useState } from "react"
import moment from "moment"
import { useTranslation } from "react-multi-lang"

import { apiCall } from "../utils/api"
import variables from "../configs/variables"

import BatteryStatusCard from "../components/BatteryStatusCard"
import Clock from "../components/Clock"
import EnergyCard from "../components/EnergyCard"
import LineChart from "../components/LineChart"

import { ReactComponent as EditIcon } from "../assets/icons/edit.svg"

const { colors } = variables

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
        tickCallback: (val, index) => val + "%",
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
                { kwh: 7.5, percentage: 15, type: "grid" },
                { kwh: 30, percentage: 60, type: "solar" },
                { kwh: 12.5, percentage: 25, type: "battery" },
            ],
            kwh: 0
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
                { kwh: 0, percentage: 25, type: "battery" },
            ],
            kwh: 50
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
        // [prices, setPrices]
        //     = useState({ onPeak: 0, midPeak: 0, offPeak: 0, superOffPeak: 0 }),
        // [timeOfUse, setTimeOfUse] = useState([
        //     {
        //         end: "05:00",
        //         name: "superOffPeak",
        //         price: 1.2,
        //         start: "00:00"
        //     },
        //     {
        //         end: "11:00",
        //         name: "offPeak",
        //         price: 1.8,
        //         start: "05:00"
        //     },
        //     {
        //         end: "17:00",
        //         name: "midPeak",
        //         price: 2.2,
        //         start: "11:00"
        //     },
        //     {
        //         end: "23:00",
        //         name: "onPeak",
        //         price: 3.5,
        //         start: "17:00"
        //     },
        //     {
        //         end: "24:00",
        //         name: "superOffPeak",
        //         price: 1.2,
        //         start: "23:00"
        //     }
        // ]),
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
            console.log(item.name)
            dataset.data.push(duration)
            dataset.backgroundColor.push(colors[item.name])
            prices[item.name] = item.touRate

            if (currentTime >= startTime) {
                currentPeriod = item.name
            }
        })
        setClockDataset(dataset)
        // setCurrentPeriod(currentPeriod)
        // setCurrentTime(timeOfUse.currentPeakType || "")
        setPrices(prices)

    }, [timeOfUse])

    const urlPrefix = `/api/${props.gatewayID}/devices`
    const
        callTodayCards = (startTime) => {
            apiCall({
                onComplete: () => setInfoLoading(false),
                onError: error => setInfoError(error),
                onStart: () => setInfoLoading(true),
                onSuccess: rawData => {
                    if (!rawData?.data) return
                    const { data } = rawData,
                        { onPeak } = data.energySources

                    setOnPeak(r => ({
                        ...r,
                        types: [
                            {
                                kwh: onPeak?.gridProducedLifetimeEnergyACDiff,
                                percentage: onPeak?.gridProducedEnergyPercentAC,
                                type: "grid"
                            },
                            {
                                kwh: onPeak?.pvProducedLifetimeEnergyACDiff,
                                percentage: onPeak?.pvProducedEnergyPercentAC,
                                type: "solar"
                            },
                            {
                                kwh: onPeak?.batteryProducedLifetimeEnergyACDiff,
                                percentage: onPeak?.batteryProducedEnergyPercentAC,
                                type: "battery"
                            },
                        ],
                        kwh: onPeak?.allProducedLifetimeEnergyACDiff
                    }))

                    const { offPeak } = data.energySources
                    setOffPeak(r => ({
                        ...r,
                        types: [
                            {
                                kwh: offPeak?.gridProducedLifetimeEnergyACDiff,
                                percentage: offPeak?.gridProducedEnergyPercentAC,
                                type: "grid"
                            },
                            {
                                kwh: offPeak?.pvProducedLifetimeEnergyACDiff,
                                percentage: offPeak?.pvProducedEnergyPercentAC,
                                type: "solar"
                            },
                            {
                                kwh: offPeak?.batteryProducedLifetimeEnergyACDiff,
                                percentage: offPeak?.batteryProducedEnergyPercentAC,
                                type: "battery"
                            },
                        ],
                        kwh: offPeak?.allProducedLifetimeEnergyACDiff
                    }))

                    const { timeOfUse } = data
                    let periods = []
                    Object.keys(timeOfUse).forEach(key => {
                        if (typeof timeOfUse[key] != "object" || !timeOfUse[key]?.length) return
                        timeOfUse[key].forEach(item => {
                            periods.push({ name: key, ...item })
                        })
                    })
                    setCurrentTime(moment().format("hh:mm A"))
                    setCurrentPeriod(timeOfUse.currentPeakType || "")
                    setTimeOfUse(periods)
                },
                url: `${urlPrefix}/time-of-use-info?startTime=${startTime}`
            })
        },
        callYesterdayCards = (preStartTime) => {
            apiCall({
                onComplete: () => setPreInfoLoading(false),
                onError: error => setPreInfoError(error),
                onStart: () => setPreInfoLoading(true),
                onSuccess: rawData => {
                    if (!rawData?.data) return

                    const { data } = rawData,
                        { onPeak } = data.energySources

                    setPreOnPeak(r => ({
                        ...r,
                        types: [
                            {
                                kwh: onPeak?.gridProducedLifetimeEnergyACDiff,
                                percentage: onPeak?.gridProducedEnergyPercentAC,
                                type: "grid"
                            },
                            {
                                kwh: onPeak?.pvProducedLifetimeEnergyACDiff,
                                percentage: onPeak?.pvProducedEnergyPercentAC,
                                type: "solar"
                            },
                            {
                                kwh: onPeak?.batteryProducedLifetimeEnergyACDiff,
                                percentage: onPeak?.batteryProducedEnergyPercentAC,
                                type: "battery"
                            },
                        ],
                        kwh: onPeak?.allProducedLifetimeEnergyACDiff
                    }))

                    const { offPeak } = data.energySources
                    setPreOffPeak(r => ({
                        ...r,
                        types: [
                            {
                                kwh: offPeak?.gridProducedLifetimeEnergyACDiff,
                                percentage: offPeak?.gridProducedEnergyPercentAC,
                                type: "grid"
                            },
                            {
                                kwh: offPeak?.pvProducedLifetimeEnergyACDiff,
                                percentage: offPeak?.pvProducedEnergyPercentAC,
                                type: "solar"
                            },
                            {
                                kwh: offPeak?.batteryProducedLifetimeEnergyACDiff,
                                percentage: offPeak?.batteryProducedEnergyPercentAC,
                                type: "battery"
                            },
                        ],
                        kwh: offPeak?.allProducedLifetimeEnergyACDiff
                    }))
                    const { timeOfUse } = data
                    let periods = []
                    Object.keys(timeOfUse).forEach(key => {
                        if (typeof timeOfUse[key] != "object" || !timeOfUse[key]?.length) return
                        timeOfUse[key].forEach(item => {
                            periods.push({ name: key, ...item })
                        })
                    })
                    setTimeOfUse(periods)
                },
                url: `${urlPrefix}/time-of-use-info?startTime=${preStartTime}`
            })
        },
        callLineChartUsage = (startTime, endTime) => {
            const lineChartUsageUrl = `${urlPrefix}/solar/energy-usage?`
                + new URLSearchParams({
                    startTime, endTime, resolution: lineChartUsageRes
                }).toString()

            apiCall({
                onComplete: () => setLineChartUsageLoading(false),
                onError: error => setLineChartUsageError(error),
                onStart: () => setLineChartUsageLoading(true),
                onSuccess: rawData => {
                    if (!rawData || !rawData.data) return
                    const
                        { data } = rawData,
                        { timestamps } = data,
                        labels = timestamps.map(t => t * 1000)
                    setLineChartUsage({
                        data: data.loadPvConsumedEnergyPercentACs,
                        labels
                    })
                },
                url: lineChartUsageUrl
            })
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
        // if (startTime && endTime) {
        //     callTodayCards(startTime, endTime)

        //     if (tab === "today") {
        //     } else {
        //     }
        // }
    }, [props.gatewayID, tab])

    useEffect(() => {
        if (!props.gatewayID) return
        const startTime = moment().startOf("day").toISOString()

        apiCall({
            onComplete: () => setBatteryStatusLoading(false),
            onError: error => setBatteryStatusError(error),
            onStart: () => setBatteryStatusLoading(true),
            onSuccess: rawData => {
                if (!rawData || !rawData.data) return

                const { data } = rawData

                setBatteryStatus({
                    direction:
                        data.batteryChargingFrom ? "chargingFrom" : "dischargingTo",
                    target: (data.batteryChargingFrom || data.batteryDischargingTo),
                    power: (data.batteryProducedAveragePowerAC
                        + data.batteryConsumedAveragePowerAC || 0),
                    state: (data.batterySoC || 0)
                })
            },
            url: `/api/${props.gatewayID}/devices/battery/usage-info?startTime=${startTime}`
        })
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
                                            {pageT("superOffPeak")}{commonT("sources")}
                                        </span>
                                        <span className="inline-block">
                                            ({pageT("totalUntilNow")})
                                         </span>
                                    </h5>
                                </div>
                                <div className="h-2 bg-gray-500 w-full rounded-full" />
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
                        </> : null
                    }
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
                                                ${prices[key]}
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
                    </div>
                    <BatteryStatusCard
                        data={batteryStatus}
                    />
                </div>
                <div className="mt-20 page-header">
                    <h1>{pageT("directSolarUsage")}</h1>
                </div>
                <div className="card chart max-h-80vh h-160 relative w-full">
                    <LineChart data={chartSolarUsageSet({
                        ...lineChartUsage
                    })} id="touLineChart" />
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
                            <EnergyCard data={midPeak} title={energyCardTitle("midPeak")} />
                            <div className="card energyCard">
                                <div className="flex flex-wrap items-baseline mb-8">
                                    <h2 className="mr-2 whitespace-nowrap">{superOffPeak.kwh} {commonT("kwh")}</h2>
                                    <h5 className="font-bold">
                                        <span className="inline-block mr-1">
                                            {pageT("superOffPeak")}{commonT("sources")}
                                        </span>
                                        <span className="inline-block">
                                            ({pageT("totalUntilNow")})
                                        </span>
                                    </h5>
                                </div>
                                <div className="h-2 bg-gray-500 w-full rounded-full" />
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
                                                ${prices[key]}
                                            </div>
                                        </Frag>)}
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </>
            : null}
    </>
})