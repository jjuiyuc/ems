import {
    Button, Divider, InputAdornment, ListItem, MenuItem, Select, TextField
} from "@mui/material"
import { useTranslation } from "react-multi-lang"
import { useEffect, useMemo, useState } from "react"
import { ValidateNumPercent } from "../utils/utils"


export default function ExtraDeviceInfoForm(props) {
    const { subTitle, gridOutagePercent, setGridOutagePercent,
        chargingSource, setChargingSource, energyCapacity, setEnergyCapacity,
        voltage, setVoltage
    } = props

    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        dialogT = (string) => t("dialog." + string),
        formT = (string) => t("form." + string)


    const
        inputPercent = (e) => {
            const num = e.target.value
            const isNum = ValidateNumPercent(num)
            if (!isNum) return
            setGridOutagePercent(num)
        }
    return <>

        <h5 className="mb-5 ml-2">{subTitle}</h5>
        <TextField
            id="reservedForGridOutagePercent"
            label={formT("reservedForGridOutagePercent")}
            value={gridOutagePercent}
            onChange={inputPercent}
            InputProps={{
                endAdornment:
                    <InputAdornment position="end">%</InputAdornment>
            }}
        />
        <TextField
            id="chargingSource"
            select
            label={formT("chargingSource")}
            defaultValue=""
        >
            {chargingSource.map((option) => (
                <MenuItem key={option.value} value={option.value}>
                    {option.label}
                </MenuItem>
            ))}
        </TextField>
        <TextField
            id="energyCapacity"
            type="number"
            label={formT("energyCapacity")}
        // value={energyCapacity}
        />
        <TextField
            id="voltage"
            type="number"
            label={commonT("voltage")}
        // value={voltage}
        />
    </>
}