import moment from "moment"
import { useMemo, useState } from "react"
import { useTranslation } from "react-multi-lang"

import { API_HOST } from "../constant/env"
import LineChart from "../components/LineChart"
import PriceCard from "../components/PriceCard"
import variables from "../configs/variables"

import { ReactComponent as DemandPeak }
    from "../assets/icons/demand_charge_line.svg"

const { colors } = variables

export default function DemandCharge(props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("demandCharge." + string, params)

    const
        [currentBillCycle, setCurrentBillCycle] = useState(15),
        [realizedSavings, setRealizedSavings] = useState(30),
        [lastMonthBillCycleDemandCharge, setLastMonthBillCycleDemandCharge]
            = useState(30)

    const
        hours24 = Array.from(new Array(24).keys()),
        lineChartDateLabels = hours24.map(n =>
            moment().hour(n).startOf("h").toISOString()),
        currentHour = moment().hour(),
        lineChartDataArray = hours24.filter(v => v <= currentHour).map(() =>
            Math.floor(Math.random() * (60 - 40 + 1) + 40))
    const
        [peak, setPeak] = useState({ active: 0, current: 0, threshhold: 0 }),
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

    const updateData = data => {
        setPeak({
            active: data.gridIsPeakShaving,
            current: data.gridProducedAveragePowerAC,
            threshhold: data.gridContractPowerAC
        })
    }

    const peakShaveRate = useMemo(() => {
        if (!peak.current || !peak.threshhold) return 0

        const rawRate = peak.current / peak.threshhold

        return rawRate <= 1 ? rawRate : 1
    }, [peak.current, peak.threshhold])

    const activeIndicator = peak.active
        ? <>
            <div className="grid h-6 ml-6 mr-1 place-items-center relative w-6">
                <div className="absolute animate-ping bg-negative-main h-3
                            rounded-full w-3" />
                <div className="bg-negative-main h-2.5 rounded-full w-2.5" />
            </div>
            <h5 className="text-negative-main">{pageT("active")}</h5>
        </>
        : null

    const peakShaveColor = peak.active ? "negative" : "positive"

    return <>
        <h1 className="mb-9">{pageT("demandCharge")}</h1>
        <div className="gap-8 grid-cols-3 lg:grid ">
            <PriceCard
                price={currentBillCycle}
                title={pageT("currentBillCycle")} />
            <PriceCard
                price={realizedSavings}
                title={pageT("realizedSavings")} />
            <PriceCard
                price={lastMonthBillCycleDemandCharge}
                title={pageT("lastMonthBillCycleDemandCharge")} />
        </div>
        <div className="lg:flex items-start mb-8">
            <div className="flex-1">
                <div className="card">
                    <div className="flex flex-wrap items-baseline">
                        <h5 className="font-bold">{pageT("peakShave")}</h5>
                    </div>
                    <div className="flex justify-between items-center">
                        <div className="bg-gray-600 flex h-2 overflow-hidden
                                        rounded-full w-full">
                            <div
                                className={`bg-${peakShaveColor}-main
                                            rounded-full`}
                                style={{ width: `${peakShaveRate * 100}%` }}
                            />
                        </div>
                        <div className="bg-gray-600 grid h-20 min-w-20
                                        place-items-center rounded-full w-20
                                        ml-10 md:ml-20">
                            <DemandPeak className="text-gray-400 h-12 w-12" />
                        </div>
                    </div>
                    <div className="flex font-bold items-center -mt-8">
                        <div className="text-28px">
                            <span className={`text-${peakShaveColor}-main`}>
                                {peak.current}
                            </span> / {peak.threshhold} {commonT("kw")}
                        </div>
                        {activeIndicator}
                    </div>
                </div>
            </div>
        </div>
        <div className="card chart">
            <h4 className="mb-10">{pageT("demandDetails")}</h4>
            <div className="max-h-80vh h-160 w-full">
                <LineChart data={lineChartData} id="dcLineChart" />
            </div>
        </div>
    </>
}