import DatePicker from "react-datepicker"
import { useTranslation } from "react-multi-lang"
import moment from "moment"

export default function MonthPicker(props) {
    const
        t = useTranslation(),
        pageT = (string, params) => t("analysis." + string, params)

    const { startMonth, setStartMonth } = props

    const onChange = (date) => setStartMonth(date)
    return (
        <>
            <DatePicker
                selected={startMonth}
                onChange={onChange}
                dateFormat="yyyy/MM"
                showMonthYearPicker
                showFullMonthYearPicker
            />
        </>
    )
}