import {
    Button, Chip, DialogActions, Divider, FormControl, InputAdornment,
    Switch, TextField
} from "@mui/material"
import { useTranslation } from "react-multi-lang"
import { useEffect, useMemo, useState } from "react"
import { ValidateNumPercent } from "../utils/utils"

import DialogForm from "../components/DialogForm"
import ExtraDeviceInfoForm from "../components/ExtraDeviceInfoForm"

import { ReactComponent as NoticeIcon } from "../assets/icons/notice.svg"

export default function InfoField({
    dialogTitle = "",
    row,
    locationInfo,
    fieldDevices,
    deviceInfo,
    extraDeviceInfo,
    subdevice
}) {

    const
        deviceTypes = [
            {
                value: "hybridInverter",
                label: "hybridInverter",
            },
            {
                value: "inverter",
                label: "inverter",
            },
            {
                value: "meter",
                label: "meter",
            },
            {
                value: "pv",
                label: "pv",
            },
            {
                value: "battery",
                label: "battery",
            },
            {
                value: "pcs",
                label: "pcs",
            }
        ],

        deviceModel = [
            {
                value: "LXP-12K US-Luxpower Hybrid-Inverter",
                label: "LXP-12K US-Luxpower Hybrid-Inverter"
            }
        ],
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
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        dialogT = (string) => t("dialog." + string),
        formT = (string) => t("form." + string)

    const
        [openNotice, setOpenNotice] = useState(false),
        [gatewayID, setGatewayID] = useState(""),
        [locationName, setLocationName] = useState(""),
        [address, setAddress] = useState(""),
        [lat, setLat] = useState(""),
        [lng, setLng] = useState(""),
        [powerCompany, setPowerCompany] = useState(""),
        [deviceType, setDeviceType] = useState(""),
        [gridOutagePercent, setGridOutagePercent] = useState(""),
        [chargingSource, setChargingSource] = useState([
            {
                value: "Solar+Grid",
                label: "Solar+Grid",
            },
            {
                value: "Solar",
                label: "Solar",
            }
        ]),
        [energyCapacity, setEnergyCapacity] = useState(null),
        [voltage, setVoltage] = useState(null),
        [subPowerCapacity, setSubPowerCapacity] = useState(""),
        [fullWidth, setFullWidth] = useState(true),
        [maxWidth, setMaxWidth] = useState("md")

    // const { AreaOwner_TW, AreaMaintainer, Serenegray } = groupState

    const
        inputPercent = (e) => {
            const num = e.target.value
            const isNum = ValidateNumPercent(num)
            if (!isNum) return
            setGridOutagePercent(num)
        },
        handleClick = () => {
            setOpenNotice(true)
        }
    return <>
        <NoticeIcon
            className="mr-5"
            onClick={handleClick} />
        <DialogForm
            dialogTitle={dialogT("fieldInfo")}
            fullWidth={fullWidth}
            maxWidth={maxWidth}
            open={openNotice}
            setOpen={setOpenNotice}>
            <Divider variant="middle" />
            <div className="flex flex-col m-auto mt-4 min-w-49 w-fit">
                <TextField
                    sx={{ marginBottom: 2 }}
                    key="gatewayID"
                    label={commonT("gatewayID")}
                    value={row?.gatewayID || ""}
                    focused
                    disabled={true}
                />
                <h5 className="mb-4 ml-2">{locationInfo}</h5>
                <TextField
                    key="location-name"
                    label={commonT("locationName")}
                    value={row?.locationName || ""}
                    focused
                    disabled={true}
                />
                <TextField
                    key="address"
                    label={formT("address")}
                    value={row?.address || ""}
                    focused
                    disabled={true}
                />
                <div className="flex-nowrap">
                    <TextField
                        key="lat"
                        type="number"
                        label={formT("lat")}
                        value={row?.lat || ""}
                        focused
                        disabled={true}
                    />
                    <TextField
                        sx={{ marginLeft: "1rem" }}
                        key="lng"
                        type="number"
                        label={formT("lng")}
                        value={row?.lng || ""}
                        disabled={true}
                    />
                </div>
                <TextField
                    key="powerCompany"
                    label={formT("powerCompany")}
                    value={row?.powerCompany || ""}
                    disabled={true}
                />
                <TextField
                    key="voltageType"
                    label={formT("voltageType")}
                    value={formT(row?.voltageType) || ""}
                    disabled={true}
                />
                <TextField
                    key="touType"
                    label={formT("touType")}
                    value={formT(row?.touType) || ""}
                    disabled={true}
                />
                <Divider variant="middle" />
                <h5 className="mb-4 mt-4 ml-2">{fieldDevices}</h5>
                <TextField
                    key="deviceType"
                    label={formT("deviceType")}
                    value={formT(row?.deviceType) || ""}
                    disabled={true}
                />
                <TextField
                    key="deviceModel"
                    label={formT("deviceModel")}
                    value={row?.deviceModel || ""}
                    disabled={true}
                />
                <Divider variant="middle" />
                <h5 className="mb-4 mt-4 ml-2">{deviceInfo}</h5>
                <TextField
                    key="modbusID"
                    type="number"
                    label={formT("modbusID")}
                    value={row?.modbusID || ""}
                    disabled={true}
                />
                <TextField
                    key="UUEID"
                    label="UUEID"
                    value={row?.UUEID || ""}
                    disabled={true}
                />
                <TextField
                    key="powerCapacity"
                    type="number"
                    label={formT("powerCapacity")}
                    value={row?.powerCapacity || ""}
                    disabled={true}
                />
                <Divider variant="middle" sx={{ margin: "0 0 2rem" }} />
                {row?.deviceType === "battery"
                    ? <ExtraDeviceInfoForm
                        subTitle={extraDeviceInfo}
                        gridOutagePercent={gridOutagePercent}
                        setGridOutagePercent={setGridOutagePercent}
                        chargingSource={chargingSource}
                        setChargingSource={setChargingSource}
                        energyCapacity={energyCapacity}
                        setEnergyCapacity={setEnergyCapacity}
                        voltage={voltage}
                        setVoltage={setVoltage}
                    />
                    : null}
                {row?.deviceType === "hybridInverter"
                    ? <>
                        {row?.subDevice.map((item, i) => (
                            <>
                                <TextField
                                    key={"sub-d-t-" + i}
                                    label={formT("deviceType")}
                                    value={formT(`${item.deviceType}`)}
                                    disabled={true}
                                />
                                <TextField
                                    key={"sub-d-m-" + i}
                                    label={formT("deviceModel")}
                                    value={item.deviceModel || ""}
                                    disabled={true}
                                />

                                <h5 className="mb-5 ml-2">{formT("deviceInformation")}</h5>
                                <TextField
                                    key={"p-c-" + i}
                                    type="number"
                                    label={formT("powerCapacity")}
                                    value={item.powerCapacity || ""}
                                />
                                <Divider key={"line-" + i} variant="middle" sx={{ margin: "1rem 0 2.5rem" }} />
                            </>
                        ))}
                        <ExtraDeviceInfoForm
                            subTitle={extraDeviceInfo}
                            gridOutagePercent={gridOutagePercent}
                            setGridOutagePercent={setGridOutagePercent}
                            chargingSource={chargingSource}
                            setChargingSource={setChargingSource}
                            energyCapacity={energyCapacity}
                            setEnergyCapacity={setEnergyCapacity}
                            voltage={voltage}
                            setVoltage={setVoltage}
                        />
                    </>
                    : null}
                {row?.deviceType === "inverter"
                    ? <>
                        <TextField
                            key={"i-sub-d-t-"}
                            label={formT("deviceType")}
                            value={formT(`${row?.subDevice[1].deviceType}`)}
                            disabled={true}
                        />
                        <TextField
                            key={"i-sub-d-m-"}
                            label={formT("deviceModel")}
                            value={row?.subDevice[1].deviceModel || ""}
                            disabled={true}
                        />
                        <h5 className="mb-5 ml-2">{formT("deviceInformation")}</h5>
                        <TextField
                            key={"i-p-c-"}
                            type="number"
                            label={formT("powerCapacity")}
                            value={row?.subDevice[1].powerCapacity || ""}
                            disabled={true}
                        />
                    </>
                    : null}
                <div className="mb-5 flex items-baseline">
                    <p className="ml-1 mr-2">{formT("enableField")}</p>
                    <Switch disabled={true} />
                </div>
                <Divider variant="middle" sx={{ margin: "0 0 1rem" }} />
                <h5 className="mb-5">{commonT("group")}</h5>
                <div className="border-gray-400 border rounded-xl
                    grid grid-cols-3 gap-2 items-center mb-4 p-4">
                    <Chip label="AreaOwner_TW" variant="outlined" color="primary" />
                    <Chip label="Area Maintainer" variant="outlined" color="primary" />
                    <Chip label="Serenegray" variant="outlined" color="primary" />
                </div>
            </div>
            <DialogActions sx={{ margin: "1rem 1.5rem 1rem 0" }}>
                <Button onClick={() => { setOpenNotice(false) }}
                    size="large"
                    radius="pill"
                    variant="contained"
                    color="primary">
                    {commonT("okay")}
                </Button>
            </DialogActions>
        </DialogForm>
    </>
}