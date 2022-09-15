import { Button, Stack } from "@mui/material"
import { getLanguage, useTranslation } from "react-multi-lang"
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

export default function Analysis() {
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
    const
        hours24 = Array.from(new Array(24).keys()),
        lineChartDateLabels = hours24.map(n =>
            moment().hour(n).startOf("h").toISOString())

    const [lineChartPower, setLineChartPower] = useState({
        datasets: [
            {
                backgroundColor: colors.green.main,
                borderColor: colors.green.main,
                data: fakeData1,
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
                data: fakeData2,
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
                data: fakeData3,
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
                data: fakeData4,
                fill: {
                    above: colors.indigo["main-opacity-10"],
                    target: "origin"
                },
                id: "grid",
                pointBorderColor: colors.indigo["main-opacity-20"],
                label: commonT("grid")
            }
        ],
        labels: lineChartDateLabels,
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
        // [energySourcesTotal, setEnergySourcesTotal] = useState({
        //     types: [
        //         { kwh: 0, percentage: 0, type: "directSolarSupply" },
        //         { kwh: 0, percentage: 0, type: "importFromGrid" },
        //         { kwh: 0, percentage: 0, type: "batteryDischarge" },
        //     ],
        //     kwh: 50
        // }),
        // [energyDestinationsTotal, setEnergyDestinationsTotal] = useState({
        //     types: [
        //         { kwh: 10, percentage: 18, type: "load" },
        //         { kwh: 25, percentage: 41, type: "exportToGrid" },
        //         { kwh: 25, percentage: 41, type: "chargeToBattery" },
        //     ],
        //     kwh: 60
        // })
        [energySourcesTotal, setEnergySourcesTotal] = useState({
            types: [
                { kwh: 7.5, percentage: 15, type: "directSolarSupply" },
                { kwh: 30, percentage: 60, type: "importFromGrid" },
                { kwh: 12.5, percentage: 25, type: "batteryDischarge" },
            ],
            kwh: 50
        }),
        [energyDestinationsTotal, setEnergyDestinationsTotal] = useState({
            types: [
                { kwh: 10, percentage: 18, type: "load" },
                { kwh: 25, percentage: 41, type: "exportToGrid" },
                { kwh: 25, percentage: 41, type: "chargeToBattery" },
            ],
            kwh: 60
        }),
        [lineChartRes] = useState("hour")


    const lang = getLanguage()

    useEffect(() => {

        // if (!props.gatewayID) return

        // const
        //     startTime = moment().startOf("day").toISOString(),
        //     chartParams = resolution => new URLSearchParams({
        //         startTime,
        //         endTime: moment().endOf("day").toISOString(),
        //         resolution
        //     }).toString(),
        //     urlPrefix = `/api/${props.gatewayID}/devices`

        // apiCall({
        //     onSuccess: rawData => {
        //         if (!rawData?.data) return

        //         const { data } = rawData
        //         setTotalEnergySources({
        //             types: [
        //                 {
        //                     kwh: data.pvProducedLifetimeEnergyACDiff,
        //                     percentage: data.pvProducedEnergyPercentAC,
        //                     type: "directSolarSupply"
        //                 },
        //                 {
        //                     kwh: data.gridProducedLifetimeEnergyACDiff,
        //                     percentage: data.gridProducedEnergyPercentAC,
        //                     type: "importFromGrid"
        //                 },
        //                 {
        //                     kwh: data.batteryProducedLifetimeEnergyACDiff,
        //                     percentage: data.batteryProducedEnergyPercentAC,
        //                     type: "batteryDischarge"
        //                 },
        //             ],
        //             kwh: data.allProducedLifetimeEnergyACDiff
        //         })

        //         setEnergyDestinations({
        //             types: [
        //                 {
        //                     kwh: data.loadConsumedLifetimeEnergyACDiff,
        //                     percentage: data.loadConsumedEnergyPercentAC8,
        //                     type: "load"
        //                 },
        //                 {
        //                     kwh: data.gridConsumedLifetimeEnergyACDiff,
        //                     percentage: data.gridConsumedEnergyPercentAC,
        //                     type: "exportToGrid"
        //                 },
        //                 {
        //                     kwh: data.batteryConsumedLifetimeEnergyACDiff,
        //                     percentage: data.batteryConsumedEnergyPercentAC,
        //                     type: "chargeToBattery"
        //                 },
        //             ],
        //             kwh: 60
        //         })
        //     },
        //     url: `${urlPrefix}/energy-distribution-info?startTime=${startTime}&endTime=${endTime}`
        // })

        const oClocks = Array.from(new Array(25).keys()).map(n =>
            parseInt(moment().hour(n).startOf("h").format("x")))

        //day tab line chart
        // const lineChartUrl = `${urlPrefix}/power-state?`
        //     + chartParams(chargeVoltageRes)

        // apiCall({
        //     // onComplete: () => setChargeVoltageLoading(false),
        //     // onError: error => setChargeVoltageError(error),
        //     // onStart: () => setChargeVoltageLoading(true),
        //     onSuccess: rawData => {
        //         if (!rawData || !rawData.data) return

        //         const
        //             { data } = rawData,
        //             { onPeakTime, timestamps } = data,
        //             { end, start, timezone } = onPeakTime,
        //             labels = [
        //                 ...timestamps.map(t => t * 1000),
        //                 ...oClocks.slice(timestamps.length)
        //             ],
        //             peakStart = ConvertTimeToNumber(start, timezone),
        //             peakEnd = ConvertTimeToNumber(end, timezone)

        //         setLineChartData({
        //             data: {
        //                 load: data.loadAveragePowerACs,
        //                 solar: data.pvAveragePowerACs,
        //                 battery: data.batteryAveragePowerACs,
        //                 grid: data.gridAveragePowerACs
        //             },
        //             highPeak: { start: peakStart, end: peakEnd },
        //             labels
        //         })
        //     },
        //     url: chargeVoltageUrl
        // })
        const
            lineChart = { ...lineChartPower },
            barChart = { ...barChartData },
            labels = [
                pageT("load"),
                commonT("solar"),
                commonT("battery"),
                commonT("grid")
            ]

        labels.forEach((text, i) => lineChart.datasets[i].label = text)
        labels.forEach((text, i) => barChart.datasets[i].label = text)
        setLineChartPower(lineChart)
        setBarChartData(barChart)

    }, [lang])

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
                data={energyDestinationsTotal}
                title={pageT("energyDestinationsTotal")} />
        </div>
        {tab == "days"
            ? <div className="card mt-8">
                <h4>{pageT("realTimePowerkW")}</h4>
                <div className="max-h-80vh h-160 mt-10 relative w-full">
                    <LineChart
                        data={lineChartPower} id="analysisLineChart" />
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
}