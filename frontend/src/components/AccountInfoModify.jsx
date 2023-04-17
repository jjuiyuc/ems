import { connect } from "react-redux"
import { Button, Divider, Snackbar, TextField } from "@mui/material"
import CheckCircleIcon from "@mui/icons-material/CheckCircle"
import CloseIcon from "@mui/icons-material/Close"
import { useTranslation } from "react-multi-lang"
import { useEffect, useState } from "react"

import { apiCall } from "../utils/api"

import FinishedBox from "../components/FinishedBox"

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
        dialogT = string => t("dialog." + string),
        errorT = string => t("error." + string),
        pageT = (string, params) => t("account." + string, params)

    const
        [account, setAccount] = useState(props.username),
        [name, setName] = useState(props.name),
        [newName, setNewName] = useState(""),
        // [nameError, setNameError] = useState(false),
        [loading, setLoading] = useState(false),
        [isReset, setIsReset] = useState(false)

    const
        changeName = (e) => {
            setName(e.target.value)
        },
        submit = () => {
            const data = { name: name }

            apiCall({
                method: "put",
                data,
                onSuccess: () => {
                    setIsReset(true)
                    props.updateUserProfile(data)
                    props.updateSnackbarMsg({
                        type: "success",
                        msg: dialogT("modifyNameMsg")
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
    const nameError = name.length == 0 || name.length > 20
    return <>
        <div className="card w-fit">
            <h4 className="mb-6">
                {pageT("ModifyAccountInformation")}
            </h4>
            <Divider variant="fullWidth" sx={{ marginBottom: "1.5rem" }} />
            <form className="grid ">
                <TextField
                    label={commonT("account")}
                    defaultValue={account}
                    variant="outlined"
                    required
                    disabled
                />
                <TextField
                    error={nameError}
                    label={pageT("name")}
                    value={name || ""}
                    onChange={changeName}
                    helperText={nameError ? errorT("nameLength") : ""}
                    required
                    variant="outlined"
                />
            </form>
            <Divider variant="fullWidth" sx={{ marginTop: "0.5rem" }} />
            <div className="flex flex-row-reverse mt-6">
                <Button
                    sx={{ marginLeft: "0.5rem" }}
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