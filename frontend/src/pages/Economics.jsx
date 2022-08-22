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
        [total, setTotal] = useState(630),
        [ancillaryServices, setAncillaryServices] = useState(15),
        [demandCharge, setDemandCharge] = useState(150),
        [timeOfUseArbitrage, setTimeOfUseArbitrage] = useState(130),
        [renewableEnergyCertificate, setRenewableEnergyCertificate] = useState(150),
        [solarLocalUsage, setSolarLocalUsage] = useState(160),
        [exportToGrid, setExportToGrid] = useState(130),
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
            ancillaryServices: fakeData1,
            demandCharge: fakeData2,
            exportToGrid: fakeData3,
            labels: sevenDays,
            rec: fakeData4,
            solarLocalUsage: fakeData5,
            touArbitrage: fakeData6
        })

    const barChartProps = data => ({
        datasets: [
            {
                backgroundColor: colors.green.main,
                data: data.ancillaryServices,
                label: pageT("ancillaryServices")
            },
            {
                backgroundColor: colors.yellow.main,
                data: data.demandCharge,
                label: commonT("demandCharge")
            },
            {
                backgroundColor: colors.blue.main,
                data: data.touArbitrage,
                label: pageT("touArbitrage")
            },
            {
                backgroundColor: colors.indigo.main,
                data: data.rec,
                label: pageT("rec")
            },
            {
                backgroundColor: colors.purple.main,
                data: data.solarLocalUsage,
                label: pageT("solarLocalUsage")
            },
            {
                backgroundColor: colors.gray[300],
                data: data.exportToGrid,
                label: commonT("exportToGrid")
            }
        ],
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
            const {dataIndex, dataset} = context[0], {tooltipData} = dataset

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
                        title={commonT("demandCharge")} />
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
                        title={commonT("exportToGrid")} />
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
                <LineChart
                    data={lineChartProps(lineChartData)}
                    id="economicsLineChart" />
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
                <BarChart
                    data={barChartProps(barChartData)}
                    id="economicsBarChart" />
            </div>
        </div>
    </>
}