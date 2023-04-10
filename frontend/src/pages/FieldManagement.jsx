import {
    Button, DialogActions, Divider, FormControl, ListItem,
    MenuItem, TextField
} from "@mui/material"
import { useTranslation } from "react-multi-lang"
import { useEffect, useState } from "react"

import AddField from "../components/AddField"
import EditField from "../components/EditField"
import InfoField from "../components/InfoField"
import Table from "../components/DataTable"

export default function FieldManagement() {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        dialogT = (string) => t("dialog." + string),
        formT = (string) => t("form." + string),
        pageT = (string, params) => t("fieldManagement." + string, params)
    const
        [data, setData] = useState([
            {
                id: 1,
                locationName: "Serenegray",
                gatewayID: "0E0BA27A8175AF978C49396BDE9D7A1E",
                address: "宜蘭縣五結鄉大吉五路157巷68號",
                "lat": "24.702",
                "lng": "121.797",
                "powerCompany": "TPC",
                "voltageType": "lowVoltage",
                "touType": "twoSection",
                "deviceType": "hybridInverter",
                "deviceModel": "LXP-12K US-Luxpower Hybrid-Inverter",
                "modbusID": "1",
                "UUEID": "0E8F167E58271833EA01BAE79F2FD8C0",
                "powerCapacity": "24",
                "subDevice": [
                    {
                        deviceType: "meter",
                        deviceModel: "CMO336 CM Meter",
                        "powerCapacity": "22"
                    },
                    {
                        deviceType: "pv",
                        deviceModel: "D1K330H3A URE PV",
                        "powerCapacity": "23"
                    },
                    {
                        deviceType: "battery",
                        deviceModel: "L051100-A UZ-Energy Battery",
                        "powerCapacity": "24"
                    }
                ],
                "enableField": true,
                "group": [
                    "AreaOwnerTW", "AreaMaintainer", "Serenegray"
                ]
            },
            {
                id: 2,
                locationName: "Cht_Miaoli",
                gatewayID: "018F1623ADD8E739F7C6CBE62A7DF3C0",
                address: "苗栗",
                "lat": "",
                "lng": "",
                "powerCompany": "TPC",
                "voltageType": "",
                "touType": "",
                "enableField": false,
                "group": [
                    "AreaOwnerTW", "AreaMaintainer"
                ]
            }
        ]),
        [error, setError] = useState(null),
        [loading, setLoading] = useState(false),
        [locationName, setLocationName] = useState(data?.locationName || ""),
        [locationNameError, setLocationNameError] = useState(null),
        [gatewayID, setGatewayID] = useState(data?.gatewayID || ""),
        [gatewayIDError, setGatewayIDError] = useState(null)

    const editSave = (row) => {
        const newData = data.map((value) =>
            value.id === row.id ? row : value
        )
        setData(newData)
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
                <InfoField
                    row={row}
                    locationInfo={pageT("locationInformation")}
                    fieldDevices={pageT("fieldDevices")}
                    deviceInfo={pageT("deviceInformation")}
                    extraDeviceInfo={pageT("extraDeviceInfo")}
                    subdevice={pageT("subdevice")}
                />
                <EditField className="mr-5"
                    row={row}
                    onSave={editSave}
                />
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
            data={data}
            columns={columns}
            paginationComponentOptions={{
                rowsPerPageText: t("dataTable.rowsPerPage")
            }}
            paginationPerPage={100}
            progressPending={loading}
            theme="dark"
        />
    </>
}