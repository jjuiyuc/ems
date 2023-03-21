import {
    Button, DialogActions, Divider, FormControl, ListItem,
    MenuItem, TextField
} from "@mui/material"
import AddIcon from "@mui/icons-material/Add"
import { useTranslation } from "react-multi-lang"
import { useEffect, useMemo, useState } from "react"

import DialogForm from "../components/DialogForm"
import Table from "../components/DataTable"

import { ReactComponent as DeleteIcon } from "../assets/icons/trash_solid.svg"
import { ReactComponent as EditIcon } from "../assets/icons/edit.svg"
import { ReactComponent as NoticeIcon } from "../assets/icons/notice.svg"

export default function AccountManagementGroup() {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        dialogT = (string) => t("dialog." + string),
        pageT = (string, params) => t("accountManagementGroup." + string, params)
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
    const
        [data, setData] = useState([
            {
                id: 1,
                groupName: "AreaOwner_TW",
                groupType: "Area Owner",
                fieldList: "AreaOwner_TW-0E0BA27A8175AF978C49396BDE9D7A1E"
            },
            {
                id: 2,
                groupName: "AreaMaintainer_TW",
                groupType: "Area Maintainer",
                fieldList: "AreaMaintainer_TW-0E0BA27A8175AF978C49396BDE9D7A1E"

            },
            {
                id: 3,
                groupName: "Serenegray",
                groupType: "Field Owner",
                fieldList: "Serenegray-0E0BA27A8175AF978C49396BDE9D7A1E",
                parentGroup: "AreaOwner_TW"
            },
            {
                id: 4,
                groupName: "Cht_Miaoli",
                groupType: "Area maintainer",
                fieldList: "Cht_Miaoli-0E0BA27A8175AF978C49396BDE9D7A1E",
                parentGroup: "AreaOwner_TW"

            }
        ]),
        [error, setError] = useState(null),
        [loading, setLoading] = useState(false),
        [openAdd, setOpenAdd] = useState(false),
        [openNotice, setOpenNotice] = useState(false),
        [openEdit, setOpenEdit] = useState(false),
        [openDelete, setOpenDelete] = useState(false),
        [groupName, setGroupName] = useState(data?.groupName || ""),
        [groupNameError, setGroupNameError] = useState(null),
        [groupType, setGroupType] = useState(""),
        [groupTypeError, setGroupTypeError] = useState(null),
        [parentGroup, setParentGroup] = useState(""),
        [parentGroupError, setParentGroupError] = useState(null),
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
            cell: row => <span className="font-mono">{row.groupName}</span>,
            center: true,
            name: commonT("groupName"),
            selector: row => row.groupName
        },
        {
            cell: row => <span className="font-mono">{row.groupType}</span>,
            center: true,
            name: pageT("groupType"),
            selector: row => row.groupType
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

    return <>
        <h1 className="mb-9">{commonT("accountManagementGroup")}</h1>
        <div className="mb-9">
            <Button
                onClick={() => { setOpenAdd(true) }}
                key={"ac-b-"}
                size="x-large"
                variant="outlined"
                radius="pill"
                fontSize="large"
                color="brand"
                startIcon={<AddIcon />}>
                {commonT("add")}
            </Button>
            <DialogForm
                type={"addGroup"}
                dialogTitle={commonT("group")}
                open={openAdd}
                setOpen={setOpenAdd}>
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
                        id="add-name"
                        label={commonT("groupName")}
                        value={groupName}
                        focused
                    />
                    <TextField
                        id="add-type"
                        select
                        label={pageT("groupType")}
                        defaultValue=""
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
                        label={pageT("parentGroup")}
                        defaultValue=""
                    >
                        {parentGroupType.map((option) => (
                            <MenuItem key={option.value} value={option.value}>
                                {option.label}
                            </MenuItem>
                        ))}
                    </TextField>
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
            paginationComponentOptions={{
                rowsPerPageText: t("dataTable.rowsPerPage")
            }}
            paginationPerPage={100}
            progressPending={loading}
            theme="dark"
        />
        {/* notice */}
        <DialogForm
            type={"notice"}
            dialogTitle={commonT("group")}
            open={openNotice}
            setOpen={setOpenNotice}>
            <Divider variant="middle" />
            <div className="flex flex-col m-auto mt-4 min-w-49.5 w-fit">
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
            type={"editGroup"}
            dialogTitle={commonT("group")}
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
            type={"delete"}
            dialogTitle={dialogT("deleteMsg")}
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