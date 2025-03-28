import { connect } from "react-redux"
import { useTranslation } from "react-multi-lang"
import { useEffect, useMemo, useState } from "react"

import { apiCall } from "../utils/api"

import AddGroup from "../components/AddGroup"
import DeleteGroup from "../components/DeleteGroup"
import EditGroup from "../components/EditGroup"
import InfoGroup from "../components/InfoGroup"
import Table from "../components/DataTable"
import { ReactComponent as DeleteIcon } from "../assets/icons/trash_solid.svg"

const mapState = state => ({
    parentID: state?.user?.group?.parentID
})
const mapDispatch = dispatch => ({
    updateSnackbarMsg: value =>
        dispatch({ type: "snackbarMsg/updateSnackbarMsg", payload: value }),

})

const mockData = {
    "groups": [
        {
            "id": 1,
            "name": "Admin",
            "typeID": 1,
            "parentID": null
        },
        {
            "id": 2,
            "name": "AreaOwner_TW",
            "typeID": 2,
            "parentID": 1
        },
        {
            "id": 3,
            "name": "AreaMaintainer_TW",
            "typeID": 3,
            "parentID": 2
        },
        {
            "id": 4,
            "name": "PLACE 0",
            "typeID": 4,
            "parentID": 2
        },
        {
            "id": 5,
            "name": "PLACE 1",
            "typeID": 4,
            "parentID": 2
        },
        {
            "id": 6,
            "name": "PLACE 2",
            "typeID": 4,
            "parentID": 2
        },
    ],
    "groupTypes": [
        {
            "id": 1,
            "name": "Admin"
        },
        {
            "id": 2,
            "name": "Area owner"
        },
        {
            "id": 3,
            "name": "Area maintainer"
        },
        {
            "id": 4,
            "name": "Field owner"
        }
    ]
}
export default connect(mapState, mapDispatch)(function AccountManagementGroup(props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("accountManagementGroup." + string, params)

    const
        [groupList, setGroupList] = useState([]),
        [groupTypeDict, setGroupTypeDict] = useState({}),
        [groupDictionary, setGroupDictionary] = useState({}),
        [row, setRow] = useState(null),
        [openDelete, setOpenDelete] = useState(false),
        [loading, setLoading] = useState(false)

    const onSave = (row) => {
        const newData = groupList.map((value) =>
            value.id === row.id ? row : value
        )
        setGroupList(newData)
    }
    const handleClickDelete = row => {
        setOpenDelete(true)
        setRow(row)
    }

    const adminID = groupList[0]?.parentID
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
                    <InfoGroup {...{ row, groupTypeDict, groupDictionary }} />
                    {/* Admin has no parentID */}
                    {row.parentID === null || row.parentID === adminID
                        ? null
                        : <>
                            <EditGroup {...{ row, groupList, onSave }} />
                            <DeleteIcon onClick={() => handleClickDelete(row)} />

                        </>}
                </div>,
            center: true
        }
    ]
    const MOCK_MODE = true

    const getList = () => {
        if (MOCK_MODE) {
            const { groups, groupTypes } = mockData
            setGroupList(groups || [])
            setGroupTypeDict(
                groupTypes?.reduce((acc, cur) => {
                    acc[cur.id] = cur.name
                    return acc
                }, {}) || {}
            )
            setGroupDictionary(
                groups?.reduce((acc, cur) => {
                    acc[cur.id] = cur.name
                    return acc
                }, {}) || {}
            )
        } else {
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
                                msg: errorT("failureToGenerate")
                            })
                    }
                },
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
    }

    useEffect(() => {
        getList()
    }, [])
    return <>
        <h1 className="mb-9">{commonT("accountManagementGroup")}</h1>
        <div className="mb-9">
            <AddGroup {...{ getList, groupList, groupTypes: groupTypeDict, setGroupList }} />
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
        <DeleteGroup {...{
            row,
            openDelete,
            setOpenDelete,
            groupList,
            setGroupList
        }} />
    </>
})