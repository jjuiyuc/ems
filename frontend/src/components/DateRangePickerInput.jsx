import { TextField } from "@mui/material"
import { CalendarToday } from "@mui/icons-material"
import { DateRangePicker } from "materialui-daterange-picker"

import { Fragment as useState, useMemo } from "react"
import moment from "moment"
import { useTranslation } from "react-multi-lang"

const dateFormat = "YYYY-MM-DD"
const defaultDate = {
    startDate: moment().format(dateFormat),
    endDate: moment().format(dateFormat)
}

export default function DateRangePickerInput(props) {

    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("analysis." + string, params)

    const { dateRange, startDate, endDate, open, toggle } = props


    return <>
        <div className="flex justify-end mb-10 relative w-auto">
            <div className="flex items-center">
                <TextField
                    InputProps={{
                        endAdornment: <CalendarToday
                            className="text-gray-300" />
                    }}
                    label={pageT("startDate")}
                    onFocus={() => setOpen(true)}
                    style={{ marginBottom: 0 }}
                    type="text"
                    value={moment(dateRange.startDate).format(dateFormat)}
                    variant="outlined" />
                <span className="mx-4">{pageT("to")}</span>
                <TextField
                    InputProps={{
                        endAdornment: <CalendarToday
                            className="text-gray-300" />
                    }}
                    label={pageT("endDate")}
                    onFocus={() => setOpen(true)}
                    style={{ marginBottom: 0 }}
                    type="text"
                    value={moment(dateRange.endDate).format(dateFormat)}
                    variant="outlined" />
            </div>
            <div className="absolute mt-2 top-full">
                <DateRangePicker
                    // onChange={(range) => {
                    //     setTempDateRange(range)
                    // }}
                    open={open}
                    toggle={toggle}
                    maxDate={new Date}
                    minDate={moment().subtract(1, "y")}
                    definedRanges={[]}
                    wrapperClassName="date-range-picker"
                />
            </div>
        </div>
    </>
}
