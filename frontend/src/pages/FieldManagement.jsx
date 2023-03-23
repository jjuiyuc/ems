import {
    Button, DialogActions, Divider, FormControl, ListItem,
    MenuItem, TextField
} from "@mui/material"
import AddIcon from "@mui/icons-material/Add"
import { useTranslation } from "react-multi-lang"
import { useEffect, useMemo, useState } from "react"

import DialogForm from "../components/DialogForm"
import Table from "../components/DataTable"

import { ReactComponent as EditIcon } from "../assets/icons/edit.svg"
import { ReactComponent as NoticeIcon } from "../assets/icons/notice.svg"

export default function FieldManagement() {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        dialogT = (string) => t("dialog." + string),
        pageT = (string, params) => t("fieldManagement." + string, params)

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
                locationName: "Serenegray",
                gatewayID: "0E0BA27A8175AF978C49396BDE9D7A1E"
            },
            {
                id: 2,
                locationName: "Serenegray",
                gatewayID: "0E0BA27A8175AF978C49396BDE9D7A1E"

            }
        ]),
        [error, setError] = useState(null),
        [loading, setLoading] = useState(false),
        [openAdd, setOpenAdd] = useState(false),
        [openNotice, setOpenNotice] = useState(false),
        [openEdit, setOpenEdit] = useState(false),
        [locationName, setLocationName] = useState(data?.locationName || ""),
        [locationNameError, setLocationNameError] = useState(null),
        [gatewayID, setGatewayID] = useState(data?.gatewayID || ""),
        [gatewayIDError, setGatewayIDError] = useState(null),
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
            cell: row => <span className="font-mono">{row.locationName}</span>,
            center: true,
            name: commonT("locationName"),
            selector: row => row.locationName
        },
        {
            cell: row => <span className="font-mono">{row.gatewayID}</span>,
            center: true,
            name: pageT("gatewayID"),
            selector: row => row.gatewayID
        },
        {
            cell: (row, index) => <div className="flex w-28">
                <NoticeIcon
                    className="mr-5"
                    onClick={() => {
                        setOpenNotice(true)
                        setTarget(row)
                    }} />

                <EditIcon className="mr-5"
                    onClick={() => {
                        setOpenEdit(true)
                        setTarget({ ...row, index })
                    }} />

            </div>,
            center: true
        }
    ]

    return <>
        <h1 className="mb-9">{commonT("accountManagementGroup")}</h1>
        {/* <div className="mb-9">
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
        </div> */}
        <Table
            {...{ columns, data }}
            paginationComponentOptions={{
                rowsPerPageText: t("dataTable.rowsPerPage")
            }}
            paginationPerPage={100}
            progressPending={loading}
            theme="dark"
        />

    </>
}