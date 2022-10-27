import DatePicker from "react-datepicker"
import { useTranslation } from "react-multi-lang"
import moment from "moment"
import { useState } from "react"

export default function TimeRangePicker(props) {
    const
        t = useTranslation(),
        pageT = (string, params) => t("settings." + string, params)

    const { startTime, setStartTime, endTime, setEndTime } = props

    const onChange = (dates) => {
        const [start, end] = dates
        setStartTime(start)
        setEndTime(end)
    }
    return (
        <>
            <div>
                <h6 className="mb-1 ml-1">{pageT("startTime")}</h6>
                <DatePicker
                    dateFormat="h:mm aa"
                    showTimeSelect
                    showTimeSelectOnly
                    timeIntervals={1}
                    timeCaption="Start Time"
                    selected={startTime}
                    onChange={onChange}
                    value={startTime ? moment(startTime).format("h:mm aa") : ""}
                    selectsStart
                    startDate={startTime}
                    endDate={endTime}
                    // minDate={moment().subtract(2, "week")._d}
                    // maxDate={moment(new Date()).add(-1, 'days')._d}
                    showDisabledMonthNavigation
                    selectsRange
                />
            </div>
            <span className="mt-6 mx-4">{pageT("to")}</span>
            <div>
                <h6 className="mb-1 ml-1">{pageT("endTime")}</h6>
                <DatePicker
                    dateFormat="h:mm aa"
                    showTimeSelect
                    showTimeSelectOnly
                    timeIntervals={1}
                    timeCaption="End Time"
                    selected={endTime}
                    onChange={(date) => setEndTime(date)}
                    selectsEnd
                    startDate={startTime}
                    endDate={endTime}
                    minDate={startTime}
                    // maxDate={moment(new Date()).add(-1, 'days')._d}
                    disabled={true}
                />
            </div>
        </>
    )
}