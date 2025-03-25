import { Button, ButtonGroup } from "@mui/material"
import { NavLink } from "react-router-dom"
import { useTranslation } from "react-multi-lang"

import { ReactComponent as BatteryIcon } from "../assets/icons/battery.svg"
import { ReactComponent as GridIcon } from "../assets/icons/grid.svg"
import { ReactComponent as SolarIcon } from "../assets/icons/sunny.svg"

export default function EnergyResoucesTabs(props) {
    const
        tabs = {
            solar: { icon: SolarIcon, path: "solar", color: "yellow" },
            battery: { icon: BatteryIcon, path: "battery", color: "blue", },
            grid: { icon: GridIcon, path: "grid", color: "indigo" },
        }

    const t = useTranslation()

    return <ButtonGroup amount="3" className="mb-8" variant="contained ubiik">
        {Object.keys(tabs).map((key, i) => {
            const { color, path } = tabs[key], Icon = tabs[key].icon

            return <Button
                aria-current={key === props.current}
                component={NavLink}
                key={"ert-" + i}
                to={"/energy-resources/" + path}>
                <Icon className={`h-5 mr-2 text-${color}-main w-5`} />
                {t(`common.${key}`)}
            </Button>
        })}
    </ButtonGroup>
}
