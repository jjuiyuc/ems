import { Button, Stack } from "@mui/material"
import { useState } from "react"
import { useTranslation } from "react-multi-lang"
import moment from "moment"

import { API_HOST } from "../constant/env"
import BarChart from "../components/BarChart"
import EnergySolarSubCard from "../components/EnergySolarSubCard"
import LineChart from "../components/LineChart"
import Spinner from "../components/Spinner"

import variables from "../configs/variables"

import { ReactComponent as EconomicsIcon } from "../assets/icons/economics.svg"

const { colors } = variables

const barChartColors = {
    ancillaryServices: colors.green.main,
    demandCharge: colors.yellow.main,
    exportToGrid: colors.gray[300],
    touArbitrage: colors.blue.main,
    rec: colors.indigo.main,
    solarLocalUsage: colors.purple.main,
}

const fakeDataArray = amount => Array.from(new Array(amount).keys())
    .map(() => Math.floor(Math.random() * (40 - 10 + 1) + 10))

export default function Economics(props) {

    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("economics." + string, params),
        monthtlyTabs = ["thisMonth", "perviousMonth", "thisMonthLastYear"],
        weeklyTabs = ["thisWeek", "perviouWeek"]

    const
        [infoError, setInfoError] = useState(""),
        [infoLoading, setInfoLoading] = useState(false),
        [total, setTotal] = useState(630),
        [preUbiik, setPreUbiik] = useState(0),
        [postUbiik, setPostUbiik] = useState(0),
        [monthtlyTab, setMonthtlyTab] = useState(monthtlyTabs[0]),
        [weeklyTab, setWeeklyTab] = useState(weeklyTabs[0])

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

    const
        [lineChartData, setLineChartData] = useState({
            ancillaryServices: fakeData1,
            demandCharge: fakeData2,
            exportToGrid: fakeData3,
            touArbitrage: fakeData4,
            rec: fakeData5,
            solarLocalUsage: fakeData6,
            total: lineChartDataArray,
            labels: sevenDays
        }),
        [barChartData, setBarChartData] = useState({
            labels: sevenDays,
            data: {
                ancillaryServices: fakeData1,
                demandCharge: fakeData2,
                touArbitrage: fakeData3,
                rec: fakeData4,
                solarLocalUsage: fakeData5,
                exportToGrid: fakeData6
            }
        })

    const barChartProps = data => ({
        datasets: Object.keys(data.data).map(key => ({
            backgroundColor: barChartColors[key],
            data: data.data[key],
            label: pageT(key)
        })),
        labels: data.labels,
        tooltipLabel: item =>
            `${item.dataset.label} ${item.parsed.y} ${commonT("kwh")}`,
        y: { max: 70, min: 0 }
    })

    const lineChartProps = data => ({
        datasets: [{
            backgroundColor: colors.primary.main,
            borderColor: colors.primary.main,
            data: data.total,
            fill: {
                above: colors.primary["main-opacity-10"],
                target: "origin"
            },
            pointBorderColor: colors.primary["main-opacity-20"],
            tooltipData: {
                total: data.total,
                ancillaryServices: data.ancillaryServices,
                demandCharge: data.demandCharge,
                touArbitrage: data.touArbitrage,
                rec: data.rec,
                solarLocalUsage: data.solarLocalUsage,
                exportToGrid: data.exportToGrid
            }
        }],
        labels: data.labels,
        tickCallback: (val, index) => "$" + val,
        tooltipAfterBody: context => {
            const { dataIndex, dataset } = context[0], { tooltipData } = dataset

            return Object.keys(tooltipData).map(key =>
                pageT(key) + " $" + tooltipData[key][dataIndex])
        },
        tooltipLabel: () => null,
        tooltipUsePointStyle: false,
        x: {
            type: "time",
            grid: { lineWidth: 0 },
            time: {
                displayFormats: {
                    day: "MMM D"
                },
                tooltipFormat: "MMM D",
                unit: "day"
            }
        },
        y: { max: 80, min: 0 }
    })

    return <>
        <h1 className="mb-9">{pageT("economics")}</h1>
        <div className="font-bold gap-5 grid md:grid-cols-2 mt-4 mb-8">
            <EnergySolarSubCard
                icon={EconomicsIcon}
                subTitle={pageT("thisCalendarMonth")}
                title={pageT("preUbiik")}
                value={"$" + preUbiik}

            />
            <EnergySolarSubCard
                icon={EconomicsIcon}
                title={pageT("postUbiik")}
                subTitle={pageT("thisCalendarMonth")}
                value={"$" + postUbiik}

            />
            {infoLoading
                ? <div className="absolute bg-black-main-opacity-95 grid inset-0
                                place-items-center rounded-3xl">
                    <Spinner />
                </div>
                : null}
        </div>

        <div className="card chart">
            <div className="items-center grid-cols-1fr-auto-1fr mb-8 lg:grid">
                <h4 className="mb-4 lg:mb-0">{pageT("monthlyStackedRevenue")}</h4>
                <Stack direction="row" spacing={1.5} className="flex-wrap lg:flex">
                    {monthtlyTabs.map((t, i) =>
                        <Button
                            className="mb-5"
                            color="gray"
                            onClick={() => setMonthtlyTab(t)}
                            filter={monthtlyTab === t ? "selected" : ""}
                            key={"ec-m" + i}
                            radius="pill"
                            variant="contained">
                            {pageT(t)}
                        </Button>)}
                </Stack>
                <div />
            </div>
            <div className="max-h-80vh h-160 w-full mt-8">
                <LineChart
                    data={lineChartProps(lineChartData)}
                    id="economicsLineChart" />
            </div>
        </div>
        <div className="card mt-8">
            <div className="items-center grid-cols-1fr-auto-1fr mb-8 lg:grid">
                <h4 className="mb-4 lg:mb-0">{pageT("weeklyRevenueBreakdown")}</h4>
                <Stack direction="row" spacing={1.5} className="flex-wrap lg:flex">
                    {weeklyTabs.map((t, i) =>
                        <Button
                            color="gray"
                            onClick={() => setWeeklyTab(t)}
                            filter={weeklyTab === t ? "selected" : ""}
                            key={"ec-w" + i}
                            radius="pill"
                            variant="contained">
                            {pageT(t)}
                        </Button>)}
                </Stack>
            </div>
            <div className="max-h-80vh h-160 mt-8 relative w-full">
                <BarChart
                    data={barChartProps(barChartData)}
                    id="economicsBarChart" />
            </div>
        </div>
    </>
}