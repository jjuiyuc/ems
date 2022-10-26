import { useTranslation } from "react-multi-lang"
import { useEffect, useRef } from "react"
import variables from "../configs/variables"
import WaterChart from "water-chart"

const { colors } = variables

export default function BatteryStatusCard({ data }) {

    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("timeOfUse." + string, params)

    const batteryChart = useRef()

    useEffect(() => {

        if (batteryChart.current.stateValue == data.state) return

        batteryChart.current.stateValue = data.state
        batteryChart.current.innerHTML = ""

        new WaterChart({
            container: "#batteryChart",
            fillOpacity: .4,
            margin: 6,
            maxValue: 100,
            minValue: 0,
            series: [data.state],
            stroke: colors.gray[400],
            strokeWidth: 2,
            textColor1: "white",
            textPositionY: .45,
            textSize: .3,
            textUnitSize: "32px",
            waveColor1: colors.primary.main,
            waveColor2: colors.primary.main
        })
    }, [data.state])
    return <div className="card">
        <div className="header">
            <h4>{pageT("batteryStatus")}</h4>
        </div>
        <div className="flex flex-wrap items-center justify-around">
            <div className="h-48 relative w-48">
                <div className="absolute bg-gray-800 h-44 m-2
                rounded-full w-44" />
                <svg
                    className="h-48 relative w-48"
                    id="batteryChart"
                    ref={batteryChart} />
            </div>
            <div className="column-separator grid grid-cols-3 my-6
    mw-88 gap-x-5 sm:gap-x-10">
                <div>
                    <h3 className="text-22px">{data.state}%</h3>
                    <span className="text-13px">
                        {commonT("stateOfCharge")}
                    </span>
                </div>
                <div>
                    <h3 className="text-22px">{data.power} {commonT("kw")}</h3>
                    <span className="text-13px">
                        {commonT("batteryPower")}
                    </span>
                </div>
                <div>
                    <h3 className="text-22px">{data.target ? commonT(data.target) : "-"}</h3>
                    <span className="text-13px">
                        {pageT(data.direction)}
                    </span>
                </div>
            </div>
        </div>
    </div>
}