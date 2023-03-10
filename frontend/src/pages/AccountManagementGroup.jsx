import { Button } from "@mui/material"
import AddIcon from "@mui/icons-material/Add"
import DataTable from "react-data-table-component"
import moment from "moment"
import { useTranslation } from "react-multi-lang"
import { useEffect, useMemo, useState } from "react"

import Table from "../components/DataTable"

import { ReactComponent as DeleteIcon } from "../assets/icons/trash_solid.svg"
import { ReactComponent as EditIcon } from "../assets/icons/edit.svg"
import { ReactComponent as NoticeIcon } from "../assets/icons/notice.svg"

export default function AccountManagementGroup() {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("accountManagementGroup." + string, params)
    const
        [data, setData] = useState([
            {
                id: 1,
                groupName: "AreaOwner_TW",
                groupType: "Area Owner"
            },
            {
                id: 2,
                groupName: "AreaMaintainer_TW",
                groupType: "Area Maintainer"
            },
            {
                id: 3,
                groupName: "Serenegray",
                groupType: "Field Owner"
            },
            {
                id: 4,
                groupName: "Cht_Miaoli",
                groupType: "Area maintainer"
            }
        ]),
        [error, setError] = useState(null),
        [loading, setLoading] = useState(false)

    const columns = [
        {
            cell: row => <span className="font-mono">{row.groupName}</span>,
            center: true,
            name: pageT("groupName"),
            selector: row => row.groupName
        },
        {
            cell: row => <span className="font-mono">{row.groupType}</span>,
            center: true,
            name: pageT("groupType"),
            selector: row => row.groupType
        },
        {
            cell: row => <div className="flex w-28">
                <NoticeIcon className="mr-5" />
                {row.groupType === "Area Owner"
                    ? null
                    : <>
                        <EditIcon className="mr-5" />
                        <DeleteIcon />
                    </>}
            </div>,
            center: true,
        }
    ]
    return <>
        <h1 className="mb-9">{commonT("accountManagementGroup")}</h1>
        <div className="mb-9">
            <Button
                // onClick={}
                key={"ac-b-"}
                size="x-large"
                variant="outlined"
                radius="pill"
                fontSize="large"
                color="brand"
                startIcon={<AddIcon />}>
                {commonT("add")}
            </Button>
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