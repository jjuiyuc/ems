import { connect } from "react-redux"
import moment from "moment"
import { useState, useEffect } from "react"
import { useTranslation } from "react-multi-lang"

import { apiCall } from "../utils/api"
import { ConvertTimeToNumber } from "../utils/utils"
import variables from "../configs/variables"

import EnergyResourcesTabs from "../components/EnergyResourcesTabs"
import EnergyGridCard from "../components/EnergyGridCard"
import LineChart from "../components/LineChart"
import Spinner from "../components/Spinner"

import { ReactComponent as GridExportIcon } from "../assets/icons/grid_export.svg"

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
const mapState = state => ({ gatewayID: state.gateways.active.gatewayID })

export default connect(mapState)(function EnergyResourcesGrid(props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = string => t("energyResources.grid." + string)

    const
        [infoError, setInfoError] = useState(""),
        [infoLoading, setInfoLoading] = useState(false),
        [todayGrid, setTodayGrid] = useState([
            { kwh: 0, type: "exportToGrid" },
            { kwh: 0, type: "importFromGrid" },
            { kwh: 0, type: "netImport" }
        ]),
        [thisMonth, setThisMonth] = useState(0),
        [lineChartGridPower, setLineChartGridPower] = useState(null),
        [lineChartGridPowerError, setLineChartGridPowerError] = useState(""),
        [lineChartGridPowerLoading, setLineChartGridPowerLoading] = useState(false),
        [lineChartGridPowerRes] = useState("5minute")

    const chartGridPowerSet = ({ data, highPeak, labels }) => ({
        beforeDraw: drawHighPeak(highPeak.start, highPeak.end),
        datasets: [{
            backgroundColor: colors.indigo.main,
            borderColor: colors.indigo.main,
            data: data || [],
            fill: {
                above: colors.indigo["main-opacity-10"],
                target: "origin"
            },
            pointBorderColor: colors.primary["main-opacity-20"],
            hoverRadius: 2,
            pointHoverBorderWidth: 2,
            radius: 2
        }],
        labels,
        tickCallback: (val, index) => val + commonT("kw"),
        tooltipLabel: item =>
            `${item.parsed.y} ${commonT("kw")}`,
        y: { max: 40, min: 0 },
        x: {
            max: moment().add(1, "day").startOf("day"),
            min: moment().startOf("day")
        }
    })

    useEffect(() => {
        if (!props.gatewayID) return
        const
            startTime = moment().startOf("day").toISOString(),
            endTime = moment().toISOString(),
            chartParams = resolution => new URLSearchParams({
                startTime,
                endTime,
                resolution
            }).toString(),
            urlPrefix = `/api/${props.gatewayID}/devices/grid`

        apiCall({
            onComplete: () => setInfoLoading(false),
            onError: error => setInfoError(error),
            onStart: () => setInfoLoading(true),
            onSuccess: rawData => {
                if (!rawData || !rawData.data) return

                const { data } = rawData
                setTodayGrid([
                    {
                        kwh: data.gridConsumedLifetimeEnergyACDiff,
                        type: "exportToGrid"
                    },
                    {
                        kwh: data.gridProducedLifetimeEnergyACDiff,
                        type: "importFromGrid"
                    },
                    {
                        kwh: data.gridLifetimeEnergyACDiff,
                        type: "netImport"
                    }
                ])
                setThisMonth(data.gridLifetimeEnergyACDiffOfMonth || 0)
            },
            url: `${urlPrefix}/energy-info?startTime=${startTime}&endTime=${endTime}`
        })

        const oClocks = Array.from(new Array(25).keys()).map(n =>
            parseInt(moment().hour(n).startOf("h").format("x")))

        const GridPowerUrl = `${urlPrefix}/power-state?`
            + chartParams(lineChartGridPowerRes)
        apiCall({
            onComplete: () => setLineChartGridPowerLoading(false),
            onError: error => setLineChartGridPowerError(error),
            onStart: () => setLineChartGridPowerLoading(true),
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

                setLineChartGridPower({
                    data: data.gridAveragePowerACs,
                    highPeak: { start: peakStart, end: peakEnd },
                    labels
                })
            },
            url: GridPowerUrl
        })
    }, [props.gatewayID])

    const gridPowerChart = lineChartGridPower
        ? <LineChart
            data={chartGridPowerSet(lineChartGridPower)}
            id="erbChargeVoltage" />
        : null

    const infoErrorBox = <ErrorBox
        error={infoError}
        margin="mb-8"
        message={pageT("infoError")} />

    return <>
        <h1 className="mb-9">{t("navigator.energyResources")}</h1>
        <EnergyResourcesTabs current="grid" />
        <div className="lg:grid grid-cols-auto-19rem gap-x-5">
            {infoErrorBox}
            <EnergyGridCard data={todayGrid} title={commonT("today")} />
            <div className="card mt-8 lg:m-0">
                <h5 className="font-bold mb-8">{commonT("thisMonth")}</h5>
                <h6 className="font-bold text-white">{pageT("netImport")}</h6>
                <div className="flex justify-between items-center mt-3.5">
                    <h3>{thisMonth} {commonT("kwh")}</h3>
                    <div
                        className="bg-gray-400-opacity-20 grid h-12 w-12
                            place-items-center rounded-full">
                        <GridExportIcon className="text-gray-400 w-8 h-8" />
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
        <div className="card chart mt-8">
            <h4 className="mb-10">{pageT("gridPowerImport")}</h4>
            <div className="max-h-80vh h-160 w-full">
                <ErrorBox error={lineChartGridPowerError} message={pageT("chartError")} />
                <LoadingBox loading={lineChartGridPowerLoading} />
                {gridPowerChart}
            </div>
        </div>
    </>
})