import { Button, Autocomplete, Stack, TextField } from "@mui/material"
import { Fragment as Frag, useEffect, useState } from "react"
import { useTranslation } from "react-multi-lang"

import Clock from "./Clock"
import TimeRangePicker from "./TimeRangePicker"
import variables from "../configs/variables"

import { ReactComponent as TimerIcon } from "../assets/icons/timer.svg"

const { colors } = variables

export default function TimeOfUseCard(props) {
    const { data } = props
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("settings." + string, params),
        dayTabs = ["weekdays", "saturday", "sundayHoliday"]

    const
        [dayTab, setDayTab] = useState(dayTabs[0]),
        [prices, setPrices]
            = useState({ onPeak: 0, midPeak: 0, offPeak: 0, superOffPeak: 0 }),
        [startTime, setStartTime] = useState(null),
        [endTime, setEndTime] = useState(null)

    return <div className="card">
        <div className="flex justify-between sm:col-span-2 items-center">
            <div className="flex items-center">
                <div
                    className="bg-gray-400-opacity-20 grid h-12 w-12
                            place-items-center rounded-full">
                    <TimerIcon className="h-8 text-gray-400 w-8" />
                </div>
                <h2 className="font-bold ml-4">{pageT("timeOfUse")}</h2>
                <h6 className="border-solid rounded-lg border-gray-500 border
                    px-4 py-2 opacity-60 ml-4 mr-2">低壓-單向</h6>
                <h6 className="border-solid rounded-lg border-gray-500 border
                    px-4 py-2 opacity-60">兩段式</h6>
            </div>
            <Button
                // onClick={() => setTab(t)}
                key={"s-b-"}
                radius="pill"
                variant="contained">
                {pageT("save")}
            </Button>
        </div>
        <div className="flex items-center">
            <div className="pr-6">
                <Autocomplete
                    id="country-select-demo"
                    sx={{ width: 220 }}
                    //   options={}
                    // autoHighlight
                    //   getOptionLabel={(option) => option.label}
                    //   renderOption={(props, option) => (
                    //     <Box component="li" sx={{ '& > img': { mr: 2, flexShrink: 0 } }} {...props}>
                    //     </Box>
                    //   )}
                    renderInput={(params) => (
                        <TextField
                        />
                    )}
                />
            </div>
            <div className="flex-wrap lg:flex border-l border-gray-400 border-solid pl-6">
                <Stack direction="row" spacing={1.5}>
                    {dayTabs.map((t, i) =>
                        <Button
                            className="py-0.5"
                            color="gray"
                            onClick={() => setDayTab(t)}
                            filter={dayTab === t ? "selected" : ""}
                            key={"st-d" + i}
                            radius="pill"
                            variant="contained">
                            {pageT(t)}
                        </Button>)}
                </Stack>
            </div>
        </div>
        <div className="flex">
            <Clock dataset={data} id="touClock" />
            {/* <div className="grid grid-cols-3-auto
                                        items-center mx-8 my-4 text-white"> */}
            <div className="flex mb-12 mt-4">
                <div className="flex items-center mx-2.5 text-white">
                    <div
                        className="bg-blue-main h-2 rounded-full mr-3 w-2" />
                    <h5 className="font-bold">{pageT("onPeak")}</h5>
                    <TimeRangePicker
                        {...{ startTime, setStartTime, endTime, setEndTime }}
                    />
                </div>
            </div>
        </div>
    </div>
}