import { Button } from "@mui/material"
import AddIcon from "@mui/icons-material/Add"
import moment from "moment"
import { useTranslation } from "react-multi-lang"
import { useEffect, useMemo, useState } from "react"

import DialogBox from "../components/DialogBox"
import DialogForm from "../components/DialogForm"
import Table from "../components/DataTable"

import { ReactComponent as DeleteIcon } from "../assets/icons/trash_solid.svg"

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
                <DialogForm
                    type={"notice"}
                    dialogTitle={commonT("group")}
                    okayButton={commonT("okay")}
                />
                {row.groupType === "Area Owner"
                    ? null
                    : <>
                        <DialogForm
                            type={"editGroup"}
                            dialogTitle={commonT("group")}
                            leftButtonName={commonT("cancel")}
                            rightButtonName={commonT("save")}
                        />
                        <DialogBox
                            type={"delete"}
                            leftButtonName={commonT("cancel")}
                            rightButtonName={commonT("delete")}
                        />
                    </>}
            </div>,
            center: true
        }
    ]
    return <>
        <h1 className="mb-9">{commonT("accountManagementGroup")}</h1>
        <div className="mb-9">
            <DialogForm
                type={"addGroup"}
                triggerName={commonT("add")}
                dialogTitle={commonT("group")}
                leftButtonName={commonT("cancel")}
                rightButtonName={commonT("add")}
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