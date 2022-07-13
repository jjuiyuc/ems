import { Button, Stack, TextField, Box } from "@mui/material"
import { DateRangePickerDay } from '@mui/x-date-pickers-pro'
import { LocalizationProvider } from '@mui/x-date-pickers-pro'
import { AdapterDateFns } from '@mui/x-date-pickers-pro/AdapterDateFns'
import { StaticDateRangePicker } from '@mui/x-date-pickers-pro/StaticDateRangePicker'

import { Fragment as Frag, useEffect, useRef, useState } from "react"
import moment from "moment"
import { useTranslation } from "react-multi-lang"

import AnalysisCard from "../components/AnalysisCard"
import variables from "../configs/variables"

const { colors } = variables

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
    const [value, setValue] = useState([null, null])
    const
        hours24 = Array.from(new Array(24).keys()),
        lineChartDateLabels = hours24.map(n => {
            const time = moment().hour(n).minute(0).second(0)

            return time.format("hh A")
        }),
        currentHour = moment().hour(),
        lineChartDataArray = hours24.filter(v => v <= currentHour).map(() =>
            Math.floor(Math.random() * (60 - 40 + 1) + 40))

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
                { kwh: 25, percentage: 41, type: "batteryDischarge" },
            ],
            kwh: 60
        }),
        [prices, setPrices]
            = useState({ onPeak: 0, midPeak: 0, offPeak: 0, superOffPeak: 0 }),

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
                {tab === "custom"
                    ? <>
                        {/* <DateRangePickerDay /> */}
                        <LocalizationProvider dateAdapter={AdapterDateFns}>
                            <StaticDateRangePicker
                                displayStaticWrapperAs="desktop"
                                value={value}
                                onChange={(newValue) => {
                                    setValue(newValue);
                                }}
                                renderInput={(startProps, endProps) => (
                                    <>
                                        <TextField {...startProps} />
                                        <Box sx={{ mx: 2 }}> to </Box>
                                        <TextField {...endProps} />
                                    </>
                                )}
                            />
                        </LocalizationProvider>
                    </>
                    : null}
            </Stack>
        </div>
        <div className="gap-8 grid md:grid-cols-2 items-start">
            <AnalysisCard data={totalEnergySources} title={analysisCardTitle("totalEnergySources")} />
            <AnalysisCard data={energyDestinations} title={analysisCardTitle("energyDestinations")} />
        </div>
    </>
}