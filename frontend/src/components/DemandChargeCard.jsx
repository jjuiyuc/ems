import { connect } from "react-redux"
import { Button, TextField, InputAdornment } from "@mui/material"
import { useEffect, useState } from "react"
import { useTranslation } from "react-multi-lang"

import { apiCall } from "../utils/api"
import { validateNum } from "../utils/utils"

import { ReactComponent as DemandChargeIcon } from "../assets/icons/demand_charge_line.svg"

const mapState = state => ({
    gatewayID: state.gateways.active.gatewayID
})
const mapDispatch = dispatch => ({
    updateSnackbarMsg: value =>
        dispatch({ type: "snackbarMsg/updateSnackbarMsg", payload: value }),

})
export default connect(mapState, mapDispatch)(function DemandChargeCard(props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        errorT = string => t("error." + string)

    const
        [maxDemandCapacity, setMaxDemandCapacity] = useState(null),
        [loading, setLoading] = useState(false)
    const
        inputNum = (e) => {
            const num = e.target.value
            const isNum = validateNum(num)
            if (!isNum) return
            setMaxDemandCapacity(num)
        },
        handleSave = () => {
            const newData = maxDemandCapacity
            setMaxDemandCapacity(newData)
        }
    const getData = () => {

        const gatewayID = props.gatewayID

        apiCall({
            onComplete: () => setLoading(false),
            onStart: () => setLoading(true),
            onError: (err) => {
                switch (err) {
                    case 60019:
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

                setMaxDemandCapacity(data.maxDemandCapacity || null)

            },
            url: `/api/device-management/gateways/${gatewayID}/meter-settings`
        })
    }
    useEffect(() => {
        getData()
    }, [props.gatewayID])

    const submit = async () => {

        const gatewayID = props.gatewayID

        const data = { maxDemandCapacity: parseInt(maxDemandCapacity) }

        await apiCall({
            method: "put",
            data,
            onSuccess: () => {
                handleSave()
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
                    case 60033:
                        props.updateSnackbarMsg({
                            type: "error",
                            msg: errorT("updateMeterSettingsError")
                        })
                        break
                    default: setOtherError(err)
                        props.updateSnackbarMsg({
                            type: "error",
                            msg: errorT("failureToSave")
                        })
                }
            },
            url: `/api/device-management/gateways/${gatewayID}/meter-settings`
        })
    }
    return <div className="card mt-8">
        <div className="flex justify-between sm:col-span-2 items-center">
            <div className="flex items-center">
                <div
                    className="bg-gray-400-opacity-20 grid h-12 w-12
                place-items-center rounded-full">
                    <DemandChargeIcon className="h-8 text-gray-400 w-8" />
                </div>
                <h2 className="font-bold ml-4">{commonT("demandCharge")}</h2>
            </div>
            <Button
                onClick={submit}
                key={"s-b-s-d"}
                radius="pill"
                variant="contained">
                {commonT("save")}
            </Button>
        </div>
        <div className="flex items-center mt-12">
            <h5 className="mr-8">{props.title}</h5>
            <div className="mt-6">
                <TextField
                    label={props.label}
                    id="outlined-end-adornment"
                    value={maxDemandCapacity}
                    onChange={inputNum}
                    InputProps={{
                        endAdornment:
                            <InputAdornment position="end">
                                {commonT("kw")}
                            </InputAdornment>
                    }}
                    focused
                />
            </div>
        </div>
    </div>
})