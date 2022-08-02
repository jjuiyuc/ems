import { Button, Stack } from "@mui/material"
import { useEffect, useState } from "react"
import { useTranslation } from "react-multi-lang"
import moment from "moment"

import { API_HOST } from "../constant/env"
import PriceCard from "../components/PriceCard"
import LineChart from "../components/LineChart"
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
        lineChartDateLabels = hours24.map(n => {
            const time = moment().hour(n).minute(0).second(0)

            return time.format("hh A")
        }),
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
            x: { grid: { lineWidth: 0 } },
            y: { max: 80, min: 0 }
        })
    const
        [tab, setTab] = useState("monthlyStackedRevenue"),
        tabs = ["thisMonth", "perviousMonth", "thisMonthLastYear"]

    return <>
        <h1 className="mb-9">{pageT("economics")}</h1>
        <div className="flex max-w-7xl">
            <div className="card w-3/12 mb-8">
                <h5 className="font-bold">{pageT("total")}</h5>
                <h2 className="font-bold mb-1 pt-4">{pageT("february")} 2022</h2>
                <h2 className="font-bold mb-1 pt-4">${total}</h2>
                <div className="bg-primary-main-opacity-20 w-20 h-20 rounded-full relative">
                    <EconomicsIcon className="text-brand-main w-12 h-12 ml-3 absolute" />
                </div>
            </div>
            <div className="flex flex-wrap ml-5">
                <div className="lg:grid grid-cols-3 auto-cols-max">
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
            <div className="grid items-center grid-cols-1fr-auto-1fr">
                <h4>{pageT("monthlyStackedRevenue")}</h4>
                <Stack direction="row" spacing={1.5}>
                    {tabs.map((t, i) =>
                        <Button
                            color="purple"
                            onClick={() => setTab(t)}
                            filter={tab === t ? "selected" : ""}
                            key={"a-t-" + i}
                            radius="pill"
                            variant="contained">
                            {pageT(t)}
                        </Button>)}
                </Stack>
                <div />
            </div>
            <div className="max-h-80vh h-160 w-full">
                <LineChart data={lineChartData} id="dcLineChart" />
            </div>
        </div>


    </>
}