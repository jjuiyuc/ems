import { TextField } from "@mui/material"
import DatePicker from "react-datepicker"
import { useTranslation } from "react-multi-lang"
import moment from "moment"
import { useState } from "react"
import "../assets/css/timeRangePicker.css"

export default function TimeRangePicker(props) {
    const
        t = useTranslation(),
        pageT = (string, params) => t("settings." + string, params)

    const { startTime, setStartTime, endTime, setEndTime, basicPrice, rate,
        setBasicPrice, setRate } = props


    return (
        <div className="time-range-picker flex items-center">
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
                    style={{ paddingBottom: "16.5px", paddingTop: "16.5px" }}
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
            <div className="ml-6 mr-4">
                <TextField
                    id="outlined-basic"
                    variant="outlined"
                    value={basicPrice}
                    onChange={(e) => { setBasicPrice(e.target.value) }}
                />
            </div>
            <div>
                <TextField
                    id="outlined-basic"
                    label="rate"
                    value={rate}
                    onChange={(e) => { setRate(e.target.value) }}
                />
            </div>
        </div>
    )
}