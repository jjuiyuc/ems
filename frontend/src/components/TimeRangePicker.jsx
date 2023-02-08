import { TextField } from "@mui/material"
import DatePicker from "react-datepicker"
import { useTranslation } from "react-multi-lang"
import moment from "moment"

import { validateNumTwoDecimalPlaces } from "../utils/utils"
import "../assets/css/timeRangePicker.css"

export default function TimeRangePicker(props) {
    const
        t = useTranslation(),
        pageT = (string, params) => t("settings." + string, params),
        errorT = (string) => t("error." + string)

    const { startTime, setStartTime, endTime, setEndTime, basicPrice, rate,
        setBasicPrice, setRate } = props
    const
        inputPrice = (e) => {
            const num = e.target.value
            const isNum = validateNumTwoDecimalPlaces(num)
            if (!isNum) return
            setBasicPrice(num)
        },
        inputRate = (e) => {
            const num = e.target.value
            const isNum = validateNumTwoDecimalPlaces(num)
            if (!isNum) return
            setRate(num)
        }
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
            <span className="mt-6">{pageT("to")}</span>
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
            <div>
                <h6 className="mb-1 ml-1">{pageT("belowBaseline")}</h6>
                <TextField
                    className="react-datepicker__input-container"
                    id="outlined-basic"
                    variant="outlined"
                    value={basicPrice}
                    onChange={inputPrice}
                />
            </div>
            <div>
                <h6 className="mb-1 ml-1">{pageT("aboveBaseline")}</h6>
                <TextField
                    id="outlined-basic"
                    variant="outlined"
                    value={rate}
                    onChange={inputRate}
                />
            </div>
        </>
    )
}