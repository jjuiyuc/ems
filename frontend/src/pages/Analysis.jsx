import { connect } from "react-redux"
import { Button } from "@mui/material"
import { useTranslation } from "react-multi-lang"
import moment from "moment"
import ReportProblemIcon from "@mui/icons-material/ReportProblem"
import { useEffect, useState, useRef } from "react"

import { apiCall } from "../utils/api"
import variables from "../configs/variables"

import AlertBox from "../components/AlertBox"
import AnalysisCard from "../components/AnalysisCard"
import BarChart from "../components/BarChart"
import DateRangePicker from "../components/DateRangePicker"
import PrevDatePicker from "../components/PrevDatePicker"
import LineChart from "../components/LineChart"
import LoadingBox from "../components/LoadingBox"
import MonthPicker from "../components/MonthPicker"

const { colors } = variables

const MOCK_MODE = true

const mockAnalysisData = {
    energySources: {
        allProducedLifetimeEnergyACDiff: 100,
        gridProducedLifetimeEnergyACDiff: 40,
        pvProducedLifetimeEnergyACDiff: 40,
        batteryProducedLifetimeEnergyACDiff: 20,
        gridProducedEnergyPercentAC: 40,
        pvProducedEnergyPercentAC: 40,
        batteryProducedEnergyPercentAC: 20,

        loadConsumedLifetimeEnergyACDiff: 80,
        gridConsumedLifetimeEnergyACDiff: 10,
        batteryConsumedLifetimeEnergyACDiff: 10,
        loadConsumedEnergyPercentAC: 80,
        gridConsumedEnergyPercentAC: 10,
        batteryConsumedEnergyPercentAC: 10,

        allConsumedLifetimeEnergyACDiff: 100,
    },
    powerState: {
        timestamps: [1742749150, 1742752750, 1742756350],
        loadAveragePowerACs: [10, 20, 30],
        pvAveragePowerACs: [5, 15, 25],
        batteryAveragePowerACs: [3, 6, 9],
        gridAveragePowerACs: [2, 4, 6]
    },
    accumulatedPowerState: {
        timestamps: [1742749150, 1742752750, 1742756350],
        loadConsumedLifetimeEnergyACDiffs: [10, 20, 30],
        pvProducedLifetimeEnergyACDiffs: [5, 15, 25],
        batteryLifetimeEnergyACDiffs: [3, 6, 9],
        gridLifetimeEnergyACDiffs: [2, 4, 6]
    },
    selfSupplyRate: {
        timestamps: [1742749150, 1742752750, 1742756350],
        loadSelfConsumedEnergyPercentACs: [50, 60, 70]
    }
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

const mapState = state => ({ gatewayID: state?.gateways?.active?.gatewayID })

export default connect(mapState)(function Analysis(props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("analysis." + string, params)

    const defaultMonth = moment().get("date") == 1
        ? moment().subtract(1, "month").startOf("month")._d
        : new Date()

    const defaultDate = moment().subtract(1, "day").startOf("day")._d

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
        tickCallback: val => parseFloat(val.toFixed(3)) + commonT("kw"),
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
        tickCallback: val => parseFloat(val.toFixed(3)) + " " + commonT("kwh"),
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
            borderWidth: 2,
            hoverRadius: 2,
            pointHoverBorderWidth: 2,
            radius: 2,
        }],
        labels,
        tickCallback: val => parseFloat(val.toFixed(2)) + "%",
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
        [tab, setTab] = useState("day"),
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
        [barChartData, setBarChartData] = useState(null),
        [barChartDataError, setBarChartDataError] = useState(""),
        [barChartDataLoading, setBarChartDataLoading] = useState(false),
        [lineChartSupply, setLineChartSupply] = useState(null),
        [lineChartSupplyError, setLineChartSupplyError] = useState(""),
        [lineChartSupplyLoading, setLineChartSupplyLoading] = useState(false),
        [startDate, setStartDate] = useState(null),
        [endDate, setEndDate] = useState(null),
        [startMonth, setStartMonth] = useState(defaultMonth),
        [prevDate, setPrevDate] = useState(defaultDate),
        timeID = useRef(true)

    const urlPrefix = `/api/${props.gatewayID}/devices`
    const
        callTodayCards = (startTime, endTime) => {
            if (MOCK_MODE) {
                const data = mockAnalysisData.energySources
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
                return
            }
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
        callLineChartPower = (startTime, endTime) => {
            if (MOCK_MODE) {
                const {
                    timestamps,
                    loadAveragePowerACs,
                    pvAveragePowerACs,
                    batteryAveragePowerACs,
                    gridAveragePowerACs
                } = mockAnalysisData.powerState

                const labels = timestamps.map(t => t * 1000)

                setLineChartPower({
                    data: {
                        load: loadAveragePowerACs,
                        solar: pvAveragePowerACs,
                        battery: batteryAveragePowerACs,
                        grid: gridAveragePowerACs
                    },
                    labels
                })

                setLineChartPowerLoading(false)
            } else {
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

                        const { data } = rawData
                        const labels = data.timestamps.map(t => t * 1000)

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
            }
        },
        callBarChartData = (startTime, endTime) => {
            if (MOCK_MODE) {
                const {
                    timestamps,
                    loadConsumedLifetimeEnergyACDiffs,
                    pvProducedLifetimeEnergyACDiffs,
                    batteryLifetimeEnergyACDiffs,
                    gridLifetimeEnergyACDiffs
                } = mockAnalysisData.accumulatedPowerState

                const labels = timestamps.map(t => t * 1000)

                setBarChartData({
                    data: {
                        load: loadConsumedLifetimeEnergyACDiffs,
                        solar: pvProducedLifetimeEnergyACDiffs,
                        battery: batteryLifetimeEnergyACDiffs,
                        grid: gridLifetimeEnergyACDiffs
                    },
                    labels
                })

                setBarChartDataLoading(false)
                return
            }
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
            if (MOCK_MODE) {
                const {
                    timestamps,
                    loadSelfConsumedEnergyPercentACs
                } = mockAnalysisData.selfSupplyRate

                const labels = timestamps.map(t => t * 1000)

                setLineChartSupply({
                    data: loadSelfConsumedEnergyPercentACs,
                    labels
                })

                setLineChartSupplyLoading(false)
                return
            }
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

        if (tab === "day") {
            startTime = moment().startOf("day").toISOString()
            endTime = moment().toISOString()

        } else if (tab === "week") {
            startTime = moment().startOf("week").toISOString()
            endTime = moment().startOf("day").toISOString()

            if (moment().get("day") == 0) {
                startTime = moment().subtract(1, "week").startOf("week").toISOString()
            }
        } else if (tab === "month") {
            startTime = startMonth ? moment(startMonth).startOf("month").toISOString() : ""
            endTime = startMonth ? moment(startMonth).endOf("month").toISOString() : ""

        } else if (tab === "year") {
            startTime = moment().startOf("year").toISOString()
            endTime = moment().startOf("month").toISOString()

            if (moment().get("month") == 0) {
                startTime = moment().subtract(1, "year").startOf("year").toISOString()
            }
        } else if (tab === "custom") {
            startTime = startDate ? moment(startDate).toISOString() : ""
            endTime = endDate ? moment(endDate).add(1, "day").startOf("day").toISOString() : ""
        }
        if (startTime && endTime) {
            callTodayCards(startTime, endTime)

            if (tab === "day") {
                callLineChartPower(startTime, endTime)
            } else {
                callBarChartData(startTime, endTime)
                callLineChartSupply(startTime, endTime)
            }
        }
    }, [props.gatewayID, tab, startDate, endDate, startMonth])

    const tabs = ["day", "week", "month", "year", "custom"]

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
            ? <div className="flex justify-end mb-10 mr-36 relative w-auto">
                <div className="flex items-center">
                    <MonthPicker
                        {...{ startMonth, setStartMonth }}
                    />
                </div>
            </div>
            : null}
        {tab === "custom"
            ? <div className="flex justify-end mb-10 relative w-auto">
                <div className="flex">
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
            <LoadingBox loading={infoLoading} />
        </div>
        {tab == "day"
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
            </>
            : null
        }
        {tab !== "day"
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