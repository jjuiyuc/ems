import { connect } from "react-redux"
import { Button, ToggleButtonGroup, ToggleButton } from "@mui/material"
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

const mapState = state => ({ gatewayID: state.gateways.active.gatewayID })

export default connect(mapState)(function Economics(props) {

    const
        t = useTranslation(),
        commonT = string => t("common." + string),
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
        [lineChartCostsError, setLineChartCostsError] = useState(""),
        [lineChartCostsLoading, setLineChartCostsLoading] = useState(false)

    const handleFormat = (event, newFormats) => {
        setFormats(newFormats)
    }
    const cardsData = {
        preUbiik: [
            {
                title: pageT("preUbiik"),
                value: "$" + `${preUbiikThisMonth}`
            },
            {
                title: pageT("lastMonth"),
                value: "$" + `${preUbiikLastMonth}`
            },
            {
                title: pageT("sameMonthLastYear"),
                value: "$" + `${preUbiikSameMonthLastYear}`
            }
        ],
        postUbiik: [
            {
                title: pageT("postUbiik"),
                value: "$" + `${postUbiikThisMonth}`
            },
            {
                title: pageT("lastMonth"),
                value: "$" + `${postUbiikLastMonth}`
            },
            {
                title: pageT("sameMonthLastYear"),
                value: "$" + `${postUbiikSameMonthLastYear}`
            }
        ]
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
                hoverRadius: 0,
                pointHoverBorderWidth: 0,
                radius: 0,
                borderWidth: 1,
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
                data: data?.preUbiikSameMonthLastYear || [],
                fill: {
                    above: colors.yellow["main-opacity-10"],
                    target: "origin"
                },
                id: "preUbiikSameMonthLastYear",
                hoverRadius: 0,
                pointHoverBorderWidth: 0,
                radius: 0,
                borderWidth: 1,
                label: pageT("preUbiik") + " - " + pageT("sameMonthLastYear")
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
                    data: data?.preUbiikThisMonth || [],
                    fill: {
                        above: colors.blue["main-opacity-10"],
                        target: "origin"
                    },
                    id: "preUbiikThisMonth",
                    hoverRadius: 0,
                    pointHoverBorderWidth: 0,
                    radius: 0,
                    borderWidth: 1,
                    label: pageT("preUbiik") + " - " + pageT("thisCalendarMonth")
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
                    hoverRadius: 0,
                    borderWidth: 2,
                    pointHoverBorderWidth: 0,
                    radius: 0,
                    label: pageT("postUbiik") + " - " + pageT("thisCalendarMonth")
                },
                ...lastMonth,
                ...sameMonthLastYear
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
                },
                type: "timeseries"
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
    const urlPrefix = `/api/${props.gatewayID}/devices`
    const
        callCards = (startTime, endTime) => {
            apiCall({
                onComplete: () => setInfoLoading(false),
                onError: error => setInfoError(error),
                onStart: () => setInfoLoading(true),
                onSuccess: rawData => {
                    if (!rawData?.data) return
                    const { data } = rawData,
                        preAndPost = data.energyCosts
                    // console.log(data)
                    setPreUbiikThisMonth(preAndPost?.preUbiikThisMonth || 0)
                    setPostUbiikThisMonth(preAndPost?.postUbiikThisMonth || 0)
                    setPreUbiikLastMonth(preAndPost?.preUbiikLastMonth || 0)
                    setPostUbiikLastMonth(preAndPost?.postUbiikLastMonth || 0)
                    setPreUbiikSameMonthLastYear(preAndPost?.preUbiikTheSameMonthLastYear || 0)
                    setPostUbiikSameMonthLastYear(preAndPost?.postUbiikTheSameMonthLastYear || 0)
                },
                url: `${urlPrefix}/tou/energy-cost?startTime=${startTime}&endTime=${endTime}`
            })
        },
        callLineChartCosts = (startTime, endTime) => {
            apiCall({
                onComplete: () => setLineChartCostsLoading(false),
                onError: error => setLineChartCostsError(error),
                onStart: () => setLineChartCostsLoading(true),
                onSuccess: rawData => {
                    if (!rawData || !rawData.data) return

                    const
                        { data } = rawData,
                        costs = data.energyDailyCosts,
                        { timestamps } = data,
                        labels = timestamps.map(t => t * 1000)

                    console.log(data)

                    setLineChartCosts({
                        data: {
                            preUbiikLastMonth: costs?.preUbiikLastMonth,
                            postUbiikLastMonth: costs?.postUbiikLastMonth

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
            callCards(startTime, endTime)
            callLineChartCosts(startTime, endTime)
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
                    subTitle={pageT("thisCalendarMonth")}
                    data={cardsData.preUbiik}
                    tabs={formats}
                />
                <EconomicsCard
                    icon={EconomicsIcon}
                    subTitle={pageT("thisCalendarMonth")}
                    data={cardsData.postUbiik}
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
                <LineChart data={chartCostComparisonSet({
                    formats,
                    ...lineChartCosts
                })}
                    id="ecoCost" />
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
})