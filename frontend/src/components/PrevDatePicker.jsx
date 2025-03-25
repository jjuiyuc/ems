import DatePicker from "react-datepicker"
import moment from "moment"
import ArrowBackIosIcon from "@mui/icons-material/ArrowBackIos"
import ArrowForwardIosIcon from "@mui/icons-material/ArrowForwardIos"
import "../assets/css/dateRangePicker.css"

export default function PrevDatePicker(props) {

    const
        minDate = moment().subtract(1, "month").add(-1, "days")._d,
        maxDate = moment().startOf("day").subtract(1, "days")._d

    const { prevDate, setPrevDate } = props

    const
        onLeftClick = () => {
            if (prevDate >= minDate) {
                const newDate = moment(prevDate).subtract(1, "day")._d
                setPrevDate(newDate)
            }
        },
        onRightClick = () => {
            if (prevDate < maxDate) {
                const newDate = moment(prevDate).add(1, "day")._d
                setPrevDate(newDate)
            }
        }
    return (
        <div className="flex items-center">
            <ArrowBackIosIcon
                onClick={onLeftClick}
                className={prevDate <= minDate ? "opacity-30" : ""}
            />
            <DatePicker
                dateFormat="yyyy/MM/dd"
                selected={prevDate}
                onChange={(date) => setPrevDate(date)}
                value={prevDate ? moment(prevDate).format("yyyy/MM/DD") : ""}
                minDate={minDate}
                maxDate={maxDate}
                monthsShown={2}
                showDisabledMonthNavigation
            />
            <ArrowForwardIosIcon
                onClick={onRightClick}
                className={"ml-1" + (prevDate >= maxDate ? " opacity-30" : "")}
            />
        </div>
    )
}