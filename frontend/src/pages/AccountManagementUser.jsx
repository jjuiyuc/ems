import {
    Button, DialogActions, Divider, FormControl,
    InputLabel, InputAdornment, IconButton, MenuItem,
    OutlinedInput, TextField
} from "@mui/material"
import AddIcon from "@mui/icons-material/Add"
import Visibility from "@mui/icons-material/Visibility"
import VisibilityOff from "@mui/icons-material/VisibilityOff"

import { useTranslation } from "react-multi-lang"
import { useEffect, useMemo, useState } from "react"

import { ValidateEmail } from "../utils/utils"

import DialogForm from "../components/DialogForm"
import Table from "../components/DataTable"

import { ReactComponent as DeleteIcon } from "../assets/icons/trash_solid.svg"
import { ReactComponent as EditIcon } from "../assets/icons/edit.svg"

export default function AccountManagementUser() {

    const groupData = [
        {
            value: "Area Owner_TW",
            label: "Area Owner_TW",
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
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        dialogT = (string) => t("dialog." + string),
        pageT = (string, params) => t("accountManagementUser." + string, params)
    const
        [data, setData] = useState([
            {
                id: 1,
                account: "XXXXX@ubiik.com",
                password: "xxxxxll",
                name: "XXXXX",
                group: "Area Owner_TW"
            },
            {
                id: 2,
                account: "YYYYY@ubiik.com",
                password: "x77xxxxll",
                name: "YYYYY",
                group: "Area Maintainer_TW"
            },
            {
                id: 3,
                account: "serenegray@ubiik.com",
                password: "9977xxxxll",
                name: "Serenegray",
                group: "Serenegray"
            },
            {
                id: 4,
                account: "cht_miaoli@ubiik.com",
                password: "kk977xxxxll",
                name: "Cht_Miaoli",
                group: "Cht_Miaoli"
            }
        ]),
        [openAdd, setOpenAdd] = useState(false),
        [openEdit, setOpenEdit] = useState(false),
        [openDelete, setOpenDelete] = useState(false),
        [account, setAccount] = useState(data?.account || ""),
        [accountError, setAccountError] = useState(null),
        [password, setPassword] = useState(data?.password || ""),
        [passwordError, setPasswordError] = useState(false),
        [showPassword, setShowPassword] = useState(false),
        [name, setName] = useState(data?.name || ""),
        [nameError, setNameError] = useState(null),
        [group, setGroup] = useState(data?.group || ""),
        [groupError, setGroupError] = useState(null),
        [target, setTarget] = useState({}),
        [error, setError] = useState(null),
        [otherError, setOtherError] = useState(""),
        [loading, setLoading] = useState(false)
    const
        changeAccount = (e) => {
            setAccount(e.target.value)
            setAccountError(null)
            setOtherError("")
        },
        changePassword = (e) => {
            setTarget(r => ({ ...r, password: e.target.value }))
            setPasswordError(false)
            setOtherError("")
        },
        changeName = (e) => {
            setTarget(r => ({ ...r, name: e.target.value }))
            setNameError(null)
            setOtherError("")
        },
        changeGroup = (e) => {
            setTarget(r => ({ ...r, group: e.target.value }))
            setGroupError(null)
            setOtherError("")
        }
    const
        handleClickShowPassword = () => setShowPassword((show) => !show),
        handleMouseDownPassword = (event) => {
            event.preventDefault()
        }
    const editSave = () => {
        setData(r => {
            const newData = [...r]
            newData[target.index].password = target.password
            newData[target.index].name = target.name
            newData[target.index].group = target.group
            return newData
        })
    }
    const submit = async () => {
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

    const columns = [
        {
            cell: row => <span className="font-mono">{row.account}</span>,
            center: true,
            name: pageT("account"),
            selector: row => row.account
        },
        {
            cell: row => <span className="font-mono">{row.name}</span>,
            center: true,
            name: pageT("name"),
            selector: row => row.name
        },
        {
            cell: row => <span className="font-mono">{row.group}</span>,
            center: true,
            name: commonT("group"),
            selector: row => row.group
        },
        {
            cell: (row, index) => <div className="flex w-28">
                <EditIcon className="mr-5"
                    onClick={() => {
                        setOpenEdit(true)
                        setTarget({ ...row, index })
                    }} />
                {row.group === "Area Owner_TW"
                    ? null
                    : <DeleteIcon onClick={() => {
                        setOpenDelete(true)
                        setTarget(row)
                    }} />
                }
            </div>,
            center: true,
        }
    ]
    return <>
        <h1 className="mb-9">{commonT("accountManagementUser")}</h1>
        <div className="mb-9">
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
                open={openAdd}
                setOpen={setOpenAdd}>
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
                        label={pageT("account")}
                        onChange={changeAccount}
                        value={account}
                        error={accountError !== null}
                        helperText={accountError ? errorT(accountError.type) : ""}
                        type="email"
                        focused />
                    <FormControl sx={{ mb: "2rem", minWidth: 120 }} variant="outlined">
                        <InputLabel htmlFor="outlined-adornment-password">
                            {pageT("password")}
                        </InputLabel>
                        <OutlinedInput
                            id="add-password"
                            type={showPassword ? "text" : "password"}
                            label={pageT("password")}
                            autoComplete="current-password"
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
                        />
                    </FormControl>
                    <TextField
                        id="add-name"
                        label={pageT("name")}
                        value={name}
                        onChange={changeName}>
                    </TextField>
                    <TextField
                        id="add-group"
                        select
                        onChange={changeGroup}
                        label={commonT("group")}
                        defaultValue="">
                        {groupData.map((option) => (
                            <MenuItem key={option.value} value={option.value}>
                                {option.label}
                            </MenuItem>
                        ))}
                    </TextField>
                    {otherError
                        ? <div className="box mb-8 negative text-center text-red-400">
                            {otherError}
                        </div>
                        : null}
                </FormControl>
                <Divider variant="middle" />
                <DialogActions sx={{ margin: "1rem 0.5rem 1rem 0" }}>
                    <Button onClick={() => { setOpenAdd(false) }}
                        radius="pill"
                        variant="outlined"
                        color="gray">
                        {commonT("cancel")}
                    </Button>
                    <Button onClick={() => { setOpenAdd(false) }}
                        radius="pill"
                        variant="contained"
                        color="primary">
                        {commonT("add")}
                    </Button>
                </DialogActions>
            </DialogForm>
        </div>
        <Table
            {...{ columns, data }}
            customStyles={{
                headRow: {
                    style: {
                        backgroundColor: "#12c9c990",
                        fontWeight: 600,
                        fontSize: "16px",
                        borderRadius: ".45rem .45rem 0 0 "
                    }
                }
            }}
            noDataComponent={t("dataTable.noDataMsg")}
            pagination={true}
            paginationComponentOptions={{
                rowsPerPageText: t("dataTable.rowsPerPage")
            }}
            paginationPerPage={100}
            progressPending={loading}
            theme="dark"
        />
        {/* edit */}
        <DialogForm
            dialogTitle={pageT("user")}
            open={openEdit}
            setOpen={setOpenEdit}
        >
            <Divider variant="middle" />
            <FormControl sx={{
                display: "flex",
                flexDirection: "column",
                margin: "auto",
                width: "fit-content",
                mt: "1rem",
                minWidth: 120
            }}>
                <h5 id="account"
                    className="ml-3 mb-8"
                    label={pageT("account")}>
                    {target?.account || ""}
                </h5>
                <FormControl sx={{ mb: "2rem", minWidth: 120 }} variant="outlined">
                    <InputLabel htmlFor="outlined-adornment-password">
                        {pageT("password")}
                    </InputLabel>
                    <OutlinedInput
                        id="edit-password"
                        type={showPassword ? "text" : "password"}
                        label={pageT("password")}
                        value={target?.password || ""}
                        onChange={changePassword}
                        autoComplete="current-password"
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
                    />
                </FormControl>
                <TextField
                    id="edit-name"
                    label={pageT("name")}
                    onChange={changeName}
                    value={target?.name || ""}
                />
                <TextField
                    id="edit-group"
                    select
                    label={commonT("group")}
                    onChange={changeGroup}
                    value={target?.group || ""}
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
                <Button onClick={() => { setOpenEdit(false) }}
                    radius="pill"
                    variant="outlined"
                    color="gray">
                    {commonT("cancel")}
                </Button>
                <Button onClick={() => {
                    setOpenEdit(false)
                    editSave()
                }}
                    radius="pill"
                    variant="contained"
                    color="primary">
                    {commonT("save")}
                </Button>
            </DialogActions>
        </DialogForm>
        {/* delete */}
        <DialogForm
            dialogTitle={dialogT("deleteMsg")}
            open={openDelete}
            setOpen={setOpenDelete}>
            <div className="flex">
                <h5 className="ml-6 mr-2">{pageT("account")} :</h5>
                {target?.account || ""}
            </div>
            <DialogActions sx={{ margin: "0.5rem 0.5rem 0.5rem 0" }}>
                <Button onClick={() => { setOpenDelete(false) }}
                    radius="pill"
                    variant="outlined"
                    color="gray">
                    {commonT("cancel")}
                </Button>
                <Button onClick={() => { setOpenDelete(false) }} autoFocus
                    radius="pill"
                    variant="contained"
                    color="negative"
                    sx={{ color: "#ffffff" }}>
                    {commonT("delete")}
                </Button>
            </DialogActions>
        </DialogForm>
    </>
}
// flex flex-col mt-4 min-w-49 w-fit min-w-xs