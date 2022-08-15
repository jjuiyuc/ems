import { connect } from "react-redux"
import { Link, NavLink } from "react-router-dom"
import { Button, ButtonGroup, Fade, Tooltip } from "@mui/material"
import { useTranslation } from "react-multi-lang"

import { ReactComponent as BatteryIcon } from "../assets/icons/battery.svg"
import { ReactComponent as GridIcon } from "../assets/icons/grid.svg"
import { ReactComponent as SolarIcon } from "../assets/icons/sunny.svg"

export default function EnergyResoucesTabs(props) {
    const toggle = () =>
        props.updateEnergyResoucesTabsStatus(isExpanded ? "collapse" : "expand")

    const
        { status } = props,
        isExpanded = status === "expand",
        transitionClasses = "duration-300 transition "
            + (isExpanded ? "opacity-100" : "opacity-0")
    const
        tabs = {
            solar: { icon: SolarIcon, path: "solar", color: "yellow" },
            battery: { icon: BatteryIcon, path: "battery", color: "blue", },
            grid: { icon: GridIcon, path: "#", color: "indigo" },
        }

    const t = useTranslation(),
        commonT = string => t("common." + string)


    return <ButtonGroup amount="3" className="mb-8" variant="contained ubiik">
        {Object.keys(tabs).map((key, i) => {
            const { color, path } = tabs[key], Icon = tabs[key].icon

            return <Button
                aria-current={key === props.current}
                href={"/energy-resources/" + path}
                key={"ert-" + i}>

                <Icon className={`h-5 mr-2 text-${color}-main w-5`} />
                {t(`common.${key}`)}
            </Button>
        })}
    </ButtonGroup>
}
