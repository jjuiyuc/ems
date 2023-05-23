import { connect } from "react-redux"
import { Button, DialogActions, Divider, TextField } from "@mui/material"
import { useTranslation } from "react-multi-lang"
import { useState } from "react"

import { apiCall } from "../utils/api"

import DialogForm from "../components/DialogForm"
import { ReactComponent as EditIcon } from "../assets/icons/edit.svg"

const mapDispatch = dispatch => ({
    updateSnackbarMsg: value =>
        dispatch({ type: "snackbarMsg/updateSnackbarMsg", payload: value }),

})
export default connect(null, mapDispatch)(function EditGroup(props) {
    const { row, groupList, onSave = () => { } } = props
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        errorT = string => t("error." + string)

    const
        [openEdit, setOpenEdit] = useState(false),
        [groupName, setGroupName] = useState(row.name),
        [groupNameError, setGroupNameError] = useState(null),
        [isGroupNameError, setIsGroupNameError] = useState(false),
        [otherError, setOtherError] = useState(""),
        [fullWidth, setFullWidth] = useState(true),
        [maxWidth, setMaxWidth] = useState("sm")

    const submitDisabled = !groupName.length || isGroupNameError
    const
        changeGroupName = (e) => {
            const
                groupNameTarget = e.target.value,
                groupNameError = groupNameTarget.length == 0 || groupNameTarget.length > 20
            setGroupName(groupNameTarget)
            setIsGroupNameError(groupNameError)
        },
        handleClick = () => {
            setOpenEdit(true)
        },
        handleSave = () => {
            if (!isGroupNameError) {
                const newName = groupName

                onSave({
                    ...row,
                    name: newName,
                })
                setOpenEdit(false)
            }
        }
    const submit = async () => {

        const groupID = row.id

        const data = { name: groupName }

        await apiCall({
            method: "put",
            data,
            onSuccess: () => {
                handleSave()
                props.updateSnackbarMsg({
                    type: "success",
                    msg: t("dialog.modifyNameMsg")
                })
            },
            onError: (err) => {
                switch (err) {
                    case 60003:
                        setGroupNameError({ type: "groupNameExistsOnTheSameLevel" })
                        props.updateSnackbarMsg({
                            type: "error",
                            msg: errorT("groupNameExistsOnTheSameLevel")
                        })
                        break
                    case 60006:
                        setGroupNameError({ type: "modifyOwnAccountGroupNotAllow" })
                        props.updateSnackbarMsg({
                            type: "error",
                            msg: errorT("modifyOwnAccountGroupNotAllow")
                        })
                        break
                    default: setOtherError(err)
                        props.updateSnackbarMsg({
                            type: "error",
                            msg: errorT("failureToSave")
                        })
                }
            },
            url: `/api/account-management/groups/${groupID}`
        })
    }
    return <>
        <EditIcon
            className="mr-5"
            onClick={handleClick} />
        <DialogForm
            dialogTitle={commonT("group")}
            fullWidth={fullWidth}
            maxWidth={maxWidth}
            open={openEdit}
            setOpen={setOpenEdit}>
            <Divider variant="middle" />
            <div className="flex flex-col m-auto mt-4 min-w-49 w-fit">
                <TextField
                    id="edit-name"
                    label={commonT("groupName")}
                    onChange={changeGroupName}
                    value={groupName || ""}
                    focused>
                </TextField>
            </div>
            <DialogActions sx={{ margin: "1rem 1.5rem 1rem 0" }}>
                <Button
                    onClick={() => { setOpenEdit(false) }}
                    radius="pill"
                    variant="outlined"
                    color="gray">
                    {commonT("cancel")}
                </Button>
                <Button
                    onClick={submit}
                    disabled={submitDisabled}
                    size="large"
                    radius="pill"
                    variant="contained"
                    color="primary">
                    {commonT("save")}
                </Button>
            </DialogActions>
        </DialogForm>
    </>
})