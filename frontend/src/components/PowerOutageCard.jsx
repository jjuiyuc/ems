import { Fragment as Frag, useEffect, useState, useRef } from "react"
import { Button, Box, FormControl, Stack, InputLabel, Select, MenuItem, TextField } from "@mui/material"
import { useTranslation } from "react-multi-lang"

import DatePeriodPicker from "./DatePeriodPicker"

import { ReactComponent as AddIcon } from "../assets/icons/add.svg"
import { ReactComponent as DeleteIcon } from "../assets/icons/delete.svg"
import { ReactComponent as PowerOutageIcon } from "../assets/icons/power_outage.svg"

const maxLength = 2

const defaultPolicyConfig = {
    preNotifiedOutagePeriod: {
        name: "preNotifiedOutagePeriod",
        tempName: "preNotifiedOutagePeriod",
        extensible: true,
        deletable: false
    }
}
const defaultPolicyTime = {
    preNotifiedOutagePeriod: [
        { startTime: "", endTime: "" },
    ]
}
export default function PowerOutageCard(props) {
    const { data } = props
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("settings." + string, params)

    const
        [policyConfig, setPolicyConfig] = useState(defaultPolicyConfig),
        [policyTime, setPolicyTime] = useState(defaultPolicyTime),
        [startDate, setStartDate] = useState(null),
        [endDate, setEndDate] = useState(null)

    const
        handleChange = (e) => {
            setTariff(e.target.value)
        }
    // console.log(Object.keys(policyConfig))

    return <div className="card mb-8">
        <div className="flex justify-between sm:col-span-2 items-center">
            <div className="flex items-center">
                <div
                    className="bg-gray-400-opacity-20 grid h-12 w-12
                            place-items-center rounded-full">
                    <PowerOutageIcon className="text-gray-400 w-10 h-10" />
                </div>
                <h2 className="font-bold ml-4">{pageT("powerOutage")}</h2>
            </div>
            <Button
                // onClick={() => setTab(t)}
                key={"s-b-"}
                radius="pill"
                variant="contained">
                {pageT("save")}
            </Button>
        </div>
        <div className="flex items-start mt-12">
            <div className="mb-12 mt-4">
                {Object.keys(policyConfig).map((policy) => {
                    const timeGroup = policyTime[policy]
                    return (
                        <div className="mb-12 ml-12" key={policy}>
                            <div className="flex items-center text-white mb-4">
                                {/* <div className="bg-blue-main h-2 rounded-full mr-3 w-2" /> */}
                                <h5 className="font-bold">{pageT(policyConfig[policy].name)}</h5>
                            </div>
                            {timeGroup.map(({ startDate, endDate }, index) => {
                                return (
                                    <>
                                        <div className="time-range-picker grid
                                        grid-cols-settings-input gap-x-4 items-center mt-4">
                                            <DatePeriodPicker
                                                key={index}
                                                startDate={startDate}
                                                endDate={endDate}
                                                setStartDate={(date) => {
                                                    const newPolicyTime = {
                                                        ...policyTime,
                                                        [policy]: timeGroup.map((row, i) =>
                                                            i === index
                                                                ? { ...row, startDate: date }
                                                                : row)
                                                    }
                                                    setPolicyTime(newPolicyTime)
                                                }}
                                                setEndDate={(date) => {
                                                    const newPolicyTime = {
                                                        ...policyTime,
                                                        [policy]: timeGroup.map((row, i) =>
                                                            i === index
                                                                ? { ...row, endTime: date }
                                                                : row)
                                                    }
                                                    setPolicyTime(newPolicyTime)
                                                }}
                                            />
                                            {index ?
                                                <div className="ml-2 mt-4 h-4 w-4 flex cursor-pointer">
                                                    <DeleteIcon
                                                        onClick={() => {
                                                            const newPolicyTime = {
                                                                ...policyTime,
                                                                [policy]: timeGroup.filter((_, i) => i !== index)
                                                            }
                                                            setPolicyTime(newPolicyTime)
                                                        }}
                                                    />
                                                </div> : <div></div>}
                                        </div>
                                    </>
                                )
                            })}
                            {policyConfig[policy].extensible && timeGroup.length < maxLength ?
                                <div className="flex ml-4 mt-4">
                                    <AddIcon className="w-4 h-4 mt-0.5" />
                                    <button
                                        className="ml-1"
                                        onClick={() => {
                                            const newPolicyTime = {
                                                ...policyTime,
                                                [policy]: [
                                                    ...timeGroup,
                                                    { startDate: "", endDate: "" }
                                                ]
                                            }
                                            setPolicyTime(newPolicyTime)
                                        }}>
                                        {pageT("addTimeRange")}
                                    </button>
                                </div> : null}
                        </div>)
                })}
            </div>
        </div>
    </div>
}