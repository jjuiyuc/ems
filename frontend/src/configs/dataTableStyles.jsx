import variables from "./variables"

const { colors } = variables

const DATATABLE_STYLES = {
    headRow: {
        style: {
            backgroundColor: colors.primary["main-opacity-90"],
            border: "1px solid " + colors.gray[400],
            borderRadius: ".25rem .25rem 0 0",
            fontSize: "1rem",
            fontWeight: 600
        }
    },
    noData: {
        style: {
            backgroundColor: colors.gray[200],
            border: "1px solid " + colors.gray[100],
            borderRadius: ".25rem",
            color: colors.gray[600],
            padding: "1.5rem"
        }
    },
    pagination: {
        style: {
            backgroundColor: colors.gray[600],
            borderColor: colors.gray[400],
            borderStyle: "solid",
            borderWidth: "0 1px 1px 1px",
            borderRadius: "0 0 .25rem .25rem",
            fontSize: ".85rem"
        }
    },
    rows: {
        style: {
            backgroundColor: colors.gray[600],
            borderColor: colors.gray[400],
            borderStyle: "solid",
            borderWidth: "0 1px 1px",
            fontSize: "1rem"
        }
    },
    table: { style: { backgroundColor: "transparent" } },
    tableWrapper: { style: { backgroundColor: "transparent" } }
}

export default DATATABLE_STYLES