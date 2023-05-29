import { connect } from "react-redux"
import {
    Button, DialogActions, Divider, InputAdornment, IconButton,
    MenuItem, Switch, TextField
} from "@mui/material"
import Visibility from "@mui/icons-material/Visibility"
import VisibilityOff from "@mui/icons-material/VisibilityOff"
import { useTranslation } from "react-multi-lang"
import { useState } from "react"

import { apiCall } from "../utils/api"

import DialogForm from "../components/DialogForm"
import { ReactComponent as EditIcon } from "../assets/icons/edit.svg"

export default function EditUser(props) {
    const { row, groupDict } = props
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        errorT = string => t("error." + string),
        pageT = (string, params) => t("accountManagementUser." + string, params)

    const
        [openEdit, setOpenEdit] = useState(false),
        [unlock, setUnlock] = useState(true),
        [password, setPassword] = useState(""),
        [passwordError, setPasswordError] = useState(false),
        [showPassword, setShowPassword] = useState(false),
        [name, setName] = useState(row.name),
        [nameError, setNameError] = useState(false),
        [group, setGroup] = useState(row?.groupID),
        [groupError, setGroupError] = useState(null),
        [otherError, setOtherError] = useState("")

    const submitDisabled = !password.length || group == null || passwordError || nameError || groupError
    const
        handleClick = () => {
            setOpenEdit(true)
        },
        handleSwitch = () => {
            setUnlock(!unlock)
        },
        changePassword = (e) => {
            setPassword(e.target.value)
        },
        passwordLengthError = password.length == 0 || password.length < 8 || password.length > 50,
        validateCurPassword = () => setPasswordError(!validatePassword(password)),
        changeName = (e) => {
            const
                nameTarget = e.target.value,
                nameError = nameTarget.length == 0 || nameTarget.length > 20
            setName(nameTarget)
            setNameError(nameError)
        },
        changeGroup = (e) => {
            setGroup(e.target.value)
        }
    const
        handleClickShowPassword = () => setShowPassword((show) => !show),
        handleMouseDownPassword = (event) => {
            event.preventDefault()
        }
    const
        submit = async () => {

            const userID = row.id

            const data = {
                password: password,
                name: name,
                groupID: parseInt(group)
            }
            await apiCall({
                method: "put",
                data,
                onSuccess: () => {
                    setOpenEdit(false)
                    props.updateSnackbarMsg({
                        type: "success",
                        msg: t("dialog.addedSuccessfully")
                    })

                },
                onError: err => {
                    switch (err) {
                        case 60012:
                            // setAccountError(true)
                            props.updateSnackbarMsg({
                                type: "error",
                                msg: errorT("emailExist")
                            })
                            break
                        case 60013:
                            // setAccountError(true)
                            props.updateSnackbarMsg({
                                type: "error",
                                msg: errorT("failureToCreate")
                            })
                            break
                        default: setOtherError(err)
                    }
                },
                url: `/api/account-management/users/${userID}`
            })
        }
    return <>
        <EditIcon className="mr-4"
            onClick={handleClick} />
        <DialogForm
            dialogTitle={pageT("user")}
            fullWidth={true}
            maxWidth={"md"}
            open={openEdit}
            setOpen={setOpenEdit}
        >
            <Divider variant="middle" />
            <div className="flex flex-col m-auto mt-4 min-w-49 w-fit">
                <div className="mb-5 flex items-baseline">
                    <p className="ml-1 mr-2">{pageT("unlockUser")}</p>
                    <Switch
                        checked={unlock}
                        onChange={handleSwitch}
                    />
                </div>
                <h5 id="account"
                    className="ml-3 mb-8"
                    label={pageT("account")}>
                    {row?.username || ""}
                </h5>
                <TextField
                    id="edit-password"
                    type={showPassword ? "text" : "password"}
                    label={pageT("password")}
                    value={password}
                    onChange={changePassword}
                    error={passwordError}
                    helperText={passwordError ? errorT("passwordFormat") : ""
                        || passwordLengthError ? errorT("passwordLength") : ""}
                    autoComplete="password"
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
                    }
                    }
                />
                <TextField
                    id="edit-name"
                    label={pageT("name")}
                    onChange={changeName}
                    value={name || ""}
                />
                <TextField
                    id="edit-group"
                    select
                    label={commonT("group")}
                    onChange={changeGroup}
                    value={group}>
                    {Object.entries(groupDict).map(([key, value]) =>
                        <MenuItem key={"e-g-" + key} value={key}>
                            {value}
                        </MenuItem>)}
                </TextField>
                {otherError
                    ? <div className="box mb-8 negative text-center text-red-400">
                        {otherError}
                    </div>
                    : null}
            </div>
            <Divider variant="middle" />
            <DialogActions sx={{ margin: "1rem 0.5rem 1rem 0" }}>
                <Button onClick={() => { setOpenEdit(false) }}
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
                    {commonT("save")}
                </Button>
            </DialogActions>
        </DialogForm>
    </>
}