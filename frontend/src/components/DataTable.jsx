import DataTable from "react-data-table-component"
import { useTranslation } from "react-multi-lang"

import DATATABLE_STYLES from "../configs/dataTableStyles"

export default function Table(props) {
    const t = useTranslation(),
        tableT = string => t("dataTable." + string)

    return <DataTable
        {...props}
        customStyles={DATATABLE_STYLES}
        noDataComponent={tableT("noDataMsg")}
        paginationRowsPerPageOptions={[10, 25, 50, 100]} />
}
