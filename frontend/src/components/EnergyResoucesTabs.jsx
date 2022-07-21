import { Button, ButtonGroup} from "@mui/material"
import { useTranslation } from "react-multi-lang"

import { ReactComponent as BatteryIcon } from "../assets/icons/battery.svg"
import { ReactComponent as GridIcon } from "../assets/icons/grid.svg"
import { ReactComponent as SolarIcon } from "../assets/icons/sunny.svg"

export default function EnergyResoucesTabs (props) {
    const t = useTranslation()

    const tabs = {
        solar: {color: "yellow", icon: SolarIcon},
        battery: {color: "blue", icon: BatteryIcon},
        grid: {color: "indigo", icon: GridIcon}
    }

    return <ButtonGroup amount="3" className="mb-8" variant="contained ubiik">
    {Object.keys(tabs).map((key, i) => {
        const {color} = tabs[key], Icon = tabs[key].icon

        return <Button
                aria-current={key === props.current}
                key={"ert-" + i}>
            <Icon className={`h-5 mr-2 text-${color}-main w-5`} />
            {t(`common.${key}`)}
        </Button>
    })}
    </ButtonGroup>
}