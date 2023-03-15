import { Button } from "@mui/material"
import AddIcon from "@mui/icons-material/Add"
import DataTable from "react-data-table-component"
import moment from "moment"
import { useTranslation } from "react-multi-lang"
import { useEffect, useMemo, useState } from "react"

import DialogForm from "../components/DialogForm"
import Table from "../components/DataTable"

import { ReactComponent as DeleteIcon } from "../assets/icons/trash_solid.svg"
import { ReactComponent as EditIcon } from "../assets/icons/edit.svg"

export default function AccountManagementUser() {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("accountManagementUser." + string, params)
    const
        [data, setData] = useState([
            {
                id: 1,
                account: "XXXXX@ubiik.com",
                name: "XXXXX",
                group: "Area Owner_TW"
            },
            {
                id: 2,
                account: "YYYYY@ubiik.com",
                name: "YYYYY",
                group: "AreaMaintainer_TW"
            },
            {
                id: 3,
                account: "serenegray@ubiik.com",
                name: "Serenegray",
                group: "Serenegray"
            },
            {
                id: 4,
                account: "cht_miaoli@ubiik.com",
                name: "Cht_Miaoli",
                group: "Cht_Miaoli"
            }
        ]),
        [error, setError] = useState(null),
        [loading, setLoading] = useState(false)

    const columns = [
        {
            cell: row => <span className="font-mono">{row.account}</span>,
            center: true,
            name: pageT("account"),
            selector: row => row.account
        },
        {
            cell: row => <span className="font-mono">{row.name}</span>,
            center: true,
            name: pageT("name"),
            selector: row => row.name
        },
        {
            cell: row => <span className="font-mono">{row.group}</span>,
            center: true,
            name: commonT("group"),
            selector: row => row.group
        },
        {
            cell: row => <div className="flex w-28">
                <EditIcon className="mr-5" />
                {row.group === "Area Owner_TW"
                    ? null
                    : <DeleteIcon />
                }
            </div>,
            center: true,
        }
    ]
    return <>
        <h1 className="mb-9">{commonT("accountManagementUser")}</h1>
        <div className="mb-9">
            <DialogForm
                type={"addUser"}
                triggerName={commonT("add")}
                dialogTitle={pageT("user")}
                account={pageT("account")}
                password={pageT("password")}
                name={pageT("name")}
                group={commonT("group")}
                leftButtonName={commonT("cancel")}
                rightButtonName={commonT("add")}
            />
        </div>
        <Table
            {...{ columns, data }}
            customStyles={{
                headRow: {
                    style: {
                        backgroundColor: "#12c9c990",
                        fontWeight: 600,
                        fontSize: "16px",
                        borderRadius: ".45rem .45rem 0 0 "
                    }
                }
            }}
            noDataComponent={t("dataTable.noDataMsg")}
            pagination={true}
            paginationComponentOptions={{
                rowsPerPageText: t("dataTable.rowsPerPage")
            }}
            paginationPerPage={100}
            progressPending={loading}
            theme="dark"
        />
    </>

}