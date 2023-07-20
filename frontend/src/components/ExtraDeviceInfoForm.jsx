
import { InputAdornment, MenuItem, TextField } from "@mui/material"
import { useTranslation } from "react-multi-lang"
import { useEffect, useMemo } from "react"
import { validateNumPercent } from "../utils/utils"

export default function ExtraDeviceInfoForm(props) {
    const { subTitle, gridOutagePercent, setGridOutagePercent, chargingSource,
        setChargingSource, energyCapacity, setEnergyCapacity, voltage, setVoltage
    } = props

    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        formT = (string) => t("form." + string)

    const chargingSourceOptions = [
        {
            "id": 1,
            "name": "solarGrid"
        },
        {
            "id": 2,
            "name": "solar"
        }
    ]
    const
        inputPercent = (e) => {
            const num = e.target.value
            const isNum = validateNumPercent(num)
            if (!isNum) return
            setGridOutagePercent(num)
        },
        changeChargingSource = (e) => {
            setChargingSource(e.target.value)
        },
        changeEnergyCapacity = (e) => {
            setEnergyCapacity(e.target.value)
        },
        changeVoltage = (e) => {
            setVoltage(e.target.value)
        }
    return <>
        <h5 className="mb-5 ml-2">{subTitle}</h5>
        <TextField
            key="r-g-o-p"
            label={formT("reservedForGridOutagePercent")}
            value={gridOutagePercent}
            onChange={inputPercent}
            InputProps={{
                endAdornment:
                    <InputAdornment position="end">%</InputAdornment>
            }}
        />
        <TextField
            key="charging-source"
            select
            label={formT("chargingSource")}
            onChange={changeChargingSource}
            value={chargingSource}
            defaultValue=""

        >
            {chargingSourceOptions.map(({ id, name }) => (
                <MenuItem key={"option-c-f-" + id} value={name}>
                    {formT(`${name}`)}
                </MenuItem>
            ))}
        </TextField>
        <TextField
            key="energy-capacity"
            type="number"
            label={formT("energyCapacity")}
            onChange={changeEnergyCapacity}
            value={energyCapacity}
        />
        <TextField
            key="voltage"
            type="number"
            label={commonT("voltage")}
            onChange={changeVoltage}
            value={voltage}
        />
    </>
}