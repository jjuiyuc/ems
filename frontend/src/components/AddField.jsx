import { connect } from "react-redux"
import {
    Button, Checkbox, DialogActions, Divider, FormControlLabel, FormGroup,
    MenuItem, Switch, TextField
} from "@mui/material"
import AddIcon from "@mui/icons-material/Add"
import { useTranslation } from "react-multi-lang"
import { useEffect, useMemo, useState } from "react"

import { apiCall } from "../utils/api"
import { validateLat, validateLng, validateNumTwoDecimalPlaces } from "../utils/utils"

import DialogForm from "../components/DialogForm"
import ExtraDeviceInfoForm from "../components/ExtraDeviceInfoForm"
import SubDeviceForm from "../components/SubDeviceForm"

const mapDispatch = (dispatch) => ({
    updateSnackbarMsg: (value) =>
        dispatch({ type: "snackbarMsg/updateSnackbarMsg", payload: value }),
})
const TYPE_HYBRID_INVERTER = "Hybrid-Inverter"
const TYPE_INVERTER = "Inverter"
const TYPE_BATTERY = "Battery"

export default connect(null, mapDispatch)(function AddField(props) {
    const { getList } = props
    const powerCompanyOptions = [
        {
            id: 1,
            name: "TPC",
        },
    ],
        voltageTypeOptions = [
            {
                id: 1,
                name: "Low voltage",
            },
            {
                id: 2,
                name: "High voltage",
            },
        ],
        touTypeOptions = [
            {
                id: 1,
                name: "Two-section",
            },
        ]
    const t = useTranslation(),
        commonT = (string) => t("common." + string),
        errorT = (string) => t("error." + string),
        formT = (string) => t("form." + string),
        pageT = (string, params) => t("fieldManagement." + string, params)

    const
        [modelList, setModelList] = useState([]),
        [gatewayID, setGatewayID] = useState(""),
        [gatewayIDError, setGatewayIDError] = useState(false),
        [locationName, setLocationName] = useState(""),
        [address, setAddress] = useState(""),
        [lat, setLat] = useState(""),
        [lng, setLng] = useState(""),
        [powerCompany, setPowerCompany] = useState(""),
        [voltageType, setVoltageType] = useState(""),
        [touType, setTouType] = useState(""),
        [deviceType, setDeviceType] = useState([]),
        [deviceModel, setDeviceModel] = useState([]),
        [deviceInfo, setDeviceInfo] = useState({
            modbusID: [null, null, null],
            uueID: [null, null, null],
            powerCapacity: [null, null, null],
        }),
        [uueIDError, setUueIDError] = useState(false),
        [subDeviceInfo, setSubDeviceInfo] = useState({
            subDeviceModel: ["", "", ""],
            subPowerCapacity: [null, null, null],
        }),
        [extraDeviceInfo, setExtraDeviceInfo] = useState({
            gridOutagePercent: "",
            chargingSource: "",
            energyCapacity: null,
            voltage: null,
        }),
        [deviceInfoCount, setDeviceInfoCount] = useState(1),
        [showAddIcon, setShowAddIcon] = useState(true),
        [hybridInverterSelected, setHybridInverterSelected] = useState(false),
        [enable, setEnable] = useState(false),
        [fullWidth, setFullWidth] = useState(true),
        [maxWidth, setMaxWidth] = useState("lg"),
        [openAdd, setOpenAdd] = useState(false),
        [loading, setLoading] = useState(false),
        [infoError, setInfoError] = useState("")

    const getModelList = () => {
        apiCall({
            onError: (error) => setInfoError(error),
            onSuccess: (rawData) => {
                if (!rawData?.data) return

                const { data } = rawData

                setModelList(data.models)
            },
            url: `/api/device-management/devices/models`,
        })
    }
    const deviceTypeOptions = useMemo(() => {
        const filteredData = modelList.filter(({ type }) => type !== "PV")
        const allTypes = [...new Set(filteredData.map(({ type }) => type))]

        const typeOptions = allTypes.map((type) => {
            const id = filteredData.find(({ type: t }) => t === type)?.id

            if (type === TYPE_HYBRID_INVERTER) {
                const otherSelected = deviceType.some((t) => t !== TYPE_HYBRID_INVERTER)
                if (otherSelected) return { id, type, disabled: true }
            } else {
                const hybridInverterSelected = deviceType.includes(TYPE_HYBRID_INVERTER)
                if (hybridInverterSelected) return { id, type, disabled: true }
            }
            return { id, type, disabled: false }
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
        changeGatewayID = (e) => {
            setGatewayIDError(false)
            setGatewayID(e.target.value)
        },
        changeLocationName = (e) => {
            setLocationName(e.target.value)
        },
        changeAddress = (e) => {
            setAddress(e.target.value)
        },
        changeLat = (e) => {
            const num = e.target.value
            const isNum = validateLat(num)
            if (num === "" || isNum) {
                setLat(num)
            }
        },
        changeLng = (e) => {
            const num = e.target.value
            const isNum = validateLng(num)
            if (num === "" || isNum) {
                setLng(num)
            }
        },
        changePowerCompany = (e) => {
            setPowerCompany(e.target.value)
        },
        changeVoltageType = (e) => {
            setVoltageType(e.target.value)
        },
        changeTouType = (e) => {
            setTouType(e.target.value)
        },
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
            if (value === TYPE_HYBRID_INVERTER) {
                setHybridInverterSelected(!alreadyChecked)
            }
            if (alreadyChecked && !deviceType.includes(TYPE_HYBRID_INVERTER)) {
                setHybridInverterSelected(false)
                setDeviceInfo((prevInfo) => {
                    const newDeviceInfo = { ...prevInfo }
                    newDeviceInfo.uueID[deviceType.indexOf(value) + 1] = prevInfo.uueID[0]
                    return newDeviceInfo
                })
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
        changeDeviceInfo = (index, devices, value) => {
            setDeviceInfo((prevInfo) => {
                const newDeviceInfo = { ...prevInfo }
                newDeviceInfo[devices][index] = value
                return newDeviceInfo
            })
        },
        changeModbusID = (e, index) => {
            const num = e.target.value.replace(/\D/g, "")
            changeDeviceInfo(index, "modbusID", num)
        },
        changeUueID = (e, index) => {
            const uueIDTarget = e.target.value,
                uueIDLengthError = uueIDTarget.length < 32
            setUueIDError(uueIDLengthError)
            changeDeviceInfo(index, "uueID", uueIDTarget)

        },
        handleKeyDown = (e) => {
            if (
                !(
                    e.key === "Backspace" ||
                    e.key === "Delete" ||
                    e.key === "Tab" ||
                    e.key === "." ||
                    (e.key >= "0" && e.key <= "9")
                )
            ) {
                e.preventDefault()
            }
        },
        changePowerCapacity = (e, index) => {
            const num = e.target.value

            const isNum = validateNumTwoDecimalPlaces(num)
            if (num === "" || isNum) {
                changeDeviceInfo(index, "powerCapacity", num)
            }
        },
        addDeviceInfoGroup = () => {
            if (deviceInfoCount < 4) {
                setDeviceInfoCount((prevCount) => prevCount + 1)
            }
            if (deviceInfoCount >= 3) {
                setShowAddIcon(false)
            }
        },
        changeSubDeviceInfo = (index, field, value) => {
            setSubDeviceInfo((prevInfo) => {
                const newSubDeviceInfo = { ...prevInfo }
                newSubDeviceInfo[field][index] = value
                return newSubDeviceInfo
            })
        },
        handleSwitch = () => {
            setEnable((preState) => !preState)
        }
    const
        validateGatewayID = async () => {
            await apiCall({
                method: "get",
                onComplete: () => setLoading(false),
                onStart: () => setLoading(true),
                onError: (err) => {
                    switch (err) {
                        case 60024:
                            setGatewayIDError(true)
                            props.updateSnackbarMsg({
                                type: "error",
                                msg: errorT("gatewayIDInvalid"),
                            })
                            break
                        case 60025:
                            setGatewayIDError(true)
                            props.updateSnackbarMsg({
                                type: "error",
                                msg: errorT("gatewayIDisUsed"),
                            })
                            break
                        default:
                            setGatewayIDError(true)
                            props.updateSnackbarMsg({
                                type: "error",
                                msg: errorT("error"),
                            })
                    }
                },
                onSuccess: () => {
                    props.updateSnackbarMsg({
                        type: "success",
                        msg: t("dialog.verifySuccessfully"),
                    })
                },
                url: `/api/device-management/gateways/${gatewayID}/validity`,
            })
        },
        validateUueID = async (uueID) => {
            await apiCall({
                method: "get",
                onComplete: () => setLoading(false),
                onStart: () => setLoading(true),
                onError: (err) => {
                    switch (err) {
                        case 60026:
                            setUueIDError(true)
                            props.updateSnackbarMsg({
                                type: "error",
                                msg: errorT("uueIDInvalid"),
                            })
                            break
                        case 60027:
                            setUueIDError(true)
                            props.updateSnackbarMsg({
                                type: "error",
                                msg: errorT("uueIDIsUsed"),
                            })
                            break
                        default:
                            setUueIDError(true)
                            props.updateSnackbarMsg({
                                type: "error",
                                msg: errorT("error"),
                            })
                    }
                },
                onSuccess: () => {
                    props.updateSnackbarMsg({
                        type: "success",
                        msg: t("dialog.verifySuccessfully"),
                    })
                },
                url: `/api/device-management/devices/${uueID}/validity`,
            })
        }

    const submit = async () => {
        const devices = deviceInfo.uueID.map((uueID, index) => {
            const isBattery = deviceType[index] === "Battery"
            const hasDeviceInfo = deviceModel[index] || deviceInfo.modbusID[index] || uueID || deviceInfo.powerCapacity[index]

            const device = hasDeviceInfo
                ? {
                    modelID: parseInt(deviceModel[index])
                        ? parseInt(deviceModel[index])
                        : parseInt(deviceModel[0]),
                    modbusID: parseInt(deviceInfo.modbusID[index]),
                    uueID: uueID,
                    powerCapacity: parseFloat(deviceInfo.powerCapacity[index]),
                }
                : null

            if (index === 0 && subDeviceInfo.subDeviceModel.some((subDeviceModel) => subDeviceModel)) {
                const subDevices = subDeviceInfo.subDeviceModel.map((subDeviceModel, subIndex) => {
                    const isSubBattery = subDeviceModel && deviceType[subIndex] === "Battery"
                    const hasSubDeviceInfo = subDeviceModel || subDeviceInfo.subPowerCapacity[subIndex]

                    const subDevice = hasSubDeviceInfo
                        ? {
                            modelID: subDeviceModel ? parseInt(subDeviceModel) : null,
                            powerCapacity: parseFloat(subDeviceInfo.subPowerCapacity[subIndex])
                                ? parseFloat(subDeviceInfo.subPowerCapacity[subIndex])
                                : null,
                        }
                        : null

                    if (isSubBattery && subDevice) {
                        subDevice.extraInfo = {
                            reservedForGridOutagePercent: parseInt(extraDeviceInfo.gridOutagePercent),
                            chargingSources: extraDeviceInfo.chargingSource,
                            energyCapacity: parseFloat(extraDeviceInfo.energyCapacity),
                            voltage: parseFloat(extraDeviceInfo.voltage),
                        }
                    }
                    return subDevice
                })
                device.subDevices = subDevices.filter((subDevice) => subDevice !== null)
            }

            if (isBattery && device) {
                device.extraInfo = {
                    reservedForGridOutagePercent: parseInt(extraDeviceInfo.gridOutagePercent),
                    chargingSources: extraDeviceInfo.chargingSource,
                    energyCapacity: parseFloat(extraDeviceInfo.energyCapacity),
                    voltage: parseFloat(extraDeviceInfo.voltage),
                }
            }
            return device
        })

        const data = {
            gatewayID: gatewayID,
            locationName: locationName,
            address: address,
            lat: parseFloat(lat),
            lng: parseFloat(lng),
            powerCompany: powerCompany,
            voltageType: voltageType,
            touType: touType,
            devices: devices.filter((device) => device !== null),
            enable: enable,
        }
        await apiCall({
            method: "post",
            data,
            onSuccess: () => {
                setOpenAdd(false)
                getList()
                props.updateSnackbarMsg({
                    type: "success",
                    msg: t("dialog.addedSuccessfully"),
                })
                setGatewayID("")
                setLocationName("")
                setAddress("")
                setLat("")
                setLng("")
                setPowerCompany("")
                setVoltageType("")
                setTouType("")
                setDeviceType([])
                setDeviceModel([])
                setDeviceInfo({
                    modbusID: [null, null, null],
                    uueID: [null, null, null],
                    powerCapacity: [null, null, null],
                })
                setSubDeviceInfo({
                    subDeviceModel: ["", "", ""],
                    subPowerCapacity: [null, null, null],
                })
                setExtraDeviceInfo({
                    gridOutagePercent: "",
                    chargingSource: "",
                    energyCapacity: null,
                    voltage: null,
                })
                setDeviceInfoCount(1)
                setShowAddIcon(true)
                setHybridInverterSelected(false)
                setEnable(false)
            },
            onError: (err) => {
                switch (err) {
                    case 400:
                        props.updateSnackbarMsg({
                            type: "error",
                            msg: errorT("invalidParameter"),
                        })
                        break
                    case 60028:
                        props.updateSnackbarMsg({
                            type: "error",
                            msg: errorT("failureToCreate"),
                        })
                        break
                    case 60029:
                        props.updateSnackbarMsg({
                            type: "error",
                            msg: errorT("failureToCreate"),
                        })
                        break
                    default:
                        props.updateSnackbarMsg({
                            type: "error",
                            msg: errorT("failureToCreate"),
                        })
                }
            },
            url: "/api/device-management/gateways",
        })
    },
        cancelClick = () => {
            setOpenAdd(false)
            setGatewayID("")
            setLocationName("")
            setAddress("")
            setLat("")
            setLng("")
            setPowerCompany("")
            setVoltageType("")
            setTouType("")
            setDeviceType([])
            setDeviceModel([])
            setDeviceInfo({
                modbusID: [null, null, null],
                uueID: [null, null, null],
                powerCapacity: [null, null, null],
            })
            setSubDeviceInfo({
                subDeviceModel: ["", "", ""],
                subPowerCapacity: [null, null, null],
            })
            setExtraDeviceInfo({
                gridOutagePercent: "",
                chargingSource: "",
                energyCapacity: null,
                voltage: null,
            })
            setDeviceInfoCount(1)
            setShowAddIcon(true)
            setHybridInverterSelected(false)
            setEnable(false)
        }
    useEffect(() => {
        getModelList()
    }, [deviceType, openAdd])

    return (
        <>
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
                            error={gatewayIDError}
                            onChange={changeGatewayID}
                            value={gatewayID}
                        />
                        <Button
                            onClick={validateGatewayID}
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
                        onChange={changeLocationName}
                        value={locationName}
                    />
                    <TextField
                        id="address"
                        label={formT("address")}
                        onChange={changeAddress}
                        value={address}
                    />
                    <div className="flex-nowrap">
                        <TextField
                            id="lat"
                            type="number"
                            label={formT("lat")}
                            onChange={changeLat}
                            value={lat}
                        />
                        <TextField
                            id="lng"
                            type="number"
                            label={formT("lng")}
                            onChange={changeLng}
                            value={lng}
                            sx={{ marginLeft: "1rem" }}
                        />
                    </div>
                    <TextField
                        id="powerCompany"
                        select
                        label={formT("powerCompany")}
                        onChange={changePowerCompany}
                        defaultValue=""
                    >
                        {powerCompanyOptions.map(({ id, name }) => (
                            <MenuItem key={"option-p-c-" + id} value={name}>
                                {formT(name)}
                            </MenuItem>
                        ))}
                    </TextField>
                    <TextField
                        id="voltageType"
                        select
                        label={formT("voltageType")}
                        onChange={changeVoltageType}
                        defaultValue=""
                    >
                        {voltageTypeOptions.map(({ id, name }) => (
                            <MenuItem key={"option-v-t-" + id} value={name}>
                                {formT(name)}
                            </MenuItem>
                        ))}
                    </TextField>
                    <TextField
                        id="touType"
                        select
                        label={formT("touType")}
                        onChange={changeTouType}
                        defaultValue=""
                    >
                        {touTypeOptions.map(({ id, name }) => (
                            <MenuItem key={"option-t-t-" + id} value={name}>
                                {formT(name)}
                            </MenuItem>
                        ))}
                    </TextField>
                    <Divider variant="middle" sx={{ margin: "0 0 2rem" }} />
                    <h4 className="mb-5 ml-2">{pageT("fieldDevices")}</h4>
                    <h5 className="mb-5 ml-2">{formT("deviceType")}</h5>
                    <div
                        className="border-gray-400 border rounded-xl
                    grid grid-cols-2 gap-2 items-center mb-8 p-4"
                    >
                        <FormGroup>
                            {deviceTypeOptions.map(({ id, type, disabled }) => (
                                <FormControlLabel
                                    key={"option-d-t-" + type + id}
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
                            ))}
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
                                            <MenuItem key={"option-d-m-" + id + name} value={id}>
                                                {name}
                                            </MenuItem>
                                        ))}
                                </TextField>
                                {Array.from(Array(deviceInfoCount)).map((_, i) => (
                                    <div className="flex flex-col">
                                        <div className="grid grid-cols-1fr-auto items-center mb-5 ml-2">
                                            <h5 className="">
                                                {hybridInverterSelected
                                                    ? formT("deviceInfo") + ` ${i + 1}`
                                                    : formT("deviceInfo") + ` ${index + 1}`}
                                            </h5>
                                            {i === deviceInfoCount - 1 &&
                                                hybridInverterSelected &&
                                                showAddIcon && <AddIcon onClick={addDeviceInfoGroup} />}
                                        </div>
                                        <TextField
                                            id={`modbusID-${i}`}
                                            type="text"
                                            label={formT("modbusID")}
                                            onChange={hybridInverterSelected
                                                ? (e) => changeModbusID(e, i)
                                                : (e) => changeModbusID(e, index)}
                                            value={hybridInverterSelected
                                                ? deviceInfo.modbusID[i]
                                                : deviceInfo.modbusID[index]}
                                        />
                                        <div className="grid grid-cols-1fr-auto items-center mb-8">
                                            <TextField
                                                focused
                                                sx={{ marginBottom: 0 }}
                                                id={`uueID-${i}`}
                                                label="uueID"
                                                error={uueIDError}
                                                onChange={hybridInverterSelected
                                                    ? (e) => changeUueID(e, i)
                                                    : (e) => changeUueID(e, index)}
                                                value={hybridInverterSelected
                                                    ? deviceInfo.uueID[i]
                                                    : deviceInfo.uueID[index]}
                                            />
                                            <Button
                                                onClick={hybridInverterSelected
                                                    ? () => validateUueID(deviceInfo.uueID[i])
                                                    : () => validateUueID(deviceInfo.uueID[index])}
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
                                            label={formT("powerCapacity")}
                                            onKeyDown={(e) => hybridInverterSelected ? handleKeyDown(e, i) : handleKeyDown(e, index)}
                                            onChange={(e) => changePowerCapacity(e, hybridInverterSelected ? i : index)}
                                            value={hybridInverterSelected ? deviceInfo.powerCapacity[i] : deviceInfo.powerCapacity[index]}
                                        />
                                        <Divider variant="middle" sx={{ margin: "0 0 2rem" }} />
                                    </div>
                                ))}
                            </>
                        )
                    })}
                    {deviceType.includes(TYPE_HYBRID_INVERTER)
                        ? <>
                            <div className="pl-10 flex flex-col">
                                <SubDeviceForm
                                    title={pageT("subdevice")}
                                    mainDeviceType={deviceType}
                                    subDeviceInfo={subDeviceInfo}
                                    changeSubDeviceInfo={changeSubDeviceInfo}
                                    handleKeyDown={handleKeyDown}
                                />
                                <ExtraDeviceInfoForm
                                    subTitle={pageT("extraDeviceInfo")}
                                    gridOutagePercent={extraDeviceInfo.gridOutagePercent}
                                    setGridOutagePercent={(value) =>
                                        setExtraDeviceInfo((prevInfo) => ({
                                            ...prevInfo,
                                            gridOutagePercent: value,
                                        }))
                                    }
                                    chargingSource={extraDeviceInfo.chargingSource}
                                    setChargingSource={(value) =>
                                        setExtraDeviceInfo((prevInfo) => ({
                                            ...prevInfo,
                                            chargingSource: value,
                                        }))
                                    }
                                    energyCapacity={extraDeviceInfo.energyCapacity}
                                    setEnergyCapacity={(value) =>
                                        setExtraDeviceInfo((prevInfo) => ({
                                            ...prevInfo,
                                            energyCapacity: value,
                                        }))
                                    }
                                    voltage={extraDeviceInfo.voltage}
                                    setVoltage={(value) =>
                                        setExtraDeviceInfo((prevInfo) => ({
                                            ...prevInfo,
                                            voltage: value,
                                        }))
                                    }
                                    handleKeyDown={handleKeyDown}
                                />
                            </div>
                        </>
                        : null}
                    {deviceType.includes(TYPE_INVERTER)
                        ? <>
                            <div className="pl-10 flex flex-col">
                                <SubDeviceForm
                                    title={pageT("subdevice")}
                                    mainDeviceType={deviceType}
                                    subDeviceInfo={subDeviceInfo}
                                    changeSubDeviceInfo={changeSubDeviceInfo}
                                    handleKeyDown={handleKeyDown}
                                />
                            </div>
                        </>
                        : null}
                    {deviceType.includes(TYPE_BATTERY)
                        ? <>
                            <div className="pl-10 flex flex-col">
                                <ExtraDeviceInfoForm
                                    subTitle={pageT("extraDeviceInfo")}
                                    gridOutagePercent={extraDeviceInfo.gridOutagePercent}
                                    setGridOutagePercent={(value) =>
                                        setExtraDeviceInfo((prevInfo) => ({
                                            ...prevInfo,
                                            gridOutagePercent: value,
                                        }))
                                    }
                                    chargingSource={extraDeviceInfo.chargingSource}
                                    setChargingSource={(value) =>
                                        setExtraDeviceInfo((prevInfo) => ({
                                            ...prevInfo,
                                            chargingSource: value,
                                        }))
                                    }
                                    energyCapacity={extraDeviceInfo.energyCapacity}
                                    setEnergyCapacity={(value) =>
                                        setExtraDeviceInfo((prevInfo) => ({
                                            ...prevInfo,
                                            energyCapacity: value,
                                        }))
                                    }
                                    voltage={extraDeviceInfo.voltage}
                                    setVoltage={(value) =>
                                        setExtraDeviceInfo((prevInfo) => ({
                                            ...prevInfo,
                                            voltage: value,
                                        }))
                                    }
                                    handleKeyDown={handleKeyDown}
                                />
                            </div>
                        </>
                        : null}
                    <div className="mb-8 flex items-baseline">
                        <p className="ml-1 mr-2">{formT("enableField")}</p>
                        <Switch checked={enable} onChange={handleSwitch} />
                    </div>
                </div>
                <Divider variant="middle" />
                <DialogActions sx={{ margin: "1rem 0.8rem 1rem 0" }}>
                    <Button onClick={cancelClick}
                        sx={{ marginRight: "0.4rem" }}
                        size="large"
                        radius="pill"
                        variant="outlined"
                        color="gray">
                        {commonT("cancel")}
                    </Button>
                    <Button onClick={() => submit()}
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
    )
})