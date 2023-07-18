import { connect } from "react-redux"
import { Button, Chip, DialogActions, Divider, Switch, TextField } from "@mui/material"
import { useTranslation } from "react-multi-lang"
import { Fragment, useEffect, useState } from "react"

import { apiCall } from "../utils/api"

import DialogForm from "../components/DialogForm"
import InfoExtraDeviceForm from "../components/InfoExtraDeviceForm"
import { ReactComponent as NoticeIcon } from "../assets/icons/notice.svg"

const mapDispatch = dispatch => ({
    updateSnackbarMsg: value =>
        dispatch({ type: "snackbarMsg/updateSnackbarMsg", payload: value }),

})
export default connect(null, mapDispatch)(function InfoField(props) {
    const { row } = props

    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        formT = (string) => t("form." + string),
        errorT = string => t("error." + string),
        pageT = (string, params) => t("fieldManagement." + string, params)

    const
        [openNotice, setOpenNotice] = useState(false),
        [gatewayID, setGatewayID] = useState(""),
        [locationName, setLocationName] = useState(""),
        [address, setAddress] = useState(""),
        [lat, setLat] = useState(""),
        [lng, setLng] = useState(""),
        [powerCompany, setPowerCompany] = useState(""),
        [voltageType, setVoltageType] = useState(""),
        [touType, setTouType] = useState(""),
        [modelType, setModelType] = useState(""),
        [modelNameDict, setModelNameDict] = useState({}),
        [deviceList, setDeviceList] = useState([]),
        [enable, setEnable] = useState(false),
        [groupDict, setGroupDict] = useState({}),
        [loading, setLoading] = useState(false),
        [infoError, setInfoError] = useState(""),
        [fetched, setFetched] = useState(false)

    const getModelList = () => {
        apiCall({
            onError: error => setInfoError(error),
            onSuccess: rawData => {
                if (!rawData?.data) return

                const { data } = rawData

                setModelNameDict(data.models?.reduce((acc, cur) => {
                    acc[cur.id] = cur.name
                    return acc
                }, {}) || {})
            },
            url: `/api/device-management/devices/models`
        })
    }
    const
        iconOnClick = async () => {
            setOpenNotice(true)

            const gatewayID = row.gatewayID
            await apiCall({
                onComplete: () => {
                    setLoading(false)
                    setFetched(true)
                },
                onStart: () => setLoading(true),
                onError: (err) => {
                    switch (err) {
                        case 60019:
                            props.updateSnackbarMsg({
                                type: "error",
                                msg: errorT("noDataMsg")
                            })
                            break
                        default:
                            props.updateSnackbarMsg({
                                type: "error",
                                msg: errorT("failureToGenerate")
                            })
                    }
                },
                onSuccess: rawData => {

                    if (!rawData?.data) return

                    const { data } = rawData

                    setGatewayID(data.gatewayID || "")
                    setLocationName(data.locationName || "")
                    setAddress(data.address || "")
                    setLat(data.lat || "")
                    setLng(data.lng || "")
                    setPowerCompany(data.powerCompany || "")
                    setVoltageType(data.voltageType || "")
                    setTouType(data.touType || "")
                    setModelType(data?.devices?.modelType || "")
                    setDeviceList(data?.devices || [])
                    setEnable(data.enable)
                    setGroupDict(data.groups?.reduce((acc, cur) => {
                        if (cur.check) {
                            acc[cur.id] = cur.name
                        }
                        return acc
                    }, {}))
                },
                url: `/api/device-management/gateways/${gatewayID}`
            })
        }

    useEffect(() => {
        if (openNotice && fetched == false)
            getModelList()
    }, [fetched, openNotice])
    return <>
        <NoticeIcon
            className="mr-5"
            onClick={iconOnClick} />
        <DialogForm
            dialogTitle={pageT("fieldInfo")}
            fullWidth={true}
            maxWidth="md"
            open={openNotice}
            setOpen={setOpenNotice}>
            <Divider variant="middle" />
            <div className="flex flex-col m-auto mt-4 min-w-49 w-fit">
                <TextField
                    sx={{ marginBottom: 2 }}
                    key="gateway-id"
                    label={commonT("gatewayID")}
                    value={gatewayID}
                    focused
                    disabled={true}
                />
                <h5 className="mb-4 ml-2">{pageT("locationInformation")}</h5>
                <TextField
                    key="location-name"
                    label={commonT("locationName")}
                    value={locationName}
                    focused
                    disabled={true}
                />
                <TextField
                    key="address"
                    label={formT("address")}
                    value={address}
                    focused
                    disabled={true}
                />
                <div className="flex-nowrap">
                    <TextField
                        key="lat"
                        type="number"
                        label={formT("lat")}
                        value={lat}
                        focused
                        disabled={true}
                    />
                    <TextField
                        sx={{ marginLeft: "1rem" }}
                        key="lng"
                        type="number"
                        label={formT("lng")}
                        value={lng}
                        disabled={true}
                    />
                </div>
                <TextField
                    key="power-company"
                    label={formT("powerCompany")}
                    value={powerCompany}
                    disabled={true}
                />
                <TextField
                    key="v-t"
                    label={formT("voltageType")}
                    value={voltageType}
                    disabled={true}
                />
                <TextField
                    key="tou-t"
                    label={formT("touType")}
                    value={touType}
                    disabled={true}
                />
                <Divider variant="middle" sx={{ margin: "0.8rem 0 2rem" }} />
                {deviceList.map((item, index) => {
                    let extraContent = null
                    return (
                        <Fragment key={"f-d-" + index}>
                            <h5 className="mb-4 ml-2">
                                {pageT("fieldDevices") + " " + (index + 1)}
                            </h5>
                            <TextField
                                label={formT("deviceType")}
                                value={item?.modelType}
                                disabled={true}
                            />
                            <TextField
                                label={formT("deviceModel")}
                                value={modelNameDict?.[item.modelID]}
                                disabled={true}
                            />
                            <h5 className="mb-4 ml-2">
                                {formT("deviceInfo") + ` ${index + 1}`}
                            </h5>
                            <TextField
                                key="m-id"
                                type="number"
                                label={formT("modbusID")}
                                value={item.modbusID}
                                disabled={true}
                            />
                            <TextField
                                key="uueid"
                                label="UUEID"
                                value={item?.uueID}
                                disabled={true}
                            />
                            <TextField
                                key="p-c"
                                type="number"
                                label={formT("powerCapacity")}
                                value={item.powerCapacity}
                                disabled={true}
                            />
                            {/* Battery */}
                            {item?.modelType === "Battery" ? <>
                                <InfoExtraDeviceForm
                                    key="extra-info"
                                    subTitle={pageT("extraDeviceInfo")}
                                    voltage={item.extraInfo?.voltage}
                                    energyCapacity={item.extraInfo?.energyCapacity}
                                    chargingSource={item.extraInfo?.chargingSources}
                                    gridOutagePercent={item.extraInfo?.reservedForGridOutagePercent}
                                />
                            </>
                                : null}
                            {item.subDevices?.map((subItem, subIndex) => {

                                let subDeviceContent = null
                                //Inverter - sub: PV
                                if (item?.modelType === "Inverter" && subItem?.modelType === "PV") {
                                    subDeviceContent = <>
                                        <TextField
                                            key="i-sub-d-t-"
                                            label={formT("deviceType")}
                                            value={subItem?.modelType}
                                            disabled={true}
                                        />
                                        <TextField
                                            key="i-sub-d-m-"
                                            label={formT("deviceModel")}
                                            value={modelNameDict?.[subItem.modelID]}
                                            disabled={true}
                                        />
                                        <h5 className="mb-5 ml-2">
                                            {formT("deviceInfo") + ` ${subIndex + 1}`}
                                        </h5>
                                        <TextField
                                            key="i-sub-p-c-"
                                            type="number"
                                            label={formT("powerCapacity")}
                                            value={subItem?.powerCapacity}
                                            disabled={true}
                                        />
                                    </>
                                }
                                //Battery
                                if (subItem?.modelType === "Battery") {
                                    subDeviceContent = <InfoExtraDeviceForm
                                        key="b-sub-extra-info"
                                        subTitle={pageT("extraDeviceInfo")}
                                        voltage={subItem.extraInfo?.voltage}
                                        energyCapacity={subItem.extraInfo?.energyCapacity}
                                        chargingSource={subItem.extraInfo?.chargingSources}
                                        gridOutagePercent={subItem.extraInfo?.reservedForGridOutagePercent}
                                    />
                                }
                                //Hybrid-Inverter
                                if (item?.modelType === "Hybrid-Inverter") {
                                    subDeviceContent = <>
                                        <TextField
                                            key="h-sub-d-t-"
                                            label={formT("deviceType")}
                                            value={subItem?.modelType}
                                            disabled={true}
                                        />
                                        <TextField
                                            key="h-sub-d-m-"
                                            label={formT("deviceModel")}
                                            value={modelNameDict?.[subItem.modelID]}
                                            disabled={true}
                                        />
                                        <h5 className="mb-5 ml-2">
                                            {formT("deviceInformation") + " " + (subIndex + 1)}
                                        </h5>
                                        <TextField
                                            key="h-p-c-"
                                            type="number"
                                            label={formT("powerCapacity")}
                                            value={subItem?.powerCapacity}
                                            disabled={true}
                                        />
                                        {/* sub: Battery */}
                                        {subItem?.modelType === "Battery" &&
                                            <InfoExtraDeviceForm
                                                key="h-extra-i-"
                                                subTitle={pageT("extraDeviceInfo")}
                                                voltage={subItem.extraInfo?.voltage}
                                                energyCapacity={subItem.extraInfo?.energyCapacity}
                                                chargingSource={subItem.extraInfo?.chargingSources}
                                                gridOutagePercent={subItem.extraInfo?.reservedForGridOutagePercent}
                                            />}
                                    </>
                                }
                                return <div className="pl-10 flex flex-col">
                                    <h5 className="mb-4 ml-2">
                                        {pageT("subdevice") + " " + (subIndex + 1)}
                                    </h5>
                                    {subDeviceContent}
                                </div>
                            })}
                            <Divider variant="middle" sx={{ margin: "0.8rem 0 1rem" }} />
                        </Fragment>
                    )
                })}
                <div className="flex items-baseline">
                    <p className="ml-1 mr-2">{formT("enableField")}</p>
                    <Switch
                        checked={enable}
                        disabled={true} />
                </div>
                <Divider variant="middle" sx={{ margin: "1rem 0" }} />
                <h5 className="mb-5">{commonT("group")}</h5>
                <div className="border-gray-400 border rounded-xl
                    grid grid-cols-3 gap-2 items-center mb-4 p-4">
                    {Object.entries(groupDict).map(([key, value]) =>
                        <Chip key={"g-t-p-" + key} label={value}
                            variant="outlined" color="primary"
                        />
                    )}
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
})