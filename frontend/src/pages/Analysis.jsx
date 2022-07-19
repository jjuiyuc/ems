import { Button, Stack, TextField, Box } from "@mui/material"
// import { DateRangePickerDay, LocalizationProvider } from '@mui/x-date-pickers-pro'
// import { AdapterDateFns } from '@mui/x-date-pickers-pro/AdapterDateFns'
// import { StaticDateRangePicker } from '@mui/x-date-pickers-pro/StaticDateRangePicker'
// import { DateRangePicker } from "materialui-daterange-picker"
import { CalendarToday, Today } from "@mui/icons-material"
import { DateRangePicker } from "materialui-daterange-picker"

import { Fragment as Frag, useEffect, useRef, useState } from "react"
import moment from "moment"
import { useTranslation } from "react-multi-lang"

import AnalysisCard from "../components/AnalysisCard"
import variables from "../configs/variables"
import "../assets/css/datePicker.scss"

const { colors } = variables

const dateFormat = "YYYY-MM-DD"

export default function Analysis() {

    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        errorT = string => t("error." + string),
        formT = string => t("form." + string),
        pageT = (string, params) => t("analysis." + string, params)

    const analysisCardTitle = () => {
        <>
            <span className="inline-block mr-1">
                {pageT("totalEnergySources")}
            </span>
            <span className="inline-block mr-1">
                {pageT("energyDestinations")}
            </span>

        </>
    }

    const [open, setOpen] = useState(false);
    const [dateRange, setDateRange] = useState({});

    const toggle = () => setOpen(!open);

    const
        [totalEnergySources, setTotalEnergySources] = useState({
            types: [
                { kwh: 7.5, percentage: 15, type: "directSolarSupply" },
                { kwh: 30, percentage: 60, type: "importFromGrid" },
                { kwh: 12.5, percentage: 25, type: "batteryDischarge" },
            ],
            kwh: 50
        }),

        [energyDestinations, setEnergyDestinations] = useState({
            types: [
                { kwh: 10, percentage: 18, type: "load" },
                { kwh: 25, percentage: 41, type: "exportFromGrid" },
                { kwh: 25, percentage: 41, type: "chargeToBattery" },
            ],
            kwh: 60
        }),

        [tab, setTab] = useState("days")





    return <>
        <div className="page-header">
            <h1>{pageT("analysis")}</h1>
            <Stack direction="row" justifyContent="flex-end" spacing={1.5}>
                <Button
                    onClick={() => setTab("days")}
                    filter={tab === "days" ? "selected" : ""}
                    radius="pill"
                    variant="contained">
                    {pageT("days")}
                </Button>
                <Button
                    onClick={() => setTab("weeks")}
                    filter={tab === "weeks" ? "selected" : ""}
                    radius="pill"
                    variant="contained">
                    {pageT("weeks")}
                </Button>
                <Button
                    onClick={() => setTab("month")}
                    filter={tab === "month" ? "selected" : ""}
                    radius="pill"
                    variant="contained">
                    {pageT("month")}
                </Button>
                <Button
                    onClick={() => setTab("year")}
                    filter={tab === "year" ? "selected" : ""}
                    radius="pill"
                    variant="contained">
                    {pageT("year")}
                </Button>
                <Button
                    onClick={() => setTab("custom")}
                    filter={tab === "custom" ? "selected" : ""}
                    radius="pill"
                    variant="contained">
                    {pageT("custom")}
                </Button>
                {/* {tab === "custom" */}
                {/* ?  */}
                {/* : null} */}
            </Stack>
        </div>
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
                    onChange={(range) => setDateRange(range)}
                    open={open}
                    toggle={toggle}
                    maxDate={Today}
                    wrapperClassName="date-range-picker"
                />
                {open &&
                    <div className="date-range-picker-wrapper flex">
                        <Button
                            onClick={() => setTab("cancel")}
                            className="date-range-button"
                            radius="pill"
                            variant="contained">
                            {pageT("cancel")}
                        </Button>
                        <Button
                            onClick={() => setTab("apply")}
                            className="date-range-button"
                            radius="pill"
                            variant="contained">
                            {pageT("apply")}
                        </Button>
                    </div>}
            </div>
        </div>
        <div className="gap-8 grid md:grid-cols-2 items-start">
            <AnalysisCard data={totalEnergySources} title={analysisCardTitle("totalEnergySources")} />
            <AnalysisCard data={energyDestinations} title={analysisCardTitle("energyDestinations")} />
        </div>
    </>
}