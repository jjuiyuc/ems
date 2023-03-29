import {
    Button, Divider, InputAdornment, ListItem, MenuItem, Select, TextField
} from "@mui/material"
import { useTranslation } from "react-multi-lang"
import { useEffect, useMemo, useState } from "react"
import { validateNumTwoDecimalPlaces } from "../utils/utils"


export default function SubDeviceForm(props) {
    const { title, mainDeviceType } = props
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
                        value: "D1K330H4B URE PV",
                        label: "D1K330H4B URE PV"
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
                ]
            }
        ]
    const DeviceModelOption = subDeviceData
        .map(option => Object.values(option)[1]).flat()

    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        dialogT = (string) => t("dialog." + string),
        formT = (string) => t("form." + string)

    const
        [deviceModel, setDeviceModel] = useState(""),
        [subPowerCapacity, setSubPowerCapacity] = useState("")

    const
        inputPowerCapacity = (e) => {
            const num = e.target.value
            const isNum = validateNumTwoDecimalPlaces(num)
            if (!isNum) return
            setSubPowerCapacity(num)
            console.log(subDeviceData[1])
        }
    return <>
        <h5 className="mb-5 ml-2">{title}</h5>
        {mainDeviceType.value === "hybridInverter"
            ? <>
                {subDeviceData.map((item, i) => (
                    <>
                        <TextField
                            key={"sub-d-t-" + i}
                            label={formT("deviceType")}
                            value={formT(`${item.deviceType}`)} />
                        <TextField
                            key={"sub-d-m-" + i}
                            select
                            label={formT("deviceModel")}
                            defaultValue="">
                            {DeviceModelOption.map((option, i) => (
                                <MenuItem
                                    key={"option-d-m-" + i}
                                    value={option.value}>
                                    {option.label}
                                </MenuItem>
                            ))}
                        </TextField>
                        <h5 className="mb-5 ml-2">{formT("deviceInformation")}</h5>
                        <TextField
                            key={"p-c-" + i}
                            type="number"
                            label={formT("powerCapacity")}
                            onChange={inputPowerCapacity}
                        // value={subPowerCapacity}
                        />
                        <Divider key={"line-" + i} variant="middle" sx={{ margin: "1rem 0 2.5rem" }} />
                    </>
                ))}
            </>
            : null}
        {/* {mainDeviceType.value === "inverter"
            ? <>
                <TextField
                    key={"i-sub-d-t-" + i}
                    label={formT("deviceType")}
                    value={formT(`${subDeviceData[1].deviceType}`)} />
                <TextField
                    key={"i-sub-d-m-" + i}
                    select
                    label={formT("deviceModel")}
                    defaultValue="">
                    {DeviceModelOption.map((option, i) => (
                        <MenuItem
                            key={"option-d-m-" + i}
                            value={option.value}>
                            {option.label}
                        </MenuItem>
                    ))}
                </TextField>
                <h5 className="mb-5 ml-2">{formT("deviceInformation")}</h5>
                <TextField
                    key={"i-p-c-" + i}
                    type="number"
                    label={formT("powerCapacity")}
                    onChange={inputPowerCapacity}
                // value={subPowerCapacity}
                />
            </>
            : null} */}

    </>
}

