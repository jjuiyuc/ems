import { connect } from "react-redux"
import { Button, Slider, Switch } from "@mui/material"
import { useEffect, useState } from "react"
import { useTranslation } from "react-multi-lang"

import { apiCall } from "../utils/api"

import DemandChargeCard from "../components/DemandChargeCard"
import PowerOutageCard from "../components/PowerOutageCard"
import SettingCard from "../components/SettingCard"

import { ReactComponent as BatteryIcon } from "../assets/icons/battery.svg"

const mapState = state => ({
    gatewayID: state.gateways.active.gatewayID
})
const mapDispatch = dispatch => ({
    updateSnackbarMsg: value =>
        dispatch({ type: "snackbarMsg/updateSnackbarMsg", payload: value }),

})
export default connect(mapState, mapDispatch)(function Settings(props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        errorT = string => t("error." + string),
        pageT = (string, params) => t("settings." + string, params)
    const
        [reservedForGridOutage, setReservedForGridOutage] = useState(1),
        [availableRegularUsage, setAvailableRegularUsage] = useState(100 - Number(reservedForGridOutage)),
        [backupReserve, setBackupReserve] = useState(100),
        [grid, setGrid] = useState(false),
        [loading, setLoading] = useState(false),
        [otherError, setOtherError] = useState("")
    const
        handleSlider = (e) => {
            setReservedForGridOutage(Number(e.target.value))
            setAvailableRegularUsage(100 - Number(e.target.value))
        },
        handleSwitch = () => {
            setGrid(preState => !preState)
        }
    const getData = () => {

        const gatewayID = props.gatewayID

        apiCall({
            onComplete: () => setLoading(false),
            onStart: () => setLoading(true),
            onError: (err) => {
                switch (err) {
                    case 60030:
                        props.updateSnackbarMsg({
                            type: "error",
                            msg: errorT("failureToGenerate")
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

                setReservedForGridOutage(data.reservedForGridOutagePercent || 1)
                setAvailableRegularUsage(100 - Number(data.reservedForGridOutagePercent) || 99)
                setGrid(data.chargingSources.includes("Grid") ? true : false)

            },
            url: `/api/device-management/gateways/${gatewayID}/battery-settings`
        })
    }
    useEffect(() => {
        getData()
    }, [props.gatewayID])

    const submit = async () => {

        const gatewayID = props.gatewayID

        const chargingSources = grid ? "Solar + Grid" : "Solar"

        const data = {
            reservedForGridOutagePercent: parseInt(reservedForGridOutage),
            chargingSources: chargingSources
        }
        await apiCall({
            method: "put",
            data,
            onSuccess: () => {
                props.updateSnackbarMsg({
                    type: "success",
                    msg: t("dialog.modifySuccessfully")
                })
            },
            onError: (err) => {
                switch (err) {
                    case 60022:
                        props.updateSnackbarMsg({
                            type: "error",
                            msg: errorT("fieldDisabled")
                        })
                        break
                    case 60031:
                        props.updateSnackbarMsg({
                            type: "error",
                            msg: errorT("updateBatterySettingsError")
                        })
                        break
                    default: setOtherError(err)
                        props.updateSnackbarMsg({
                            type: "error",
                            msg: errorT("failureToSave")
                        })
                }
            },
            url: `/api/device-management/gateways/${gatewayID}/battery-settings`
        })
    }
    return <>
        <h1 className="mb-8">{pageT("settings")}</h1>
        <div className="card mb-8">
            <div className="flex justify-between sm:col-span-2 items-center">
                <div className="flex items-center">
                    <div
                        className="bg-gray-400-opacity-20 grid h-12 w-12
                            place-items-center rounded-full">
                        <BatteryIcon className="h-8 text-gray-400 w-8" />
                    </div>
                    <h2 className="font-bold ml-4">{commonT("battery")}</h2>
                </div>
                <Button
                    onClick={submit}
                    key={"s-b-s-b"}
                    radius="pill"
                    variant="contained">
                    {commonT("save")}
                </Button>
            </div>
            <div className="lg:grid grid-cols-3 mt-12">
                <div className="col-span-2">
                    <h5 className="font-bold">{pageT("backupReserve")}</h5>
                    <div className=" border-r border-gray-400 border-solid pr-12">
                        <div className="flex items-center ">
                            <SettingCard
                                data={reservedForGridOutage}
                                title={pageT("reservedForGridOutage")} />
                            <h4 className="mx-6">+</h4>
                            <SettingCard
                                data={availableRegularUsage}
                                title={pageT("availableRegularUsage")} />
                            <h4 className="mx-6">=</h4>
                            <SettingCard
                                data={backupReserve}
                                title={pageT("backupReserve")} />
                        </div>
                        <div>
                            <Slider defaultValue={1} min={1} value={reservedForGridOutage}
                                onChange={handleSlider} />
                            <div className="flex justify-between">
                                <p className="text-11px">{pageT("reservedForOutage")}</p>
                                <p className="text-11px">{pageT("regularUsage")}</p>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="lg:grid grid-cols-auto pl-12">
                    <h5 className="font-bold mt-9">{pageT("chargingSources")}</h5>
                    <div className="grid grid-cols-2 gap-5">
                        <div className="subCard bg-gray-700">
                            <p className="text-13px ml-2">{commonT("grid")}</p>
                            <Switch
                                checked={grid}
                                onChange={handleSwitch}
                            />
                        </div>
                        <div className="subCard bg-gray-700">
                            <p className="text-13px ml-2">{commonT("solar")}</p>
                            <Switch
                                checked={true}
                                disabled={true}
                            />
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <PowerOutageCard />
        <DemandChargeCard
            title={pageT("maximumDemandCapacity")}
            label={pageT("maximumDemand")}
        />
    </>
})