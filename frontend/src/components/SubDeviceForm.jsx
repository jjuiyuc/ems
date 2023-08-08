import { Divider, MenuItem, TextField } from "@mui/material"
import { useTranslation } from "react-multi-lang"
import { useEffect, useMemo, useState } from "react"

import { apiCall } from "../utils/api"
import { validateNumTwoDecimalPlaces } from "../utils/utils"

export default function SubDeviceForm(props) {
    const { title, mainDeviceType, subDeviceInfo, changeSubDeviceInfo,
        handleKeyDown } = props
    const
        t = useTranslation(),
        formT = (string) => t("form." + string)

    const
        [subDevicesList, setSubDevicesList] = useState([]),
        [infoError, setInfoError] = useState("")

    const
        changeSubDeviceModel = (e, index) => {
            const { value } = e.target
            changeSubDeviceInfo(index, "subDeviceModel", value)
        },
        changePowerCapacity = (e, index) => {
            const num = e.target.value

            const isNum = validateNumTwoDecimalPlaces(num)
            if (num === "" || isNum) {
                changeSubDeviceInfo(index, "subPowerCapacity", num)
            }
        }

    const filteredSubDevicesList = useMemo(() => {
        let filteredList = []
        if (mainDeviceType.includes("Hybrid-Inverter")) {
            filteredList = subDevicesList.filter((item) =>
                ["Meter", "PV", "Battery"].includes(item.type)
            )
        } else if (mainDeviceType.includes("Inverter")) {
            filteredList = subDevicesList.filter((item) => item.type === "PV")
        }
        return filteredList
    }, [mainDeviceType, subDevicesList])

    const getList = () => {
        apiCall({
            onError: error => setInfoError(error),
            onSuccess: rawData => {
                if (!rawData?.data) return

                const { data } = rawData

                setSubDevicesList(data.subDevices || [])
            },
            url: `/api/device-management/devices/sub-devices/models`
        })
    }
    useEffect(() => {
        getList()
    }, [])

    return <>
        <h5 className="mb-5 ml-2">{title}</h5>
        {mainDeviceType.includes("Hybrid-Inverter") &&
            filteredSubDevicesList?.map((item, i) => (
                <>
                    <TextField
                        key={"sub-d-t-" + i}
                        label={formT("deviceType")}
                        value={item.type} />
                    <TextField
                        key={"sub-d-m-" + i}
                        select
                        label={formT("deviceModel")}
                        value={subDeviceInfo.subDeviceModel[i]}
                        onChange={(e) => changeSubDeviceModel(e, i)}
                        defaultValue="">
                        {item.models.map(({ id, name }, i) => (
                            <MenuItem key={`o-sub-d-m-${id}-${i}`} value={id}>
                                {name}
                            </MenuItem>
                        ))}
                    </TextField>
                    <h5 className="mb-5 ml-2">{formT("deviceInfo")}</h5>
                    <TextField
                        key={"p-c-" + i}
                        type="number"
                        label={formT("powerCapacity")}
                        onChange={(e) => changePowerCapacity(e, i)}
                        onKeyDown={(e) => handleKeyDown(e, i)}
                        value={subDeviceInfo.subPowerCapacity[i]}
                    />
                    <Divider key={"line-" + i} variant="middle" sx={{ margin: "1rem 0 2.5rem" }} />
                </>
            ))}
        {mainDeviceType.includes("Inverter") &&
            filteredSubDevicesList?.map((item, i) => (
                <>
                    <TextField
                        key={"i-sub-d-t-" + i}
                        label={formT("deviceType")}
                        value={item.type} />
                    <TextField
                        key={"i-sub-d-m-" + i}
                        select
                        label={formT("deviceModel")}
                        value={subDeviceInfo.subDeviceModel[i]}
                        onChange={(e) => changeSubDeviceModel(e, i)}
                        defaultValue="">
                        {item.models.map(({ id, name }) => (
                            <MenuItem key={`o-sub-d-m-${id}`} value={id}>
                                {name}
                            </MenuItem>
                        ))}
                    </TextField>
                    <h5 className="mb-5 ml-2">{formT("deviceInfo")}</h5>
                    <TextField
                        key={"i-p-c-"}
                        type="number"
                        label={formT("powerCapacity")}
                        onChange={(e) => changePowerCapacity(e, i)}
                        value={subDeviceInfo.subPowerCapacity[i]}
                    />
                </>
            ))}
    </>
}

