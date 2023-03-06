import DataTable from "react-data-table-component"
import moment from "moment"
import { useTranslation } from "react-multi-lang"
import { useEffect, useMemo, useState } from "react"

export default function AccountManagementGroup() {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("accountManagementGroup." + string, params)
    return <>
        <h1 className="mb-9">{commonT("accountManagementGroup")}</h1>
    </>


}