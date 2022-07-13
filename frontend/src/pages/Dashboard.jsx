import { useState, useMemo } from "react"
import { useTranslation } from "react-multi-lang"

import ElectricalGridDiagram from "../components/ElectricalGridDiagram"

import { ReactComponent as HomeIcon } from "../assets/icons/home.svg"
import { ReactComponent as BatteryIcon } from "../assets/icons/battery.svg"
import { ReactComponent as GridIcon } from "../assets/icons/grid.svg"
import { ReactComponent as SolarIcon } from "../assets/icons/sunny.svg"

import "../assets/css/dashboard.scss"

const iconMap = {
    battery: {color: "blue", icon: BatteryIcon},
    grid: {color: "indigo", icon: GridIcon},
    home: {color: "green", icon: HomeIcon},
    solar: {color: "yellow", icon: SolarIcon}
}

const IconSet = props => {
    const {size} = props, {color, icon} = iconMap[props.icon], Icon = icon

    return <div className={`h-${size} bg-${color}-main-opacity-20 grid
                            place-items-center rounded-full w-${size}`}>
        <Icon className={`h-3/5 w-3/5 text-${color}-main`} />
    </div>
}
const Card = props => {
    const
        cardClasses = "card narrow px-6"
            + (props.className ? " " + props.className : ""),
        {data, title} = props

    return <div className={cardClasses}>
        <div className="grid sm:grid-cols-2 text-center
                        gap-x-4 sm:gap-x-6
                        gap-y-6 sm:gap-y-8">
            <div className="flex flex-wrap sm:col-span-2 items-center">
                <IconSet icon={props.icon} size="10" />
                <h5 className="font-bold ml-3">{title}</h5>
            </div>
        {data.map((item, i) =>
            <div key={"d-c-d-" + i}>
                <h2 className="font-bold text-3xl mb-1">{item.value}</h2>
                <p className="lg:test text-sm">{item.name}</p>
            </div>)}
        </div>
    </div>
}
const CardOnDiagram = props =>
    <div className={`arrow-${props.arrow} card-diagram bg-gray-900 flex
                    items-start px-4 py-3 rounded-lg`}>
        <IconSet icon={props.icon} size="6" />
        <div className="ml-2 mt-0.5">
            <h6 className="font-bold text-gray-200">{props.value}</h6>
            <span className="text-11px text-gray-300">{props.title}</span>
        </div>
    </div>

export default function Dashboard() {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("dashboard." + string, params)

    const
        [current, setCurrent] = useState(13),
        [threshhold, setThreshhold] = useState(15)

    const
        peakShaveRate = useMemo(() => {
            const rawRate = current / threshhold
            return rawRate <= 1 ? rawRate : 1
        }, [current, threshhold])

    const
        batteryData = [
            {name: pageT("stateOfCharge"), value: "80%"},
            {name: pageT("batteryPower"), value: `20 ${commonT("kw")}`},
            {name: pageT("importFromGrid"), value: `10 ${commonT("kw")}`},
            {name: pageT("dischargingTo"), value: "-"},
        ],
        loadData = [
            {name: commonT("solar"), value: `20 ${commonT("kw")}`},
            {name: pageT("batteryDischarge"), value: "-"},
            {name: pageT("importFromGrid"), value: `10 ${commonT("kw")}`}
        ],
        solarData = [
            {name: commonT("solar"), value: `20 ${commonT("kw")}`},
            {name: pageT("batteryDischarge"), value: "-"},
            {name: pageT("importFromGrid"), value: `10 ${commonT("kw")}`},
        ]

    const
        BatteryCard = props =>
            <CardOnDiagram
                arrow={props.arrow}
                icon="battery"
                title={commonT("battery")}
                value={`20 ${commonT("kw")}`} />,
        GridCard = props =>
            <CardOnDiagram
                arrow={props.arrow}
                icon="grid"
                title={commonT("grid")}
                value={`10 ${commonT("kw")}`} />,
        LoadCard = props =>
            <CardOnDiagram
                arrow={props.arrow}
                icon="home"
                title={pageT("load")}
                value={`30 ${commonT("kw")}`} />,
        SolarCard = props =>
            <CardOnDiagram
                arrow={props.arrow}
                icon="solar"
                title={commonT("solar")}
                value={`40 ${commonT("kw")}`} />

    const peakShaveColor = current > threshhold ? "negative" : "positive"

    return <>
        <h1 className="mb-9">{pageT("dashboard")}</h1>
        <div className="lg:flex items-start">
            <div className="flex-1">
                <div className="card">
                    <div className="flex flex-wrap items-baseline mb-6">
                        <h5 className="font-bold">{pageT("peakShave")}</h5>
                    </div>
                    <div className="bg-gray-600 flex h-2 overflow-hidden
                                    rounded-full w-full ">
                        <div
                            className={`bg-${peakShaveColor}-main rounded-full`}
                            style={{ width: `${peakShaveRate * 100}%` }}
                        />
                    </div>
                    <div className="mt-2 text-28px font-bold">
                        <span className={`text-${peakShaveColor}-main`}>
                            {current}
                        </span> / {threshhold} {commonT("kw")}
                    </div>
                </div>
                <div className="grid-cols-3-auto items-stretch
                                block md:grid lg:block xl:grid
                                mx-5 lg:mx-9 my-10 lg:my-18">
                    <div className="grid-cols-2 mb-4
                                    grid md:hidden lg:grid xl:hidden">
                        <div className="flex"><LoadCard arrow="bottom" /></div>
                        <div className="flex justify-end">
                            <GridCard arrow="bottom" />
                        </div>
                    </div>
                    <div className="flex-col justify-between
                                    hidden md:flex lg:hidden xl:flex">
                        <div><LoadCard arrow="right" /></div>
                        <div><GridCard arrow="right" /></div>
                    </div>
                    <ElectricalGridDiagram
                        className="h-auto w-full
                                    md:py-8 lg:py-0 xl:py-8" />
                    <div className="grid-cols-2 mt-4
                                    grid md:hidden lg:grid xl:hidden">
                        <div className="flex"><BatteryCard arrow="top" /></div>
                        <div className="flex justify-end">
                            <SolarCard arrow="top" />
                        </div>
                    </div>
                    <div className="flex-col justify-between
                                    hidden md:flex lg:hidden xl:flex">
                        <div><BatteryCard arrow="left" /></div>
                        <div><SolarCard arrow="left" /></div>
                    </div>
                </div>
            </div>
            <div className="lg:ml-5 gap-5 grid md:grid-cols-2 lg:block lg:w-88">
                <Card data={loadData} icon="home" title={pageT("load")} />
                <Card
                    className="lg:mt-5"
                    data={batteryData}
                    icon="battery"
                    title={commonT("battery")} />
                <Card
                    className="lg:mt-5"
                    data={solarData}
                    icon="solar"
                    title={commonT("solar")} />
            </div>
        </div>
    </>
}