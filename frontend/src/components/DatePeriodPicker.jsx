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
                    showTimeSelect
                    timeFormat="HH:mm"
                    timeIntervals={15}
                    dateFormat="yyyy/MM/dd HH:mm"
                    selected={startDate}
                    onChange={(date) => setStartDate(date)}
                    value={startDate ? moment(startDate).format("yyyy/MM/DD HH:mm") : ""}
                    selectsStart
                    startDate={startDate}
                    endDate={endDate}
                    minDate={moment(new Date())._d}
                />
            </div>
            <span className="mt-6">{pageT("to")}</span>
            <div>
                <h6 className="mb-1 ml-1">{pageT("endDate")}</h6>
                <DatePicker
                    showTimeSelect
                    timeFormat="HH:mm"
                    timeIntervals={15}
                    dateFormat="yyyy/MM/dd HH:mm"
                    selected={endDate}
                    onChange={(date) => setEndDate(date)}
                    selectsEnd
                    endDate={endDate}
                    startDate={startDate}
                    value={endDate ? moment(endDate).format("yyyy/MM/DD HH:mm") : ""}
                    minDate={startDate}
                    minTime={moment(startDate).add(30, "minutes")._d}
                    maxTime={moment().endOf("day")._d}
                    disabled={!startDate}
                />
            </div>
            <div>
                <h6 className="mb-1 ml-1">{pageT("type")}</h6>
                <TextField
                    id="p-o-type"
                    // select
                    variant="outlined"
                // value={rate}
                // onChange={inputRate}
                />
            </div>
        </>
    )
}