
import { InputAdornment, MenuItem, TextField } from "@mui/material"
import { useTranslation } from "react-multi-lang"
import { useEffect, useMemo, useState } from "react"
import { validateNumPercent } from "../utils/utils"

export default function ExtraDeviceInfoForm(props) {
    const { subTitle, voltage, energyCapacity, chargingSource, gridOutagePercent,
        setVoltage, setEnergyCapacity, setChargingSource, setGridOutagePercent
    } = props

    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        formT = (string) => t("form." + string)

    const
        inputPercent = (e) => {
            const num = e.target.value
            const isNum = validateNumPercent(num)
            if (!isNum) return
            setGridOutagePercent(num)
        }
    return <>
        <h5 className="mb-5 ml-2">{props.subTitle}</h5>
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
            // value={chargingSource}
            defaultValue=""

        >
            {chargingSource.map((option) => (
                <MenuItem key={option.value} value={option.value}>
                    {option.label}
                </MenuItem>
            ))}
        </TextField>
        <TextField
            key="energy-capacity"
            type="number"
            label={formT("energyCapacity")}
        // value={energyCapacity}
        />
        <TextField
            key="voltage"
            type="number"
            label={commonT("voltage")}
        // value={voltage}
        />
    </>
}