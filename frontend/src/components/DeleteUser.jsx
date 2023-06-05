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
export default connect(null, mapDispatch)(function DeleteUser(props) {
    const { getList, openDelete, row, setOpenDelete } = props
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        errorT = string => t("error." + string),
        pageT = (string, params) => t("accountManagementUser." + string, params)

    const
        [deleteError, setDeleteError] = useState(null),
        [otherError, setOtherError] = useState("")

    const
        submit = async () => {
            await apiCall({
                method: "delete",
                data: null,
                onSuccess: () => {
                    setOpenDelete(false)
                    getList()
                    props.updateSnackbarMsg({
                        type: "success",
                        msg: t("dialog.deletedSuccessfully")
                    })
                },
                onError: (err) => {
                    switch (err) {
                        case 60015:
                            setDeleteError({ type: "deleteOwnAccountNotAllow" })
                            props.updateSnackbarMsg({
                                type: "error",
                                msg: errorT("deleteOwnAccountNotAllow")
                            })
                            break
                        case 60016:
                            setDeleteError({ type: "deleteAccountUserError" })
                            props.updateSnackbarMsg({
                                type: "error",
                                msg: errorT("deleteAccountUserError")
                            })
                            break
                        default: setOtherError(err)
                            props.updateSnackbarMsg({
                                type: "error",
                                msg: errorT("failureToDelete")
                            })
                    }
                },
                url: `/api/account-management/users/${row.id}`
            })
    }
    return (
        <DialogForm
            dialogTitle={t("dialog.deleteMsg")}
            fullWidth={true}
            maxWidth="sm"
            open={openDelete}
            setOpen={setOpenDelete}>
            <div className="flex">
                <h5 className="ml-6 mr-2">{pageT("account")} :</h5>
                {row?.username || ""}
            </div>
            <DialogActions sx={{ margin: "0.5rem 0.5rem 0.5rem 0" }}>
                <Button onClick={() => { setOpenDelete(false) }}
                    radius="pill"
                    variant="outlined"
                    color="gray">
                    {commonT("cancel")}
                </Button>
                <Button onClick={submit} autoFocus
                    radius="pill"
                    variant="contained"
                    color="negative"
                    sx={{ color: "#ffffff" }}>
                    {commonT("delete")}
                </Button>
            </DialogActions>
        </DialogForm>
    )
})