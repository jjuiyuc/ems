import {
    Button, Dialog, DialogTitle, DialogActions, FormControl,
    InputLabel, InputAdornment, IconButton, ListItem, MenuItem,
    OutlinedInput, TextField
} from "@mui/material"
import Visibility from "@mui/icons-material/Visibility"
import VisibilityOff from "@mui/icons-material/VisibilityOff"

import { useState } from "react"
import { useTranslation } from "react-multi-lang"

import { ValidateEmail } from "../utils/utils"

import { ReactComponent as EditIcon } from "../assets/icons/edit.svg"

export default function Account({
    type = "",
    triggerName = "",
    dialogTitle = "",
    leftButtonName = "",
    rightButtonName = "",
    okayButton = "",
    closeOutside = false
}) {
    const t = useTranslation(),
        commonT = string => t("common." + string),
        dialogT = (string) => t("dialog." + string),
        pageT = (string, params) => t("account." + string, params),
        errorT = (string) => t("error." + string),

    const
        [open, setOpen] = useState(false),
        [fullWidth, setFullWidth] = useState(true),
        [maxWidth, setMaxWidth] = useState("sm"),

        [account, setAccount] = useState(""),
        [accountError, setAccountError] = useState(null),
        [password, setPassword] = useState(""),
        [passwordError, setPasswordError] = useState(false),
        [showPassword, setShowPassword] = useState(false),
        [name, setName] = useState(""),
        [nameError, setNameError] = useState(null),
        [otherError, setOtherError] = useState(""),
        [defaultValue, setDefaultValue] = useState("KKK")

    const
        handleClickOpen = () => {
            setOpen(true)
        },
        handleClose = () => {
            setOpen(false)
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

            // const onError = (err) => {
            //     switch (err) {
            //         case 20004:
            //             setEmailError({ type: "emailNotExist" })
            //             break
            //         case 20006:
            //             setEmailError({ type: "userLocked" })
            //             break
            //         case 20007:
            //             setPasswordError(true)
            //             break
            //         default: setOtherError(err)
            //     }
            // }
        }

    return <>
        {type === "addUser"
            ? <>
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
            : null}
        {type === "notice"
            ? <>
                <div>
                    <NoticeIcon
                        className="mr-5"
                        onClick={handleClickOpen} />
                    <Dialog
                        fullWidth={fullWidth}
                        maxWidth={maxWidth}
                        open={open}
                        onClose={handleClose}
                    >
                        <DialogTitle id="form-dialog-title">
                            {dialogTitle}
                        </DialogTitle>

                        <div className="flex flex-col m-auto mt-4 min-w-49.5 w-fit">
                            <div className="grid grid-cols-1fr-auto">
                                <h5 className="ml-6 mt-2">{dialogT("groupName")} :</h5>
                                <ListItem
                                    id="name"
                                    label={dialogT("groupName")}>
                                    Serenegray
                                </ListItem>
                                <h5 className="ml-6 mt-2">{dialogT("groupType")} :</h5>
                                <ListItem
                                    id="group-type"
                                    label={dialogT("groupType")}
                                >
                                    Field owner
                                </ListItem>
                                <h5 className="ml-6 mt-2">{dialogT("parentGroup")} :</h5>
                                <ListItem
                                    id="parent-group-type"
                                    label={dialogT("parentGroup")}
                                >
                                    AreaOwner_TW
                                </ListItem>
                                <h5 className="ml-6 mt-2">{dialogT("fieldList")} :</h5>
                                <ListItem
                                    id="field-list"
                                    label={dialogT("fieldList")}
                                >
                                    Serenegray-0E0BA27A8175AF978C49396BDE9D7A1E
                                </ListItem>
                            </div>
                        </div>
                        <DialogActions sx={{ margin: "1rem 0.5rem 1rem 0" }}>
                            <Button onClick={handleClose}
                                radius="pill"
                                variant="contained"
                                color="primary">
                                {okayButton}
                            </Button>
                        </DialogActions>
                    </Dialog>
                </div>
            </>
            : null}
        {type === "editUser"
            ? <>
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
            : null}
    </>
}