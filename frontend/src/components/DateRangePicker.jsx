import DatePicker, { CalendarContainer } from "react-datepicker"
import { useTranslation } from "react-multi-lang"
import moment from "moment"
import { useState } from "react"

// import "../assets/css/clock.scss"

export default function DateRangePicker(props) {
    const
        t = useTranslation(),
        pageT = (string, params) => t("analysis." + string, params)

    const
        [startDate, setStartDate] = useState(null),
        [endDate, setEndDate] = useState(null)
    const onChange = (dates) => {
        const [start, end] = dates;
        setStartDate(start);
        setEndDate(end);
    }


    return (
        <>
            <DatePicker
                dateFormat="yyyy/MM/dd"
                selected={startDate}
                onChange={onChange}
                value={startDate ? moment(startDate).format("yyyy/MM/DD") : ""}
                selectsStart
                startDate={startDate}
                endDate={endDate}
                minDate={moment().subtract(2, "month")._d}
                maxDate={new Date()}
                showDisabledMonthNavigation
                monthsShown={2}
                selectsRange
            />
            <span className="mx-4">{pageT("to")}</span>
            <DatePicker
                dateFormat="yyyy/MM/dd"
                selected={endDate}
                onChange={(date) => setEndDate(date)}
                selectsEnd
                startDate={startDate}
                endDate={endDate}
                minDate={startDate}
                maxDate={new Date()}
                disabled={true}
            />
        </>
    )
}