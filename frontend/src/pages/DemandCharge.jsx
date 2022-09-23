import { connect } from "react-redux"
import moment from "moment"
import { useMemo, useState, useEffect } from "react"
import { useTranslation } from "react-multi-lang"

import { apiCall } from "../utils/api"
import variables from "../configs/variables"

import LineChart from "../components/LineChart"
import PriceCard from "../components/PriceCard"
import { ReactComponent as DemandPeak }
    from "../assets/icons/demand_charge_line.svg"

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

const mapState = state => ({ gatewayID: state.gateways.active.gatewayID })

export default connect(mapState)(function DemandCharge(props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("demandCharge." + string, params)
    const
        [infoError, setInfoError] = useState(""),
        [infoLoading, setInfoLoading] = useState(false),
        [currentBillingCycle, setCurrentBillingCycle] = useState(0),
        [realizedSavings, setRealizedSavings] = useState(0),
        [lastMonthBillingCycle, setLastMonthBillingCycle]
            = useState(0),
        [peak, setPeak] = useState({ active: 0, current: 0, threshhold: 0 }),
        [lineChartDemand, setLineChartDemand] = useState(null),
        [lineChartDemandError, setLineChartDemandError] = useState(""),
        [lineChartDemandLoading, setLineChartDemandLoading] = useState(false),
        [lineChartDemandRes] = useState("")

    const chartDemandDetailsSet = ({ data, labels }) => ({
        datasets: [{
            backgroundColor: colors.primary.main,
            borderColor: colors.primary.main,
            data: data || [],
            fill: {
                above: colors.primary["main-opacity-10"],
                target: "origin"
            },
            pointBorderColor: colors.primary["main-opacity-20"]
        }],
        labels,
        tickCallback: (val, index) => val + commonT("kw"),
        tooltipLabel: item => `${item.parsed.y}` + commonT("kw"),
        x: { grid: { lineWidth: 0 } },
        y: { max: 80, min: 0 },
        beforeDraw: chart => {
            if (!peak.threshhold) return
            const
                xEnd = chart.scales.x.right,
                xStart = chart.scales.x.left,
                y = chart.scales["y"].getPixelForValue(peak.threshhold)

            let ctx = chart.ctx

            ctx.beginPath()
            ctx.moveTo(xStart, y)
            ctx.lineTo(xEnd, y)
            ctx.lineWidth = 1
            ctx.strokeStyle = colors.negative.main
            ctx.stroke()
        }
    })
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
            <h5 className="text-negative-main">{commonT("active")}</h5>
        </>
        : null

    const peakShaveColor = peak.active ? "negative" : "positive"

    useEffect(() => {
        if (!props.gatewayID) return

        const
            startTime = moment().startOf("day").toISOString(),
            endTime = moment().endOf("day").toISOString(),
            urlPrefix = `/api/${props.gatewayID}/devices`

        apiCall({
            onComplete: () => setInfoLoading(false),
            onError: error => setInfoError(error),
            onStart: () => setInfoLoading(true),
            onSuccess: rawData => {
                if (!rawData || !rawData.data) return

                const { data } = rawData

                setCurrentBillingCycle(data.gridPowerCost || 0)
                setRealizedSavings(data.gridPowerCostSavings || 0)
                setLastMonthBillingCycle(data.gridPowerCostLastMonth || 0)
                setPeak({
                    active: data.gridIsPeakShaving,
                    current: data.gridProducedAveragePowerAC,
                    threshhold: data.gridContractPowerAC
                })
            },
            url: `${urlPrefix}/charge-info?startTime=${startTime}`
        })

        apiCall({
            onComplete: () => setLineChartDemandLoading(false),
            onError: error => setLineChartDemandError(error),
            onStart: () => setLineChartDemandLoading(true),
            onSuccess: rawData => {
                if (!rawData || !rawData.data) return

                const
                    { data } = rawData,
                    { timestamps } = data,
                    labels = timestamps.map(t => t * 1000)
                setLineChartDemand({
                    data: data.gridLifetimeEnergyACDiffToPowers,
                    labels
                })
            },
            url: `${urlPrefix}/demand-state?startTime=${startTime}&endTime=${endTime}`
        })
    }, [props.gatewayID])

    return <>
        <h1 className="mb-9">{commonT("demandCharge")}</h1>
        <div className="gap-8 grid-cols-3 lg:grid ">
            <PriceCard
                price={currentBillingCycle}
                title={pageT("currentBillingCycle")} />
            <PriceCard
                price={realizedSavings}
                title={pageT("realizedSavings")} />
            <PriceCard
                price={lastMonthBillingCycle}
                title={pageT("lastMonthBillingCycle")} />
            {infoLoading
                ? <div className="absolute bg-black-main-opacity-95 grid inset-0
                                place-items-center rounded-3xl">
                    <Spinner />
                </div>
                : null}
        </div>
        <div className="lg:flex items-start mb-8">
            <div className="flex-1">
                <div className="card">
                    <div className="flex flex-wrap items-baseline">
                        <h5 className="font-bold">{commonT("peakShave")}</h5>
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
                <LineChart data={chartDemandDetailsSet({
                    ...lineChartDemand
                })} id="dcLineChart" />
                <ErrorBox
                    error={lineChartDemandError}
                    message={pageT("chartError")} />
                <LoadingBox loading={lineChartDemandLoading} />
            </div>
        </div>
    </>
})