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

    const onChange = (dates) => {
        const [start, end] = dates
        setStartDate(start)
        setEndDate(end)
    }
    // const filterPassedTime = (time) => {
    //     const currentDate = new Date()
    //     const selectedDate = new Date(time)

    //     return currentDate.getTime() < selectedDate.getTime();
    // }
    return (
        <>
            <div>
                <h6 className="mb-1 ml-1">{pageT("startTime")}</h6>
                <DatePicker
                    dateFormat="yyyy/MM/dd h:mm"
                    showTimeSelect
                    selected={startDate}
                    // filterTime={filterPassedTime}
                    onChange={(date) => setStartDate(date)}
                    value={startDate ? moment(startDate).format("yyyy/MM/DD h:mm") : ""}
                    startDate={startDate}
                />
            </div>
            <span className="mt-6">{pageT("to")}</span>
            <div>
                <h6 className="mb-1 ml-1">{pageT("endTime")}</h6>
                <DatePicker
                    dateFormat="yyyy/MM/DD h:mm"
                    showTimeSelect
                    selected={endDate}
                    // filterTime={filterPassedTime}
                    onChange={(date) => setEndDate(date)}
                    endDate={endDate}
                    value={endDate ? moment(endDate).format("yyyy/MM/DD h:mm") : ""}
                    minDate={startDate}
                    disabled={!startDate}
                />
            </div>
        </>
    )
}