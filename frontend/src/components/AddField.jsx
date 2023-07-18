import {
    Button, Checkbox, DialogActions, Divider, FormControlLabel, FormGroup,
    MenuItem, Switch, TextField
} from "@mui/material"
import AddIcon from "@mui/icons-material/Add"
import { useTranslation } from "react-multi-lang"
import { useEffect, useMemo, useState } from "react"
import { validateNumPercent } from "../utils/utils"

import { apiCall } from "../utils/api"

import DialogForm from "../components/DialogForm"
import ExtraDeviceInfoForm from "../components/ExtraDeviceInfoForm"
import SubDeviceForm from "../components/SubDeviceForm"

export default function AddField({
    extraDeviceInfo,
    subdevice
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
        ]
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        formT = (string) => t("form." + string),
        pageT = (string, params) => t("fieldManagement." + string, params)

    const
        [gatewayID, setGatewayID] = useState(""),
        [locationName, setLocationName] = useState(""),
        [address, setAddress] = useState(""),
        [lat, setLat] = useState(""),
        [lng, setLng] = useState(""),
        [modelList, setModelList] = useState([]),
        [deviceType, setDeviceType] = useState([]),
        [deviceModel, setDeviceModel] = useState([]),
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
        [deviceInfoCount, setDeviceInfoCount] = useState(1),
        [showAddIcon, setShowAddIcon] = useState(true),
        [isHybridInverterSelected, setIsHybridInverterSelected] = useState(false),
        [fullWidth, setFullWidth] = useState(true),
        [maxWidth, setMaxWidth] = useState("lg"),
        [openAdd, setOpenAdd] = useState(false),
        [fetched, setFetched] = useState(false)

    const getModelList = () => {
        apiCall({
            onError: error => setInfoError(error),
            onSuccess: rawData => {
                if (!rawData?.data) return

                const { data } = rawData

                setModelList(data.models)
            },
            url: `/api/device-management/devices/models`
        })
    }
    const deviceTypeOptions = useMemo(() => {
        const filteredData = modelList.filter(({ type }) => type !== "PV")
        const allTypes = [...new Set(filteredData.map(({ type }) => type))]
        // console.log(allTypes)

        const typeOptions = allTypes.map((type) => {
            if (type === "Hybrid-Inverter") {
                const otherSelected = deviceType.some((type) => type !== "Hybrid-Inverter")
                if (otherSelected) return { type, disabled: true }
            } else {
                const hybridInverterSelected = deviceType.includes("Hybrid-Inverter")
                if (hybridInverterSelected) return { type, disabled: true }
            }
            return { type, disabled: false }
        })
        return typeOptions
    }, [modelList, deviceType])

    const deviceModelOptions = useMemo(() => {
        return deviceType.flatMap((type) => {
            const filteredOptions = modelList.filter((model) => model.type === type)
            return filteredOptions
        })
    }, [modelList, deviceType])

    const
        changeDeviceType = (e) => {
            const { value } = e.target
            const alreadyChecked = deviceType.includes(value)

            const newDeviceType = alreadyChecked
                ? deviceType.filter((v) => v !== value)
                : [...deviceType, value]
            setDeviceType(newDeviceType)

            if (alreadyChecked) {
                const newDeviceModel = deviceModel.filter((type) => {
                    const device = modelList.find(({ type }) => type === type)
                    return newDeviceType.includes(device.type)
                })
                setDeviceModel(newDeviceModel)
            }
            if (value === "Hybrid-Inverter") {
                setIsHybridInverterSelected(!alreadyChecked)
            }
        },
        changeDeviceModel = (index) => (e) => {
            const { value } = e.target
            setDeviceModel((prevDeviceModel) => {
                const updatedDeviceModel = [...prevDeviceModel]
                updatedDeviceModel[index] = value
                return updatedDeviceModel
            })
        },
        inputPercent = (e) => {
            const num = e.target.value
            const isNum = validateNumPercent(num)
            if (!isNum) return
            setGridOutagePercent(num)
        },
        addDeviceInfoGroup = () => {
            if (deviceInfoCount < 3) {
                setDeviceInfoCount((prevCount) => prevCount + 1)
            }
            if (deviceInfoCount + 1 >= 3) {
                setShowAddIcon(false)
            }
        }
    useEffect(() => {
        // if (openAdd && fetched == false)
        getModelList()
    }, [fetched, openAdd])

    // console.log(modelList)
    // console.log(deviceModelOptions)
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
            dialogTitle={pageT("addField")}
            open={openAdd}
            setOpen={setOpenAdd}
            fullWidth={fullWidth}
            maxWidth={maxWidth}>
            <Divider variant="middle" />
            <div className="flex flex-col m-auto mt-4 min-w-49 w-fit">
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
                <h4 className="mb-5 ml-2">{pageT("locationInfo")}</h4>
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
                <Divider variant="middle" sx={{ margin: "0 0 2rem" }} />
                <h4 className="mb-5 ml-2">{pageT("fieldDevices")}</h4>
                <h5 className="mb-5 ml-2">{formT("deviceType")}</h5>
                <div className="border-gray-400 border rounded-xl
                    grid grid-cols-2 gap-2 items-center mb-8 p-4">
                    <FormGroup>
                        {deviceTypeOptions.map(({ i, type, disabled }) =>
                            <FormControlLabel
                                key={"option-d-t-" + type + i}
                                control={
                                    <Checkbox
                                        checked={deviceType.includes(type)}
                                        value={type}
                                        onChange={changeDeviceType}
                                        disabled={disabled}
                                    />
                                }
                                label={type}
                            />
                        )}
                    </FormGroup>
                </div>
                {deviceType.map((type, index) => {
                    return (
                        <>
                            <TextField
                                key={index}
                                id={`d-m-${index}`}
                                select
                                label={formT("deviceModel") + ` ${index + 1}`}
                                onChange={changeDeviceModel(index)}
                                value={deviceModel[index] || ""}
                            >
                                {deviceModelOptions
                                    .filter((model) => model.type === type)
                                    .map(({ id, name }) => (
                                        <MenuItem key={id} value={id}>
                                            {name}
                                        </MenuItem>
                                    ))}
                            </TextField>
                            {Array.from({ length: deviceInfoCount }).map((_, i) => (
                                <div className="flex flex-col">
                                    <div className="grid grid-cols-1fr-auto items-center mb-5 ml-2">
                                        <h5 className="">{formT("deviceInfo") + ` ${i + 1}`}</h5>
                                        {i === deviceInfoCount - 1 && isHybridInverterSelected && showAddIcon && (
                                            <AddIcon onClick={addDeviceInfoGroup} />
                                        )}
                                    </div>
                                    <TextField
                                        id={`modbusID-${i}`}
                                        type="number"
                                        label={formT("modbusID")}
                                    // value={modbusID}
                                    />
                                    <div className="grid grid-cols-1fr-auto items-center mb-8">
                                        <TextField
                                            sx={{ marginBottom: 0 }}
                                            id={`UUEID-${i}`}
                                            label="UUEID"
                                        // value={UUEID}
                                        />
                                        <Button
                                            // onClick={}
                                            sx={{ marginLeft: "0.3rem" }}
                                            radius="pill"
                                            variant="contained"
                                            color="primary"
                                        >
                                            {commonT("verify")}
                                        </Button>
                                    </div>
                                    <TextField
                                        id={`powerCapacity-${i}`}
                                        type="number"
                                        label={formT("powerCapacity")}
                                    // value={powerCapacity}
                                    />
                                    <Divider variant="middle" sx={{ margin: "0 0 2rem" }} />
                                </div>
                            ))}
                        </>
                    )
                })}
                {/* <Divider variant="middle" sx={{ margin: "0 0 2rem" }} /> */}
                {deviceType.includes("Hybrid-Inverter")
                    ? <>
                        <div className="pl-10 flex flex-col">
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
                        </div>
                    </>
                    : null}
                {deviceType.includes("Inverter")
                    ? <>
                        <div className="pl-10 flex flex-col">
                            <SubDeviceForm
                                title={subdevice}
                                mainDeviceType={deviceType}
                            />
                        </div>
                    </>
                    : null}
                {deviceType.includes("Battery")
                    ? <>
                        <div className="pl-10 flex flex-col">
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
                        </div>
                    </>
                    : null}
                <div className="mb-8 flex items-baseline">
                    <p className="ml-1 mr-2">{formT("enableField")}</p>
                    <Switch />
                </div>
            </div>
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