import DatePicker from "react-datepicker"
import { useTranslation } from "react-multi-lang"
import moment from "moment"
import "../assets/css/monthPicker.css"

export default function MonthPicker(props) {
    const
        t = useTranslation(),
        pageT = (string, params) => t("analysis." + string, params)

    const { startMonth, setStartMonth } = props

    const onChange = (date) => setStartMonth(date)
    return (
        <>
            <div>
                <h6 className="mb-1 ml-1">{pageT("selectMonth")}</h6>
                <DatePicker
                    selected={startMonth}
                    onChange={onChange}
                    dateFormat="yyyy/MM"
                    maxDate={moment(new Date())._d}
                    showMonthYearPicker
                    showFullMonthYearPicker
                />
            </div>
        </>
    )
}