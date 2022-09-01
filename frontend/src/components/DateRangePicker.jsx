import DatePicker from "react-datepicker"
import { useTranslation } from "react-multi-lang"
import moment from "moment"
import { useState } from "react"

export default function DateRangePicker(props) {
    const
        t = useTranslation(),
        pageT = (string, params) => t("analysis." + string, params)

    const
        [startDate, setStartDate] = useState(null),
        [endDate, setEndDate] = useState(null)
    const onChange = (dates) => {
        const [start, end] = dates
        setStartDate(start)
        setEndDate(end)
    }
    return (
        <>
            <div>
                <h6 className="mb-1 ml-1">{pageT("startDate")}</h6>
                <DatePicker
                    dateFormat="yyyy/MM/dd"
                    selected={startDate}
                    onChange={onChange}
                    value={startDate ? moment(startDate).format("yyyy/MM/DD") : ""}
                    selectsStart
                    startDate={startDate}
                    endDate={endDate}
                    minDate={moment().subtract(2, "week")._d}
                    maxDate={moment(new Date()).add(-1, 'days')._d}
                    showDisabledMonthNavigation
                    monthsShown={2}
                    selectsRange
                />
            </div>
            <span className="mt-6 mx-4">{pageT("to")}</span>
            <div>
                <h6 className="mb-1 ml-1">{pageT("endDate")}</h6>
                <DatePicker
                    dateFormat="yyyy/MM/dd"
                    selected={endDate}
                    onChange={(date) => setEndDate(date)}
                    selectsEnd
                    startDate={startDate}
                    endDate={endDate}
                    minDate={startDate}
                    maxDate={moment(new Date()).add(-1, 'days')._d}
                    disabled={true}
                />
            </div>
        </>
    )
}