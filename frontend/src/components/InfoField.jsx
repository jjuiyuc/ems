import {
    Button, Chip, DialogActions, Divider, FormControl, InputAdornment, ListItem,
    MenuItem, Switch, TextField
} from "@mui/material"
import { useTranslation } from "react-multi-lang"
import { useEffect, useMemo, useState } from "react"
import { ValidateNumPercent } from "../utils/utils"

import DialogForm from "../components/DialogForm"
import ExtraDeviceInfoForm from "../components/ExtraDeviceInfoForm"
import SubDeviceForm from "../components/SubDeviceForm"

import { ReactComponent as NoticeIcon } from "../assets/icons/notice.svg"

export default function InfoField({
    children = null,
    dialogTitle = "",
    openNotice,
    setOpenNotice,
    target,
    setTarget,
    onClick,
    locationInfo,
    fieldDevices,
    deviceInfo,
    extraDeviceInfo,
    subdevice,
    closeOutside = false
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
        ]
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        dialogT = (string) => t("dialog." + string),
        formT = (string) => t("form." + string),
        pageT = (string, params) => t("fieldManagement." + string, params)


    const
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
        [fullWidth, setFullWidth] = useState(true),
        [maxWidth, setMaxWidth] = useState("md")

    const
        inputPercent = (e) => {
            const num = e.target.value
            const isNum = ValidateNumPercent(num)
            if (!isNum) return
            setGridOutagePercent(num)
        }
    return <>
        <NoticeIcon
            className="mr-5"
            onClick={onClick
                // () => {

                //     console.log(target?.deviceType)}
            } />
        <DialogForm
            dialogTitle={commonT("group")}
            fullWidth={fullWidth}
            maxWidth={maxWidth}
            open={openNotice}
            setOpen={setOpenNotice}>
            <Divider variant="middle" />
            <div className="flex flex-col m-auto mt-4 min-w-49 w-fit">
                <div className="grid grid-cols-1fr-auto items-center mb-4">
                    <TextField
                        sx={{ marginBottom: 0 }}
                        id="gatewayID"
                        label={commonT("gatewayID")}
                        value={target?.gatewayID || ""}
                        focused
                        disabled={true}
                    />
                </div>
                <h5 className="mb-4 ml-2">{locationInfo}</h5>
                <TextField
                    id="location-name"
                    label={commonT("locationName")}
                    value={target?.locationName || ""}
                    focused
                    disabled={true}
                />
                <TextField
                    id="address"
                    label={formT("address")}
                    value={target?.address || ""}
                    focused
                    disabled={true}
                />
                <div className="flex-nowrap">
                    <TextField
                        id="lat"
                        type="number"
                        label={formT("lat")}
                        value={target?.lat || ""}
                        focused
                        disabled={true}
                    />
                    <TextField
                        sx={{ marginLeft: "1rem" }}
                        id="lng"
                        type="number"
                        label={formT("lng")}
                        value={target?.lng || ""}
                        disabled={true}
                    />
                </div>
                <TextField
                    id="powerCompany"
                    label={formT("powerCompany")}
                    value={target?.powerCompany || ""}
                    disabled={true}
                />
                <TextField
                    id="voltageType"
                    label={formT("voltageType")}
                    value={target?.voltageType || ""}
                    disabled={true}
                />
                <TextField
                    id="touType"
                    label={formT("touType")}
                    value={target?.touType || ""}
                    disabled={true}
                />
                <Divider variant="middle" />
                <h5 className="mb-4 mt-4 ml-2">{fieldDevices}</h5>
                <TextField
                    id="deviceType"
                    label={formT("deviceType")}
                    value={formT(target?.deviceType) || ""}
                    disabled={true}
                />
                <TextField
                    id="deviceModel"
                    label={formT("deviceModel")}
                    value={target?.deviceModel || ""}
                    disabled={true}
                />
                <Divider variant="middle" />
                <h5 className="mb-4 mt-4 ml-2">{deviceInfo}</h5>
                <TextField
                    id="modbusID"
                    type="number"
                    label={formT("modbusID")}
                    value={target?.modbusID || ""}
                    disabled={true}
                />
                <TextField
                    id="UUEID"
                    label="UUEID"
                    value={target?.UUEID || ""}
                    disabled={true}
                />
                <TextField
                    id="powerCapacity"
                    type="number"
                    label={formT("powerCapacity")}
                    value={target?.powerCapacity || ""}
                    disabled={true}
                />
                <Divider variant="middle" sx={{ margin: "0 0 2rem" }} />
                {deviceType.value === "battery"
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
                {target?.deviceType === "hybridInverter"
                    ? <>
                        <SubDeviceForm
                            title={subdevice}
                            mainDeviceType={target?.deviceType}
                        />
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
                {deviceType.value === "inverter"
                    ? <>
                        <SubDeviceForm
                            title={subdevice}
                            mainDeviceType={deviceType}
                        />
                    </>
                    : null}
                <div className="mb-8 flex items-baseline">
                    <p className="ml-1 mr-2">{formT("enableField")}</p>
                    <Switch />
                </div>
            </div>
            <DialogActions sx={{ margin: "1rem 0.5rem 1rem 0" }}>
                <Button onClick={() => { setOpenNotice(false) }}
                    radius="pill"
                    variant="contained"
                    color="primary">
                    {commonT("okay")}
                </Button>
            </DialogActions>
        </DialogForm>
    </>
}