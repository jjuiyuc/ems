import { connect } from "react-redux"
import {
    Button, DialogActions, Divider, FormControl, TextField
} from "@mui/material"
import LockOpenIcon from "@mui/icons-material/LockOpen"

import { useTranslation } from "react-multi-lang"
import { useEffect, useMemo, useState } from "react"

import { apiCall } from "../utils/api"

import AddUser from "../components/AddUser"
import EditUser from "../components/EditUser"
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
            grow: 1.1

        },
        {
            cell: row => <span className="font-mono">{row.name}</span>,
            center: true,
            name: pageT("name"),
            selector: row => row.name,
            grow: 0.8
        },
        {
            cell: row => <span className="font-mono">
                {`${row.groupName + " " + `(${row.groupParentName})`}`}
            </span>,
            center: true,
            name: commonT("group"),
            selector: row => `${row.groupName + row.groupParentName}`,
            grow: 1.1

        },
        {
            cell: (row, index) => <div className="flex w-24">
                <EditUser className="mr-4"
                    {...{ row, groupDict }}
                />
                {row.group === "Area Owner_TW"
                    ? null
                    : <DeleteIcon onClick={() => {
                        setOpenDelete(true)
                        setTarget(row)
                    }} />
                }
                <LockOpenIcon className="ml-4" />
            </div>,
            center: true,
            grow: 0.4
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
            pagination={true}
            paginationComponentOptions={{
                rowsPerPageText: t("dataTable.rowsPerPage")
            }}
            paginationPerPage={100}
            progressPending={loading}
            theme="dark"
        />

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