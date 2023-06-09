import { Button, DialogActions, Divider, InputAdornment, IconButton, MenuItem, TextField } from "@mui/material"
import AddIcon from "@mui/icons-material/Add"
import Visibility from "@mui/icons-material/Visibility"
import VisibilityOff from "@mui/icons-material/VisibilityOff"
import { connect } from "react-redux"
import { useTranslation } from "react-multi-lang"
import { useState } from "react"

import { apiCall } from "../utils/api"
import { validateEmail } from "../utils/utils"
import { validatePassword } from "../utils/utils"

import DialogForm from "../components/DialogForm"
const mapDispatch = dispatch => ({
    updateSnackbarMsg: value =>
        dispatch({ type: "snackbarMsg/updateSnackbarMsg", payload: value }),

})
export default connect(null, mapDispatch)(function AddUser(props) {
    const { getList, groupDictionary } = props

    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        errorT = string => t("error." + string),
        pageT = (string, params) => t("accountManagementUser." + string, params)

    const
        [openAdd, setOpenAdd] = useState(false),
        [account, setAccount] = useState(""),
        [accountError, setAccountError] = useState(false),
        [password, setPassword] = useState(""),
        [passwordError, setPasswordError] = useState(false),
        [showPassword, setShowPassword] = useState(false),
        [name, setName] = useState(""),
        [nameError, setNameError] = useState(false),
        [group, setGroup] = useState(null),
        [otherError, setOtherError] = useState("")

    const submitDisabled = !password.length || group == null || accountError || passwordError || nameError
    const
        changeAccount = (e) => {
            setAccount(e.target.value)
            setAccountError(!validateEmail(e.target.value))
        },
        changePassword = (e) => {
            setPassword(e.target.value)
            setPasswordError(!validatePassword(e.target.value))
        },
        passwordLengthError = password.length == 0 || password.length < 8 || password.length > 50,
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

            const data = {
                username: account,
                password: password,
                name: name,
                groupID: parseInt(group)
            }
            await apiCall({
                method: "post",
                data,
                onSuccess: () => {
                    setOpenAdd(false)
                    getList()
                    props.updateSnackbarMsg({
                        type: "success",
                        msg: t("dialog.addedSuccessfully")
                    })
                    setAccount("")
                    setPassword("")
                    setName("")
                    setGroup("")
                },
                onError: err => {
                    switch (err) {
                        case 60012:
                            setAccountError(true)
                            props.updateSnackbarMsg({
                                type: "error",
                                msg: errorT("emailExist")
                            })
                            break
                        case 60013:
                            setAccountError(true)
                            props.updateSnackbarMsg({
                                type: "error",
                                msg: errorT("failureToCreate")
                            })
                            break
                        default: setOtherError(err)
                    }
                },
                url: "/api/account-management/users"
            })
        },
        cancelClick = () => {
            setOpenAdd(false)
            setAccount("")
            setPassword("")
            setName("")
            setGroup("")
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
            dialogTitle={pageT("addUser")}
            fullWidth={true}
            maxWidth="md"
            open={openAdd}
            setOpen={setOpenAdd}>
            <Divider variant="middle" />
            <div className="flex flex-col m-auto mt-4 min-w-49 w-fit">
                <TextField
                    id="add-account"
                    label={pageT("account")}
                    onChange={changeAccount}
                    value={account}
                    error={accountError}
                    helperText={accountError ? errorT("emailFormat") : ""}
                    type="email" />
                <TextField
                    id="add-password"
                    type={showPassword ? "text" : "password"}
                    label={pageT("password")}
                    value={password}
                    onChange={changePassword}
                    error={passwordError}
                    helperText={(passwordError ? errorT("passwordFormat") : "")
                        || (passwordLengthError ? errorT("passwordLength") : "")
                    }
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
                    }}
                />
                <TextField
                    id="add-name"
                    label={pageT("name")}
                    value={name || ""}
                    onChange={changeName}
                    error={nameError}
                    helperText={nameError ? errorT("nameLength") : ""}
                >
                </TextField>
                <TextField
                    id="add-group"
                    select
                    onChange={changeGroup}
                    label={commonT("group")}
                    defaultValue="">
                    {Object.entries(groupDictionary).map(([key, value]) =>
                        <MenuItem key={"a-g-" + key} value={key}>
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
                <Button onClick={cancelClick}
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