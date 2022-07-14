import { Button, Stack, TextField, Box } from "@mui/material"
// import { DateRangePickerDay, LocalizationProvider } from '@mui/x-date-pickers-pro'
// import { AdapterDateFns } from '@mui/x-date-pickers-pro/AdapterDateFns'
// import { StaticDateRangePicker } from '@mui/x-date-pickers-pro/StaticDateRangePicker'
// import { DateRangePicker } from "materialui-daterange-picker"
import { CalendarToday } from "@mui/icons-material"
import { DateRangePicker } from "materialui-daterange-picker"

import { Fragment as Frag, useEffect, useRef, useState } from "react"
import moment from "moment"
import { useTranslation } from "react-multi-lang"

import AnalysisCard from "../components/AnalysisCard"
import BarChart from "../components/BarChart"
import variables from "../configs/variables"

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
    const [value, setValue] = useState([null, null])

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
        [prices, setPrices]
            = useState({ onPeak: 0, midPeak: 0, offPeak: 0, superOffPeak: 0 }),

        [tab, setTab] = useState("days")

    const
        hours24 = Array.from(new Array(24).keys()),
        BarChartDateLabels = hours24.map(n => {
            const time = moment().hour(n).minute(0).second(0)

            return time.format("hh A")
        }),
        currentHour = moment().hour(),
        BarChartDataArray = hours24.filter(v => v <= currentHour).map(() =>
            Math.floor(Math.random() * (60 - 40 + 1) + 40))


    const [barChartData, setBarChartData] = useState({
        datasets: [{
            backgroundColor: "#12c9c9",
            borderColor: "#12c9c9",
            borderWidth: 1,
            data: BarChartDataArray,
            fill: {
                above: "rgba(18, 201, 201, .2)",
                target: "origin"
            },
            hoverRadius: 3,
            pointBorderColor: "rgba(18, 201, 201, .2)",
            pointHoverBorderWidth: 6,
            pointBorderWidth: 0,
            radius: 3,
            tension: 0
        }],
        labels: BarChartDateLabels,
        tooltipCallbacks: {
            label: item => `${item.parsed.y}%`,
            labelPointStyle: context => {
                const
                    color = context.dataset.backgroundColor
                        .replace("#", "%23"),
                    image = new Image(8, 8)

                image.className = "test"
                image.src = `data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' height='8' width='8'%3E%3Ccircle cx='4' cy='4' r ='4' fill='${color}' /%3E%3C/svg%3E`

                return { pointStyle: image }
            }
        },
        tickCallback: function (val, index) {
            return val + commonT("kwh")
        }
    })



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

                {/* <LocalizationProvider dateAdapter={AdapterDateFns}>
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
                    </LocalizationProvider> */}
                {/* : null} */}
            </Stack>
        </div>
        <div className="flex justify-end mb-10 relative w-auto">
            <div className="flex items-center">
                <TextField
                    InputProps={{endAdornment: <CalendarToday
                                                className="text-gray-300" />}}
                    label={pageT("startDate")}
                    onFocus={() => setOpen(true)}
                    style={{marginBottom: 0}}
                    type="text"
                    value={moment(dateRange.startDate).format(dateFormat)}
                    variant="outlined" />
                <span className="mx-4">{pageT("to")}</span>
                <TextField
                    InputProps={{endAdornment: <CalendarToday
                                                className="text-gray-300" />}}
                    label={pageT("endDate")}
                    onFocus={() => setOpen(true)}
                    style={{marginBottom: 0}}
                    type="text"
                    value={moment(dateRange.endDate).format(dateFormat)}
                    variant="outlined" />
            </div>
            <div className="absolute mt-2 top-full">
                <DateRangePicker
                    onChange={(range) => setDateRange(range)}
                    open={open}
                    toggle={toggle} />
            </div>
        </div>
        <div className="gap-8 grid md:grid-cols-2 items-start">
            <AnalysisCard data={totalEnergySources} title={analysisCardTitle("totalEnergySources")} />
            <AnalysisCard data={energyDestinations} title={analysisCardTitle("energyDestinations")} />
        </div>
    </>
}