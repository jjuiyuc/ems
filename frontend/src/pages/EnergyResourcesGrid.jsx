import { connect } from "react-redux"
import moment from "moment"
import { useEffect, useState } from "react"
import { useTranslation } from "react-multi-lang"

import { apiCall } from "../utils/api"

import EnergyResoucesTabs from "../components/EnergyResoucesTabs"
import EnergyGridCard from "../components/EnergyGridCard"
// import EnergySolarSubCard from "../components/EnergySolarSubCard"
import LineChart from "../components/LineChart"

export default function EnergyResourcesGrid(props) {
    const
        [gridPower, setGridPower] = useState(0),
        [todayGrid, setTodayGrid]
            = useState({
                types: [
                    { kwh: 0, type: "exportToGrid" },
                    { kwh: 100, type: "importToGrid" },
                    { kwh: -10, type: "netExport" }
                ]
            }),
        [thisMonthNetExport, setThisMonthNetExport] = useState(100)

    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = string => t("energyResources.grid." + string)


    return <>
        <h1 className="mb-9">{t("navigator.energyResources")}</h1>
        <EnergyResoucesTabs current="grid" />
        <EnergyGridCard
            data={todayGrid}
            title={commonT("today")} />
    </>
}