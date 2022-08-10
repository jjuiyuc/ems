import { connect } from "react-redux"
import moment from "moment"
import ReportProblemIcon from "@mui/icons-material/ReportProblem"
import { useEffect, useState } from "react"
import { useTranslation } from "react-multi-lang"

import { apiCall } from "../utils/api"
import variables from "../configs/variables"

import EnergyResoucesTabs from "../components/EnergyResoucesTabs"
import EnergySolarCard from "../components/EnergySolarCard"
import EnergySolarSubCard from "../components/EnergySolarSubCard"
import LineChart from "../components/LineChart"

import { ReactComponent as SolarIcon } from "../assets/icons/sunny.svg"
import { ReactComponent as EconomicsIcon } from "../assets/icons/economics.svg"
import { ReactComponent as UpIcon } from "../assets/icons/up.svg"
import { ReactComponent as Co2Icon } from "../assets/icons/co2.svg"
const { colors } = variables

export default function EnergyResoucesSolar(props) {
    const
        [solarPower, setSolarPower] = useState(0),
        [totalSolarEnergyDestinations, setTotalSolarEnergyDestinations]
            = useState({
                types: [
                    { kwh: 450, percentage: 50, type: "directUsage" },
                    { kwh: 350, percentage: 40, type: "chargeToBattery" },
                    { kwh: 50, percentage: 10, type: "exportToGrid" },
                ],
                kwh: 60
            }),
        [directUsage, setDirectUsage] = useState(50),
        [chargeToBattery, setChargeToBattery] = useState(0),
        [exportToGrid, setExportToGrid] = useState(5),
        [economics, setEconomics] = useState(85),
        [cO2Reduction, setCO2Reduction] = useState(0)

    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = string => t("energyResources.solar." + string)

    const cardsData = {
        economics: [{
            title: pageT("economics"),
            value: `$${economics}`
        }],
        cO2Reduction: [{
            title: pageT("cO2Reduction"),
            value: `${cO2Reduction} ${pageT("tons")}`
        }]
    }
    const
        hours24 = Array.from(new Array(24).keys()),
        lineChartDateLabels = hours24.map(n =>
            moment().hour(n).startOf("h").toISOString()),
        currentHour = moment().hour(),
        lineChartDataArray = hours24.filter(v => v <= currentHour).map(() =>
            Math.floor(Math.random() * (60 - 40 + 1) + 40))

    const [lineChartData, setLineChartData] = useState({
        datasets: [{
            backgroundColor: colors.yellow.main,
            borderColor: colors.yellow.main,
            data: lineChartDataArray,
            fill: {
                above: colors.yellow["main-opacity-10"],
                target: "origin"
            },
            pointBorderColor: colors.primary["main-opacity-20"]
        }],
        labels: lineChartDateLabels,
        tickCallback: (val, index) => val + commonT("kw"),
        tooltipLabel: item => `${item.parsed.y}` + commonT("kw"),
        x: { grid: { lineWidth: 0 } },
        y: { max: 80, min: 0 }
    })

    return <>
        <h1 className="mb-9">{t("navigator.energyResources")}</h1>
        <EnergyResoucesTabs current="solar" />
        <EnergySolarCard
            data={totalSolarEnergyDestinations}
            title={pageT("totalSolarEnergyDestinations")} />
        <div className="font-bold gap-5 grid lg:grid-cols-2 mt-4">
            <EnergySolarSubCard
                data={cardsData.economics}
                icon={EconomicsIcon}
                title={pageT("economics")} />
            <EnergySolarSubCard
                data={cardsData.cO2Reduction}
                icon={Co2Icon}
                title={pageT("cO2Reduction")} />
        </div>
        <div className="card chart mt-8">
            <h4 className="mb-10">{pageT("realTimeSolarGeneration")}</h4>
            <div className="max-h-80vh h-160 w-full">
                <LineChart data={lineChartData} id="ersLineChart" />
            </div>
        </div>
    </>
}