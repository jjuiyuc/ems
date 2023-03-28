import {
    Button, Divider, InputAdornment, ListItem, MenuItem, TextField
} from "@mui/material"
import { useTranslation } from "react-multi-lang"
import { useEffect, useState } from "react"
import { validateNumTwoDecimalPlaces } from "../utils/utils"


export default function SubDeviceForm({
    title,
    children = null
}) {

    const
        subDeviceData = [
            {
                deviceType: "meter",
                deviceModel: [
                    {
                        value: "CMO336 CM Meter",
                        label: "CMO336 CM Meter"
                    },
                    {
                        value: "CMO337 CM Meter",
                        label: "CMO337 CM Meter"
                    }
                ],
            },
            {
                deviceType: "pv",
                deviceModel: [
                    {
                        value: "D1K330H3A URE PV",
                        label: "D1K330H3A URE PV"
                    },
                    {
                        value: "D1K330H4B URE PV ",
                        label: "D1K330H4B URE PV "
                    }
                ],
            },
            {
                deviceType: "battery",
                deviceModel: [
                    {
                        value: "L051100-A UZ-Energy Battery",
                        label: "L051100-A UZ-Energy Battery"
                    },
                    {
                        value: "L051101-B UZ-Energy Battery",
                        label: "L051101-B UZ-Energy Battery"
                    }
                ],
            }
        ],
        subDeviceModel = [
            {
                value: "CMO336 CM Meter",
                label: "CMO336 CM Meter"
            },
            {
                value: "D1K330H3A URE PV ",
                label: "D1K330H3A URE PV "
            },
            {
                value: "L051100-A UZ-Energy Battery",
                label: "L051100-A UZ-Energy Battery"
            }
        ]
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        dialogT = (string) => t("dialog." + string),
        formT = (string) => t("form." + string)

    const
        [deviceType, setDeviceType] = useState(""),
        [chargingSource, setChargingSource] = useState([
            {
                value: "Solar+Grid",
                label: "Solar+Grid",
            },
            {
                value: "Solar",
                label: "Solar",
            }
        ])
    const
        inputPowerCapacity = (e) => {
            const num = e.target.value
            const isNum = validateNumTwoDecimalPlaces(num)
            if (!isNum) return
            // setBasicPrice(num)
        }
    return <>
        <h5 className="mt-8 mb-5 ml-2">{title}</h5>

        {subDeviceData.map((item, i) => (
            <>
                <TextField
                    key={"sub-deviceType-" + i}
                    label={formT("deviceType")}
                    value={formT(`${item.deviceType}`)} />
                <TextField
                    key={"sub-deviceModel-" + i}
                    select
                    label={formT("deviceModel")}
                    defaultValue="">
                    {subDeviceModel.map((option, i) => (
                        <MenuItem
                            key={"option-d-m-" + i}
                            value={option.value}>
                            {option.label}
                        </MenuItem>
                    ))}
                </TextField>
                <h5 className="mb-5 ml-2">{formT("deviceInformation")}</h5>
                <TextField
                    id={"p-c-" + i}
                    type="number"
                    label={formT("powerCapacity")}
                // value={powerCapacity}
                />
                <Divider id={"line-" + i} variant="middle" sx={{ margin: "1rem 0 2.5rem" }} />
            </>
        ))}



    </>
}