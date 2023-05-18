import {
    Button, DialogActions, Divider, FormControl, ListItem, TextField
} from "@mui/material"
import { useTranslation } from "react-multi-lang"
import { useEffect, useMemo, useState } from "react"

import { apiCall } from "../utils/api"

import AddGroup from "../components/AddGroup"
import DialogForm from "../components/DialogForm"
import Table from "../components/DataTable"

import { ReactComponent as DeleteIcon } from "../assets/icons/trash_solid.svg"
import { ReactComponent as EditIcon } from "../assets/icons/edit.svg"
import { ReactComponent as NoticeIcon } from "../assets/icons/notice.svg"


export default function AccountManagementGroup(props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        dialogT = (string) => t("dialog." + string),
        pageT = (string, params) => t("accountManagementGroup." + string, params)

    const
        [groupList, setGroupList] = useState([]),
        [groupTypes, setGroupTypes] = useState({}),
        [loading, setLoading] = useState(false),
        [infoError, setInfoError] = useState(""),
        [fullWidth, setFullWidth] = useState(true),
        [maxWidth, setMaxWidth] = useState("sm"),
        [openNotice, setOpenNotice] = useState(false),
        [openEdit, setOpenEdit] = useState(false),
        [openDelete, setOpenDelete] = useState(false),
        [target, setTarget] = useState({})

    const handleChange = (e) => {
        setTarget(r => ({ ...r, groupName: e.target.value }))
    }
    const editSave = () => {
        setData(r => {
            const newData = [...r]
            newData[target.index].groupName = target.groupName
            return newData
        })
    }
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
                    {groupTypes[row.typeID] || ""}
                </span>,
            center: true,
            name: pageT("groupType"),
            selector: row => row.typeID
        },
        {
            cell: (row, index) => <div className="flex w-28">
                <NoticeIcon
                    className="mr-5"
                    onClick={() => {
                        setOpenNotice(true)
                        setTarget(row)
                    }} />
                {row.groupType === "Area Owner"
                    ? null
                    : <>
                        <EditIcon className="mr-5"
                            onClick={() => {
                                setOpenEdit(true)
                                setTarget({ ...row, index })
                            }} />
                        <DeleteIcon onClick={() => {
                            setOpenDelete(true)
                            setTarget(row)
                        }} />
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
                setGroupTypes(data.groupTypes?.reduce((acc, cur) => {
                    acc[cur.id] = cur.name
                    return acc
                }, {}) || {})
            },
            url: `/api/account-management/groups`
        })
    }
    useEffect(() => getList(), [])
    return <>
        <h1 className="mb-9">{commonT("accountManagementGroup")}</h1>
        <div className="mb-9">
            <AddGroup {...{ getList, groupList, groupTypes }} />
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
        {/* notice */}
        <DialogForm
            dialogTitle={commonT("group")}
            fullWidth={fullWidth}
            maxWidth={maxWidth}
            open={openNotice}
            setOpen={setOpenNotice}>
            <Divider variant="middle" />
            <div className="flex flex-col m-auto mt-4 min-w-49 w-fit">
                <div className="grid grid-cols-1fr-auto">
                    <h5 className="ml-6 mt-2">{commonT("groupName")} :</h5>
                    <ListItem
                        id="name"
                        label={commonT("groupName")}>
                        {target?.groupName || ""}
                    </ListItem>
                    <h5 className="ml-6 mt-2">{pageT("groupType")} :</h5>
                    <ListItem
                        id="group-type"
                        label={pageT("groupType")}
                    >
                        {target?.groupType || ""}
                    </ListItem>
                    {target?.parentGroup
                        ? <> <h5 className="ml-6 mt-2">{pageT("parentGroup")} :</h5>
                            <ListItem
                                id="parent-group-type"
                                label={pageT("parentGroup")}
                            >
                                {target?.parentGroup || ""}
                            </ListItem></>
                        : null}
                    <h5 className="ml-6 mt-2">{pageT("fieldList")} :</h5>
                    <ListItem
                        id="field-list"
                        label={pageT("fieldList")}
                    >
                        {target?.fieldList || ""}
                    </ListItem>
                </div>
            </div>
            <DialogActions sx={{ margin: "1rem 0.5rem 1rem 0" }}>
                <Button onClick={() => { setOpenNotice(false) }}
                    radius="pill"
                    variant="contained"
                    color="primary">
                    {commonT("okay")}
                </Button>
            </DialogActions>
        </DialogForm>
        {/* edit */}
        <DialogForm
            dialogTitle={commonT("group")}
            fullWidth={fullWidth}
            maxWidth={maxWidth}
            open={openEdit}
            setOpen={setOpenEdit}>
            <Divider variant="middle" />
            <FormControl sx={{
                display: "flex",
                flexDirection: "column",
                margin: "auto",
                width: "fit-content",
                mt: 2,
                minWidth: 120
            }}>
                <TextField
                    id="edit-name"
                    label={commonT("groupName")}
                    onChange={handleChange}
                    value={target?.groupName || ""}
                    focused>
                </TextField>
            </FormControl>
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
                <h5 className="ml-6 mr-2">{commonT("groupName")} :</h5>
                {target?.groupName || ""}
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