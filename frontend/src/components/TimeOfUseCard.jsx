import { Button, Box, FormControl, Stack, InputLabel, Select, MenuItem } from "@mui/material"
import { Fragment as Frag, useEffect, useState } from "react"
import { useTranslation } from "react-multi-lang"

import Clock from "./Clock"
import TimeRangePicker from "./TimeRangePicker"
import variables from "../configs/variables"

import { ReactComponent as TimerIcon } from "../assets/icons/timer.svg"
import { clamp } from "date-fns"

const { colors } = variables
const categories = [
    { label: 1, value: 1 },
    { label: 2, value: 2 },
    { label: 3, value: 3 },
    { label: 4, value: 4 },
    { label: 5, value: 5 }
]

const defaultPolicyConfig = {
    onPeak: {
        name: "onPeak",
        extensible: false,
    },
    midPeak: {
        name: "midPeak",
        extensible: true,

    },
    offPeak: {
        name: "offPeak",
        extensible: true,

    },
    superOffPeak: {
        name: "superOffPeak",
        extensible: false,
    },
}
const defaultPolicyPrice = {
    onPeak: [
        { startTime: null, endTime: null, basicPrice: null, rate: null }
    ],
    midPeak: [
        { startTime: null, endTime: null, basicPrice: 47, rate: 2 }
    ],
    offPeak: [
        { startTime: null, endTime: null, basicPrice: 47, rate: 3 }
    ],
    superOffPeak: [
        { startTime: null, endTime: null, basicPrice: 47, rate: 4 }
    ]

}
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
        [policyConfig, setPolicyConfig] = useState(defaultPolicyConfig),
        [policyPrice, setPolicyPrice] = useState(defaultPolicyPrice)

    const [age, setAge] = useState('')

    const handleChange = (event) => {
        setAge(event.target.value)
    }

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
                <Box sx={{ minWidth: 120 }}>
                    <FormControl fullWidth>
                        <InputLabel id="demo-simple-select-label">Age</InputLabel>
                        <Select
                            labelId="demo-simple-select-label"
                            id="demo-simple-select"
                            value={age}
                            label="Age"
                            onChange={handleChange}
                        >
                            {categories.map(({ label, value }, i) =>
                                <MenuItem key={i} value={value}>{label}</MenuItem>)}
                        </Select>
                    </FormControl>
                </Box>
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
        <div className="flex items-start">
            <Clock size={{ height: "auto", width: "clamp(12rem,24vw,27.5rem)", "aspect-ratio": "1 / 1" }} dataset={data} id="touClock" />

            <div className="mb-12 mt-4">
                {Object.keys(policyConfig).map((policy) => {
                    const priceGroup = policyPrice[policy]
                    return (
                        <div key={policy}>
                            <div className="flex items-center mx-2.5 text-white">
                                <div
                                    className="bg-blue-main h-2 rounded-full mr-3 w-2" />
                                <h5 className="font-bold">{pageT(policyConfig[policy].name)}</h5>
                            </div>
                            {priceGroup.map(({ startTime, endTime, basicPrice, rate }, index) => {
                                return (
                                    <TimeRangePicker
                                        key={index}
                                        startTime={startTime}
                                        endTime={endTime}
                                        basicPrice={basicPrice}
                                        rate={rate}
                                        setStartTime={(time) => {
                                            const newPolicyPrice = {
                                                ...policyPrice,
                                                [policy]: priceGroup.map((row, i) =>
                                                    i === index
                                                        ? { ...row, startTime: time }
                                                        : row
                                                )
                                            }
                                            setPolicyPrice(newPolicyPrice)
                                        }}
                                        setEndTime={(time) => {
                                            const newPolicyPrice = {
                                                ...policyPrice,
                                                [policy]: priceGroup.map((row, i) =>
                                                    i === index
                                                        ? { ...row, endTime: time }
                                                        : row)
                                            }
                                            setPolicyPrice(newPolicyPrice)
                                        }}
                                        setBasicPrice={(price) => {
                                            const newPolicyPrice = {
                                                ...policyPrice,
                                                [policy]: priceGroup.map((row, i) =>
                                                    i === index
                                                        ? { ...row, basicPrice: price }
                                                        : row)
                                            }
                                            setPolicyPrice(newPolicyPrice)
                                        }}
                                        setRate={(price) => {
                                            const newPolicyPrice = {
                                                ...policyPrice,
                                                [policy]: priceGroup.map((row, i) =>
                                                    i == index
                                                        ? { ...row, rate: price }
                                                        : row)
                                            }
                                            setPolicyPrice(newPolicyPrice)
                                        }}
                                    />
                                )
                            })}
                            <button
                                onClick={() => {
                                    const newPolicyPrice = {
                                        ...policyPrice,
                                        [policy]: [
                                            ...priceGroup,
                                            { startTime: null, endTime: null, basicPrice: null, rate: null }
                                        ]
                                    }
                                    setPolicyPrice(newPolicyPrice)
                                }}
                                disabled={!policyConfig[policy].extensible}
                            >
                                Add Time Range
                            </button>
                        </div>)
                })}
            </div>
        </div>
    </div>
}
