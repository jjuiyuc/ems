import { useMemo } from 'react'
import { MenuItem, OutlinedInput, Select, TextField } from "@mui/material"
import DatePicker from "react-datepicker"
import { useTranslation } from "react-multi-lang"
import moment from "moment"

import { validateNumTwoDecimalPlaces } from "../utils/utils"
import "../assets/css/timeRangePicker.css"

const
    getMoment = string => {
        const [hour, minute] = string.split(":")
        return moment().hour(parseInt(hour)).minute(parseInt(minute)).second(0)
    },
    compareTimeString = (t1, t2) => {
        return getMoment(t1).diff(getMoment(t2)) > 0
    }
const
    hourArr = new Array(24)
        .fill(null)
        .map((_, i) => i.toString())
        .map(
            (h) => h.length === 1
                ? `0${h}`
                : h),
    minArr = ["00", "30"],
    timeArr = [...hourArr.map((h) => minArr.map(m => `${h}:${m}`)).flat()],
    startTimeOptions = timeArr.map((time) => ({ label: time, value: time }))

export default function TimeRangePicker(props) {
    const
        t = useTranslation(),
        pageT = (string, params) => t("settings." + string, params),
        errorT = (string) => t("error." + string)

    const { startTime, setStartTime, endTime, setEndTime, basicPrice, rate,
        setBasicPrice, setRate } = props

    const endTimeOptions = useMemo(() =>
        [...timeArr, "24:00"]
            .filter((time) => compareTimeString(time, startTime))
            .map((t) => ({ label: t, value: t }))
        , [startTime]
    )
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
                <Select
                    sx={{ minWidth: 90 }}
                    id="outlined-basic"
                    variant="outlined"
                    size="medium"
                    defaultValue={""}
                    value={startTime}
                    onChange={(e) => {
                        const newStartTime = e.target.value
                        setStartTime(newStartTime)
                        if (compareTimeString(newStartTime, endTime)) setEndTime("")
                    }}
                >
                    {startTimeOptions.map((option) => (
                        <MenuItem key={option.value} value={option.value}>
                            {option.label}
                        </MenuItem>
                    ))}
                </Select>
            </div>
            <span className="mt-6">{pageT("to")}</span>
            <div>
                <h6 className="mb-1 ml-1">{pageT("endTime")}</h6>
                <Select
                    sx={{ minWidth: 90 }}
                    className="react-datepicker__input-container"
                    id="outlined-basic"
                    variant="outlined"
                    size="medium"
                    defaultValue={""}
                    value={endTime}
                    onChange={(e) => setEndTime(e.target.value)}
                >
                    {endTimeOptions.map((option) => (
                        <MenuItem key={option.value} value={option.value}>
                            {option.label}
                        </MenuItem>
                    ))}
                </Select>
            </div>
            <div>
                <h6 className="mb-1 ml-1">{pageT("belowBaseline")}</h6>
                <TextField
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