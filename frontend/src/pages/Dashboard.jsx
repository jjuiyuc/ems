import { useState } from "react"
import { useTranslation } from "react-multi-lang"
import variables from "../configs/variables"
import { ReactComponent as HomeIcon } from "../assets/icons/home.svg"
import { ReactComponent as BatteryIcon } from "../assets/icons/battery.svg"
import { ReactComponent as SolarIcon } from "../assets/icons/sunny.svg"

export default function Dashboard() {
    const
        t = useTranslation(),
        commonT = string => t("common." + string)

    const
        [current, setCurrent] = useState(13),
        [threshhold, setThreshhold] = useState(15)

    return <>
        <h1>Dashboard</h1>
        <div className="flex items-start">
            <div className="card w-full mt-8">
                <div className="flex flex-wrap items-baseline mb-6">
                    <h5 className="font-bold">Peak Shave</h5>
                </div>
                <div className="flex h-2 overflow-hidden rounded-full w-full bg-gray-600">
                    <div
                        className="bg-positive-main h-full"
                        style={{ width: '40%' }}
                    />
                </div>
                <div className="text-28px font-bold">
                    <span className={current > threshhold ? "text-negative-main" : "text-positive-main"}>
                        {current}
                    </span>/ {threshhold} {commonT("kw")}
                </div>
            </div>
            <div className="card-column mt-3">
                <div className="card w-96 m-5">
                    <div className="flex flex-wrap items-center mb-8">
                        <div className="h-10 w-10 bg-green-main-opacity-80 rounded-full">
                            <HomeIcon className="h-8 w-8 ml-1 mt-0.5 p-1 opacity-80 text-green-main" />
                        </div>
                        <h5 className="font-bold ml-3">Load</h5>
                    </div>
                    <div className="grid grid-cols-2 twocolumns gap-x-5 sm:gap-x-10 mb-8">
                        <div className="text-center">
                            <h2 className="font-bold text-white text-3xl mb-1">20kW</h2>
                            <p className="lg:test text-white text-sm">Solar</p>
                        </div>
                        <div className="text-center">
                            <h2 className="font-bold text-white text-3xl mb-1">-</h2>
                            <p className="lg:test text-white text-sm">Battery Discharge</p>
                        </div>
                    </div>
                    <div className="grid grid-cols-2 twocolumns gap-x-5 sm:gap-x-10 mb-8">
                        <div className="text-center">
                            <h2 className="font-bold text-white text-3xl mb-1">10kW</h2>
                            <p className="lg:test text-white text-sm">Import From Grid</p>
                        </div>
                    </div>
                </div>
                <div className="card w-96 m-5">
                    <div className="flex flex-wrap items-center mb-8">
                        <div className="h-10 w-10 bg-blue-main-opacity-80 rounded-full">
                            <BatteryIcon className="h-8 w-8 ml-1 mt-0.5 p-0.5 opacity-80 text-blue-main" />
                        </div>
                        <h5 className="font-bold ml-3">Battery</h5>
                    </div>
                    <div className="grid grid-cols-2 twocolumns gap-x-5 sm:gap-x-10 mb-8">
                        <div className="text-center">
                            <h2 className="font-bold text-white text-3xl mb-1">80%</h2>
                            <p className="lg:test text-white text-sm">State of Charge</p>
                        </div>
                        <div className="text-center">
                            <h2 className="font-bold text-white text-3xl mb-1">20kW</h2>
                            <p className="lg:test text-white text-sm">Battery Power</p>
                        </div>
                    </div>
                    <div className="grid grid-cols-2 twocolumns gap-x-5 sm:gap-x-10 mb-8">
                        <div className="text-center">
                            <h2 className="font-bold text-white text-3xl mb-1">10kW</h2>
                            <p className="lg:test text-white text-sm">Import From Grid</p>
                        </div>
                        <div className="text-center">
                            <h2 className="font-bold text-white text-3xl mb-1">-</h2>
                            <p className="lg:test text-white text-sm">Discharging To</p>
                        </div>
                    </div>
                </div>
                <div className="card w-96 m-5">
                    <div className="flex flex-wrap items-center mb-8">
                        <div className="w-10 h-10 bg-yellow-main-opacity-80 rounded-full">
                            <SolarIcon className="h-8 w-8 ml-1 mt-1 p-0.5 opacity-80 text-yellow-main" />
                        </div>
                        <h5 className="font-bold ml-3">Solar</h5>
                    </div>
                    <div className="grid grid-cols-2 twocolumns gap-x-5 sm:gap-x-10 mb-8">
                        <div className="text-center">
                            <h2 className="font-bold text-white text-3xl mb-1">20kW</h2>
                            <p className="lg:test text-white text-sm">Solar</p>
                        </div>
                        <div className="text-center">
                            <h2 className="font-bold text-white text-3xl mb-1">-</h2>
                            <p className="lg:test text-white text-sm">Battery Discharge</p>
                        </div>
                    </div>
                    <div className="grid grid-cols-2 twocolumns gap-x-5 sm:gap-x-10 mb-8">
                        <div className="text-center">
                            <h2 className="font-bold text-white text-3xl mb-1">10kW</h2>
                            <p className="lg:test text-white text-sm">Import From Grid</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </>


}