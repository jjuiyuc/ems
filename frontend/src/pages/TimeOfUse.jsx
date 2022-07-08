import { Button, Stack } from "@mui/material"
import { Fragment as Frag, useEffect, useRef, useState } from "react"
import moment from "moment"
import { useTranslation } from "react-multi-lang"
import WaterChart from "water-chart"

import Clock from "../components/Clock"
import EnergyCard from "../components/EnergyCard"
import LineChart from "../components/LineChart"
import variables from "../configs/variables"

import { ReactComponent as EditIcon } from "../assets/icons/edit.svg"

const { colors } = variables

export default function TimeOfUse() {
    const batteryChart = useRef()

    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        errorT = string => t("error." + string),
        formT = string => t("form." + string),
        pageT = (string, params) => t("timeOfUse." + string, params)

    const energyCardTitle = source => <>
        <span className="inline-block mr-1">
            {pageT("source", { type: pageT(source) })}
        </span>
        <span className="inline-block">
            ({pageT("totalUntilNow")})
        </span>
    </>

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
        [battery, setBattery] = useState({
            source: "solar",
            level: 52.5,
            power: 20,
        }),
        [clockDataset, setClockDataset] = useState({
            data: [], backgroundColor: []
        }),
        [currentPeriod, setCurrentPeriod] = useState(""),
        [currentTime, setCurrentTime] = useState(""),
        [lineChartData, setLineChartData] = useState({
            datasets: [{
                backgroundColor: "#12c9c9",
                borderColor: "#12c9c9",
                borderWidth: 1,
                data: lineChartDataArray,
                fill: {
                    above: "rgba(18, 201, 201, .2)",
                    target: "origin"
                },
                hoverRadius: 3,
                pointBorderColor: "rgba(18, 201, 201, .2)",
                pointHoverBorderWidth: 6,
                pointBorderWidth: 0,
                radius: 3,
                tension: 0
            }],
            labels: lineChartDateLabels,
            tooltipCallbacks: {
                label: item => `${item.parsed.y}%`,
                labelPointStyle: context => {
                    const
                        color = context.dataset.backgroundColor
                            .replace("#", "%23"),
                        image = new Image(8, 8)

                    image.className = "test"
                    image.src = `data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' height='8' width='8'%3E%3Ccircle cx='4' cy='4' r ='4' fill='${color}' /%3E%3C/svg%3E`

                    return { pointStyle: image }
                }
            },
            tickCallback: function (val, index) {
                return val + '%'
            }
        }),
        [midPeak, setMidPeak] = useState({
            types: [
                { kwh: 7.5, percentage: 15, type: "grid" },
                { kwh: 30, percentage: 60, type: "solar" },
                { kwh: 12.5, percentage: 25, type: "battery" },
            ],
            kwh: 50
        }),
        [onPeak, setOnPeak] = useState({
            types: [
                { kwh: 5, percentage: 10, type: "grid" },
                { kwh: 52, percentage: 50, type: "solar" },
                { kwh: 20, percentage: 40, type: "battery" },
            ],
            kwh: 50
        }),
        [offPeak, setOffPeak] = useState({
            types: [
                { kwh: 10, percentage: 18, type: "grid" },
                { kwh: 25, percentage: 41, type: "solar" },
                { kwh: 25, percentage: 41, type: "battery" },
            ],
            kwh: 60
        }),
        [prices, setPrices]
            = useState({ onPeak: 0, midPeak: 0, offPeak: 0, superOffPeak: 0 }),
        [superOffPeak, setSuperOffPeak] = useState({
            types: [
                { kwh: 21, percentage: 35, type: "grid" },
                { kwh: 24, percentage: 40, type: "solar" },
                { kwh: 15, percentage: 25, type: "battery" },
            ],
            kwh: 60
        }),
        [tab, setTab] = useState("today"),
        [timeOfUse, setTimeOfUse] = useState([
            {
                end: "05:00",
                name: "superOffPeak",
                price: 1.2,
                start: "00:00"
            },
            {
                end: "11:00",
                name: "offPeak",
                price: 1.8,
                start: "05:00"
            },
            {
                end: "17:00",
                name: "midPeak",
                price: 2.2,
                start: "11:00"
            },
            {
                end: "23:00",
                name: "onPeak",
                price: 3.5,
                start: "17:00"
            },
            {
                end: "24:00",
                name: "superOffPeak",
                price: 1.2,
                start: "23:00"
            }
        ])

    const getMoment = string => {
        const [hour, minute] = string.split(":")

        return moment().hour(parseInt(hour)).minute(parseInt(minute)).second(0)
    }

    useEffect(() => {
        const
            currentTime = moment(),
            dataset = { data: [], backgroundColor: [] },
            prices = { onPeak: 0, midPeak: 0, offPeak: 0, superOffPeak: 0 }

        let currentPeriod = ""

        if (timeOfUse.length === 0) return

        timeOfUse.forEach(item => {
            const
                { end, start } = item,
                endTime = getMoment(end),
                startTime = getMoment(start),
                duration = moment.duration(endTime.diff(startTime)).as("hours")

            dataset.data.push(duration)
            dataset.backgroundColor.push(colors[item.name])
            prices[item.name] = item.price

            if (currentTime >= startTime) {
                currentPeriod = item.name
            }
        })

        setClockDataset(dataset)
        setCurrentPeriod(currentPeriod)
        setCurrentTime(currentTime.format("hh:mm A"))
        setPrices(prices)
    }, [timeOfUse])

    useEffect(() => {
        if (!batteryChart.current) return

        batteryChart.current.innerHTML = ""

        new WaterChart({
            container: "#batteryChart",
            fillOpacity: .4,
            margin: 6,
            maxValue: 100,
            minValue: 0,
            series: [battery.level],
            stroke: colors.gray[400],
            strokeWidth: 2,
            textColor1: "white",
            textPositionY: .45,
            textSize: .3,
            textUnitSize: "32px",
            waveColor1: colors.primary.main,
            waveColor2: colors.primary.main
        })
    }, [battery, tab])

    return <>
        <div className="page-header">
            <h1>{pageT("timeOfUse")}</h1>
            <Stack direction="row" justifyContent="flex-end" spacing={1.5}>
                <Button
                    onClick={() => setTab("today")}
                    filter={tab === "today" ? "selected" : ""}
                    radius="pill"
                    variant="contained">
                    {pageT("today")}
                </Button>
                <Button
                    filter={tab === "yesterday" ? "selected" : ""}
                    onClick={() => setTab("yesterday")}
                    radius="pill"
                    variant="contained">
                    {pageT("yesterday")}
                </Button>
            </Stack>
        </div>
        <div className="gap-8 grid md:grid-cols-2 items-start">
            <EnergyCard data={onPeak} title={energyCardTitle("onPeak")} />
            <EnergyCard data={offPeak} title={energyCardTitle("offPeak")} />
            <EnergyCard data={midPeak} title={energyCardTitle("midPeak")} />
            <EnergyCard
                data={superOffPeak}
                title={energyCardTitle("superOffPeak")} />
            <div className="card">
                <div className="header -mr-4">
                    <h4>{pageT("timeOfUse")}</h4>
                    <Button
                        color="brand"
                        radius="pill"
                        size="small"
                        variant="text">
                        <EditIcon className="h-4 mr-1 w-4" />
                        {pageT("editTimeOfUse")}
                    </Button>
                </div>
                <div className="flex flex-wrap items-center justify-around">
                    <div className="flex flex-wrap items-center justify-center">
                        <Clock dataset={clockDataset} id="touClock" />
                        <div className="grid grid-cols-3-auto gap-y-2
                                        items-center mx-8 my-4 text-white">
                            {Object.keys(prices).map((key, i) =>
                                <Frag key={"t-p-" + i}>
                                    <div
                                        className="h-2 rounded-full mr-3 w-2"
                                        style={{ background: colors[key] }} />
                                    <div className="text-11px">
                                        {pageT(key)}
                                    </div>
                                    <div className="font-bold ml-2 text-base">
                                        ${prices[key]}
                                    </div>
                                </Frag>)}
                        </div>
                    </div>
                    <div className="my-6 subCard w-56">
                        <h6 className="font-bold text-11px text-gray-300">
                            {pageT("current")}
                        </h6>
                        <div className="flex flex-col items-center justify-center">
                            <div
                                className="font-bold mb-3 mt-2 px-5 py-2 rounded-full text-gray-900 text-sm"
                                style={{ background: colors[currentPeriod] }}>
                                {pageT(currentPeriod)}
                            </div>
                            <div className="font-bold text-2xl">
                                {currentTime}
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            {tab === "today"
                ? <div className="card">
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
                        <div className="grid grid-cols-3 my-6 mw-88 three-columns
                                    gap-x-5 sm:gap-x-10">
                            <div>
                                <h3>{battery.level}%</h3>
                                <span className="text-13px">
                                    {pageT("stateOfCharge")}
                                </span>
                            </div>
                            <div>
                                <h3>{battery.power} {commonT("kw")}</h3>
                                <span className="text-13px">
                                    {pageT("batteryPower")}
                                </span>
                            </div>
                            <div>
                                <h3>{commonT(battery.source)}</h3>
                                <span className="text-13px">
                                    {pageT("chargingFrom")}
                                </span>
                            </div>
                        </div>
                    </div>
                </div>
                : null}
        </div>
        {tab === "today"
            ? <>
                <div className="mt-20 page-header">
                    <h1>{pageT("directSolarUsage")}</h1>
                </div>
                <div className="card chart max-h-80vh h-160 relative w-full">
                    <LineChart data={lineChartData} id="touLineChart" />
                </div>
            </>
            : null}
    </>
}