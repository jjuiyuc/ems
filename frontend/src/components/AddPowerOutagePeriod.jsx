import { Button, DialogActions, Divider, MenuItem, TextField } from "@mui/material"
import AddIcon from "@mui/icons-material/Add"
import { connect } from "react-redux"
import DatePicker from "react-datepicker"
import { useTranslation } from "react-multi-lang"
import { useEffect, useState } from "react"
import moment from "moment"
import "../assets/css/timeRangePicker.css"

import { apiCall } from "../utils/api"

import DialogForm from "./DialogForm"

import { ReactComponent as Add_Icon } from "../assets/icons/add.svg"
import { ReactComponent as DeleteIcon } from "../assets/icons/delete.svg"

const mapState = state => ({
    gatewayID: state.gateways.active.gatewayID
})
const mapDispatch = dispatch => ({
    updateSnackbarMsg: value =>
        dispatch({ type: "snackbarMsg/updateSnackbarMsg", payload: value }),

})
const maxLength = 3

const defaultPolicyConfig = {
    preNotifiedOutagePeriod: {
        name: "preNotifiedOutagePeriod",
        tempName: "preNotifiedOutagePeriod",
        extensible: true,
        deletable: false
    }
}
const defaultPolicyTime = {
    preNotifiedOutagePeriod: [
        { startDate: "", endDate: "", type: "" },
    ]
}
export default connect(mapState, mapDispatch)(function AddPowerOutagePeriod(props) {
    const { getList } = props

    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        errorT = string => t("error." + string),
        pageT = (string, params) => t("settings." + string, params)

    const powerOutageTypes = [
        {
            "id": 1,
            "name": "advanceBlackout"
        },
        {
            "id": 2,
            "name": "evCharge"
        }
    ]
    const
        [openAdd, setOpenAdd] = useState(false),
        [policyConfig, setPolicyConfig] = useState(defaultPolicyConfig),
        [policyTime, setPolicyTime] = useState(defaultPolicyTime),
        [startDate, setStartDate] = useState(null),
        [endDate, setEndDate] = useState(null),
        [type, setType] = useState(null),
        [typeDict, setTypeDict] = useState({}),
        [timeError, setTimeError] = useState(false),
        [otherError, setOtherError] = useState("")

    const timeChangeError = startDate < moment().toDate() || startDate >= endDate
    const submitDisabled = type == null || timeChangeError || timeError == true

    const
        generateTypeDict = () => {
            setTypeDict(powerOutageTypes.reduce((acc, cur) => {
                acc[cur.id] = cur.name
                return acc
            }, {}) || {})
        },
        changeType = (e) => {
            setType(e.target.value)
        }
    useEffect(() => {
        generateTypeDict()
    }, [])
    const
        submit = async () => {
            const gatewayID = props.gatewayID

            const data = {
                "periods": [
                    {
                        "startTime": moment(startDate).toISOString(),
                        "endTime": moment(endDate).toISOString(),
                        "type": type
                    },
                ]
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
                    setStartDate("")
                    setEndDate("")
                    setType("")
                },
                onError: err => {
                    switch (err) {
                        case 60033:
                            setAccountError(true)
                            props.updateSnackbarMsg({
                                type: "error",
                                msg: errorT("emailExist")
                            })
                            break
                        case 60034:
                            setAccountError(true)
                            props.updateSnackbarMsg({
                                type: "error",
                                msg: errorT("failureToCreate")
                            })
                            break
                        default: setOtherError(err)
                    }
                },
                url: `/api/device-management/gateways/${gatewayID}/power-outage-periods`
            })
        },

        cancelClick = () => {
            setOpenAdd(false)
            setStartDate("")
            setEndDate("")
            setType("")
        }
    return <>
        <Button
            onClick={() => { setOpenAdd(true) }}
            size="medium"
            variant="outlined"
            radius="pill"
            fontSize="medium"
            color="brand"
            startIcon={<AddIcon />}>
            {commonT("add")}
        </Button>
        <DialogForm
            dialogTitle={pageT("addPowerOutagePeriod")}
            fullWidth={true}
            maxWidth="lg"
            open={openAdd}
            setOpen={setOpenAdd}>
            <Divider variant="middle" />
            <div className="flex flex-col m-auto mt-4 min-w-49 w-fit">
                <div className="flex items-start mt-12">
                    <div className="mb-2">
                        {Object.keys(policyConfig).map((policy) => {
                            const timeGroup = policyTime[policy]
                            return (
                                <div className="mb-12 ml-12" key={policy}>
                                    <div className="flex items-center text-white mb-4">
                                        <h5 className="font-bold">{pageT(policyConfig[policy].name)}</h5>
                                    </div>
                                    {timeGroup.map(({ startDate, endDate }, index) => {
                                        return (
                                            <div key={`${policy}-${index}`}
                                                className="time-range-picker grid
                                        grid-cols-settings-input-col5 gap-x-4 items-center mt-4">
                                                <div>
                                                    <h6 className="mb-1 ml-1">{pageT("startDate")}</h6>
                                                    <DatePicker
                                                        showTimeSelect
                                                        timeFormat="HH:mm"
                                                        timeIntervals={15}
                                                        selected={startDate}
                                                        onChange={(date) => {
                                                            const newPolicyTime = {
                                                                ...policyTime,
                                                                [policy]: timeGroup.map((row, i) =>
                                                                    i === index
                                                                        ? { ...row, startDate: date }
                                                                        : row)
                                                            }
                                                            setPolicyTime(newPolicyTime)
                                                            setStartDate(date)
                                                            if (timeChangeError) {
                                                                setTimeError(true)
                                                            } else {
                                                                setTimeError(false)
                                                            }
                                                        }}
                                                        value={startDate ? moment(startDate).format("yyyy/MM/DD HH:mm") : ""}
                                                        selectsStart
                                                        startDate={startDate}
                                                        endDate={endDate}
                                                        minDate={moment(new Date())._d}
                                                    />
                                                </div>
                                                <span className="mt-6">{pageT("to")}</span>
                                                <div>
                                                    <h6 className="mb-1 ml-1">{pageT("endDate")}</h6>
                                                    <DatePicker
                                                        showTimeSelect
                                                        timeFormat="HH:mm"
                                                        timeIntervals={15}
                                                        selected={endDate}
                                                        onChange={(date) => {
                                                            const newPolicyTime = {
                                                                ...policyTime,
                                                                [policy]: timeGroup.map((row, i) =>
                                                                    i === index
                                                                        ? { ...row, endDate: date }
                                                                        : row)
                                                            }
                                                            setPolicyTime(newPolicyTime)
                                                            setEndDate(date)
                                                            if (timeChangeError) {
                                                                setTimeError(true)
                                                            } else {
                                                                setTimeError(false)
                                                            }
                                                        }}
                                                        selectsEnd
                                                        endDate={endDate}
                                                        startDate={startDate}
                                                        value={endDate ? moment(endDate).format("yyyy/MM/DD HH:mm") : ""}
                                                        minDate={startDate}
                                                        minTime={moment(startDate).add(15, "minutes")._d}
                                                        maxTime={endDate <= startDate ? moment(startDate).add(15, "minutes").endOf("day")._d : moment().startOf("day")._d}
                                                        disabled={!startDate}
                                                    />
                                                </div>
                                                <div className="flex flex-col m-auto min-w-49 w-fit">
                                                    <h6 className="mb-1 ml-1">{pageT("type")}</h6>
                                                    <TextField
                                                        id="p-o-type"
                                                        select
                                                        variant="outlined"
                                                        onChange={changeType}
                                                        value={type}
                                                        defaultValue="">
                                                        {Object.entries(typeDict).map(([key, value]) =>
                                                            <MenuItem key={"type-o-" + key} value={value}>
                                                                {pageT(`${value}`)}
                                                            </MenuItem>)}
                                                    </TextField>
                                                </div>
                                                {index ?
                                                    <div className="ml-2 mt-4 h-4 w-4 flex cursor-pointer">
                                                        <DeleteIcon
                                                            onClick={() => {
                                                                const newPolicyTime = {
                                                                    ...policyTime,
                                                                    [policy]: timeGroup.filter((_, i) => i !== index)
                                                                }
                                                                setPolicyTime(newPolicyTime)
                                                            }}
                                                        />
                                                    </div> : <div></div>}
                                            </div>
                                        )
                                    })}
                                    {policyConfig[policy].extensible && timeGroup.length < maxLength ?
                                        <button
                                            className="flex ml-4 mt-4"
                                            onClick={() => {
                                                const newPolicyTime = {
                                                    ...policyTime,
                                                    [policy]: [
                                                        ...timeGroup,
                                                        { startDate: "", endDate: "", type: "" }
                                                    ]
                                                }
                                                setPolicyTime(newPolicyTime)
                                            }}>
                                            <Add_Icon className="w-4 h-4 mt-0.7 mr-1" />
                                            {pageT("addDateRange")}
                                        </button> : null}
                                </div>)
                        })}
                    </div>
                </div>
                {otherError
                    ? <div className="box mb-8 negative text-center text-red-400">
                        {otherError}
                    </div>
                    : null}
            </div>
            <Divider variant="middle" />
            <DialogActions sx={{ margin: "1rem 0.5rem 1rem 0" }}>
                <Button onClick={cancelClick}
                    radius="pill"
                    variant="outlined"
                    color="gray">
                    {commonT("cancel")}
                </Button>
                <Button
                    onClick={submit}
                    disabled={submitDisabled}
                    radius="pill"
                    variant="contained"
                    color="primary">
                    {commonT("add")}
                </Button>
            </DialogActions>
        </DialogForm>
    </>
})