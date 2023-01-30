import { TextField } from "@mui/material"
import DatePicker from "react-datepicker"
import { useTranslation } from "react-multi-lang"
import moment from "moment"
import "../assets/css/timeRangePicker.css"

export default function DatePeriodPicker(props) {
    const
        t = useTranslation(),
        pageT = (string, params) => t("settings." + string, params),
        errorT = (string) => t("error." + string)

    const { startDate, setStartDate, endDate, setEndDate } = props

    return (
        <>
            <div>
                <h6 className="mb-1 ml-1">{pageT("startDate")}</h6>
                <DatePicker
                    dateFormat="yyyy/MM/dd h:mm A"
                    showTimeSelect
                    selected={startDate}
                    onChange={(date) => setStartDate(date)}
                    value={startDate ? moment(startDate).format("yyyy/MM/DD h:mm A") : ""}
                    selectsStart
                    startDate={startDate}
                    endDate={endDate}
                />
            </div>
            <span className="mt-6">{pageT("to")}</span>
            <div>
                <h6 className="mb-1 ml-1">{pageT("endDate")}</h6>
                <DatePicker
                    dateFormat="yyyy/MM/DD h:mm A"
                    showTimeSelect
                    selected={endDate}
                    onChange={(date) => setEndDate(date)}
                    selectsEnd
                    endDate={endDate}
                    startDate={startDate}
                    value={endDate ? moment(endDate).format("yyyy/MM/DD h:mm A") : ""}
                    minDate={startDate}
                    disabled={!startDate}
                />
            </div>
        </>
    )
}