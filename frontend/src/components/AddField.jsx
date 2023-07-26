import { connect } from "react-redux"
import {
    Button, Checkbox, DialogActions, Divider, FormControlLabel, FormGroup,
    MenuItem, Switch, TextField
} from "@mui/material"
import AddIcon from "@mui/icons-material/Add"
import { useTranslation } from "react-multi-lang"
import { useEffect, useMemo, useState } from "react"

import { apiCall } from "../utils/api"
import { validateNumTwoDecimalPlaces } from "../utils/utils"

import DialogForm from "../components/DialogForm"
import ExtraDeviceInfoForm from "../components/ExtraDeviceInfoForm"
import SubDeviceForm from "../components/SubDeviceForm"

const mapDispatch = dispatch => ({
    updateSnackbarMsg: value =>
        dispatch({ type: "snackbarMsg/updateSnackbarMsg", payload: value }),

})
const TYPE_HYBRID_INVERTER = "Hybrid-Inverter"
const TYPE_INVERTER = "Inverter"
const TYPE_BATTERY = "Battery"

export default connect(null,mapDispatch)(function AddField(props) {
    const { getList } = props
    const
        powerCompanyOptions = [
            {
                "id": 1,
                "name": "TPC"
            }
        ],
        voltageTypeOptions = [
            {
                "id": 1,
                name: "Low voltage",
            },
            {
                "id": 2,
                name: "High voltage",
            }
        ],
        touTypeOptions = [
            {
                "id": 1,
                name: "Two-section",
            }
        ]
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        errorT = string => t("error." + string),
        formT = (string) => t("form." + string),
        pageT = (string, params) => t("fieldManagement." + string, params)

    const
        [modelList, setModelList] = useState([]),
        [gatewayID, setGatewayID] = useState(""),
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
            powerCapacity: [null, null, null]
        }),
        [uueIDError, setUueIDError] = useState(false),
        [subDeviceInfo, setSubDeviceInfo] = useState({
            subDeviceModel: ["", "", ""],
            subPowerCapacity: [null, null, null]
        }),
        [extraDeviceInfo, setExtraDeviceInfo] = useState({
            gridOutagePercent: "",
            chargingSource: "",
            energyCapacity: null,
            voltage: null,
        }),
        [deviceInfoCount, setDeviceInfoCount] = useState(1),
        [showAddIcon, setShowAddIcon] = useState(true),
        [isHybridInverterSelected, setIsHybridInverterSelected] = useState(false),
        [fullWidth, setFullWidth] = useState(true),
        [maxWidth, setMaxWidth] = useState("lg"),
        [openAdd, setOpenAdd] = useState(false),
        [fetched, setFetched] = useState(false),
        [loading, setLoading] = useState(false),
        [infoError, setInfoError] = useState("")

    const
        getModelList = () => {
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

        const typeOptions = allTypes.map((type) => {
            if (type === TYPE_HYBRID_INVERTER) {
                const otherSelected = deviceType.some((type) => type !== TYPE_HYBRID_INVERTER)
                if (otherSelected) return { type, disabled: true }
            } else {
                const hybridInverterSelected = deviceType.includes(TYPE_HYBRID_INVERTER)
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
        changeGatewayID = (e) => {
            setGatewayID(e.target.value)
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
        changeDeviceInfo = (index, devices,value) => {
            setDeviceInfo(prevInfo => {
                const newDeviceInfo = { ...prevInfo }
                newDeviceInfo[devices][index] = value
                return newDeviceInfo
            })
        },
        uueIDLengthError = deviceInfo.uueID.length < 32 ||deviceInfo.uueID.length >32,
         changeUueID = (e, index) => {
            const uueIDTarget = e.target.value,
                uueIDLengthError = uueIDTarget.length < 32 ||uueIDTarget.length ==0
            setUueIDError(uueIDLengthError)
            changeDeviceInfo(index, "uueID", uueIDTarget)
        },
        changePowerCapacity = (e, index) => {
            const num = e.target.value
            const isNum = validateNumTwoDecimalPlaces(num)
            if (!isNum) return
            changeDeviceInfo(index, "powerCapacity", num)
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
            setSubDeviceInfo(prevInfo => {
                const newSubDeviceInfo = { ...prevInfo }
                newSubDeviceInfo[field][index] = value
                return newSubDeviceInfo
            })
        }
        // useEffect(()=>{},[gatewayID])
       const validateGatewayID = async () => {

            // const gatewayID = gatewayID

        await apiCall({
            method: "get",
            // data,
            onComplete: () => {
                setLoading(false)
                setFetched(true)
            },
            onStart: () => setLoading(true),
            onError: (err) => {
                switch (err) {
                    case 60024:
                        props.updateSnackbarMsg({
                            type: "error",
                            msg: errorT("gatewayIDInvalid")
                        })
                        break
                    case 60025:
                        props.updateSnackbarMsg({
                            type: "error",
                            msg: errorT("syncError")
                        })
                        break
                    default:
                        props.updateSnackbarMsg({
                            type: "error",
                            msg: errorT("error")
                        })
                }
            },
            onSuccess: () => {

                props.updateSnackbarMsg({
                    type: "success",
                    msg: t("dialog.syncSuccessfully")
                })

                // if (!rawData?.data) return

            },
            url: `/api/device-management/gateways/${gatewayID}/validity`
        })
    }
    const
        submit = async () => {

            const data = {
                gatewayID: gatewayID,
                locationName: locationName,
                address: address,
                lat: parseInt(lat),
                lng: parseInt(lng),
                powerCompany: powerCompany,
                voltageType: voltageType,
                touType: touType,
            }
            await apiCall({
                method: "post",
                data,
                onSuccess: () => {
                    setOpenAdd(false)
                    getList()
                    props.updateSnackbarMsg({
                        type: "success",
                        msg: t("dialog.addedSuccessfully")
                    })
                    setAccount("")
                    setPassword("")
                    setName("")
                    setGroup("")
                },
                onError: err => {
                    switch (err) {
                        case 60012:
                            setAccountError(true)
                            props.updateSnackbarMsg({
                                type: "error",
                                msg: errorT("emailExist")
                            })
                            break
                        case 60013:
                            setAccountError(true)
                            props.updateSnackbarMsg({
                                type: "error",
                                msg: errorT("failureToCreate")
                            })
                            break
                        default: setOtherError(err)
                    }
                },
                url: "/api/device-management/gateways"
            })
        },
        cancelClick = () => {
            setOpenAdd(false)
            setAccount("")
            setPassword("")
            setName("")
            setGroup("")
        }
    useEffect(() => {
        // if (openAdd && fetched == false)
        getModelList()
    }, [fetched, openAdd])

    console.log(gatewayID)
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
                    value={locationName}
                />
                <TextField
                    id="address"
                    label={formT("address")}
                    value={address}
                />
                <div className="flex-nowrap">
                    <TextField
                        id="lat"
                        type="number"
                        label={formT("lat")}
                        value={lat}
                    />
                    <TextField
                        id="lng"
                        type="number"
                        label={formT("lng")}
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
                                        <MenuItem key={"option-d-m-" + id + name} value={id}>
                                            {name}
                                        </MenuItem>
                                    ))}
                            </TextField>
                            {Array.from(Array(deviceInfoCount)).map((_, i) => (
                                <div className="flex flex-col">
                                    <div className="grid grid-cols-1fr-auto items-center mb-5 ml-2">
                                        <h5 className="">{formT("deviceInfo") + ` ${i + 1}`}</h5>
                                        {(i === deviceInfoCount - 1) && isHybridInverterSelected && showAddIcon && (
                                            <AddIcon onClick={addDeviceInfoGroup} />
                                        )}
                                    </div>
                                    <TextField
                                        id={`modbusID-${i}`}
                                        type="number"
                                        label={formT("modbusID")}
                                        value={deviceInfo.modbusID[i]}
                                        onChange={(e) => changeSubDeviceModel(e, i)}

                                    />
                                    <div className="grid grid-cols-1fr-auto items-center mb-8">
                                        <TextField
                                            sx={{ marginBottom: 0 }}
                                            id={`uueID-${i}`}
                                            label="UUEID"
                                            error={uueIDLengthError}
                                            onChange={(e) => changeUueID(e, i)}
                                            value={deviceInfo.uueID[i]}
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
                                        onChange={(e) => changePowerCapacity(e, i)}
                                        value={deviceInfo.powerCapacity[i]}
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
})