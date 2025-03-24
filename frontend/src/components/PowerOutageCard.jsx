import { connect } from "react-redux"
import { useEffect, useState } from "react"
import { useTranslation } from "react-multi-lang"
import moment from "moment"

import { apiCall } from "../utils/api"

import AddPowerOutagePeriod from "./AddPowerOutagePeriod"
import DeletePeriod from "./DeletePeriod"
import Table from "./DataTable"

import { ReactComponent as DeleteIcon } from "../assets/icons/trash_solid.svg"
import { ReactComponent as PowerOutageIcon } from "../assets/icons/power_outage.svg"

const mapState = state => ({
    gatewayID: state.gateways.active.gatewayID
})
const mapDispatch = dispatch => ({
    updateSnackbarMsg: value =>
        dispatch({ type: "snackbarMsg/updateSnackbarMsg", payload: value }),

})
export default connect(mapState, mapDispatch)(function PowerOutageCard(props) {
    const
        t = useTranslation(),
        pageT = (string, params) => t("settings." + string, params)

    const
        [periodList, setPeriodList] = useState([]),
        [row, setRow] = useState(null),
        [openDelete, setOpenDelete] = useState(false),
        [loading, setLoading] = useState(false),
        [otherError, setOtherError] = useState("")

    const handleClickDelete = row => {
        setOpenDelete(true)
        setRow(row)
    }
    const columns = [
        {
            cell: row => <span className="font-mono">{moment(row.startTime).format("YYYY/MM/DD HH:mm")}</span>,
            center: true,
            name: pageT("startDate"),
            selector: row => row.startTime,
            grow: 0.6

        },
        {
            cell: row => <span className="font-mono">{moment(row.endTime).format("YYYY/MM/DD HH:mm")}</span>,
            center: true,
            name: pageT("endDate"),
            selector: row => row.endTime,
            grow: 0.6
        },
        {
            cell: row => <span className="font-mono">{pageT(`${row.type}`)}</span>,
            center: true,
            name: pageT("type"),
            selector: row => row.type,
            grow: 0.5

        },
        {
            cell: (row) => <div className="flex w-24">
                {row.ongoing === true
                    ? <div className="bg-gray-600 w-6 h-6"></div>
                    : <DeleteIcon onClick={() => handleClickDelete(row)} />
                }
            </div>,
            center: true,
            grow: 0.2
        }
    ]
    const getList = () => {

        const gatewayID = props.gatewayID

        apiCall({
            onComplete: () => setLoading(false),
            onStart: () => setLoading(true),
            onError: (err) => {
                switch (err) {
                    case 60034:
                        props.updateSnackbarMsg({
                            type: "error",
                            msg: errorT("failureToGenerate")
                        })
                        break
                    default:setOtherError(err)
                }
            },
            onSuccess: rawData => {
                if (!rawData?.data) return

                const { data } = rawData

                setPeriodList(data.periods || [])
            },
            url: `/api/device-management/gateways/${gatewayID}/power-outage-periods`
        })
    }
    useEffect(() => {
        getList()
    }, [props.gatewayID])
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
            <AddPowerOutagePeriod {...{ getList, periodList }} />
        </div>
        <div className="flex flex-col mt-4 min-w-49 w-full">
            <Table
                {...{ columns, data: periodList }}
                pagination={false}
                progressPending={loading}
                theme="dark"
            />
        </div>
        <DeletePeriod {...{ row, getList, openDelete, setOpenDelete }} />
    </div>
})