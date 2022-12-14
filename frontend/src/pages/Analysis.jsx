import { connect } from "react-redux"
import { Button } from "@mui/material"
import { useTranslation } from "react-multi-lang"
import moment from "moment"
import ReportProblemIcon from "@mui/icons-material/ReportProblem"
import { useEffect, useState } from "react"

import { apiCall } from "../utils/api"
import variables from "../configs/variables"

import AlertBox from "../components/AlertBox"
import AnalysisCard from "../components/AnalysisCard"
import BarChart from "../components/BarChart"
import DateRangePicker from "../components/DateRangePicker"
import LineChart from "../components/LineChart"
import MonthPicker from "../components/MonthPicker"
import Spinner from "../components/Spinner"
import "../assets/css/dateRangePicker.css"

const { colors } = variables
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

export default connect(mapState)(function Analysis(props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("analysis." + string, params)

    const chartRealTimePowerSet = ({ data, labels }) => ({
        datasets: [
            {
                backgroundColor: colors.green.main,
                borderColor: colors.green.main,
                data: data?.load || [],
                fill: {
                    above: colors.green["main-opacity-10"],
                    below: colors.green["main-opacity-10"],
                    target: "origin"
                },
                id: "load",
                pointBorderColor: colors.green["main-opacity-20"],
                hoverRadius: 0,
                pointHoverBorderWidth: 0,
                radius: 0,
                label: pageT("load")
            },
            {
                backgroundColor: colors.yellow.main,
                borderColor: colors.yellow.main,
                data: data?.solar || [],
                fill: {
                    above: colors.yellow["main-opacity-10"],
                    below: colors.yellow["main-opacity-10"],
                    target: "origin"
                },
                id: "solar",
                pointBorderColor: colors.yellow["main-opacity-20"],
                hoverRadius: 0,
                pointHoverBorderWidth: 0,
                radius: 0,
                label: commonT("solar")
            },
            {
                backgroundColor: colors.blue.main,
                borderColor: colors.blue.main,
                data: data?.battery || [],
                fill: {
                    above: colors.blue["main-opacity-10"],
                    below: colors.blue["main-opacity-10"],
                    target: "origin"
                },
                id: "battery",
                pointBorderColor: colors.blue["main-opacity-20"],
                hoverRadius: 0,
                pointHoverBorderWidth: 0,
                radius: 0,
                label: commonT("battery")
            },
            {
                backgroundColor: colors.indigo.main,
                borderColor: colors.indigo.main,
                data: data?.grid || [],
                fill: {
                    above: colors.indigo["main-opacity-10"],
                    below: colors.indigo["main-opacity-10"],
                    target: "origin"
                },
                id: "grid",
                pointBorderColor: colors.indigo["main-opacity-20"],
                hoverRadius: 0,
                pointHoverBorderWidth: 0,
                radius: 0,
                label: commonT("grid")
            }
        ],
        labels,
        legend: true,
        tickCallback: (val, index) => val + commonT("kw"),
        tooltipLabel: item => `${item.dataset.label} ${item.parsed.y} `
            + commonT("kw"),
        x: { grid: { lineWidth: 0 } }

    })
    const chartAccumulatedPowerSet = ({ data, labels }) => ({
        datasets: [
            {
                backgroundColor: colors.green.main,
                data: data?.load || [],
                label: pageT("load")
            },
            {
                backgroundColor: colors.yellow.main,
                data: data?.solar || [],
                label: commonT("solar")
            },
            {
                backgroundColor: colors.blue.main,
                data: data?.battery || [],
                label: commonT("battery")
            },
            {
                backgroundColor: colors.indigo.main,
                data: data?.grid || [],
                label: commonT("grid")
            }
        ],
        labels,
        legend: true,
        tooltipLabel: item =>
            `${item.dataset.label} ${item.parsed.y} ${commonT("kwh")}`,
        x: {
            time: tab === "year"
                ? {
                    displayFormats: {
                        month: "YYYY MMM",
                    },
                    tooltipFormat: "YYYY MMM",
                    unit: "month"
                }
                : {
                    displayFormats: {
                        day: "MMM D",
                    },
                    tooltipFormat: "MMM D",
                    unit: "day"
                },
            type: "timeseries"
        },
        xTickSource: "labels"
    })
    const chartSelfSupplySet = ({ data, labels }) => ({
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
        tooltipLabel: item => item.parsed.y + "%",
        x: {
            grid: { lineWidth: 0 },
            time: tab === "year"
                ? {
                    displayFormats: {
                        month: "YYYY MMM",
                    },
                    tooltipFormat: "YYYY MMM",
                    unit: "month"
                }
                : {
                    displayFormats: {
                        day: "MMM D",
                    },
                    tooltipFormat: "MMM D",
                    unit: "day"
                },
            type: "timeseries"
        },
        y: { min: 0 },
        xTickSource: "labels"
    })
    const
        [tab, setTab] = useState("days"),
        [infoError, setInfoError] = useState(""),
        [infoLoading, setInfoLoading] = useState(false),
        [preInfoError, setPreInfoError] = useState(""),
        [preInfoLoading, setPreInfoLoading] = useState(false),
        [energySourcesTotal, setEnergySourcesTotal] = useState({
            types: [
                { kwh: 0, percentage: 0, type: "importFromGrid" },
                { kwh: 0, percentage: 0, type: "directSolarSupply" },
                { kwh: 0, percentage: 0, type: "batteryDischarge" },
            ],
            kwh: 0
        }),
        [energyDestinations, setEnergyDestinations] = useState({
            types: [
                { kwh: 0, percentage: 0, type: "load" },
                { kwh: 0, percentage: 0, type: "exportToGrid" },
                { kwh: 0, percentage: 0, type: "chargeToBattery" },
            ],
            kwh: 0
        }),
        [preEnergySourcesTotal, setPreEnergySourcesTotal] = useState({
            types: [
                { kwh: 0, percentage: 0, type: "importFromGrid" },
                { kwh: 0, percentage: 0, type: "directSolarSupply" },
                { kwh: 0, percentage: 0, type: "batteryDischarge" },
            ],
            kwh: 0
        }),
        [preEnergyDestinations, setPreEnergyDestinations] = useState({
            types: [
                { kwh: 0, percentage: 0, type: "load" },
                { kwh: 0, percentage: 0, type: "exportToGrid" },
                { kwh: 0, percentage: 0, type: "chargeToBattery" },
            ],
            kwh: 0
        }),
        [lineChartPower, setLineChartPower] = useState(null),
        [lineChartPowerError, setLineChartPowerError] = useState(""),
        [lineChartPowerLoading, setLineChartPowerLoading] = useState(false),
        [lineChartPowerRes] = useState("5minute"),
        [preLineChartPower, setPreLineChartPower] = useState(null),
        [preLineChartPowerError, setPreLineChartPowerError] = useState(""),
        [preLineChartPowerLoading, setPreLineChartPowerLoading] = useState(false),
        [preLineChartPowerRes] = useState("5minute"),
        [barChartData, setBarChartData] = useState(null),
        [barChartDataError, setBarChartDataError] = useState(""),
        [barChartDataLoading, setBarChartDataLoading] = useState(false),
        [lineChartSupply, setLineChartSupply] = useState(null),
        [lineChartSupplyError, setLineChartSupplyError] = useState(""),
        [lineChartSupplyLoading, setLineChartSupplyLoading] = useState(false),
        [startDate, setStartDate] = useState(null),
        [endDate, setEndDate] = useState(null),
        [startMonth, setStartMonth] = useState(new Date())

    const urlPrefix = `/api/${props.gatewayID}/devices`
    const
        callTodayCards = (startTime, endTime) => {
            apiCall({
                onComplete: () => setInfoLoading(false),
                onError: error => setInfoError(error),
                onStart: () => setInfoLoading(true),
                onSuccess: rawData => {
                    if (!rawData?.data) return

                    const { data } = rawData
                    setEnergySourcesTotal({
                        types: [
                            {
                                kwh: data.gridProducedLifetimeEnergyACDiff,
                                percentage: data.gridProducedEnergyPercentAC,
                                type: "importFromGrid"
                            },
                            {
                                kwh: data.pvProducedLifetimeEnergyACDiff,
                                percentage: data.pvProducedEnergyPercentAC,
                                type: "directSolarSupply"
                            },
                            {
                                kwh: data.batteryProducedLifetimeEnergyACDiff,
                                percentage: data.batteryProducedEnergyPercentAC,
                                type: "batteryDischarge"
                            },
                        ],
                        kwh: data.allProducedLifetimeEnergyACDiff
                    })
                    setEnergyDestinations({
                        types: [
                            {
                                kwh: data.loadConsumedLifetimeEnergyACDiff,
                                percentage: data.loadConsumedEnergyPercentAC,
                                type: "load"
                            },
                            {
                                kwh: data.gridConsumedLifetimeEnergyACDiff,
                                percentage: data.gridConsumedEnergyPercentAC,
                                type: "exportToGrid"
                            },
                            {
                                kwh: data.batteryConsumedLifetimeEnergyACDiff,
                                percentage: data.batteryConsumedEnergyPercentAC,
                                type: "chargeToBattery"
                            },
                        ],
                        kwh: data.allConsumedLifetimeEnergyACDiff
                    })
                },
                url: `${urlPrefix}/energy-distribution-info?startTime=${startTime}&endTime=${endTime}`
            })
        },
        callYesterdayCards = (preStartTime, preEndTime) => {
            apiCall({
                onComplete: () => setPreInfoLoading(false),
                onError: error => setPreInfoError(error),
                onStart: () => setPreInfoLoading(true),
                onSuccess: rawData => {
                    if (!rawData?.data) return

                    const { data } = rawData
                    setPreEnergySourcesTotal({
                        types: [
                            {
                                kwh: data.gridProducedLifetimeEnergyACDiff,
                                percentage: data.gridProducedEnergyPercentAC,
                                type: "importFromGrid"
                            },
                            {
                                kwh: data.pvProducedLifetimeEnergyACDiff,
                                percentage: data.pvProducedEnergyPercentAC,
                                type: "directSolarSupply"
                            },
                            {
                                kwh: data.batteryProducedLifetimeEnergyACDiff,
                                percentage: data.batteryProducedEnergyPercentAC,
                                type: "batteryDischarge"
                            },
                        ],
                        kwh: data.allProducedLifetimeEnergyACDiff
                    })
                    setPreEnergyDestinations({
                        types: [
                            {
                                kwh: data.loadConsumedLifetimeEnergyACDiff,
                                percentage: data.loadConsumedEnergyPercentAC,
                                type: "load"
                            },
                            {
                                kwh: data.gridConsumedLifetimeEnergyACDiff,
                                percentage: data.gridConsumedEnergyPercentAC,
                                type: "exportToGrid"
                            },
                            {
                                kwh: data.batteryConsumedLifetimeEnergyACDiff,
                                percentage: data.batteryConsumedEnergyPercentAC,
                                type: "chargeToBattery"
                            },
                        ],
                        kwh: data.allConsumedLifetimeEnergyACDiff
                    })
                },
                url: `${urlPrefix}/energy-distribution-info?startTime=${preStartTime}&endTime=${preEndTime}`
            })
        },
        callLineChartPower = (startTime, endTime) => {
            const lineChartPowerUrl = `${urlPrefix}/power-state?`
                + new URLSearchParams({
                    startTime, endTime, resolution: lineChartPowerRes
                }).toString()

            apiCall({
                onComplete: () => setLineChartPowerLoading(false),
                onError: error => setLineChartPowerError(error),
                onStart: () => setLineChartPowerLoading(true),
                onSuccess: rawData => {
                    if (!rawData || !rawData.data) return

                    const
                        { data } = rawData,
                        { timestamps } = data,
                        labels = timestamps.map(t => t * 1000)
                    setLineChartPower({
                        data: {
                            load: data.loadAveragePowerACs,
                            solar: data.pvAveragePowerACs,
                            battery: data.batteryAveragePowerACs,
                            grid: data.gridAveragePowerACs
                        },
                        labels
                    })
                },
                url: lineChartPowerUrl
            })
        },
        callPreLineChartPower = (preStartTime, preEndTime) => {
            const preLineChartPowerUrl = `${urlPrefix}/power-state?`
                + new URLSearchParams({
                    startTime: preStartTime, endTime: preEndTime,
                    resolution: preLineChartPowerRes
                }).toString()

            apiCall({
                onComplete: () => setPreLineChartPowerLoading(false),
                onError: error => setPreLineChartPowerError(error),
                onStart: () => setPreLineChartPowerLoading(true),
                onSuccess: rawData => {
                    if (!rawData || !rawData.data) return

                    const
                        { data } = rawData,
                        { timestamps } = data,
                        labels = timestamps.map(t => t * 1000)
                    setPreLineChartPower({
                        data: {
                            load: data.loadAveragePowerACs,
                            solar: data.pvAveragePowerACs,
                            battery: data.batteryAveragePowerACs,
                            grid: data.gridAveragePowerACs
                        },
                        labels
                    })
                },
                url: preLineChartPowerUrl
            })
        },
        callBarChartData = (startTime, endTime) => {
            const barChartDataUrl = `${urlPrefix}/accumulated-power-state?`
                + new URLSearchParams({
                    startTime, endTime, resolution: tab == "year" ? "month" : "day"
                }).toString()
            apiCall({
                onComplete: () => setBarChartDataLoading(false),
                onError: error => setBarChartDataError(error),
                onStart: () => setBarChartDataLoading(true),
                onSuccess: rawData => {
                    if (!rawData || !rawData.data) return

                    const
                        { data } = rawData,
                        { timestamps } = data,
                        labels = timestamps.map(t => (t * 1000))

                    setBarChartData({
                        data: {
                            load: data.loadConsumedLifetimeEnergyACDiffs,
                            solar: data.pvProducedLifetimeEnergyACDiffs,
                            battery: data.batteryLifetimeEnergyACDiffs,
                            grid: data.gridLifetimeEnergyACDiffs
                        },
                        labels
                    })
                },
                url: barChartDataUrl
            })
        },
        callLineChartSupply = (startTime, endTime) => {
            const lineChartSupplyUrl = `${urlPrefix}/power-self-supply-rate?`
                + new URLSearchParams({
                    startTime, endTime, resolution: tab == "year" ? "month" : "day"
                }).toString()

            apiCall({
                onComplete: () => setLineChartSupplyLoading(false),
                onError: error => setLineChartSupplyError(error),
                onStart: () => setLineChartSupplyLoading(true),
                onSuccess: rawData => {
                    if (!rawData || !rawData.data) return
                    const
                        { data } = rawData,
                        { timestamps } = data,
                        labels = timestamps.map(t => t * 1000)
                    setLineChartSupply({
                        data: data.loadSelfConsumedEnergyPercentACs,
                        labels
                    })
                },
                url: lineChartSupplyUrl
            })
        }
    useEffect(() => {
        if (!props.gatewayID) return

        let startTime = "", endTime = ""
        let preStartTime = "", preEndTime = ""
        if (tab === "days") {
            startTime = moment().startOf("day").toISOString()
            endTime = moment().toISOString()
            preStartTime = moment().subtract(1, "day").startOf("day").toISOString()
            preEndTime = moment().subtract(1, "day").endOf("day").toISOString()

        } else if (tab === "weeks") {
            startTime = moment().startOf("week").toISOString()
            endTime = moment().startOf("day").toISOString()

            if (moment().get("day") == 0) {
                startTime = moment().subtract(1, "week").startOf("week").toISOString()
            }
        } else if (tab === "month") {
            startTime = moment().startOf("month").toISOString()
            endTime = moment().startOf("day").toISOString()
            if (moment().get("date") == 1) {
                startTime = moment().subtract(1, "month").startOf("month").toISOString()
            }
        } else if (tab === "year") {
            startTime = moment().startOf("year").toISOString()
            endTime = moment().startOf("month").toISOString()
            if (moment().get("date") == 1 && moment().get("month") == 0) {
                startTime = moment().subtract(1, "year").startOf("year").toISOString()
            }
        } else if (tab === "custom") {
            startTime = startDate ? moment(startDate).toISOString() : ""
            endTime = endDate ? moment(endDate).add(1, "day").startOf("day").toISOString() : ""
        }
        if (startTime && endTime) {
            callTodayCards(startTime, endTime)

            if (tab === "days") {
                callLineChartPower(startTime, endTime)
            } else {
                callBarChartData(startTime, endTime)
                callLineChartSupply(startTime, endTime)
            }
        }
        if (preStartTime && preEndTime) {
            callYesterdayCards(preStartTime, preEndTime)

            if (tab === "days") {
                callPreLineChartPower(preStartTime, preEndTime)
            }
        }
    }, [props.gatewayID, tab, startDate, endDate])

    const tabs = ["days", "weeks", "month", "year", "custom"]

    return <>
        <div className="page-header flex flex-wrap justify-between">
            <h1>{pageT("analysis")}</h1>
            <div className="flex flex-wrap">
                {tabs.map((t, i) =>
                    <Button
                        sx={{ margin: "0.5rem 0 0 0.5rem" }}
                        onClick={() => setTab(t)}
                        filter={tab === t ? "selected" : ""}
                        key={"a-t-" + i}
                        radius="pill"
                        variant="contained">
                        {pageT(t)}
                    </Button>)}
            </div>
        </div>
        {tab === "month"
            ? <div className="flex justify-end mb-10 mr-56 relative w-auto">
                <div className="flex items-center">
                    <MonthPicker
                        {...{ startMonth, setStartMonth }}
                    />
                </div>
            </div>
            : null}
        {tab === "custom"
            ? <div className="flex justify-end mb-10 relative w-auto">
                <div className="flex items-center">
                    <DateRangePicker
                        {...{ startDate, setStartDate, endDate, setEndDate }}
                    />
                </div>
            </div>
            : null}
        <div className="gap-8 grid md:grid-cols-2 items-start">
            <AnalysisCard
                data={energySourcesTotal}
                title={pageT("energySourcesTotal")} />
            <AnalysisCard
                data={energyDestinations}
                title={pageT("energyDestinationsTotal")} />
            {infoLoading
                ? <div className="absolute bg-black-main-opacity-95 grid inset-0
                                place-items-center rounded-3xl">
                    <Spinner />
                </div>
                : null}
        </div>
        {tab == "days"
            ? <>
                <div className="card mt-8">
                    <h4>{pageT("todayPowerkW")}</h4>
                    <div className="max-h-80vh h-160 mt-10 relative w-full">
                        <LineChart
                            data={chartRealTimePowerSet({
                                ...lineChartPower
                            })} id="analysisLineChart" />
                        <ErrorBox
                            error={lineChartPowerError}
                            message={pageT("chartError")} />
                        <LoadingBox loading={lineChartPowerLoading} />
                    </div>
                </div>
                <h2 className="mt-10">{pageT("yesterday")}</h2>
                <div className="gap-8 grid md:grid-cols-2 items-start mt-8">
                    <AnalysisCard
                        data={preEnergySourcesTotal}
                        title={pageT("energySourcesTotal")} />
                    <AnalysisCard
                        data={preEnergyDestinations}
                        title={pageT("energyDestinationsTotal")} />
                    {infoLoading
                        ? <div className="absolute bg-black-main-opacity-95 grid inset-0
                                place-items-center rounded-3xl">
                            <Spinner />
                        </div>
                        : null}
                </div>
                <div className="card mt-8">
                    <h4>{pageT("YesterdayPowerkW")}</h4>
                    <div className="max-h-80vh h-160 mt-10 relative w-full">
                        <LineChart
                            data={chartRealTimePowerSet({
                                ...preLineChartPower
                            })} id="analysisLineChartYesterday" />
                        <ErrorBox
                            error={preLineChartPowerError}
                            message={pageT("chartError")} />
                        <LoadingBox loading={preLineChartPowerLoading} />
                    </div>
                </div>
            </>
            : null}

        {tab !== "days"
            ? <>
                <div className="card mt-8">
                    <h4>{pageT("accumulatedKwh")}</h4>
                    <div className="max-h-80vh h-160 mt-8 relative w-full">
                        <BarChart data={chartAccumulatedPowerSet({
                            ...barChartData
                        })} id="analysisBarChart" />
                        <ErrorBox
                            error={barChartDataError}
                            message={pageT("chartError")} />
                        <LoadingBox loading={barChartDataLoading} />
                    </div>
                </div>
                <div className="card chart mt-8">
                    <h4 className="mb-10">{pageT("selfSupplyRate")}</h4>
                    <div className="max-h-80vh h-160 w-full">
                        <LineChart data={chartSelfSupplySet({
                            ...lineChartSupply
                        })} id="anLineChart" />
                        <ErrorBox
                            error={lineChartSupplyError}
                            message={pageT("chartError")} />
                        <LoadingBox loading={lineChartSupplyLoading} />
                    </div>
                </div>
            </>
            : null}
    </>
})