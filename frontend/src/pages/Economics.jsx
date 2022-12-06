import { connect } from "react-redux"
import { Button, Stack, ToggleButtonGroup, ToggleButton, Alert } from "@mui/material"
import ReportProblemIcon from "@mui/icons-material/ReportProblem"
import { useState } from "react"
import { useTranslation } from "react-multi-lang"
import moment from "moment"

import { apiCall } from "../utils/api"
import variables from "../configs/variables"

import AlertBox from "../components/AlertBox"
import BarChart from "../components/BarChart"
import EconomicsCard from "../components/EconomicsCard"
import LineChart from "../components/LineChart"
import InfoReminder from "../components/InfoReminder"
import Spinner from "../components/Spinner"

import { ReactComponent as EconomicsIcon } from "../assets/icons/economics.svg"

const { colors } = variables
const ErrorBox = ({ error, margin = "", message }) => error
    ? <AlertBox
        boxClass={`${margin} negative`}
        content={<>
            <span className="font-mono ml-2">{error}</span>
            <span className="ml-2">{message}</span>
        </>}
        icon={ReportProblemIcon}
        iconColor="negative-main" />
    : null
const LoadingBox = ({ loading }) => loading
    ? <div className="grid h-24 place-items-center"><Spinner /></div>
    : null

const fakeDataArray = amount => Array.from(new Array(amount).keys())
    .map(() => Math.floor(Math.random() * (40 - 10 + 1) + 10))

const
    currentHour = moment().hour(),
    hours24 = Array.from(new Array(24).keys()),
    lineChartDataArray = hours24.filter(v => v <= currentHour).map(() =>
        Math.floor(Math.random() * (60 - 40 + 1) + 40)),
    days = 7,
    sevenDays = Array.from(new Array(days).keys()).map(n =>
        moment().subtract(days - n, "d").startOf("day").toISOString()),
    fakeData1 = fakeDataArray(days),
    fakeData2 = fakeDataArray(days),
    fakeData3 = fakeDataArray(days),
    fakeData4 = fakeDataArray(days),
    fakeData5 = fakeDataArray(days),
    fakeData6 = fakeDataArray(days)

export default function Economics(props) {

    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("economics." + string, params)

    const
        [formats, setFormats] = useState([]),
        [infoError, setInfoError] = useState(""),
        [infoLoading, setInfoLoading] = useState(false),
        [total, setTotal] = useState(630),
        [preUbiik, setPreUbiik] = useState(0),
        [postUbiik, setPostUbiik] = useState(0),
        [lastMonth, setLastMonth] = useState(0),
        [sameMonthLastYear, setSameMonthLastYear] = useState(0)

    const handleFormat = (event, newFormats) => {
        setFormats(newFormats)
    }

    const chartCostComparisonSet = (formats) => {
        const lastMonth = formats?.includes("lastMonth")
            ? ([{
                backgroundColor: colors.purple["main"],
                borderColor: colors.purple["main"],
                borderDash: [5, 5],
                data: fakeData5,
                fill: {
                    above: colors.purple["main-opacity-10"],
                    target: "origin"
                },
                id: "preUbiik",
                hoverRadius: 0,
                pointHoverBorderWidth: 0,
                radius: 0,
                borderWidth: 1,
                label: pageT("preUbiik") + " - " + pageT("lastMonth")
            },
            {
                backgroundColor: colors.purple.main,
                borderColor: colors.purple.main,
                data: fakeData6,
                fill: {
                    above: colors.purple["main-opacity-20"],
                    target: "origin"
                },
                id: "postUbiik",
                hoverRadius: 0,
                borderWidth: 2,
                pointHoverBorderWidth: 0,
                radius: 0,
                label: pageT("postUbiik") + " - " + pageT("lastMonth")
            }])
            : []

        const sameMonthLastYear = formats?.includes("sameMonthLastYear")
            ? ([{
                backgroundColor: colors.yellow["main"],
                borderColor: colors.yellow["main"],
                borderDash: [5, 5],
                data: fakeData3,
                fill: {
                    above: colors.yellow["main-opacity-10"],
                    target: "origin"
                },
                id: "preUbiik",
                hoverRadius: 0,
                pointHoverBorderWidth: 0,
                radius: 0,
                borderWidth: 1,
                label: pageT("preUbiik") + " - " + pageT("sameMonthLastYear")
            },
            {
                backgroundColor: colors.yellow.main,
                borderColor: colors.yellow.main,
                data: fakeData4,
                fill: {
                    above: colors.yellow["main-opacity-20"],
                    target: "origin"
                },
                id: "postUbiik",
                hoverRadius: 0,
                borderWidth: 2,
                pointHoverBorderWidth: 0,
                radius: 0,
                label: pageT("postUbiik") + " - " + pageT("sameMonthLastYear")
            }])
            : []
        return ({
            datasets: [
                {
                    backgroundColor: colors.blue["main"],
                    borderColor: colors.blue["main"],
                    borderDash: [5, 5],
                    data: fakeData2,
                    fill: {
                        above: colors.blue["main-opacity-10"],
                        target: "origin"
                    },
                    id: "preUbiik",
                    hoverRadius: 0,
                    pointHoverBorderWidth: 0,
                    radius: 0,
                    borderWidth: 1,
                    label: pageT("preUbiik") + " - " + pageT("thisCalendarMonth")
                },
                {
                    backgroundColor: colors.blue.main,
                    borderColor: colors.blue.main,
                    data: fakeData1,
                    fill: {
                        above: colors.blue["main-opacity-20"],
                        target: "origin"
                    },
                    id: "postUbiik",
                    hoverRadius: 0,
                    borderWidth: 2,
                    pointHoverBorderWidth: 0,
                    radius: 0,
                    label: pageT("postUbiik") + " - " + pageT("thisCalendarMonth")
                },
                ...lastMonth,
                ...sameMonthLastYear
            ],
            labels: sevenDays,
            tickCallback: val => "$" + val,
            tooltipLabel: item =>
                item.dataset.label + " $" + item.parsed.y,
            x: {
                time: {
                    "displayFormats": {
                        "day": "MMM D"
                    },
                    "tooltipFormat": "MMM D",
                    "unit": "day"
                }
            }
        })
    }
    const chartSavedCostSet = (formats) => {
        const lastMonth = formats?.includes("lastMonth")
            ? ([{
                backgroundColor: colors.purple["main-opacity-20"],
                borderColor: colors.purple.main,
                data: fakeData1,
                id: "preUbiik",
                borderWidth: 1,
                label: pageT("preUbiik") + " - " + pageT("lastMonth")
            },
            {
                backgroundColor: colors.purple.main,
                data: fakeData2,
                id: "postUbiik",
                borderWidth: 1,
                label: pageT("postUbiik") + " - " + pageT("lastMonth")
            }])
            : []

        const sameMonthLastYear = formats?.includes("sameMonthLastYear")
            ? ([{
                backgroundColor: colors.yellow["main-opacity-20"],
                borderColor: colors.yellow.main,
                data: fakeData3,
                id: "preUbiik",
                borderWidth: 1,
                label: pageT("preUbiik") + " - " + pageT("sameMonthLastYear")
            },
            {
                backgroundColor: colors.yellow.main,
                data: fakeData4,
                id: "postUbiik",
                borderWidth: 1,
                label: pageT("postUbiik") + " - " + pageT("sameMonthLastYear")
            }])
            : []
        return ({
            datasets: [
                {
                    backgroundColor: colors.blue["main-opacity-20"],
                    borderColor: colors.blue["main"],
                    data: fakeData5,
                    id: "preUbiik",
                    borderWidth: 1,
                    label: pageT("preUbiik") + " - " + pageT("thisCalendarMonth")
                },
                {
                    backgroundColor: colors.blue.main,
                    data: fakeData6,
                    id: "postUbiik",
                    borderWidth: 1,
                    label: pageT("postUbiik") + " - " + pageT("thisCalendarMonth")
                },
                ...lastMonth,
                ...sameMonthLastYear
            ],
            labels: sevenDays,
            tickCallback: val => "$" + val,
            tooltipLabel: item =>
                item.dataset.label + " $" + item.parsed.y,
            x: {
                time: {
                    "displayFormats": {
                        "day": "MMM D"
                    },
                    "tooltipFormat": "MMM D",
                    "unit": "day"
                }
            }
        })
    }
    return <>
        <div className="page-header flex flex-wrap justify-between">
            <h1 className="mb-9">{pageT("economics")}</h1>
            <div className="flex flex-wrap">
                <ToggleButtonGroup
                    value={formats}
                    onChange={handleFormat}
                    size="large"
                    aria-label="text formatting">
                    <ToggleButton value="lastMonth" aria-label="lastMonth" color="primary">
                        {pageT("lastMonth")}
                    </ToggleButton>
                    <ToggleButton value="sameMonthLastYear" aria-label="sameMonthLastYear" color="primary">
                        {pageT("sameMonthLastYear")}
                    </ToggleButton>
                </ToggleButtonGroup>
            </div>
        </div>
        <div className="font-bold mt-4 mb-8 relative">
            <div className="lg:grid-cols-2 grid gap-5">
                <EconomicsCard
                    icon={EconomicsIcon}
                    title={pageT("preUbiik")}
                    subTitle={pageT("thisCalendarMonth")}
                    leftTitle={pageT("lastMonth")}
                    rightTitle={pageT("sameMonthLastYear")}
                    value={"$" + preUbiik}
                    leftValue={"$" + lastMonth}
                    rightValue={"$" + sameMonthLastYear}
                    tabs={formats}
                />
                <EconomicsCard
                    icon={EconomicsIcon}
                    title={pageT("postUbiik")}
                    subTitle={pageT("thisCalendarMonth")}
                    leftTitle={pageT("lastMonth")}
                    rightTitle={pageT("sameMonthLastYear")}
                    value={"$" + postUbiik}
                    leftValue={"$" + lastMonth}
                    rightValue={"$" + sameMonthLastYear}
                    tabs={formats}
                />
            </div>
            {infoLoading
                ? <div className="absolute bg-black-main-opacity-95 grid inset-0
                                place-items-center rounded-3xl">
                    <Spinner />
                </div>
                : null}
        </div>
        <div className="card chart mt-8 mb-8">
            <h4 className="mb-9">{pageT("energyCostComparison")}</h4>
            {/* <ErrorBox
                error={chargeVoltageError}
                message={pageT("chartError")} /> */}
            <div className="max-h-80vh h-160 relative w-full">
                <LineChart data={chartCostComparisonSet(formats)}
                    id="ecoCost" />

                {/* // data={chartChargeVoltageSet({
                    //     ...chargeVoltage,
                    //     unit: { charge: "%", voltage: " " + commonT("v") },
                    // })}
                    // id="erbChargeVoltage" /> */}
            </div>
            {/* <LoadingBox loading={chargeVoltageLoading} /> */}
        </div>

        <div className="card mt-8">
            <h4 className="mb-4 lg:mb-0">{pageT("savedEnergyCost")}</h4>
            <div className="max-h-80vh h-160 mt-8 relative w-full">
                <BarChart
                    data={chartSavedCostSet(formats)}
                    id="economicsBarChart" />
            </div>
        </div>
        <div className="mt-8">
            <InfoReminder />
        </div>
    </>
}