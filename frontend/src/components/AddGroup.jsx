import { connect } from "react-redux"
import { Button, DialogActions, Divider, FormControl, MenuItem, TextField } from "@mui/material"
import AddIcon from "@mui/icons-material/Add"
import { useTranslation } from "react-multi-lang"
import { useMemo, useState } from "react"

import { apiCall } from "../utils/api"

import DialogForm from "../components/DialogForm"
const mapDispatch = dispatch => ({
    updateSnackbarMsg: value =>
        dispatch({ type: "snackbarMsg/updateSnackbarMsg", payload: value }),

})
export default connect(null, mapDispatch)(function AddGroup(props) {
    const { getList, groupList, groupTypes } = props

    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        errorT = string => t("error." + string),
        pageT = (string, params) => t("accountManagementGroup." + string, params)

    const
        [openAdd, setOpenAdd] = useState(false),
        [groupName, setGroupName] = useState(""),
        [isGroupNameError, setIsGroupNameError] = useState(false),
        [groupType, setGroupType] = useState(null),
        [groupTypeError, setGroupTypeError] = useState(false),
        [parentGroup, setParentGroup] = useState(null),
        [parentGroupError, setParentGroupError] = useState(false),
        [fullWidth, setFullWidth] = useState(true),
        [maxWidth, setMaxWidth] = useState("lg")

    const submitDisabled = !groupName.length || groupType == null || parentGroup == null || isGroupNameError || groupTypeError || parentGroupError
    const
        changeGroupName = (e) => {
            const
                groupNameTarget = e.target.value,
                groupNameError = groupNameTarget.length == 0 || groupNameTarget.length > 20
            setGroupName(groupNameTarget)
            setIsGroupNameError(groupNameError)
        },
        changeGroupType = (e) => {
            setGroupType(e.target.value)
        }
    const parentGroupTypeOptions = useMemo(() => {
        if (groupType !== null) {
            if (groupType === 1) {
                return groupList.filter(item => item.parentID === null)
            } else {
                // 其他 groupType 將 parentGroup 設為對應的 typeID - 1
                return groupList.filter(item => item.typeID === parseInt(groupType) - 1)
            }
        }
        return []
    }, [groupType, groupList])

    const
        submit = () => {
            const data = {
                name: groupName,
                typeID: parseInt(groupType),
                parentID: parseInt(parentGroup)
            }
            apiCall({
                method: "post",
                data,
                onSuccess: () => {
                    setOpenAdd(false)
                    getList()
                    props.updateSnackbarMsg({
                        type: "success",
                        msg: t("dialog.addedSuccessfully")
                    })
                    setGroupName("")
                },
                onError: () => {
                    props.updateSnackbarMsg({
                        type: "error",
                        msg: errorT("failureToSave")
                    })
                },
                url: "/api/account-management/groups"
            })
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
            dialogTitle={commonT("group")}
            fullWidth={fullWidth}
            maxWidth={maxWidth}
            open={openAdd}
            setOpen={setOpenAdd}>
            <Divider variant="middle" />
            <FormControl sx={{
                display: "flex",
                flexDirection: "column",
                margin: "auto",
                width: "fit-content",
                mt: 2,
                minWidth: 120
            }}>
                <TextField
                    id="add-name"
                    label={commonT("groupName")}
                    value={groupName}
                    onChange={changeGroupName}
                    error={isGroupNameError}
                    helperText={errorT("nameLength")}
                    required
                />
                <TextField
                    id="add-type"
                    select
                    label={pageT("groupType")}
                    onChange={changeGroupType}
                    defaultValue=""
                    required
                >
                    {Object.entries(groupTypes).map(([key, value]) =>
                        <MenuItem key={"g-t-p" + key} value={key}>
                            {value}
                        </MenuItem>)}

                </TextField>
                <TextField
                    id="add-parent-group-type"
                    select
                    label={pageT("parentGroup")}
                    onChange={e => setParentGroup(e.target.value)}
                    value={parentGroup || ""}
                    disabled={!groupType || groupType == 1}
                    defaultValue=""
                    required
                >
                    {parentGroupTypeOptions.map((option) => (
                        <MenuItem key={"p-g-t" + option.id} value={option.id}>
                            {option.name}
                        </MenuItem>
                    ))}
                </TextField>
            </FormControl>
            <Divider variant="middle" />
            <DialogActions sx={{ margin: "1rem 0.5rem 1rem 0" }}>
                <Button onClick={() => { setOpenAdd(false) }}
                    radius="pill"
                    variant="outlined"
                    color="gray">
                    {commonT("cancel")}
                </Button>
                <Button onClick={submit}
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