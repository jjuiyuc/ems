import { connect } from "react-redux"
import { ToggleButtonGroup, ToggleButton } from "@mui/material"
import ReportProblemIcon from "@mui/icons-material/ReportProblem"
import { useState, useEffect } from "react"
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
const mapState = state => ({ gatewayID: state.gateways.active.gatewayID })

export default connect(mapState)(function Economics(props) {
    const
        t = useTranslation(),
        pageT = (string, params) => t("economics." + string, params)
    const
        [formats, setFormats] = useState([]),
        [infoError, setInfoError] = useState(""),
        [infoLoading, setInfoLoading] = useState(false),
        [preUbiikThisMonth, setPreUbiikThisMonth] = useState(0),
        [postUbiikThisMonth, setPostUbiikThisMonth] = useState(0),
        [preUbiikLastMonth, setPreUbiikLastMonth] = useState(0),
        [postUbiikLastMonth, setPostUbiikLastMonth] = useState(0),
        [preUbiikSameMonthLastYear, setPreUbiikSameMonthLastYear] = useState(0),
        [postUbiikSameMonthLastYear, setPostUbiikSameMonthLastYear] = useState(0),
        [lineChartCosts, setLineChartCosts] = useState(null),
        [barChartSaved, setBarChartSaved] = useState(null)

    const handleFormat = (event, newFormats) => {
        setFormats(newFormats)
    }
    const chartCostComparisonSet = ({ formats, data, labels }) => {
        const lastMonth = formats?.includes("lastMonth")
            ? ([{
                backgroundColor: colors.purple["main"],
                borderColor: colors.purple["main"],
                borderDash: [5, 5],
                data: data?.preUbiikLastMonth || [],
                fill: {
                    above: colors.purple["main-opacity-10"],
                    target: "origin"
                },
                id: "preUbiikLastMonth",
                borderWidth: 1,
                pointBorderColor: colors.purple["main-opacity-10"],
                hoverRadius: 2,
                pointHoverBorderWidth: 2,
                radius: 2,
                label: pageT("preUbiik") + " - " + pageT("lastMonth")
            },
            {
                backgroundColor: colors.purple.main,
                borderColor: colors.purple.main,
                data: data?.postUbiikLastMonth || [],
                fill: {
                    above: colors.purple["main-opacity-20"],
                    target: "origin"
                },
                id: "postUbiikLastMonth",
                borderWidth: 2,
                pointBorderColor: colors.purple["main-opacity-20"],
                hoverRadius: 2,
                pointHoverBorderWidth: 2,
                radius: 2,
                label: pageT("postUbiik") + " - " + pageT("lastMonth")
            }])
            : []
        const sameDayLastYear = formats?.includes("sameDayLastYear")
            ? ([{
                backgroundColor: colors.yellow["main"],
                borderColor: colors.yellow["main"],
                borderDash: [5, 5],
                data: data?.preUbiikSameMonthLastYear || [],
                fill: {
                    above: colors.yellow["main-opacity-10"],
                    target: "origin"
                },
                id: "preUbiikSameMonthLastYear",
                borderWidth: 2,
                pointBorderColor: colors.yellow["main-opacity-10"],
                hoverRadius: 2,
                pointHoverBorderWidth: 2,
                radius: 2,
                label: pageT("preUbiik") + " - " + pageT("sameDayLastYear")
            },
            {
                backgroundColor: colors.yellow.main,
                borderColor: colors.yellow.main,
                data: data?.postUbiikSameMonthLastYear || [],
                fill: {
                    above: colors.yellow["main-opacity-20"],
                    target: "origin"
                },
                id: "postUbiikSameMonthLastYear",
                borderWidth: 2,
                pointBorderColor: colors.yellow["main-opacity-20"],
                hoverRadius: 2,
                pointHoverBorderWidth: 2,
                radius: 2,
                label: pageT("postUbiik") + " - " + pageT("sameDayLastYear")
            }])
            : []
        return ({
            datasets: [
                {
                    backgroundColor: colors.blue["main"],
                    borderColor: colors.blue["main"],
                    borderDash: [5, 5],
                    data: data?.preUbiikThisMonth || [],
                    fill: {
                        above: colors.blue["main-opacity-10"],
                        target: "origin"
                    },
                    id: "preUbiikThisMonth",
                    borderWidth: 2,
                    pointBorderColor: colors.blue["main-opacity-10"],
                    hoverRadius: 2,
                    pointHoverBorderWidth: 2,
                    radius: 2,
                    label: pageT("preUbiik") + " - " + pageT("today")
                },
                {
                    backgroundColor: colors.blue.main,
                    borderColor: colors.blue.main,
                    data: data?.postUbiikThisMonth || [],
                    fill: {
                        above: colors.blue["main-opacity-20"],
                        target: "origin"
                    },
                    id: "postUbiikThisMonth",
                    borderWidth: 2,
                    pointBorderColor: colors.blue["main-opacity-20"],
                    hoverRadius: 2,
                    pointHoverBorderWidth: 2,
                    radius: 2,
                    label: pageT("postUbiik") + " - " + pageT("today")
                },
                ...lastMonth,
                ...sameDayLastYear
            ],
            labels,
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
    const chartSavedCostSet = ({ formats, data, labels }) => {
        const lastMonth = formats?.includes("lastMonth")
            ? ([{
                backgroundColor: colors.purple["main"],
                borderColor: colors.purple["main"],
                data: data?.savedLastMonth || [],
                id: "savedLastMonth",
                borderWidth: 1,
                label: pageT("lastMonth")
            }])
            : []
        const sameDayLastYear = formats?.includes("sameDayLastYear")
            ? ([{
                backgroundColor: colors.yellow["main"],
                borderColor: colors.yellow.main,
                data: data?.savedTheSameMonthLastYear || [],
                id: "savedTheSameMonthLastYear",
                borderWidth: 1,
                label: pageT("sameDayLastYear")
            }])
            : []
        return ({
            datasets: [
                {
                    backgroundColor: colors.blue.main,
                    borderColor: colors.blue["main"],
                    data: data?.savedThisMonth || [],
                    id: "savedThisMonth",
                    borderWidth: 1,
                    label: pageT("today")
                },
                ...lastMonth,
                ...sameDayLastYear
            ],
            labels,
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
    const urlPrefix = `/api/${props.gatewayID}/devices`
    const
        callData = (startTime, endTime) => {
            apiCall({
                onComplete: () => setInfoLoading(false),
                onError: error => setInfoError(error),
                onStart: () => setInfoLoading(true),
                onSuccess: rawData => {
                    if (!rawData?.data) return
                    const { data } = rawData,
                        preAndPost = data.energyCosts

                    setPreUbiikThisMonth(preAndPost?.preUbiikThisMonth || 0)
                    setPostUbiikThisMonth(preAndPost?.postUbiikThisMonth || 0)
                    setPreUbiikLastMonth(preAndPost?.preUbiikLastMonth || 0)
                    setPostUbiikLastMonth(preAndPost?.postUbiikLastMonth || 0)
                    setPreUbiikSameMonthLastYear(preAndPost?.preUbiikTheSameMonthLastYear || 0)
                    setPostUbiikSameMonthLastYear(preAndPost?.postUbiikTheSameMonthLastYear || 0)

                    const
                        costs = data.energyDailyCosts,
                        { timestamps } = costs,
                        labels = timestamps.map(t => {
                            return moment(t * 1000).startOf("day")._d.getTime()
                        })
                    setLineChartCosts({
                        data: {
                            preUbiikThisMonth: costs?.preUbiikThisMonth,
                            postUbiikThisMonth: costs?.postUbiikThisMonth,
                            preUbiikLastMonth: costs?.preUbiikLastMonth,
                            postUbiikLastMonth: costs?.postUbiikLastMonth,
                            preUbiikSameMonthLastYear: costs?.preUbiikTheSameMonthLastYear,
                            postUbiikSameMonthLastYear: costs?.postUbiikTheSameMonthLastYear
                        },
                        labels
                    })
                    setBarChartSaved({
                        data: {
                            savedThisMonth: costs?.savedThisMonth,
                            savedLastMonth: costs?.savedLastMonth,
                            savedTheSameMonthLastYear: costs?.savedTheSameMonthLastYear
                        },
                        labels
                    })
                },
                url: `${urlPrefix}/tou/energy-cost?startTime=${startTime}&endTime=${endTime}`
            })
        }
    useEffect(() => {
        if (!props.gatewayID) return

        let startTime = moment().startOf("month").toISOString(),
            endTime = moment().startOf("day").toISOString()

        if (moment().get("date") == 1) {
            startTime = moment().subtract(1, "month").startOf("month").toISOString()
            endTime = moment().startOf("day").toISOString()
        }
        if (startTime && endTime) {
            callData(startTime, endTime)
        }
    }, [props.gatewayID])

    return <>
        <div className="page-header flex flex-wrap justify-between">
            <h1 className="mb-9">{pageT("economics")}</h1>
            <div className="flex flex-wrap">
                <ToggleButtonGroup
                    value={formats}
                    onChange={handleFormat}
                    size="large"
                    aria-label="text formatting">
                    <ToggleButton
                        sx={{ textTransform: "none", fontWeight: "700" }}
                        value="lastMonth"
                        aria-label="lastMonth"
                        color="primary">
                        {pageT("lastMonth")}
                    </ToggleButton>
                    <ToggleButton
                        sx={{ textTransform: "none", fontWeight: "700" }}
                        value="sameDayLastYear"
                        aria-label="sameDayLastYear"
                        color="primary">
                        {pageT("sameDayLastYear")}
                    </ToggleButton>
                </ToggleButtonGroup>
            </div>
        </div>
        <ErrorBox error={infoError} message={pageT("infoError")} />
        <div className="font-bold mt-4 mb-8 relative">
            <div className="lg:grid-cols-2 grid gap-5">
                <EconomicsCard
                    data={{
                        lastMonth: preUbiikLastMonth,
                        sameDayLastYear: preUbiikSameMonthLastYear,
                        thisMonth: preUbiikThisMonth,
                        type: "pre"
                    }}
                    tabs={formats}
                />
                <EconomicsCard
                    data={{
                        lastMonth: postUbiikLastMonth,
                        sameDayLastYear: postUbiikSameMonthLastYear,
                        thisMonth: postUbiikThisMonth,
                        type: "post"
                    }}
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
            <div className="flex justify-between">
                <h4 className="mb-9">{pageT("energyCostComparison")}</h4>
                <div className="bg-gray-700 p-4 mr-6 rounded-xl">
                    <p>{pageT("dottedLine") + pageT("preUbiik")}</p>
                    <p>{pageT("solidLine") + pageT("postUbiik")}</p>
                </div>
            </div>
            <div className="max-h-80vh h-160 relative w-full">
                <LineChart data={chartCostComparisonSet({
                    formats,
                    ...lineChartCosts
                })}
                    id="ecoCosts" />
            </div>
        </div>
        <div className="card mt-8">
            <h4 className="mb-4 lg:mb-0">{pageT("savedEnergyCost")}</h4>
            <div className="max-h-80vh h-160 mt-8 relative w-full">
                <BarChart
                    data={chartSavedCostSet({
                        formats,
                        ...barChartSaved
                    })}
                    id="ecoSavedCosts" />
            </div>
        </div>
        <div className="mt-8">
            <InfoReminder />
        </div>
    </>
})