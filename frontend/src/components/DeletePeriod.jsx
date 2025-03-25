import { connect } from "react-redux"
import { Button, DialogActions } from "@mui/material"
import { useTranslation } from "react-multi-lang"
import moment from "moment"

import { apiCall } from "../utils/api"

import DialogForm from "../components/DialogForm"

const mapState = state => ({
    gatewayID: state.gateways.active.gatewayID
})
const mapDispatch = dispatch => ({
    updateSnackbarMsg: value =>
        dispatch({ type: "snackbarMsg/updateSnackbarMsg", payload: value }),

})
export default connect(mapState, mapDispatch)(function DeletePeriod(props) {
    const { row, getList, openDelete, setOpenDelete } = props
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        errorT = string => t("error." + string),
        pageT = (string, params) => t("settings." + string, params)

    const
        submit = async () => {

            const
                gatewayID = props.gatewayID,
                periodId = row.id

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
                        case 60022:
                            props.updateSnackbarMsg({
                                type: "error",
                                msg: errorT("fieldDisabled")
                            })
                            break
                        case 60039:
                            props.updateSnackbarMsg({
                                type: "error",
                                msg: errorT("failureToDelete")
                            })
                            break
                        default:
                            props.updateSnackbarMsg({
                                type: "error",
                                msg: errorT("failureToDelete")
                            })
                    }
                },
                url: `/api/device-management/gateways/${gatewayID}/power-outage-periods/${periodId}`
            })
        }
    return <>
        <DialogForm
            dialogTitle={t("dialog.deleteMsg")}
            fullWidth={true}
            maxWidth="sm"
            open={openDelete}
            setOpen={setOpenDelete}>
            <div
                className="grid grid-cols-settings-col4 gap-x-4 items-center
              border rounded border-solid border-gray-400 mx-3.5">
                <p className="ml-6">{moment(row?.startTime).format("YYYY/MM/DD HH:mm")}</p>
                <span>{pageT("to")}</span>
                <p>{moment(row?.endTime).format("YYYY/MM/DD HH:mm")}</p>
                <p>{pageT(`${row?.type}`)}</p>
            </div>
            <DialogActions sx={{ margin: "1rem 0.5rem 1rem 0" }}>
                <Button onClick={() => { setOpenDelete(false) }}
                    radius="pill"
                    variant="outlined"
                    color="gray">
                    {commonT("cancel")}
                </Button>
                <Button
                    onClick={submit}
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