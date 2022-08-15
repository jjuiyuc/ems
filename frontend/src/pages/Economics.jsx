import { Button, Stack } from "@mui/material"
import { useState } from "react"
import { useTranslation } from "react-multi-lang"
import moment from "moment"

import { API_HOST } from "../constant/env"
import PriceCard from "../components/PriceCard"
import LineChart from "../components/LineChart"
import BarChart from "../components/BarChart"
import variables from "../configs/variables"

import { ReactComponent as EconomicsIcon } from "../assets/icons/economics.svg"

const { colors } = variables

export default function Economics(props) {

    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("economics." + string, params)
    const
        [total, setTotal] = useState(630),
        [ancillaryServices, setAncillaryServices] = useState(15),
        [demandCharge, setDemandCharge] = useState(150),
        [timeOfUseArbitrage, setTimeOfUseArbitrage] = useState(130),
        [renewableEnergyCertificate, setRenewableEnergyCertificate] = useState(150),
        [solarLocalUsage, setSolarLocalUsage] = useState(160),
        [exportToGrid, setExportToGrid] = useState(130)
    const
        hours24 = Array.from(new Array(24).keys()),
        lineChartDateLabels = hours24.map(n =>
            moment().hour(n).startOf("h").toISOString()),
        currentHour = moment().hour(),
        lineChartDataArray = hours24.filter(v => v <= currentHour).map(() =>
            Math.floor(Math.random() * (60 - 40 + 1) + 40))
    const
        [lineChartData, setLineChartData] = useState({
            datasets: [{
                backgroundColor: colors.primary.main,
                borderColor: colors.primary.main,
                data: lineChartDataArray,
                fill: {
                    above: colors.primary["main-opacity-10"],
                    target: "origin"
                },
                pointBorderColor: colors.primary["main-opacity-20"]
            }],
            labels: lineChartDateLabels,
            tickCallback: (val, index) => val + commonT("kw"),
            tooltipLabel: item => `${item.parsed.y}` + commonT("kw"),
            x: {
                type: "time",
                grid: { lineWidth: 0 },
                time: {
                    displayFormats: {
                        day: "MM DD"
                    },
                    tooltipFormat: "MM DD",
                    unit: "day"
                }
            },
            y: { max: 80, min: 0 }
        })
    const
        [monthtlyTab, setMonthtlyTab] = useState("monthlyStackedRevenue"),
        monthtlyTabs = ["thisMonth", "perviousMonth", "thisMonthLastYear"],
        [weeklyTab, setWeeklyTab] = useState("weeklyRevenueBreakdown"),
        weeklyTabs = ["thisWeek", "perviouWeek"]

    const fakeDataArray = amount => Array.from(new Array(amount).keys())
        .map(() => Math.floor(Math.random() * (40 - 10 + 1) + 10))

    const
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
        [barChartData, setBarChartData] = useState({
            datasets: [
                {
                    backgroundColor: colors.green.main,
                    data: fakeData1,
                    label: pageT("ancillaryServices")
                },
                {
                    backgroundColor: colors.yellow.main,
                    data: fakeData2,
                    label: pageT("demandCharge")
                },
                {
                    backgroundColor: colors.blue.main,
                    data: fakeData3,
                    label: pageT("tOUArbitrage")
                },
                {
                    backgroundColor: colors.indigo.main,
                    data: fakeData4,
                    label: pageT("rec")
                },
                {
                    backgroundColor: colors.purple.main,
                    data: fakeData5,
                    label: pageT("solarLocalUsage")
                },
                {
                    backgroundColor: colors.gray[300],
                    data: fakeData6,
                    label: pageT("exportToGrid")
                }
            ],
            labels: sevenDays,
            tooltipLabel: item =>
                `${item.dataset.label} ${item.parsed.y} ${commonT("kwh")}`,
            y: { max: 70, min: 0 }
        })

    return <>
        <h1 className="mb-9">{pageT("economics")}</h1>
        <div className="flex">
            <div className="card w-3/12 mb-8 grid grid-row-2">
                <div>
                    <h5 className="font-bold">{pageT("total")}</h5>
                    <h2 className="font-bold mb-1 pt-4">{pageT("february")} 2022</h2>
                    <h2 className="font-bold mb-1 pt-4">${total}</h2>
                </div>
                <div className="flex justify-end items-end">
                    <div className="items-center place-items-center grid
                        bg-primary-main-opacity-20 w-20 h-20 rounded-full">
                        <EconomicsIcon className="text-brand-main w-12 h-12" />
                    </div>
                </div>
            </div>
            <div className="flex flex-wrap ml-5">
                <div className="grid-cols-3 auto-cols-max gap-x-5 md:grid">
                    <PriceCard
                        price={ancillaryServices}
                        title={pageT("ancillaryServices")} />
                    <PriceCard
                        price={demandCharge}
                        title={pageT("demandCharge")} />
                    <PriceCard
                        price={timeOfUseArbitrage}
                        title={pageT("timeOfUseArbitrage")} />
                    <PriceCard
                        price={renewableEnergyCertificate}
                        title={pageT("renewableEnergyCertificate")} />
                    <PriceCard
                        price={solarLocalUsage}
                        title={pageT("solarLocalUsage")} />
                    <PriceCard
                        price={exportToGrid}
                        title={pageT("exportToGrid")} />
                </div>
            </div>
        </div>
        <div className="card chart">
            <div className="items-center grid-cols-1fr-auto-1fr mb-8 lg:grid">
                <h4 className="mb-4 lg:mb-0">{pageT("monthlyStackedRevenue")}</h4>
                <Stack direction="row" spacing={1.5}>
                    {monthtlyTabs.map((t, i) =>
                        <Button
                            className="md:"
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
                <LineChart data={lineChartData} id="dcLineChart" />
            </div>
        </div>
        <div className="card mt-8">
            <div className="items-center grid-cols-1fr-auto-1fr mb-8 lg:grid">
                <h4 className="mb-4 lg:mb-0">{pageT("weeklyRevenueBreakdown")}</h4>
                <Stack direction="row" spacing={1.5}>
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
                <BarChart data={barChartData} id="econoBarChart" />
            </div>
        </div>
    </>
}