import {
    Button, Dialog, DialogTitle, DialogActions, Divider, FormControl,
    InputLabel, InputAdornment, IconButton, ListItem, MenuItem,
    OutlinedInput, TextField
} from "@mui/material"
import AddIcon from "@mui/icons-material/Add"
import Visibility from "@mui/icons-material/Visibility"
import VisibilityOff from "@mui/icons-material/VisibilityOff"

import { useState } from "react"
import { useTranslation } from "react-multi-lang"

import { ValidateEmail } from "../utils/utils"

import { ReactComponent as EditIcon } from "../assets/icons/edit.svg"
import { ReactComponent as DeleteIcon } from "../assets/icons/trash_solid.svg"

export default function DialogForm({
    type = "",
    children = null,
    data = {},
    triggerName = "",
    dialogTitle = "",
    leftButtonName = "",
    rightButtonName = "",
    okayButton = "",
    open,
    setOpen,
    closeOutside = false
}) {
    const t = useTranslation(),
        commonT = string => t("common." + string),
        dialogT = (string) => t("dialog." + string),
        errorT = (string) => t("error." + string)

    const
        [fullWidth, setFullWidth] = useState(true),
        [maxWidth, setMaxWidth] = useState("sm"),
        [account, setAccount] = useState(""),
        [accountError, setAccountError] = useState(null),
        [password, setPassword] = useState(""),
        [passwordError, setPasswordError] = useState(false),
        [showPassword, setShowPassword] = useState(false),
        [name, setName] = useState(""),
        [nameError, setNameError] = useState(null),
        [group, setGroup] = useState(""),
        [groupError, setGroupError] = useState(null),
        [fieldList, setFieldList] = useState(""),
        [fieldListError, setFieldListError] = useState(null),
        [otherError, setOtherError] = useState(""),
        [defaultValue, setDefaultValue] = useState("KKK")

    const
        handleClickOpen = () => {
            setOpen(true)
        },
        handleClose = () => {
            setOpenAddGroup(false)
        },
        handleClickShowPassword = () => setShowPassword((show) => !show),
        handleMouseDownPassword = (event) => {
            event.preventDefault()
        },
        handleChange = (e) => {
            setDefaultValue(e.target.value)
        }
    const
        changeAccount = (e) => {
            setAccount(e.target.value)
            setAccountError(null)
            setOtherError("")
        },
        changePassword = (e) => {
            setPassword(e.target.value)
            setPasswordError(false)
            setOtherError("")
        },
        submit = async () => {
            const isEmail = ValidateEmail(email)

            if (!isEmail) {
                setEmailError({ type: "emailFormat" })
                return
            }

            const onError = (err) => {
                switch (err) {
                    case 20004:
                        setEmailError({ type: "emailNotExist" })
                        break
                    case 20006:
                        setEmailError({ type: "userLocked" })
                        break
                    case 20007:
                        setPasswordError(true)
                        break
                    default: setOtherError(err)
                }
            }
        }

    const groupData = [
        {
            value: "AreaOwner_TW",
            label: "AreaOwner_TW",
        },
        {
            value: "Area Maintainer_TW",
            label: "Area Maintainer_TW",
        },
        {
            value: "Serenegray",
            label: "Serenegray",
        },
        {
            value: "Cht_Miaoli",
            label: "Cht_Miaoli",
        }
    ]
    const triggerButtons = {
        addUser: <>
            <div>
                <Button
                    onClick={handleClickOpen}
                    key={"ac-b-"}
                    size="x-large"
                    variant="outlined"
                    radius="pill"
                    fontSize="large"
                    color="brand"
                    startIcon={<AddIcon />}>
                    {triggerName}
                </Button>
                <Dialog
                    fullWidth={fullWidth}
                    maxWidth={maxWidth}
                    open={open}
                    onClose={handleClose}
                >
                    <DialogTitle id="form-dialog-title">
                        {dialogTitle}
                    </DialogTitle>
                    <Divider variant="middle" />
                    <FormControl sx={{
                        display: "flex",
                        flexDirection: "column",
                        margin: "auto",
                        width: "fit-content",
                        mt: "1rem",
                        minWidth: 120
                    }}>
                        <TextField
                            id="add-account"
                            label={dialogT("account")}
                            onChange={changeAccount}
                            error={accountError !== null}
                            helperText={accountError ? errorT(accountError.type) : ""}
                            type="email"
                            focused />
                        <FormControl sx={{ mb: "2rem", minWidth: 120 }} variant="outlined">
                            <InputLabel htmlFor="outlined-adornment-password">{dialogT("password")}</InputLabel>
                            <OutlinedInput
                                id="add-password"
                                type={showPassword ? "text" : "password"}
                                endAdornment={
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
                                label={dialogT("password")}
                                autoComplete="current-password"
                            />
                        </FormControl>
                        <TextField
                            id="add-name"
                            label={dialogT("name")}>

                        </TextField>
                        <TextField
                            id="add-group"
                            select
                            onChange={handleChange}
                            label={commonT("group")}
                        >
                            {groupData.map((option) => (
                                <MenuItem key={option.value} value={option.value}>
                                    {option.label}
                                </MenuItem>
                            ))}
                        </TextField>
                    </FormControl>
                    <Divider variant="middle" />
                    <DialogActions sx={{ margin: "1rem 0.5rem 1rem 0" }}>
                        <Button onClick={handleClose}
                            radius="pill"
                            variant="outlined"
                            color="gray">
                            {leftButtonName}
                        </Button>
                        <Button onClick={handleClose}
                            radius="pill"
                            variant="contained"
                            color="primary">
                            {rightButtonName}
                        </Button>
                    </DialogActions>
                </Dialog>
            </div>
        </>,
        editUser: <>
            <div>
                <EditIcon className="mr-5" onClick={handleClickOpen} />
                <Dialog
                    fullWidth={fullWidth}
                    maxWidth={maxWidth}
                    open={open}
                    onClose={handleClose}
                >
                    <DialogTitle id="form-dialog-title">
                        {dialogTitle}
                    </DialogTitle>
                    <FormControl sx={{
                        display: "flex",
                        flexDirection: "column",
                        margin: "auto",
                        width: "fit-content",
                        mt: "1rem",
                        minWidth: 120
                    }}>
                        <TextField
                            id="edit-account"
                            label={dialogT("account")}
                            onChange={changeAccount}
                            error={accountError !== null}
                            helperText={accountError ? errorT(accountError.type) : ""}
                            type="email"
                            focused />
                        <FormControl sx={{ mb: "2rem", minWidth: 120 }} variant="outlined">
                            <InputLabel htmlFor="outlined-adornment-password">
                                {dialogT("password")}
                            </InputLabel>
                            <OutlinedInput
                                id="edit-password"
                                type={showPassword ? "text" : "password"}
                                endAdornment={
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
                                label={dialogT("password")}
                                onChange={handleChange}
                                autoComplete="current-password"
                            />
                        </FormControl>
                        <TextField
                            id="edit-name"
                            label={dialogT("name")}
                        />
                        <TextField
                            id="edit-group"
                            select
                            label={commonT("group")}
                            onChange={handleChange}
                        >
                            {groupData.map((option) => (
                                <MenuItem key={option.value} value={option.value}>
                                    {option.label}
                                </MenuItem>
                            ))}
                        </TextField>
                    </FormControl>
                    <DialogActions sx={{ margin: "1rem 0.5rem 1rem 0" }}>
                        <Button onClick={handleClose}
                            radius="pill"
                            variant="outlined"
                            color="gray">
                            {leftButtonName}
                        </Button>
                        <Button onClick={handleClose}
                            radius="pill"
                            variant="contained"
                            color="primary">
                            {rightButtonName}
                        </Button>
                    </DialogActions>
                </Dialog>
            </div>
        </>
    }
    return <>
        <div>
            <Dialog
                fullWidth={fullWidth}
                maxWidth={maxWidth}
                open={open}
                onClose={closeOutside ? handleClose : () => { }}
            >
                <DialogTitle id="form-dialog-title">
                    {dialogTitle}
                </DialogTitle>
                {children}
            </Dialog>
        </div>
    </>
}