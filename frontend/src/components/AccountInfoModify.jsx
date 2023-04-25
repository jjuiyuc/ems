import { connect } from "react-redux"
import { Button, Divider, Snackbar, TextField } from "@mui/material"
import CheckCircleIcon from "@mui/icons-material/CheckCircle"
import CloseIcon from "@mui/icons-material/Close"
import { useTranslation } from "react-multi-lang"
import { useEffect, useState } from "react"

import { apiCall } from "../utils/api"

const mapState = state => ({
    name: state.user.name,
    username: state.user.username
})
const mapDispatch = dispatch => ({
    updateSnackbarMsg: value =>
        dispatch({ type: "snackbarMsg/updateSnackbarMsg", payload: value }),
    updateUserProfile: value =>
        dispatch({ type: "user/updateUserProfile", payload: value })

})
export default connect(mapState, mapDispatch)(function AccountInfoModify(props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        errorT = string => t("error." + string),
        pageT = (string, params) => t("account." + string)

    const [name, setName] = useState(props.name || "")
    const
        changeName = (e) => {
            setName(e.target.value)
        },
        nameError = name.length == 0 || name.length > 20,
        submit = () => {
            const data = { name: name }

            apiCall({
                method: "put",
                data,
                onSuccess: () => {
                    props.updateUserProfile(data)
                    props.updateSnackbarMsg({
                        type: "success",
                        msg: t("dialog.modifyNameMsg")
                    })
                },
                onError: () => {
                    props.updateSnackbarMsg({
                        type: "error",
                        msg: errorT("failureToSave")
                    })
                },
                url: "/api/users/name"
            })
        }
    return <>
        <div className="card w-fit lg:w-88">
            <h4 className="mb-6">
                {pageT("modifyAccountInformation")}
            </h4>
            <form className="grid">
                <TextField
                    label={commonT("account")}
                    defaultValue={props.username}
                    variant="outlined"
                    disabled
                />
                <TextField
                    error={nameError}
                    label={pageT("name")}
                    value={name}
                    onChange={changeName}
                    helperText={errorT("nameLength")}
                    required
                    variant="outlined"
                />
            </form>
            <Divider variant="fullWidth" sx={{ marginTop: "0.5rem" }} />
            <div className="flex flex-row-reverse mt-6">
                <Button
                    onClick={submit}
                    disabled={name.length == 0}
                    radius="pill"
                    variant="contained"
                    color="primary">
                    {commonT("save")}
                </Button>
            </div>
        </div>
    </>
})