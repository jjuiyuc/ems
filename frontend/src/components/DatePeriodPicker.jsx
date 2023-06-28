import { MenuItem, TextField } from "@mui/material"
import DatePicker from "react-datepicker"
import { useTranslation } from "react-multi-lang"
import moment from "moment"
import "../assets/css/timeRangePicker.css"

export default function DatePeriodPicker(props) {
    const
        t = useTranslation(),
        pageT = (string, params) => t("settings." + string, params),
        errorT = (string) => t("error." + string)

    const { startDate, setStartDate, endDate, setEndDate, typeDict, type, setType } = props

    const changeType = (e) => {
        setType(e.target.value)
    }
    return (
        <>
            <div>
                <h6 className="mb-1 ml-1">{pageT("startDate")}</h6>
                <DatePicker
                    showTimeSelect
                    timeFormat="HH:mm"
                    timeIntervals={15}
                    dateFormat="yyyy/MM/dd HH:mm"
                    selected={startDate}
                    onChange={(date) => setStartDate(date)}
                    value={startDate ? moment(startDate).format("yyyy/MM/DD HH:mm") : ""}
                    selectsStart
                    startDate={startDate}
                    endDate={endDate}
                    minDate={moment(new Date())._d}
                />
            </div>
            <span className="mt-6">{pageT("to")}</span>
            <div>
                <h6 className="mb-1 ml-1">{pageT("endDate")}</h6>
                <DatePicker
                    showTimeSelect
                    timeFormat="HH:mm"
                    timeIntervals={15}
                    dateFormat="yyyy/MM/dd HH:mm"
                    selected={endDate}
                    onChange={(date) => setEndDate(date)}
                    selectsEnd
                    endDate={endDate}
                    startDate={startDate}
                    value={endDate ? moment(endDate).format("yyyy/MM/DD HH:mm") : ""}
                    minDate={startDate}
                    minTime={moment(startDate).add(1, "minutes")._d}
                    maxTime={moment().startOf("day")._d}
                    disabled={!startDate}
                />
            </div>
            <div className="flex flex-col m-auto min-w-49 w-fit">
                <h6 className="mb-1 ml-1">{pageT("type")}</h6>
                <TextField
                    id="p-o-type"
                    select
                    variant="outlined"
                    onChange={changeType}
                    defaultValue="">
                    {Object.entries(typeDict).map(([key, value]) =>
                        <MenuItem key={"type-o-" + key} value={key}>
                            {pageT(`${value}`)}
                        </MenuItem>)}
                </TextField>
            </div>
        </>
    )
}