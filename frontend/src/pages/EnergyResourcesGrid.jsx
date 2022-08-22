import { connect } from "react-redux"
import moment from "moment"
import { useState } from "react"
import { useTranslation } from "react-multi-lang"

import { apiCall } from "../utils/api"
import variables from "../configs/variables"

import EnergyResoucesTabs from "../components/EnergyResoucesTabs"
import EnergyGridCard from "../components/EnergyGridCard"
import LineChart from "../components/LineChart"

import { ReactComponent as GridExportIcon } from "../assets/icons/grid_export.svg"

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
export default function EnergyResourcesGrid(props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = string => t("energyResources.grid." + string)

    const
        [todayGrid, setTodayGrid] = useState([
            { kwh: 0, type: "exportToGrid" },
            { kwh: 100, type: "importFromGrid" },
            { kwh: -10, type: "netExport" }
        ]),
        [thisMonth, setThisMonth] = useState(100)

    const
        hours24 = Array.from(new Array(24).keys()),
        lineChartDateLabels = hours24.map(n =>
            moment().hour(n).startOf("h").toISOString()),
        currentHour = moment().hour(),
        lineChartDataArray = hours24.filter(v => v <= currentHour).map(() =>
            Math.floor(Math.random() * (60 - 40 + 1) + 40))

    const [lineChartData, setLineChartData] = useState({
        beforeDraw: drawHighPeak(7, 19),
        datasets: [{
            backgroundColor: colors.indigo.main,
            borderColor: colors.indigo.main,
            data: lineChartDataArray,
            fill: {
                above: colors.indigo["main-opacity-10"],
                target: "origin"
            },
            pointBorderColor: colors.primary["main-opacity-20"]
        }],
        labels: lineChartDateLabels,
        tickCallback: (val, index) => val + commonT("kw"),
        tooltipLabel: item =>
            `${item.parsed.y} ${commonT("kwh")}`,
        y: { max: 80, min: 0 }
    })

    return <>
        <h1 className="mb-9">{t("navigator.energyResources")}</h1>
        <EnergyResoucesTabs current="grid" />
        <div className="lg:grid grid-cols-auto-19rem gap-x-5">
            <EnergyGridCard data={todayGrid} title={commonT("today")} />
            <div className="card mt-8 lg:m-0">
                <h5 className="font-bold mb-8">{commonT("thisMonth")}</h5>
                <h6 className="font-bold text-white">{pageT("netExport")}</h6>
                <div className="flex justify-between items-center mt-3.5">
                    <h3>{thisMonth} {commonT("kwh")}</h3>
                    <div
                        className="bg-gray-400-opacity-20 grid h-12 w-12
                            place-items-center rounded-full">
                        <GridExportIcon className="text-gray-400 w-8 h-8" />
                    </div>
                </div>
            </div>
        </div>
        <div className="card chart mt-8">
            <h4 className="mb-10">{pageT("gridPowerImport")}</h4>
            <div className="max-h-80vh h-160 w-full">
                <LineChart data={lineChartData} id="ergLineChart" />
            </div>
        </div>
    </>
}