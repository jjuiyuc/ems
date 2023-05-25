import { connect } from "react-redux"
import { useTranslation } from "react-multi-lang"
import { useEffect, useMemo, useState } from "react"

import { apiCall } from "../utils/api"

import AddGroup from "../components/AddGroup"
import DeleteGroup from "../components/DeleteGroup"
import EditGroup from "../components/EditGroup"
import InfoGroup from "../components/InfoGroup"
import Table from "../components/DataTable"

const mapState = state => ({
    parentID: state.user.group.parentID
})

export default connect(mapState)(function AccountManagementGroup(props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        dialogT = (string) => t("dialog." + string),
        pageT = (string, params) => t("accountManagementGroup." + string, params)

    const
        [groupList, setGroupList] = useState([]),
        [groupTypeDict, setGroupTypeDict] = useState({}),
        [groupDictionary, setGroupDictionary] = useState({}),
        [loading, setLoading] = useState(false),
        [infoError, setInfoError] = useState(""),
        [fullWidth, setFullWidth] = useState(true),
        [maxWidth, setMaxWidth] = useState("sm")

    // const handleChange = (e) => {
    //     setTarget(r => ({ ...r, groupName: e.target.value }))
    // }
    const onSave = (row) => {
        const newData = groupList.map((value) =>
            value.id === row.id ? row : value
        )
        setGroupList(newData)
    }
    const adminID = groupList[0]?.id
    const columns = [
        {
            cell: row => <span className="font-mono">{row.name || ""}</span>,
            center: true,
            name: commonT("groupName"),
            selector: row => row.name
        },
        {
            cell: row =>
                <span className="font-mono">
                    {groupTypeDict[row.typeID] || ""}
                </span>,
            center: true,
            name: pageT("groupType"),
            selector: row => row.typeID
        },
        {
            cell: row =>
                <div className="flex w-28">
                    <InfoGroup
                        row={row}
                        groupTypeDict={groupTypeDict}
                        groupDictionary={groupDictionary} />
                    {/* Admin has no parentID */}
                    {row.parentID === null || props.parentID === adminID
                        ? null
                        : <>
                            <EditGroup
                                row={row}
                                groupList={groupList}
                                onSave={onSave}
                            />
                            <DeleteGroup
                                row={row}
                                getList={getList}
                            />
                        </>}
                </div>,
            center: true
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

                setGroupList(data.groups || [])
                setGroupTypeDict(data.groupTypes?.reduce((acc, cur) => {
                    acc[cur.id] = cur.name
                    return acc
                }, {}) || {})
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
    }, [])

    return <>
        <h1 className="mb-9">{commonT("accountManagementGroup")}</h1>
        <div className="mb-9">
            <AddGroup {...{ getList, groupList, groupTypes: groupTypeDict }} />
        </div>
        <Table
            columns={columns}
            data={groupList}
            paginationComponentOptions={{
                rowsPerPageText: t("dataTable.rowsPerPage")
            }}
            paginationPerPage={100}
            progressPending={loading}
            theme="dark"
        />
    </>
})