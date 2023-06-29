import { Fragment as Frag, useEffect, useState, useRef } from "react"
import { Button, Select, MenuItem, TextField } from "@mui/material"
import { useTranslation } from "react-multi-lang"
import moment from "moment"

import { apiCall } from "../utils/api"

import AddPowerOutagePeriod from "./AddPowerOutagePeriod"
import Table from "./DataTable"

import { ReactComponent as AddIcon } from "../assets/icons/add.svg"
import { ReactComponent as DeleteIcon } from "../assets/icons/delete.svg"
import { ReactComponent as PowerOutageIcon } from "../assets/icons/power_outage.svg"

const maxLength = 5

const defaultPolicyConfig = {
    preNotifiedOutagePeriod: {
        name: "preNotifiedOutagePeriod",
        tempName: "preNotifiedOutagePeriod",
        extensible: true,
        deletable: false
    }
}
const defaultPolicyTime = {
    preNotifiedOutagePeriod: [
        { startDate: "", endDate: "", type: "" },
    ]
}

export default function PowerOutageCard(props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("settings." + string, params)
    const data = [
        {
            "id": 1,
            "type": "advanceBlackout",
            "startTime": moment("2023-06-20T05:00:00.00Z").format("YYYY/MM/DD HH:mm"),
            "endTime": moment("2023-06-20T05:00:00.00Z").format("YYYY/MM/DD HH:mm")
        },
        {
            "id": 3,
            "type": "evCharge",
            "startTime": moment("2023-06-22T05:00:00.00Z").format("YYYY/MM/DD HH:mm"),
            "endTime": moment("2023-06-22T09:00:00.00Z").format("YYYY/MM/DD HH:mm")
        }
    ]
    const powerOutageTypes = [
        {
            "id": 1,
            "name": "advanceBlackout"
        },
        {
            "id": 2,
            "name": "evCharge"
        }
    ]
    const
        [policyConfig, setPolicyConfig] = useState(defaultPolicyConfig),
        [policyTime, setPolicyTime] = useState(defaultPolicyTime),
        [startDate, setStartDate] = useState(null),
        [endDate, setEndDate] = useState(null),
        [type, setType] = useState(null),
        [typeDict, setTypeDict] = useState({}),
        [openDelete, setOpenDelete] = useState(false),
        [row, setRow] = useState(null),
        [loading, setLoading] = useState(false)

    const handleClickDelete = row => {
        setOpenDelete(true)
        setRow(row)
    }
    const columns = [
        {
            cell: row => <span className="font-mono">{row.startTime}</span>,
            center: true,
            name: pageT("startDate"),
            selector: row => row.startTime,
            grow: 0.8

        },
        {
            cell: row => <span className="font-mono">{row.endTime}</span>,
            center: true,
            name: pageT("endDate"),
            selector: row => row.endTime,
            grow: 0.8
        },
        {
            cell: row => <span className="font-mono">{pageT(`${row.type}`)}</span>,
            center: true,
            name: pageT("type"),
            selector: row => row.type,
            grow: 0.6

        },
        {
            cell: (row) => <div className="flex w-24">
                <DeleteIcon onClick={() => handleClickDelete(row)} />
            </div>,
            center: true,
            grow: 0.2
        }
    ]
    const generateTypeDict = () => {
        setTypeDict(powerOutageTypes.reduce((acc, cur) => {
            acc[cur.id] = cur.name
            return acc
        }, {}) || {})
    }
    useEffect(() => {
        generateTypeDict()
    }, [])
    return <div className="card mb-8">
        <div className="flex justify-between sm:col-span-2 items-center">
            <div className="flex items-center mb-9">
                <div
                    className="bg-gray-400-opacity-20 grid h-12 w-12
                            place-items-center rounded-full">
                    <PowerOutageIcon className="text-gray-400 w-10 h-10" />
                </div>
                <h2 className="font-bold ml-4">{pageT("powerOutage")}</h2>
            </div>
            <AddPowerOutagePeriod />
        </div>
        <div className="flex flex-col mt-4 min-w-49 w-full">
            <Table
                {...{ columns, data }}
                pagination={false}
                progressPending={loading}
                theme="dark"
            />
        </div>
    </div>
}