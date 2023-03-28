import {
    Button, DialogActions, Divider, FormControl, ListItem,
    MenuItem, TextField
} from "@mui/material"
import AddIcon from "@mui/icons-material/Add"
import { useTranslation } from "react-multi-lang"
import { useEffect, useState } from "react"

import AddField from "../components/AddField"
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
                locationName: "Cht_Miaoli",
                gatewayID: "018F1623ADD8E739F7C6CBE62A7DF3C0"

            }
        ]),
        [error, setError] = useState(null),
        [loading, setLoading] = useState(false),

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
            selector: row => row.locationName,
            grow: 0.5

        },
        {
            cell: row =>
                <span className="font-mono">{row.gatewayID}</span>,
            center: true,
            name: commonT("gatewayID"),
            selector: row => row.gatewayID,
            grow: 1
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
            center: true,
            grow: 0.5
        }
    ]

    return <>
        <h1 className="mb-9">{commonT("fieldManagement")}</h1>
        <div className="mb-9">
            <AddField
                locationInfo={pageT("locationInformation")}
                fieldDevices={pageT("fieldDevices")}
                deviceInfo={pageT("deviceInformation")}
                extraDeviceInfo={pageT("extraDeviceInfo")}
                subdevice={pageT("subdevice")}
            />
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
    </>
}