import { connect } from "react-redux"
import LockIcon from "@mui/icons-material/Lock"
import { useTranslation } from "react-multi-lang"
import { useEffect, useState } from "react"

import { apiCall } from "../utils/api"

import AddUser from "../components/AddUser"
import DeleteUser from "../components/DeleteUser"
import EditUser from "../components/EditUser"
import Table from "../components/DataTable"
import { ReactComponent as DeleteIcon } from "../assets/icons/trash_solid.svg"

const mapState = state => ({
    username: state.user.username
})
const mapDispatch = dispatch => ({
    updateSnackbarMsg: value =>
        dispatch({ type: "snackbarMsg/updateSnackbarMsg", payload: value }),

})
export default connect(mapState, mapDispatch)(function AccountManagementUser(props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("accountManagementUser." + string, params)

    const
        [userList, setUserList] = useState([]),
        [groupDictionary, setGroupDictionary] = useState({}),
        [row, setRow] = useState(null),
        [openDelete, setOpenDelete] = useState(false),
        [loading, setLoading] = useState(false)

    const onSave = (row) => {
        const newData = userList.map((value) =>
            value.id === row.id ? row : value
        )
        setUserList(newData)
    }
    const handleClickDelete = row => {
        setOpenDelete(true)
        setRow(row)
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
                {`${row.groupName} (${row.groupParentName})`}
            </span>,
            center: true,
            name: commonT("group"),
            selector: row => row.groupName + row.groupParentName,
            grow: 1.1

        },
        {
            cell: (row) => <div className="flex w-24">
                {row.username === props.username
                    ? <div className="bg-gray-600 w-6 h-6"></div>
                    : <>
                        <EditUser {...{ row, groupDictionary, onSave, getList }} />
                        <DeleteIcon onClick={() => handleClickDelete(row)} />
                    </>
                }
                {row.lockedAt === null
                    ? <div className="ml-4 bg-gray-600 w-6 h-6"></div>
                    : <LockIcon className="ml-4" />
                }
            </div>,
            center: true,
            grow: 0.4
        }
    ]
    const getList = () => {
        apiCall({
            onComplete: () => setLoading(false),
            onStart: () => setLoading(true),
            onError: (err) => {
                switch (err) {
                    case 60011:
                        props.updateSnackbarMsg({
                            type: "error",
                            msg: errorT("failureToGenerate")
                        })
                        break
                    default:
                        props.updateSnackbarMsg({
                            type: "error",
                            msg: errorT("noDataMsg")
                        })
                }
            },
            onSuccess: rawData => {
                if (!rawData?.data) return

                const { data } = rawData

                setUserList(data.users || [])
            },
            url: `/api/account-management/users`
        })
    }
    const getGroupList = () => {
        apiCall({
            onComplete: () => setLoading(false),
            onStart: () => setLoading(true),
            onError: (err) => {
                switch (err) {
                    case 60002:
                        props.updateSnackbarMsg({
                            type: "error",
                            msg: errorT("failureToGenerate")
                        })
                        break
                    default:
                        props.updateSnackbarMsg({
                            type: "error",
                            msg: errorT("noDataMsg")
                        })
                }
            },
            onSuccess: rawData => {
                if (!rawData?.data) return

                const { data } = rawData
                setGroupDictionary(data.groups?.reduce((acc, cur) => {
                    acc[cur.id] = cur.name
                    return acc
                }, {}) || {})
            },
            url: `/api/account-management/groups`
        })
    }
    useEffect(() => {
        getList()
        getGroupList()
    }, [])

    return <>
        <h1 className="mb-9">{commonT("accountManagementUser")}</h1>
        <div className="mb-9">
            <AddUser {...{ getList, userList, groupDictionary }} />
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
        <DeleteUser {...{ row, getList, openDelete, setOpenDelete }} />
    </>
})