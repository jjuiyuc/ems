import { Fragment as Frag, useEffect, useState, useRef } from "react"
import { Button, Box, FormControl, Stack, InputLabel, Select, MenuItem, TextField } from "@mui/material"
import { useTranslation } from "react-multi-lang"

import Clock from "./Clock"
import TimeRangePicker from "./TimeRangePicker"

import { ReactComponent as AddIcon } from "../assets/icons/add.svg"
import { ReactComponent as DeleteIcon } from "../assets/icons/delete.svg"
import { ReactComponent as TimerIcon } from "../assets/icons/timer.svg"

const maxLength = 4
const maxPolicyCount = 5

const defaultPolicyConfig = {
    onPeak: {
        name: "onPeak",
        extensible: true,
        nameEditable: false,
    },
    midPeak: {
        name: "midPeak",
        extensible: true,
        nameEditable: false,
    },
    offPeak: {
        name: "offPeak",
        extensible: true,
        nameEditable: false,
    },
    superOffPeak: {
        name: "superOffPeak",
        extensible: true,
        nameEditable: false,
    },
}

const defaultPolicyPrice = {
    onPeak: [
        { startTime: null, endTime: null, basicPrice: "", rate: "" }
    ],
    midPeak: [
        { startTime: null, endTime: null, basicPrice: "47", rate: "" }
    ],
    offPeak: [
        { startTime: null, endTime: null, basicPrice: "47", rate: "3" }
    ],
    superOffPeak: [
        { startTime: null, endTime: null, basicPrice: "47", rate: "4" }
    ],
}
export default function TimeOfUseCard(props) {
    const { data } = props
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("settings." + string, params),
        dayTabs = ["weekdays", "saturday", "sundayHoliday"]

    const categories = [
        { label: pageT("summerTariff"), value: pageT("summerTariff") },
        { label: pageT("nonSummerTariff"), value: pageT("nonSummerTariff") }
    ]
    const customCount = useRef(1)
    const
        [dayTab, setDayTab] = useState(dayTabs[0]),
        [policyConfig, setPolicyConfig] = useState(defaultPolicyConfig),
        [policyPrice, setPolicyPrice] = useState(defaultPolicyPrice),
        [tariff, setTariff] = useState("")

    const
        handleChange = (e) => {
            setTariff(e.target.value)
        },
        changePolicyConfig = (e) => {
            const newPolicyConfig = {
                ...policyConfig,
                [policy]: {
                    ...policyConfig[policy],
                    tempName: e.target.value
                }
            }
            setPolicyConfig(newPolicyConfig)
        }


    console.log(Object.keys(policyConfig))

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
        <div className="flex items-center mt-12">
            <div className="pr-6">
                <Box sx={{ minWidth: 200 }}>
                    <FormControl fullWidth>
                        <InputLabel id="demo-simple-select-label">{pageT("tariff")}</InputLabel>
                        <Select
                            labelId="demo-simple-select-label"
                            id="demo-simple-select"
                            value={tariff}
                            label="Tariff"
                            onChange={handleChange}>
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
        <div className="flex items-start mt-12">
            <Clock size={{ height: "auto", width: "clamp(12rem,24vw,27.5rem)", "aspect-ratio": "1 / 1" }} dataset={data} id="touClock" />
            <div className="mb-12 mt-4">
                {Object.keys(policyConfig).map((policy) => {
                    const priceGroup = policyPrice[policy]
                    return (
                        <div className="mb-12 ml-12" key={policy}>
                            <div className="flex items-center text-white mb-4">
                                <div
                                    className="bg-blue-main h-2 rounded-full mr-3 w-2" />
                                {policyConfig[policy].nameEditable ?
                                    <TextField
                                        className=""
                                        id="outlined-basic"
                                        variant="outlined"
                                        value={policyConfig[policy].tempName}
                                        onChange={changePolicyConfig}
                                    /> :
                                    <h5 className="font-bold">{policyConfig[policy].name}</h5>}
                            </div>
                            {priceGroup.map(({ startTime, endTime, basicPrice, rate }, index) => {
                                return (
                                    <>
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
                                                            : row)
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
                                    </>
                                )
                            })}
                            {policyConfig[policy].extensible && priceGroup.length < maxLength ?
                                <div className="flex ml-4 mt-4">
                                    <AddIcon className="w-4 h-4 mt-0.5" />
                                    <button
                                        className="ml-1"
                                        onClick={() => {
                                            const newPolicyPrice = {
                                                ...policyPrice,
                                                [policy]: [
                                                    ...priceGroup,
                                                    { startTime: null, endTime: null, basicPrice: null, rate: null }
                                                ]
                                            }
                                            setPolicyPrice(newPolicyPrice)
                                        }}>
                                        {pageT("addTimeRange")}
                                    </button>
                                </div> : null}
                        </div>)
                })}
                {Object.keys(policyConfig).length < maxPolicyCount &&
                    <div className="ml-12">
                        <Button
                            onClick={() => {
                                const newKey = `custom${customCount.current}`
                                const newName = `Rate Period ${customCount.current}`
                                customCount.current++
                                const newPolicyConfig = {
                                    ...policyConfig,
                                    [newKey]: {
                                        name: newName,
                                        extensible: true,
                                        nameEditable: true,
                                        tempName: newName
                                    }
                                }
                                setPolicyConfig(newPolicyConfig)

                                const newPolicyPrice = {
                                    ...policyPrice,
                                    [newKey]: [{ startTime: null, endTime: null, basicPrice: "", rate: "" }]
                                }
                                setPolicyPrice(newPolicyPrice)
                            }}
                            key={"s-b-"}
                            radius="pill"
                            variant="outlined"
                            color="brand">
                            {pageT("addRatePeriod")}
                        </Button>
                    </div>}
            </div>
        </div>
    </div>
}