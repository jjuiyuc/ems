import { Button, Slider, Switch } from "@mui/material"
import { useState } from "react"
import { useTranslation } from "react-multi-lang"

import { ValidateNum } from "../utils/utils"

import DemandChargeCard from "../components/DemandChargeCard"
import DialogBox from "../components/DialogBox"
import SettingCard from "../components/SettingCard"
import TimeOfUseCard from "../components/TimeOfUseCard"
import variables from "../configs/variables"

import { ReactComponent as BatteryIcon } from "../assets/icons/battery.svg"

const { colors } = variables

export default function Settings(props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("settings." + string, params)
    const
        [reservedForGridOutage, setReservedForGridOutage] = useState(85),
        [availableRegularUsage, setAvailableRegularUsage] = useState(15),
        [backupReserve, setBackupReserve] = useState(100),
        [clockDataset, setClockDataset] = useState({
            data: [], backgroundColor: []
        }),
        [maxDemandCapacity, setMaxDemandCapacity] = useState("")
    // [midPeak, setMidPeak] = useState({
    //     types: [
    //         { kwh: 7.5, percentage: 15, type: "grid" },
    //         { kwh: 30, percentage: 60, type: "solar" },
    //         { kwh: 12.5, percentage: 25, type: "battery" },
    //     ],
    //     kwh: 50
    // }),
    // [onPeak, setOnPeak] = useState({
    //     types: [
    //         { kwh: 5, percentage: 10, type: "grid" },
    //         { kwh: 52, percentage: 50, type: "solar" },
    //         { kwh: 20, percentage: 40, type: "battery" },
    //     ],
    //     kwh: 50
    // }),
    // [offPeak, setOffPeak] = useState({
    //     types: [
    //         { kwh: 10, percentage: 18, type: "grid" },
    //         { kwh: 25, percentage: 41, type: "solar" },
    //         { kwh: 25, percentage: 41, type: "battery" },
    //     ],
    //     kwh: 60
    // }),
    // [superOffPeak, setSuperOffPeak] = useState({
    //     types: [
    //         { kwh: 21, percentage: 35, type: "grid" },
    //         { kwh: 24, percentage: 40, type: "solar" },
    //         { kwh: 15, percentage: 25, type: "battery" },
    //     ],
    //     kwh: 60
    // }),


    return <>
        <h1 className="mb-8">{pageT("settings")}</h1>
        <div className="card mb-8">
            <div className="flex justify-between sm:col-span-2 items-center">
                <div className="flex items-center">
                    <div
                        className="bg-gray-400-opacity-20 grid h-12 w-12
                            place-items-center rounded-full">
                        <BatteryIcon className="h-8 text-gray-400 w-8" />
                    </div>
                    <h2 className="font-bold ml-4">{commonT("battery")}</h2>
                </div>
                <Button
                    // onClick={() => setTab(t)}
                    key={"s-b-"}
                    radius="pill"
                    variant="contained">
                    {pageT("save")}
                </Button>
            </div>
            <div className="lg:grid grid-cols-3 mt-12">
                <div className="col-span-2">
                    <h5 className="font-bold">{pageT("backupReserve")}</h5>
                    <div className=" border-r border-gray-400 border-solid pr-12">
                        <div className="flex items-center ">
                            <SettingCard
                                data={reservedForGridOutage}
                                title={pageT("reservedForGridOutage")} />
                            <h4 className="mx-6">+</h4>
                            <SettingCard
                                data={availableRegularUsage}
                                title={pageT("availableRegularUsage")} />
                            <h4 className="mx-6">=</h4>
                            <SettingCard
                                data={backupReserve}
                                title={pageT("backupReserve")} />
                        </div>
                        <div>
                            <Slider defaultValue={50} aria-label="Default" valueLabelDisplay="auto" />
                            <div className="flex justify-between">
                                <p className="text-11px">{pageT("reservedForGrid")}</p>
                                <p className="text-11px">{pageT("regularUsage")}</p>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="lg:grid grid-cols-auto pl-12">
                    <h5 className="font-bold mt-9">{pageT("chargingSources")}</h5>
                    <div className="grid grid-cols-2 gap-5">
                        <div className="subCard bg-gray-700">
                            <p className="text-13px">{commonT("grid")}</p>
                            <Switch />
                        </div>
                        <div className="subCard bg-gray-700">
                            <p className="text-13px">{commonT("solar")}</p>
                            <Switch />
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <DialogBox />
        <TimeOfUseCard data={clockDataset} />
        <DemandChargeCard
            data={maxDemandCapacity}
            title={pageT("maximumDemandCapacity")}
        />
    </>
}