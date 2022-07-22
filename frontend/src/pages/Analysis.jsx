import { Button, Stack, TextField } from "@mui/material"
import { CalendarToday } from "@mui/icons-material"
import { DateRangePicker } from "materialui-daterange-picker"
import moment from "moment"
import { useState } from "react"
import { useTranslation } from "react-multi-lang"

import AnalysisCard from "../components/AnalysisCard"
import "../assets/css/datePicker.scss"

const dateFormat = "YYYY-MM-DD"
const defaultDate = {
    startDate: null,
    endDate: null
}

export default function Analysis() {
    const
        t = useTranslation(),
        pageT = (string, params) => t("analysis." + string, params)

    const
        [dateRange, setDateRange] = useState(defaultDate),
        [energyDestinations, setEnergyDestinations] = useState({
            types: [
                { kwh: 10, percentage: 18, type: "load" },
                { kwh: 25, percentage: 41, type: "exportFromGrid" },
                { kwh: 25, percentage: 41, type: "chargeToBattery" },
            ],
            kwh: 60
        }),
        [open, setOpen] = useState(false),
        [tab, setTab] = useState("days"),
        [totalEnergySources, setTotalEnergySources] = useState({
            types: [
                { kwh: 7.5, percentage: 15, type: "directSolarSupply" },
                { kwh: 30, percentage: 60, type: "importFromGrid" },
                { kwh: 12.5, percentage: 25, type: "batteryDischarge" },
            ],
            kwh: 50
        })

    const toggle = () => setOpen(!open)

    const
        endDate = dateRange.endDate
            ? moment(dateRange.endDate).format(dateFormat)
            : "",
        startDate = dateRange.startDate
            ? moment(dateRange.startDate).format(dateFormat)
            : "",
        tabs = ["days", "weeks", "month", "year", "custom"]

    return <>
        <div className="page-header">
            <h1>{pageT("analysis")}</h1>
            <Stack direction="row" justifyContent="flex-end" spacing={1.5}>
            {tabs.map((t, i) =>
                <Button
                    onClick={() => setTab(t)}
                    filter={tab === t ? "selected" : ""}
                    key={"a-t-" + i}
                    radius="pill"
                    variant="contained">
                    {pageT(t)}
                </Button>)}
            </Stack>
        </div>
    {tab === "custom"
        ? <div className="flex justify-end mb-10 relative w-auto">
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
                    value={startDate}
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
                    value={endDate}
                    variant="outlined" />
            </div>
            <div className="absolute mt-2 top-full">
                <DateRangePicker
                    onChange={range => setDateRange(range)}
                    open={open}
                    toggle={toggle}
                    maxDate={new Date()}
                    minDate={moment().subtract(1, "y")}
                    definedRanges={[]}
                    wrapperClassName="date-range-picker"
                />
            </div>
        </div>
        : null}
        <div className="gap-8 grid md:grid-cols-2 items-start">
            <AnalysisCard
                data={totalEnergySources}
                title={pageT("totalEnergySources")} />
            <AnalysisCard
                data={energyDestinations}
                title={pageT("energyDestinations")} />
        </div>
    </>
}