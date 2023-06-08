import { InputAdornment, TextField } from "@mui/material"
import { useTranslation } from "react-multi-lang"

export default function InfoExtraDeviceForm(props) {
    const { subTitle, voltage, energyCapacity, chargingSource, gridOutagePercent } = props

    const
        t = useTranslation(),
        formT = (string) => t("form." + string)

    return <>
        <h5 className="mb-5 ml-2">{subTitle}</h5>
        <TextField
            key="r-g-o-p"
            label={formT("reservedForGridOutagePercent")}
            value={gridOutagePercent}
            InputProps={{
                endAdornment:
                    <InputAdornment position="end">%</InputAdornment>
            }}
            disabled={true}
        />
        <TextField
            key="charging-source"
            type="string"
            label={formT("chargingSource")}
            value={chargingSource}
            disabled={true}
        />
        <TextField
            key="energy-capacity"
            type="number"
            label={formT("energyCapacity")}
            value={energyCapacity}
            disabled={true}
        />
        <TextField
            key="voltage"
            type="number"
            label={t("common.voltage")}
            value={voltage}
            disabled={true}
        />
    </>
}