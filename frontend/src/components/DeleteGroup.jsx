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
    const { row, getList, openDelete, setOpenDelete } = props
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        dialogT = (string) => t("dialog." + string),
        errorT = string => t("error." + string)

    const
        [groupNameError, setGroupNameError] = useState(null),
        [otherError, setOtherError] = useState("")

    const
        submit = async () => {

            const groupID = row.id

            await apiCall({
                method: "delete",
                data: null,
                onSuccess: () => {
                    setOpenDelete(false)
                    getList()
                    props.updateSnackbarMsg({
                        type: "success",
                        msg: dialogT("deletedSuccessfully")
                    })
                },
                onError: (err) => {
                    switch (err) {
                        case 60008:
                            setGroupNameError({ type: "groupHasSubGroup" })
                            props.updateSnackbarMsg({
                                type: "error",
                                msg: errorT("groupHasSubGroup")
                            })
                            break
                        case 60009:
                            setGroupNameError({ type: "groupHasUser" })
                            props.updateSnackbarMsg({
                                type: "error",
                                msg: errorT("groupHasUser")
                            })
                            break
                        default: setOtherError(err)
                            props.updateSnackbarMsg({
                                type: "error",
                                msg: errorT("failureToDelete")
                            })
                    }
                },
                url: `/api/account-management/groups/${groupID}`
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
                <Button onClick={() => { setOpenDelete(false) }}
                    radius="pill"
                    variant="outlined"
                    color="gray">
                    {commonT("cancel")}
                </Button>
                <Button onClick={submit}
                    radius="pill"
                    variant="contained"
                    color="negative"
                    sx={{ color: "#ffffff" }}>
                    {commonT("delete")}
                </Button>
            </DialogActions>
        </DialogForm>
    </>
})