import {
    Button, DialogActions, Divider, FormControl, InputAdornment, ListItem,
    MenuItem, Switch, TextField
} from "@mui/material"
import AddIcon from "@mui/icons-material/Add"
import { useTranslation } from "react-multi-lang"
import { useEffect, useMemo, useState } from "react"
import { validateNumPercent } from "../utils/utils"

import DialogForm from "../components/DialogForm"
import ExtraDeviceInfoForm from "../components/ExtraDeviceInfoForm"
import SubDeviceForm from "../components/SubDeviceForm"

export default function AddField({
    dialogTitle = "",
    locationInfo,
    fieldDevices,
    deviceInfo,
    extraDeviceInfo,
    subdevice,
    closeOutside = false
}) {

    const
        powerCompany = [
            {
                value: "TPC",
                label: "TPC",
            }
        ],
        voltageType = [
            {
                value: "lowVoltage",
                label: "lowVoltage",
            },
            {
                value: "highVoltage",
                label: "highVoltage",
            }
        ],
        touType = [
            {
                value: "twoSection",
                label: "twoSection",
            }
        ],
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
        formT = (string) => t("form." + string)

    const
        [gatewayID, setGatewayID] = useState(""),
        [locationName, setLocationName] = useState(""),
        [address, setAddress] = useState(""),
        [lat, setLat] = useState(""),
        [lng, setLng] = useState(""),
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
        [maxWidth, setMaxWidth] = useState("lg"),
        [openAdd, setOpenAdd] = useState(false)
    const
        inputPercent = (e) => {
            const num = e.target.value
            const isNum = validateNumPercent(num)
            if (!isNum) return
            setGridOutagePercent(num)
        }
    return <>
        <Button
            onClick={() => { setOpenAdd(true) }}
            size="x-large"
            variant="outlined"
            radius="pill"
            fontSize="large"
            color="brand"
            startIcon={<AddIcon />}>
            {commonT("add")}
        </Button>
        <DialogForm
            dialogTitle={dialogT("addField")}
            open={openAdd}
            setOpen={setOpenAdd}
            fullWidth={fullWidth}
            maxWidth={maxWidth}>
            <Divider variant="middle" />
            <FormControl sx={{
                display: "flex",
                flexDirection: "column",
                margin: "auto",
                width: "fit-content",
                mt: 2,
                minWidth: 120
            }}
                fullWidth={true}>
                <div className="grid grid-cols-1fr-auto items-center mb-8">
                    <TextField
                        sx={{ marginBottom: 0 }}
                        id="gatewayID"
                        label={commonT("gatewayID")}
                    // value={gatewayID}
                    />
                    <Button
                        // onClick={}
                        sx={{ marginLeft: "0.3rem" }}
                        radius="pill"
                        variant="contained"
                        color="primary">
                        {commonT("verify")}
                    </Button>
                </div>

                <h5 className="mb-5 ml-2">{locationInfo}</h5>
                <TextField
                    id="location-name"
                    label={commonT("locationName")}
                // value={locationName}
                />
                <TextField
                    id="address"
                    label={formT("address")}
                // value={address}
                />
                <div className="flex-nowrap">
                    <TextField
                        id="lat"
                        type="number"
                        label={formT("lat")}
                    // value={lat}
                    />
                    <TextField
                        id="lng"
                        type="number"
                        label={formT("lng")}
                        // value={lng}
                        sx={{ marginLeft: "1rem" }}
                    />
                </div>
                <TextField
                    id="powerCompany"
                    select
                    label={formT("powerCompany")}
                    defaultValue=""
                >
                    {powerCompany.map((option) => (
                        <MenuItem key={option.value} value={option.value}>
                            {option.label}
                        </MenuItem>
                    ))}
                </TextField>
                <TextField
                    id="voltageType"
                    select
                    label={formT("voltageType")}
                    defaultValue=""
                >
                    {voltageType.map((option) => (
                        <MenuItem key={option.value} value={formT(`${option.value}`)}>
                            {formT(`${option.label}`)}
                        </MenuItem>
                    ))}
                </TextField>
                <TextField
                    id="touType"
                    select
                    label={formT("touType")}
                    defaultValue=""
                >
                    {touType.map((option) => (
                        <MenuItem key={option.value} value={formT(`${option.value}`)}>
                            {formT(`${option.label}`)}
                        </MenuItem>
                    ))}
                </TextField>
                <Divider variant="middle" sx={{ margin: "1rem 0 2.5rem" }} />
                <h5 className="mb-5 ml-2">{fieldDevices}</h5>
                <TextField
                    id="deviceType"
                    select
                    label={formT("deviceType")}
                    defaultValue=""
                >
                    {deviceTypes.map((option) => (
                        <MenuItem
                            key={option.value}
                            value={formT(`${option.value}`)}
                            onClick={() => {
                                setDeviceType(option)
                            }}>
                            {formT(`${option.label}`)}
                        </MenuItem>

                    ))}
                </TextField>
                <TextField
                    id="deviceModel"
                    select
                    label={formT("deviceModel")}
                    defaultValue=""
                >
                    {deviceModel.map((option) => (
                        <MenuItem key={option.value} value={option.value}>
                            {option.label}
                        </MenuItem>
                    ))}
                </TextField>
                <h5 className="mb-5 ml-2">{deviceInfo}</h5>
                <TextField
                    id="modbusID"
                    type="number"
                    label={formT("modbusID")}
                // value={modbusID}
                />
                <div className="grid grid-cols-1fr-auto items-center mb-8">
                    <TextField
                        sx={{ marginBottom: 0 }}
                        id="UUEID"
                        label="UUEID"
                    // value={UUEID}
                    />
                    <Button
                        // onClick={}
                        sx={{ marginLeft: "0.3rem" }}
                        radius="pill"
                        variant="contained"
                        color="primary">
                        {commonT("verify")}
                    </Button>
                </div>
                <TextField
                    id="powerCapacity"
                    type="number"
                    label={formT("powerCapacity")}
                // value={powerCapacity}
                />
                <Divider variant="middle" sx={{ margin: "1rem 0 2rem" }} />
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
                {deviceType.value === "hybridInverter"
                    ? <>
                        <SubDeviceForm
                            title={subdevice}
                            mainDeviceType={deviceType}
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
            </FormControl>
            <Divider variant="middle" />
            <DialogActions sx={{ margin: "1rem 0.8rem 1rem 0" }}>
                <Button onClick={() => { setOpenAdd(false) }}
                    sx={{ marginRight: "0.4rem" }}
                    size="large"
                    radius="pill"
                    variant="outlined"
                    color="gray">
                    {commonT("cancel")}
                </Button>
                <Button onClick={() => { setOpenAdd(false) }}
                    sx={{ marginRight: "0.4rem" }}
                    size="large"
                    radius="pill"
                    variant="contained"
                    color="primary">
                    {commonT("add")}
                </Button>
            </DialogActions>
        </DialogForm>
    </>
}