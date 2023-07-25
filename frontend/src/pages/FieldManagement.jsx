import { useTranslation } from "react-multi-lang"
import { useEffect, useState } from "react"

import { apiCall } from "../utils/api"

import AddField from "../components/AddField"
import EditField from "../components/EditField"
import InfoField from "../components/InfoField"
import Table from "../components/DataTable"

export default function FieldManagement() {
    const
        t = useTranslation(),
        commonT = string => t("common." + string)

    const
        [fieldList, setFieldList] = useState([]),
        [loading, setLoading] = useState(false),
        [infoError, setInfoError] = useState(""),
        [fetched, setFetched] = useState(false)

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
                <InfoField {...{ row }} />
                <EditField className="mr-5" {...{ row, fieldList, setFieldList }} />
            </div>,
            center: true,
            grow: 0.5
        }
    ]
    const getList = () => {
        apiCall({
            onComplete: () => {
                setLoading(false)
                setFetched(true)
            },
            onError: error => setInfoError(error),
            onStart: () => setLoading(true),
            onSuccess: rawData => {
                if (!rawData?.data) return

                const { data } = rawData

                setFieldList(data.gateways || [])
            },
            url: `/api/device-management/gateways`
        })
    }
    useEffect(() => {
        if (fetched == false)
            getList()
    }, [fetched])

    return <>
        <h1 className="mb-9">{commonT("fieldManagement")}</h1>
        <div className="mb-9">
            <AddField {...{ getList }} />
        </div>
        <Table
            columns={columns}
            data={fieldList}
            paginationComponentOptions={{
                rowsPerPageText: t("dataTable.rowsPerPage")
            }}
            paginationPerPage={100}
            progressPending={loading}
            theme="dark"
        />
    </>
}