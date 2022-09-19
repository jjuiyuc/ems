import { connect } from "react-redux"
import { Button, Stack } from "@mui/material"
import { useTranslation } from "react-multi-lang"
import moment from "moment"
import { useEffect, useState } from "react"

import { apiCall } from "../utils/api"
import { ConvertTimeToNumber } from "../utils/utils"
import variables from "../configs/variables"

import AnalysisCard from "../components/AnalysisCard"
import DateRangePicker from "../components/DateRangePicker"
import BarChart from "../components/BarChart"
import LineChart from "../components/LineChart"
import "../assets/css/dateRangePicker.css"

const { colors } = variables

const mapState = state => ({ gatewayID: state.gateways.active.gatewayID })

export default connect(mapState)(function Analysis(props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("analysis." + string, params)

    const fakeDataArray = amount => Array.from(new Array(amount).keys())
        .map(() => Math.floor(Math.random() * (40 - 10 + 1) + 10))

    const
        days = 7,
        sevenDays = Array.from(new Array(days).keys()).map(n =>
            moment().subtract(days - n, "d").startOf("day").toISOString()),
        fakeData1 = fakeDataArray(days),
        fakeData2 = fakeDataArray(days),
        fakeData3 = fakeDataArray(days),
        fakeData4 = fakeDataArray(days)


    const chartRealTimePowerSet = ({ data, labels }) => ({
        datasets: [
            {
                backgroundColor: colors.green.main,
                borderColor: colors.green.main,
                data: data?.load || [],
                fill: {
                    above: colors.green["main-opacity-10"],
                    target: "origin"
                },
                id: "load",
                pointBorderColor: colors.green["main-opacity-20"],
                label: pageT("load")
            },
            {
                backgroundColor: colors.yellow.main,
                borderColor: colors.yellow.main,
                data: data?.solar || [],
                fill: {
                    above: colors.yellow["main-opacity-10"],
                    target: "origin"
                },
                id: "solar",
                pointBorderColor: colors.yellow["main-opacity-20"],
                label: commonT("solar")
            },
            {
                backgroundColor: colors.blue.main,
                borderColor: colors.blue.main,
                data: data?.battery || [],
                fill: {
                    above: colors.yellow["main-opacity-10"],
                    target: "origin"
                },
                id: "battery",
                pointBorderColor: colors.yellow["main-opacity-20"],
                label: commonT("battery")
            },
            {
                backgroundColor: colors.indigo.main,
                borderColor: colors.indigo.main,
                data: data?.grid || [],
                fill: {
                    above: colors.indigo["main-opacity-10"],
                    target: "origin"
                },
                id: "grid",
                pointBorderColor: colors.indigo["main-opacity-20"],
                label: commonT("grid")
            }
        ],
        labels,
        legend: true,
        tickCallback: (val, index) => val + commonT("kw"),
        tooltipLabel: item => `${item.dataset.label} ${item.parsed.y} `
            + commonT("kwh"),
        x: { grid: { lineWidth: 0 } },
        y: { max: 60, min: 0 }
    })

    const
        [barChartData, setBarChartData] = useState({
            datasets: [
                {
                    backgroundColor: colors.green.main,
                    data: fakeData1,
                    label: pageT("load")
                },
                {
                    backgroundColor: colors.yellow.main,
                    data: fakeData2,
                    label: commonT("solar")
                },
                {
                    backgroundColor: colors.blue.main,
                    data: fakeData3,
                    label: commonT("battery")
                },
                {
                    backgroundColor: colors.indigo.main,
                    data: fakeData4,
                    label: commonT("grid")
                }
            ],
            labels: sevenDays,
            tooltipLabel: item =>
                `${item.dataset.label} ${item.parsed.y} ${commonT("kwh")}`,
            y: { max: 100, min: 0 }
        }),
        [ssrLineChartData, setSsrLineChartData] = useState({
            datasets: [{
                backgroundColor: colors.primary.main,
                borderColor: colors.primary.main,
                data: fakeData1,
                percent: fakeData1,
                fill: {
                    above: colors.primary["main-opacity-10"],
                    target: "origin"
                },
                pointBorderColor: colors.primary["main-opacity-20"]
            }],
            labels: sevenDays,
            tickCallback: (val, index) => val + "%",
            tooltipLabel: item => item.dataset.percent[item.dataIndex] + "%",
            x: {
                grid: { lineWidth: 0 },
                time: {
                    displayFormats: {
                        day: "MMM D",
                    },
                    tooltipFormat: "MMM D",
                    unit: "day"
                }
            },
            y: { max: 100, min: 0 }
        })

    const
        [tab, setTab] = useState("days"),
        [infoError, setInfoError] = useState(""),
        [infoLoading, setInfoLoading] = useState(false),
        [energySourcesTotal, setEnergySourcesTotal] = useState({
            types: [
                { kwh: 0, percentage: 0, type: "directSolarSupply" },
                { kwh: 0, percentage: 0, type: "importFromGrid" },
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
        [lineChartPower, setLineChartPower] = useState(null),
        [lineChartPowerError, setLineChartPowerError] = useState(""),
        [lineChartPowerLoading, setLineChartPowerLoading] = useState(false),
        [lineChartPowertRes] = useState("hour")

    useEffect(() => {
        if (!props.gatewayID) return

        const
            startTime = moment().startOf("day").toISOString(),
            endTime = moment().toISOString(),
            chartParams = resolution => new URLSearchParams({
                startTime,
                endTime: moment().toISOString(),
                resolution
            }).toString(),
            urlPrefix = `/api/${props.gatewayID}/devices`

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
                            kwh: data.pvProducedLifetimeEnergyACDiff,
                            percentage: data.pvProducedEnergyPercentAC,
                            type: "directSolarSupply"
                        },
                        {
                            kwh: data.gridProducedLifetimeEnergyACDiff,
                            percentage: data.gridProducedEnergyPercentAC,
                            type: "importFromGrid"
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
                    kwh: 60
                })
            },
            url: `${urlPrefix}/energy-distribution-info?startTime=${startTime}&endTime=${endTime}`
        })

        const oClocks = Array.from(new Array(25).keys()).map(n =>
            parseInt(moment().hour(n).startOf("h").format("x")))

        const lineChartPowerUrl = `${urlPrefix}/power-state?`
            + chartParams(lineChartPowertRes)

        apiCall({
            onComplete: () => setLineChartPowerLoading(false),
            onError: error => setLineChartPowerError(error),
            onStart: () => setLineChartPowerLoading(true),
            onSuccess: rawData => {
                if (!rawData || !rawData.data) return

                const
                    { data } = rawData,
                    { timestamps } = data,
                    labels = [
                        ...timestamps.map(t => t * 1000),
                        ...oClocks.slice(timestamps.length)
                    ]
                console.log(data?.loadAveragePowerACs)
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
    }, [props.gatewayID])

    const tabs = ["days", "weeks", "month", "year", "custom"]

    return <>
        <div className="page-header">
            <h1>{pageT("analysis")}</h1>
            <Stack direction="row" justifyContent="flex-end" spacing={1.5}>
                {tabs.map((t, i) =>
                    <Button
                        onClick={() => setTab(t)}
                        filter={tab === t ? "selected" : ""}
                        key={"a-t-" + i}
                        radius="pill"
                        variant="contained">
                        {pageT(t)}
                    </Button>)}
            </Stack>
        </div>
        {tab === "custom"
            ? <div className="flex justify-end mb-10 relative w-auto">
                <div className="flex items-center">
                    <DateRangePicker />
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
        </div>
        {tab == "days"
            ? <div className="card mt-8">
                <h4>{pageT("realTimePowerkW")}</h4>
                <div className="max-h-80vh h-160 mt-10 relative w-full">
                    <LineChart
                        data={chartRealTimePowerSet({
                            ...lineChartPower
                        })} id="analysisLineChart" />
                </div>
            </div>
            : null}
        {tab !== "days"
            ? <>
                <div className="card mt-8">
                    <h4>{pageT("accumulatedKwh")}</h4>
                    <div className="max-h-80vh h-160 mt-8 relative w-full">
                        <BarChart data={barChartData} id="analysisBarChart" />
                    </div>
                </div>
                <div className="card chart mt-8">
                    <h4 className="mb-10">{pageT("selfSupplyRate")}</h4>
                    <div className="max-h-80vh h-160 w-full">
                        <LineChart data={ssrLineChartData} id="anLineChart" />
                    </div>
                </div>
            </>
            : null}
    </>
})