import { Fragment as Frag, useEffect, useState, useRef } from "react"
import { Button, Box, FormControl, Stack, InputLabel, Select, MenuItem, TextField } from "@mui/material"
import { useTranslation } from "react-multi-lang"
import moment from "moment"

import Clock from "./Clock"
import TimeRangePicker from "./TimeRangePicker"
import variables from "../configs/variables"

import { ReactComponent as AddIcon } from "../assets/icons/add.svg"
import { ReactComponent as DeleteIcon } from "../assets/icons/delete.svg"
import { ReactComponent as TimerIcon } from "../assets/icons/timer.svg"
import { apiCall } from "../utils/api"

const { colors } = variables

const maxLength = 4
const maxPolicyCount = 5

export default function TimeOfUseCard(props) {
    // const { data } = props
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("settings." + string, params),
        dayTabs = ["weekdays", "saturday", "sundayHoliday"]

    const defaultPolicyConfig = {
        onPeak: {
            name: pageT("onPeak"),
            tempName: "onPeak",
            extensible: true,
            nameEditable: false,
            deletable: false
        },
        midPeak: {
            name: pageT("midPeak"),
            tempName: "midPeak",
            extensible: true,
            nameEditable: false,
            deletable: false
        },
        offPeak: {
            name: pageT("offPeak"),
            tempName: "offPeak",
            extensible: true,
            nameEditable: false,
            deletable: false
        },
        superOffPeak: {
            name: pageT("superOffPeak"),
            tempName: "superOffPeak",
            extensible: true,
            nameEditable: false,
            deletable: false
        },
    }
    const defaultPolicyPrice = {
        onPeak: [
            { startTime: "", endTime: "", basicPrice: "", rate: "" },
        ],
        midPeak: [
            { startTime: "", endTime: "", basicPrice: "", rate: "" }
        ],
        offPeak: [
            { startTime: "", endTime: "", basicPrice: "", rate: "" }
        ],
        superOffPeak: [
            { startTime: "", endTime: "", basicPrice: "", rate: "" }
        ],
    }
    const categories = [
        { label: pageT("summerTariff"), value: pageT("summerTariff") },
        { label: pageT("nonSummerTariff"), value: pageT("nonSummerTariff") }
    ]
    const customCount = useRef(1)
    const
        [dayTab, setDayTab] = useState(dayTabs[0]),
        [clockDataset, setClockDataset] = useState({
            data: [], backgroundColor: []
        }),
        [policyConfig, setPolicyConfig] = useState(defaultPolicyConfig),
        [policyPrice, setPolicyPrice] = useState(defaultPolicyPrice),
        [tariff, setTariff] = useState("")
    const
        handleChange = (e) => {
            setTariff(e.target.value)
        }
    const getMoment = string => {
        const [hour, minute] = string.split(":")

        return moment().hour(parseInt(hour)).minute(parseInt(minute)).second(0)
    }

    const onSave = () => {
        apiCall()
    }

    useEffect(() => {
        const timeOfUse = Object
            .entries(policyPrice)
            .reduce((acc, cur) => {
                const [policy, priceArr] = cur
                const { name } = policyConfig[policy]
                const priceArrWithName = priceArr
                    .filter(({ startTime, endTime }) => startTime && endTime)
                    .map((obj) => ({ ...obj, policy, name }))
                return [...acc, ...priceArrWithName]
            }, [])
        console.log(timeOfUse)

        if (timeOfUse.length === 0) return
        const transformDataToGraphic = () => {
            const allTimeNodes = [
                ...new Set(timeOfUse.map(({ startTime, endTime }) =>
                    [startTime, endTime]).flat()), "00:00"]
                .sort()
            const sections = allTimeNodes.map((timeNode, i) => [timeNode, allTimeNodes[i + 1] || "24:00"])
            console.log(allTimeNodes, sections)

            const data = sections.map(([start, end]) => {
                const { overlappedNum, policies } = timeOfUse.reduce((acc, cur) => {
                    return cur.startTime <= start && cur.endTime > start
                        ? { overlappedNum: acc.overlappedNum + 1, policies: [...acc.policies, cur.policy] }
                        : acc
                }, { overlappedNum: 0, policies: [] })

                switch (overlappedNum) {
                    case 0:
                        return { startTime: start, endTime: end, condition: "empty", color: "transparent" }
                    case 1:
                        return { startTime: start, endTime: end, condition: "one", color: colors[policies[0]] }
                    default:
                        return { startTime: start, endTime: end, condition: "overlapped", color: "#eeeeee" }
                }
            })
            return data
        }
        const graphicData = transformDataToGraphic()
        const dataset = graphicData.reduce((acc, item) => {
            const
                { endTime, startTime, color } = item,
                duration = moment.duration(getMoment(endTime).diff(getMoment(startTime))).as("minutes")
            acc.data.push(duration)
            acc.backgroundColor.push(color)
            return acc
        }, { data: [], backgroundColor: [] })
        setClockDataset(dataset)
    }, [policyPrice])

    return <div className="card">
        <div className="flex justify-between sm:col-span-2 items-center">
            <div className="flex items-center">
                <div
                    className="bg-gray-400-opacity-20 grid h-12 w-12
                            place-items-center rounded-full">
                    <TimerIcon className="h-8 text-gray-400 w-8" />
                </div>
                <h2 className="font-bold ml-4">{pageT("timeOfUse")}</h2>
                {/* <h6 className="border-solid rounded-lg border-gray-500 border
                    px-4 py-2 opacity-60 ml-4 mr-2">低壓-單向</h6>
                <h6 className="border-solid rounded-lg border-gray-500 border
                    px-4 py-2 opacity-60">兩段式</h6> */}
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
            <Clock size={{
                height: "auto", width: "clamp(12rem,24vw,27.5rem)",
                aspectRatio: "1 / 1"
            }} dataset={clockDataset} id="touClock" />
            <div className="mb-12 mt-4">
                {Object.keys(policyConfig).map((policy, i) => {
                    const priceGroup = policyPrice[policy]
                    return (
                        <div className="mb-12 ml-12" key={policy + i}>
                            <div className="flex items-center text-white mb-4">
                                {policyConfig[policy].nameEditable
                                    ? <>
                                        <div className="bg-blue-main h-2 rounded-full mr-3 w-2" />
                                        <TextField
                                            className=""
                                            id="outlined-basic"
                                            variant="outlined"
                                            value={policyConfig[policy].tempName}
                                            onChange={(e) => {
                                                const newPolicyConfig = {
                                                    ...policyConfig,
                                                    [policy]: {
                                                        ...policyConfig[policy],
                                                        tempName: e.target.value
                                                    }
                                                }
                                                setPolicyConfig(newPolicyConfig)
                                            }}
                                        />
                                    </>
                                    : <>
                                        <div className="h-2 rounded-full mr-3 w-2"
                                            style={{ background: colors[policy] }} />
                                        <h5 className="font-bold">{policyConfig[policy].name}</h5>
                                    </>
                                }
                                {policyConfig[policy].deletable ?
                                    <div className="ml-2 mb-9 h-4 w-4 flex cursor-pointer">
                                        <DeleteIcon
                                            onClick={() => {
                                                const { [policy]: deleted, ...newPolicyConfig } = policyConfig
                                                setPolicyConfig(newPolicyConfig)
                                            }}
                                        />
                                    </div> : null}
                            </div>
                            {priceGroup.map(({ startTime, endTime, basicPrice, rate }, index) => {
                                const onChange = (key, value) => {
                                    setPolicyPrice((prevPolicyPrice) => {
                                        const prevPriceGroup = prevPolicyPrice[policy]
                                        return {
                                            ...prevPolicyPrice,
                                            [policy]: prevPriceGroup.map((row, i) =>
                                                i === index
                                                    ? { ...row, [key]: value }
                                                    : row)
                                        }
                                    })
                                }
                                return (
                                    <div key={`${policy}-${i}-${index}`} className="time-range-picker grid
                                        grid-cols-settings-input gap-x-4 items-center mt-4">
                                        <TimeRangePicker
                                            key={index}
                                            startTime={startTime}
                                            endTime={endTime}
                                            basicPrice={basicPrice}
                                            rate={rate}
                                            setStartTime={(time) => onChange("startTime", time)}
                                            setEndTime={(time) => onChange("endTime", time)}
                                            setBasicPrice={(price) => onChange("basicPrice", price)}
                                            setRate={(price) => onChange("rate", price)}
                                        />
                                        {index ?
                                            <div className="ml-2 mt-4 h-4 w-4 flex cursor-pointer">
                                                <DeleteIcon
                                                    onClick={() => {
                                                        const newPolicyPrice = {
                                                            ...policyPrice,
                                                            [policy]: priceGroup.filter((_, i) => i !== index)
                                                        }
                                                        setPolicyPrice(newPolicyPrice)
                                                    }}
                                                />
                                            </div> : <div></div>}
                                    </div>
                                )
                            })}
                            {policyConfig[policy].extensible && priceGroup.length < maxLength ?
                                <button
                                    className="flex ml-4 mt-4"
                                    onClick={() => {
                                        const newPolicyPrice = {
                                            ...policyPrice,
                                            [policy]: [
                                                ...priceGroup,
                                                { startTime: "", endTime: "", basicPrice: "", rate: "" }
                                            ]
                                        }
                                        setPolicyPrice(newPolicyPrice)
                                    }}>
                                    <AddIcon className="w-4 h-4 mt-0.7 mr-1" />
                                    {pageT("addTimeRange")}
                                </button> : null}
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
                                        tempName: newName,
                                        extensible: true,
                                        nameEditable: true,
                                        deletable: true
                                    }
                                }
                                setPolicyConfig(newPolicyConfig)

                                const newPolicyPrice = {
                                    ...policyPrice,
                                    [newKey]: [{ startTime: "", endTime: "", basicPrice: "", rate: "" }]
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