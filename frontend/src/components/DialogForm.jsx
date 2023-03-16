import {
    Button, Dialog, DialogTitle, DialogActions, FormControl,
    InputLabel, InputAdornment, IconButton, ListItem, MenuItem,
    OutlinedInput, TextField
} from "@mui/material"
import AddIcon from "@mui/icons-material/Add"
import Visibility from "@mui/icons-material/Visibility"
import VisibilityOff from "@mui/icons-material/VisibilityOff"

import { useState } from "react"
import { useTranslation } from "react-multi-lang"

import { ReactComponent as NoticeIcon } from "../assets/icons/notice.svg"

export default function DialogForm({
    addType = "",
    type = "",
    triggerName = "",
    dialogTitle = "",
    groupName = "",
    groupType = "",
    parentGroup = "",
    account = "",
    password = "",
    name = "",
    group = "",
    fieldList = "",
    leftButtonName = "",
    rightButtonName = "",
    okayButton = "",
    closeOutside = false
}) {
    const t = useTranslation(),
        dialog = (string) => t("dialog." + string)

    const
        [open, setOpen] = useState(false),
        [fullWidth, setFullWidth] = useState(true),
        [maxWidth, setMaxWidth] = useState("sm"),
        [showPassword, setShowPassword] = useState(false)

    const
        handleClickOpen = () => {
            setOpen(true)
        },
        handleClose = () => {
            setOpen(false)
        }
    const
        handleClickShowPassword = () => setShowPassword((show) => !show),
        handleMouseDownPassword = (event) => {
            event.preventDefault();
        }
    const typeGroup = [
        {
            value: "Area Maintainer",
            label: "Area Maintainer",
        },
        {
            value: "Field Owner",
            label: "Field Owner",
        },
    ],
        parentGroupType = [
            {
                value: "AreaOwner_TW",
                label: "AreaOwner_TW"
            }
        ]
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
    return <>
        {type === "addGroup"
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
                            mt: 2,
                            minWidth: 120
                        }}>
                            <TextField
                                id="add-name"
                                label={groupName}
                                focused
                                required
                            />
                            <TextField
                                id="add-type"
                                select
                                label={groupType}
                                required
                            >
                                {typeGroup.map((option) => (
                                    <MenuItem key={option.value} value={option.value}>
                                        {option.label}
                                    </MenuItem>
                                ))}
                            </TextField>
                            <TextField
                                id="add-parent-group-type"
                                select
                                label={parentGroup}
                                required
                            >
                                {parentGroupType.map((option) => (
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
                                label={account}
                                focused
                                required />
                            <FormControl sx={{ mb: "2rem", minWidth: 120 }} variant="outlined">
                                <InputLabel htmlFor="outlined-adornment-password">{password}</InputLabel>
                                <OutlinedInput
                                    id="outlined-adornment-password"
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
                                    label={password}
                                    autoComplete="current-password"
                                    required
                                />
                            </FormControl>
                            <TextField
                                id="add-name"
                                label={name}
                                required
                            />
                            <TextField
                                id="add-group"
                                select
                                label={group}
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
                                <h5 className="ml-6 mt-2">{groupName} :</h5>
                                <ListItem
                                    id="name"
                                    label={groupName}
                                    required>
                                    Serenegray
                                </ListItem>
                                <h5 className="ml-6 mt-2">{groupType} :</h5>
                                <ListItem
                                    id="group-type"
                                    label={groupType}
                                >
                                    Field owner
                                </ListItem>
                                <h5 className="ml-6 mt-2">{parentGroup} :</h5>
                                <ListItem
                                    id="parent-group-type"
                                    label={parentGroup}
                                >
                                    AreaOwner
                                    _TW
                                </ListItem>
                                <h5 className="ml-6 mt-2">{fieldList} :</h5>
                                <ListItem
                                    id="field-list"
                                    label={fieldList}
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
    </>
}