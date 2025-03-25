import { connect } from "react-redux"
import { Button, DialogActions } from "@mui/material"
import { useTranslation } from "react-multi-lang"
import { useState } from "react"

import { apiCall } from "../utils/api"

import DialogForm from "../components/DialogForm"

const mapDispatch = dispatch => ({
    updateSnackbarMsg: value =>
        dispatch({ type: "snackbarMsg/updateSnackbarMsg", payload: value }),

})
export default connect(null, mapDispatch)(function DeleteGroup(props) {
    const {
        row,
        openDelete,
        setOpenDelete,
        groupList,
        setGroupList
    } = props

    const t = useTranslation(),
        commonT = string => t("common." + string),
        dialogT = (string) => t("dialog." + string),
        errorT = string => t("error." + string)

    const [groupNameError, setGroupNameError] = useState(null)
    const [otherError, setOtherError] = useState("")

    const submit = () => {
        const groupID = row.id

        const hasSubGroups = groupList.some(g => g.parentID === groupID)

        if (hasSubGroups) {
            setGroupNameError({ type: "groupHasSubGroup" })
            props.updateSnackbarMsg({
                type: "error",
                msg: errorT("groupHasSubGroup")
            })
            return
        }

        const newGroupList = groupList.filter(g => g.id !== groupID)
        setGroupList(newGroupList)
        setOpenDelete(false)
        props.updateSnackbarMsg({
            type: "success",
            msg: dialogT("deletedSuccessfully")
        })
    }

    return <>
        <DialogForm
            dialogTitle={dialogT("deleteMsg")}
            fullWidth={true}
            maxWidth="sm"
            open={openDelete}
            setOpen={setOpenDelete}>
            <div className="flex">
                <h5 className="ml-6 mr-2">{commonT("groupName")} :</h5>
                {row?.name || ""}
            </div>
            <DialogActions sx={{ margin: "0.5rem 0.5rem 0.5rem 0" }}>
                <Button
                    onClick={() => { setOpenDelete(false) }}
                    radius="pill"
                    variant="outlined"
                    color="gray"
                >
                    {commonT("cancel")}
                </Button>
                <Button
                    onClick={submit}
                    radius="pill"
                    variant="contained"
                    color="negative"
                    sx={{ color: "#ffffff" }}
                >
                    {commonT("delete")}
                </Button>
            </DialogActions>
        </DialogForm>
    </>
})
