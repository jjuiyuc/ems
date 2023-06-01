import { connect } from "react-redux"
import LockIcon from "@mui/icons-material/Lock"
import { useTranslation } from "react-multi-lang"
import { useEffect, useState } from "react"

import { apiCall } from "../utils/api"

import AddUser from "../components/AddUser"
import EditUser from "../components/EditUser"
import DeleteUser from "../components/DeleteUser"
import Table from "../components/DataTable"

const mapState = state => ({
    username: state.user.username
})

export default connect(mapState)(function AccountManagementUser(props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        dialogT = (string) => t("dialog." + string),
        pageT = (string, params) => t("accountManagementUser." + string, params)
    const
        [userList, setUserList] = useState([]),
        [groupDict, setGroupDict] = useState({}),
        [loading, setLoading] = useState(false),
        [infoError, setInfoError] = useState("")

    const onSave = (row) => {
        const newData = userList.map((value) =>
            value.id === row.id ? row : value
        )
        setUserList(newData)
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
            cell: (row) => <div className="flex w-24">
                <EditUser
                    {...{ row, groupDict, onSave, getList }}
                />
                {row.username === props.username
                    ? <div className="bg-gray-600 w-6 h-6"></div>
                    : <DeleteUser {...{ row, getList }} />
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
    </>
})