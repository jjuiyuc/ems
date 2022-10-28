import DatePicker from "react-datepicker"
import { useTranslation } from "react-multi-lang"
import moment from "moment"
import { useState } from "react"

export default function TimeRangePicker(props) {
    const
        t = useTranslation(),
        pageT = (string, params) => t("settings." + string, params)

    const { startTime, setStartTime, endTime, setEndTime } = props


    return (
        <>
            <div>
                <h6 className="mb-1 ml-1">{pageT("startTime")}</h6>
                <DatePicker
                    dateFormat="h:mm aa"
                    showTimeSelect
                    showTimeSelectOnly
                    timeIntervals={30}
                    timeCaption="Start Time"
                    selected={startTime}
                    onChange={(time) => setStartTime(time)}
                />
            </div>
            <span className="mt-6 mx-4">{pageT("to")}</span>
            <div>
                <h6 className="mb-1 ml-1">{pageT("endTime")}</h6>
                <DatePicker
                    dateFormat="h:mm aa"
                    showTimeSelect
                    showTimeSelectOnly
                    timeIntervals={30}
                    timeCaption="End Time"
                    selected={endTime}
                    onChange={(time) => setEndTime(time)}
                    minTime={moment(startTime).add(1, "minute")._d}
                    maxTime={moment().endOf("day")._d}
                    disabled={!startTime}
                />
            </div>
        </>
    )
}