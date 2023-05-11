import DatePicker from "react-datepicker"
import moment from "moment"
import ArrowBackIosIcon from "@mui/icons-material/ArrowBackIos"
import ArrowForwardIosIcon from "@mui/icons-material/ArrowForwardIos"
import "../assets/css/dateRangePicker.css"
import { useState } from "react"

export default function DayPicker(props) {

    const
        minDate = moment().subtract(1, "month").add(-1, "days")._d,
        maxDate = moment(new Date()).startOf("day").subtract(1, "days")._d

    const { startDay, setStartDay, setEnableRequest } = props

    const [timeoutId, setTimeoutId] = useState(null)

    const timeoutHandler = () => {
        clearTimeout(timeoutId)

        const newTimeoutId = setTimeout(() => {
            setEnableRequest(true)
        }, 300)
        setTimeoutId(newTimeoutId)
    }
    const
        onLeftClick = () => {
            if (startDay >= minDate) {
                const newDay = moment(startDay).subtract(1, "day")._d
                setStartDay(newDay)
                timeoutHandler()
            }
        },
        onRightClick = () => {
            if (startDay < maxDate) {
                const newDay = moment(startDay).add(1, "day")._d
                setStartDay(newDay)
                timeoutHandler()
            }
        }
    return (
        <>
            <div className="flex items-center">
                <ArrowBackIosIcon onClick={onLeftClick} />
                <DatePicker
                    dateFormat="yyyy/MM/dd"
                    selected={startDay}
                    onChange={(date) => setStartDay(date)}
                    value={startDay ? moment(startDay).format("yyyy/MM/DD") : ""}
                    minDate={minDate}
                    maxDate={maxDate}
                    monthsShown={2}
                    showDisabledMonthNavigation
                />
                <ArrowForwardIosIcon onClick={onRightClick} className="ml-1" />
            </div>
        </>
    )
}