import { connect } from "react-redux"
import { Button, Divider, IconButton, InputAdornment, TextField } from "@mui/material"
import Visibility from "@mui/icons-material/Visibility"
import VisibilityOff from "@mui/icons-material/VisibilityOff"
import { useTranslation } from "react-multi-lang"
import { useState } from "react"

import { apiCall } from "../utils/api"
import { validatePassword } from "../utils/utils"

const mapDispatch = dispatch => ({
    updateSnackbarMsg: value =>
        dispatch({ type: "snackbarMsg/updateSnackbarMsg", payload: value })
})
export default connect(null, mapDispatch)(function ChangePassword(props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("account." + string, params),
        errorT = (string) => t("error." + string)

    const
        [curPassword, setCurPassword] = useState(""),
        [curPasswordError, setCurPasswordError] = useState(false),
        [newPassword, setNewPassword] = useState(""),
        [newPasswordError, setNewPasswordError] = useState(false),
        [otherError, setOtherError] = useState(""),
        [showPassword, setShowPassword] = useState(false)

    const submitDisabled = curPassword.length == 0 || newPassword.length == 0 || newPasswordError || curPasswordError

    const
        handleClickShowPassword = () => setShowPassword((show) => !show),
        handleMouseDownPassword = (event) => {
            event.preventDefault()
        },
        changePassword = (e) => {
            // setCurPasswordError(false)
            setCurPassword(e.target.value)
            setCurPasswordError(!validatePassword(e.target.value))
        },
        curPasswordLengthError = curPassword.length < 8 || curPassword.length > 50,
        changeNewPassword = (e) => {
            setNewPassword(e.target.value)
            setNewPasswordError(!validatePassword(e.target.value))
        },
        newPasswordLengthError = newPassword.length < 8 || newPassword.length > 50

    const submit = async () => {

        const onError = (err) => {
            switch (err) {
                case 20007:
                    setCurPasswordError(true)
                    props.updateSnackbarMsg({
                        type: "error",
                        msg: errorT("curPasswordNotMatch")
                    })
                    break
                case 60001:
                    setNewPasswordError(true)
                    props.updateSnackbarMsg({
                        type: "error",
                        msg: errorT("passwordUpdateError")
                    })
                    break
                default: setOtherError(err)
            }
        }
        await apiCall({
            data: {
                currentPassword: curPassword,
                newPassword: newPassword
            },
            method: "put",
            onSuccess: () => {
                props.updateSnackbarMsg({
                    type: "success",
                    msg: t("dialog.updatedPasswordMsg")

                })
                setCurPasswordError(false)
                setNewPasswordError(false)
                setCurPassword("")
                setNewPassword("")
            },
            onError,
            url: "/api/users/password"
        })
    }

    return <>
        <div className="card w-fit lg:w-88">
            <h4 className="mb-6">
                {pageT("changePassword")}
            </h4>
            <form className="grid">
                <TextField
                    id="cur-password"
                    size="medium"
                    label={pageT("currentPassword")}
                    type={showPassword ? "text" : "password"}
                    value={curPassword}
                    onChange={changePassword}
                    error={curPasswordError}
                    helperText={curPasswordLengthError ? errorT("passwordLength") : ""}
                    autoComplete="current-password"
                    InputProps={{
                        endAdornment:
                            <InputAdornment position="end">
                                <IconButton
                                    aria-label="toggle password visibility"
                                    onClick={handleClickShowPassword}
                                    onMouseDown={handleMouseDownPassword}
                                    edge="end"
                                >
                                    {showPassword
                                        ? <Visibility />
                                        : <VisibilityOff />
                                    }
                                </IconButton>
                            </InputAdornment>
                    }}
                />
                <TextField
                    id="new-password"
                    size="medium"
                    label={pageT("newPassword")}
                    type={showPassword ? "text" : "password"}
                    value={newPassword}
                    onChange={changeNewPassword}
                    error={newPasswordError}
                    helperText={newPasswordError ? errorT("passwordFormat") : ""
                        || newPasswordLengthError ? errorT("passwordLength") : ""
                    }
                    autoComplete="new-password"
                    InputProps={{
                        endAdornment:
                            <InputAdornment position="end">
                                <IconButton
                                    aria-label="toggle password visibility"
                                    onClick={handleClickShowPassword}
                                    onMouseDown={handleMouseDownPassword}
                                    edge="end"
                                >
                                    {showPassword
                                        ? <Visibility />
                                        : <VisibilityOff />
                                    }
                                </IconButton>
                            </InputAdornment>
                    }}
                />
            </form>
            <Divider variant="fullWidth" sx={{ marginTop: "0.5rem" }} />
            <div className="flex flex-row-reverse mt-6">
                <Button
                    onClick={submit}
                    disabled={submitDisabled}
                    radius="pill"
                    variant="contained"
                    color="primary">
                    {commonT("save")}
                </Button>
            </div>
        </div>
    </>
})