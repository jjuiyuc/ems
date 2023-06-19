import { connect } from "react-redux"
import {
    Button, Checkbox, DialogActions, Divider, FormGroup, FormControlLabel, Switch
} from "@mui/material"
import { useTranslation } from "react-multi-lang"
import { useEffect, useState } from "react"

import { apiCall } from "../utils/api"

import DialogForm from "../components/DialogForm"
import { ReactComponent as EditIcon } from "../assets/icons/edit.svg"

const mapState = state => ({
    typeID: state.user.group.typeID
})
const mapDispatch = dispatch => ({
    updateSnackbarMsg: value =>
        dispatch({ type: "snackbarMsg/updateSnackbarMsg", payload: value }),

})
export default connect(mapState, mapDispatch)(function EditField(props) {
    const { row, setFieldList } = props

    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        errorT = string => t("error." + string),
        formT = (string) => t("form." + string),
        pageT = (string, params) => t("fieldManagement." + string, params)

    const
        [openEdit, setOpenEdit] = useState(false),
        [openSync, setOpenSync] = useState(false),
        [enable, setEnable] = useState(false),
        [groups, setGroups] = useState([]),
        [groupDictionary, setGroupDictionary] = useState({}),
        [loading, setLoading] = useState(false),
        [otherError, setOtherError] = useState(""),
        [fetched, setFetched] = useState(false)

    const
        handleClick = () => {
            setOpenEdit(true)
        },
        handleSwitch = () => {
            setEnable(preState => !preState)
        },
        handleChange = (e) => {

            const { value } = e.target
            setGroups(groups => groups.map(group =>
                group.id === parseInt(value)
                    ? { ...group, check: !group.check }
                    : group
            ))
        },
        handleSaveEnable = () => {
            const newEnableState = enable
            onSave({ enable: newEnableState })
        },
        onSave = (row) => {
            const newData = groups.map((value) =>
                value.id === row.id ? row : value
            )
            setGroups(newData)
        },
        handleSaveGroup = () => {
            const newGroups = groups.map(group => ({ ...group }))
            onSave({ groups: newGroups })
            setGroups(newGroups)
        }

    const getDataList = async () => {

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

                setEnable(data.enable)
                setGroups(data.groups?.map((group) => ({
                    id: group.id, check: group.check
                })) || [])
                setGroupDictionary(data.groups?.reduce((acc, cur) => {
                    acc[cur.id] = cur.name
                    return acc
                }, {}) || {})

            },
            url: `/api/device-management/gateways/${gatewayID}`
        })
    }
    //Edit enable
    const submitEnable = async () => {

        const gatewayID = row.gatewayID

        const data = { enable: enable }

        await apiCall({
            method: "put",
            data,
            onSuccess: () => {

                props.updateSnackbarMsg({
                    type: "success",
                    msg: t("dialog.editedSuccessfully")
                })
                handleSaveEnable()
                setOpenEdit(false)
            },
            onError: (err) => {
                switch (err) {
                    case 60020:
                        props.updateSnackbarMsg({
                            type: "error",
                            msg: errorT("failureToEdit")
                        })
                        break
                    default: setOtherError(err)
                }
            },
            url: `/api/device-management/gateways/${gatewayID}/field-state`
        })
    }
    //Edit groups
    const submitGroup = async () => {

        const gatewayID = row.gatewayID

        const data = { groups: groups }

        await apiCall({
            method: "put",
            data,
            onSuccess: () => {

                props.updateSnackbarMsg({
                    type: "success",
                    msg: t("dialog.editedSuccessfully")
                })
                handleSaveGroup()
                getDataList()
                setOpenEdit(false)
            },
            onError: (err) => {
                switch (err) {
                    case 60021:
                        props.updateSnackbarMsg({
                            type: "error",
                            msg: errorT("failureToEdit")
                        })
                        break
                    default: setOtherError(err)
                }
            },
            url: `/api/device-management/gateways/${gatewayID}/account-groups`
        })
    }
    //Sync device settings
    const submitSync = async () => {

        const gatewayID = row.gatewayID

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
                    case 60022:
                        props.updateSnackbarMsg({
                            type: "error",
                            msg: errorT("fieldDisabled")
                        })
                        break
                    case 60023:
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
            onSuccess: (rawData) => {

                props.updateSnackbarMsg({
                    type: "success",
                    msg: t("dialog.syncSuccessfully")
                })

                setOpenSync(false)
                setOpenEdit(false)

                if (!rawData?.data) return

                const { data } = rawData
                setFieldList(data)

            },
            url: `/api/device-management/gateways/${gatewayID}/sync-device-settings`
        })
    }

    useEffect(() => {
        if (fetched == false && openEdit == true) {
            getDataList()
        }
    }, [fetched, openEdit])

    return <>
        <EditIcon
            className="mr-5"
            onClick={handleClick} />
        <DialogForm
            dialogTitle={pageT("editFieldInfo")}
            fullWidth={true}
            maxWidth="md"
            open={openEdit}
            setOpen={setOpenEdit}>
            <Divider variant="middle" />
            <div className="flex flex-col m-auto mt-4 min-w-49 w-fit">
                <Button
                    onClick={() => { setOpenSync(true) }}
                    size="medium"
                    variant="outlined"
                    sx={{ margin: "1rem 2rem" }}
                    color="brand">
                    {formT("syncDeviceSettings")}
                </Button>
                <DialogForm
                    dialogTitle={t("dialog.syncMsg")}
                    fullWidth={true}
                    maxWidth="sm"
                    open={openSync}
                    setOpen={setOpenSync}>
                    <div className="flex">
                        <h5 className="ml-6 mr-2">{commonT("gatewayID")} :</h5>
                        {row?.gatewayID || ""}
                    </div>
                    <DialogActions sx={{ margin: "0.5rem 0.5rem 0.5rem 0" }}>
                        <Button onClick={() => { setOpenSync(false) }}
                            radius="pill"
                            variant="outlined"
                            color="gray">
                            {commonT("cancel")}
                        </Button>
                        <Button onClick={submitSync} autoFocus
                            radius="pill"
                            variant="contained"
                            color="warning"
                            sx={{ color: "#000" }}>
                            {commonT("sync")}
                        </Button>
                    </DialogActions>
                </DialogForm>
                <Divider variant="middle" sx={{ margin: "2rem 0" }} />
                <div className="flex justify-between">
                    <div className="flex items-baseline">
                        <p className="ml-1 mr-2">{formT("enableField")}</p>
                        <Switch
                            checked={enable}
                            onChange={handleSwitch}
                        />
                    </div>
                    <DialogActions >
                        <Button
                            onClick={submitEnable}
                            radius="pill"
                            variant="contained"
                            color="primary">
                            {commonT("save")}
                        </Button>
                    </DialogActions>
                </div>
                <Divider variant="middle" sx={{ margin: "2rem 0" }} />
                <h5 className="mb-5">{commonT("group")}</h5>
                <div className="border-gray-400 border rounded-xl
                    grid grid-cols-2 gap-2 items-center mb-4 p-4">
                    <FormGroup>
                        {Object.entries(groupDictionary)?.map(([key, value]) =>
                            <FormControlLabel
                                key={"option-g-" + key}
                                control={
                                    <Checkbox
                                        checked={groups.some(item => item.id === parseInt(key) && item.check === true)}
                                        value={key}
                                        onChange={handleChange}
                                        disabled={parseInt(key) === props.typeID ? true : false}
                                        name={value} />
                                }
                                label={value}
                            />
                        )}
                    </FormGroup>
                </div>
                <DialogActions>
                    <Button
                        onClick={submitGroup}
                        radius="pill"
                        variant="contained"
                        color="primary">
                        {commonT("save")}
                    </Button>
                </DialogActions>
            </div>
            <DialogActions sx={{ margin: "1rem 1.5rem 1rem 0" }}>
                <Button
                    onClick={() => { setOpenEdit(false) }}
                    radius="pill"
                    variant="outlined"
                    color="gray">
                    {commonT("cancel")}
                </Button>
            </DialogActions>
        </DialogForm>
    </>
})