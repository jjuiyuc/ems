import { connect } from "react-redux"
import moment from "moment"
import { useEffect, useState } from "react"
import { useTranslation } from "react-multi-lang"

import { apiCall } from "../utils/api"
import variables from "../configs/variables"

import EnergyResourcesTabs from "../components/EnergyResourcesTabs"
import EnergySolarCard from "../components/EnergySolarCard"
import EnergySolarSubCard from "../components/EnergySolarSubCard"
import LineChart from "../components/LineChart"
import Spinner from "../components/Spinner"

import { ReactComponent as EconomicsIcon } from "../assets/icons/economics.svg"
import { ReactComponent as UpIcon } from "../assets/icons/up.svg"
import { ReactComponent as DownIcon } from "../assets/icons/down.svg"
import { ReactComponent as CO2Icon } from "../assets/icons/co2.svg"

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
const mapState = state => ({ gatewayID: state.gateways.active.gatewayID })
export default connect(mapState)(function EnergyResoucesSolar(props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = string => t("energyResources.solar." + string)

    const
        [infoError, setInfoError] = useState(""),
        [infoLoading, setInfoLoading] = useState(false),
        [totalSolarEnergyDestinations, setTotalSolarEnergyDestinations]
            = useState({
                types: [
                    { kwh: 0, percentage: 0, type: "directUsage" },
                    { kwh: 0, percentage: 0, type: "chargeToBattery" },
                    { kwh: 0, percentage: 0, type: "exportToGrid" },
                ],
                kwh: 0
            }),
        [economics, setEconomics] = useState(0),
        [co2Reduction, setCO2Reduction] = useState(0)

    const
        isEcoPositive = economics > 0,
        EcoIcon = isEcoPositive ? UpIcon : DownIcon,
        ecoValue =
            <span className={isEcoPositive ? "text-success-main" : ""}>
                {isEcoPositive ? "+" : "-"} ${Math.abs(economics)}
                <EcoIcon className="h-8 inline-block ml-2 w-8" />
            </span>

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
            backgroundColor: colors.yellow.main,
            borderColor: colors.yellow.main,
            data: lineChartDataArray,
            label: commonT("solar"),
            fill: {
                above: colors.yellow["main-opacity-10"],
                target: "origin"
            },
            pointBorderColor: colors.yellow["main-opacity-20"]
        }],
        labels: lineChartDateLabels,
        tickCallback: (val, index) => val + commonT("kw"),
        tooltipLabel: item =>
            `${item.dataset.label} ${item.parsed.y} ${commonT("kwh")}`,
        y: { max: 80, min: 0 }
    })
    useEffect(() => {
        if (!props.gatewayID) return

        const
            startTime = moment().startOf("day").toISOString(),
            urlPrefix = `/api/${props.gatewayID}/devices/solar`

        apiCall({
            onComplete: () => setInfoLoading(false),
            onError: error => setInfoError(error),
            onStart: () => setInfoLoading(true),
            onSuccess: rawData => {
                if (!rawData?.data) return

                const { data } = rawData

                setTotalSolarEnergyDestinations({
                    types: [
                        {
                            kwh: data.loadPvConsumedLifetimeEnergyACDiff,
                            percentage: data.loadPvConsumedEnergyPercentAC,
                            type: "directUsage"
                        },
                        {
                            kwh: data.batteryPvConsumedLifetimeEnergyACDiff,
                            percentage: data.batteryPvConsumedEnergyPercentAC,
                            type: "chargeToBattery"
                        },
                        {
                            kwh: data.gridPvConsumedLifetimeEnergyACDiff,
                            percentage: data.gridPvConsumedEnergyPercentAC,
                            type: "exportToGrid"
                        },
                    ],
                    kwh: data.allConsumedLifetimeEnergyACDiff
                })
                setEconomics(data.pvEnergyCostSavingsDiff || 0)
                setCO2Reduction(data.pvCo2SavingsDiff || 0)
            },
            url: `${urlPrefix}/energy-info?startTime=${startTime}`
        })
    }, [props.gatewayID])
    return <>
        <h1 className="mb-9">{t("navigator.energyResources")}</h1>
        <EnergyResourcesTabs current="solar" />
        <EnergySolarCard
            data={totalSolarEnergyDestinations}
            title={pageT("totalSolarEnergyDestinations")} />
        <div className="font-bold gap-5 grid md:grid-cols-2 mt-4">
            <EnergySolarSubCard
                icon={EconomicsIcon}
                subTitle={pageT("thisCalendarMonth")}
                title={pageT("economics")}
                value={ecoValue} />
            <EnergySolarSubCard
                icon={CO2Icon}
                title={pageT("co2Reduction")}
                value={co2Reduction + " " + pageT("tons")} />
        </div>
        <div className="card chart mt-8">
            <h4 className="mb-10">{pageT("realTimeSolarGeneration")}</h4>
            <div className="max-h-80vh h-160 relative w-full">
                <LineChart data={lineChartData} id="ersLineChart" />
            </div>
        </div>
    </>
})