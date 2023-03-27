import {
    Button, DialogActions, Divider, FormControl, InputAdornment, ListItem,
    MenuItem, TextField
} from "@mui/material"
import AddIcon from "@mui/icons-material/Add"
import { useTranslation } from "react-multi-lang"
import { useEffect, useMemo, useState } from "react"
import { ValidateNumPercent } from "../utils/utils"

import DialogForm from "../components/DialogForm"

export default function AddField({
    children = null,
    dialogTitle = "",
    locationInfo,
    fieldDevices,
    deviceInfo,
    extraDeviceInfo,
    open,
    setOpen,
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
        [deviceType, setDeviceType] = useState(""),
        [gridOutagePercent, setGridOutagePercent] = useState(""),
        [chargingSource, setChargingSource] = useState([
            {
                value: "solar+grid",
                label: "solar+grid",
            },
            {
                value: "solar",
                label: "solar",
            }
        ]),
        [fullWidth, setFullWidth] = useState(true),
        [maxWidth, setMaxWidth] = useState("lg"),
        [openAdd, setOpenAdd] = useState(false)
    const
        inputPercent = (e) => {
            const num = e.target.value
            const isNum = ValidateNumPercent(num)
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
                    id="address"
                    // value={address}
                    label={formT("address")}
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
                <Divider variant="middle" />
                <h5 className="mt-8 mb-5 ml-2">{fieldDevices}</h5>
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
                                console.log(deviceType)
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
                <Divider variant="middle" />
                {deviceType.value === "battery"
                    ? <>
                        <h5 className="mt-8 mb-5 ml-2">{extraDeviceInfo}</h5>
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
                    </>
                    : null}
            </FormControl>
            <Divider variant="middle" />
            <DialogActions sx={{ margin: "1rem 0.5rem 1rem 0" }}>
                <Button onClick={() => { setOpenAdd(false) }}
                    radius="pill"
                    variant="outlined"
                    color="gray">
                    {commonT("cancel")}
                </Button>
                <Button onClick={() => { setOpenAdd(false) }}
                    radius="pill"
                    variant="contained"
                    color="primary">
                    {commonT("add")}
                </Button>
            </DialogActions>
        </DialogForm>
    </>
}