import { connect } from "react-redux"
import {
    Button, DialogActions, Divider, FormControl,
    InputLabel, InputAdornment, IconButton, MenuItem,
    OutlinedInput, TextField
} from "@mui/material"
import Visibility from "@mui/icons-material/Visibility"
import VisibilityOff from "@mui/icons-material/VisibilityOff"

import { useTranslation } from "react-multi-lang"
import { useEffect, useMemo, useState } from "react"

import { apiCall } from "../utils/api"
import { validateEmail } from "../utils/utils"

import AddUser from "../components/AddUser"
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
            }
        ]),
        [userList, setUserList] = useState([]),
        [groupDict, setGroupDict] = useState({}),
        [loading, setLoading] = useState(false),
        [infoError, setInfoError] = useState(""),
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
        [otherError, setOtherError] = useState(""),
        [fullWidth, setFullWidth] = useState(true),
        [maxWidth, setMaxWidth] = useState("sm"),
        [target, setTarget] = useState({})

    const

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
    const columns = [
        {
            cell: row => <span className="font-mono">{row.username}</span>,
            center: true,
            name: pageT("account"),
            selector: row => row.username,
            grow: 1

        },
        {
            cell: row => <span className="font-mono">{row.name}</span>,
            center: true,
            name: pageT("name"),
            selector: row => row.name,
            grow: 0.3
        },
        {
            cell: row => <span className="font-mono">
                {`${row.groupName + " " + `(${row.groupParentName})`}`}
            </span>,
            center: true,
            name: commonT("group"),
            selector: row => `${row.groupName + row.groupParentName}`,
            grow: 1

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
            grow: 0.5
        }
    ]
    const getList = () => {
        apiCall({
            onComplete: () => setLoading(false),
            onError: error => setInfoError(error),
            onStart: () => setLoading(true),
            onSuccess: rawData => {
                if (!rawData?.data) return

                const { data } = rawData

                setUserList(data.users || [])
                setGroupDict(data.users?.reduce((acc, cur) => {
                    acc[cur.groupID] = cur.groupName
                    return acc
                }, {}) || {})
            },
            url: `/api/account-management/users`
        })
    }
    useEffect(() => {
        getList()
    }, [])

    return <>
        <h1 className="mb-9">{commonT("accountManagementUser")}</h1>
        <div className="mb-9">
            <AddUser {...{ getList, userList, groupDict }} />
        </div>
        <Table
            {...{ columns, data: userList }}
            // customStyles={{
            //     headRow: {
            //         style: {
            //             backgroundColor: "#12c9c990",
            //             fontWeight: 600,
            //             fontSize: "16px",
            //             borderRadius: ".45rem .45rem 0 0 "
            //         }
            //     }
            // }}
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
            fullWidth={fullWidth}
            maxWidth={maxWidth}
            open={openEdit}
            setOpen={setOpenEdit}
        >
            <Divider variant="middle" />
            <div className="flex flex-col m-auto mt-4 min-w-49 w-fit">
                <h5 id="account"
                    className="ml-3 mb-8"
                    label={pageT("account")}>
                    {target?.username || ""}
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
            </div>
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
            fullWidth={fullWidth}
            maxWidth={maxWidth}
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