import { Button, Slider, Switch } from "@mui/material"
import { useState } from "react"
import { useTranslation } from "react-multi-lang"

import DemandChargeCard from "../components/DemandChargeCard"
import DialogBox from "../components/DialogBox"
import PowerOutageCard from "../components/PowerOutageCard"
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
        [reservedForGridOutage, setReservedForGridOutage] = useState(1),
        [availableRegularUsage, setAvailableRegularUsage] = useState(100),
        [backupReserve, setBackupReserve] = useState(100),
        [clockDataset, setClockDataset] = useState({
            data: [], backgroundColor: []
        }),
        [maxDemandCapacity, setMaxDemandCapacity] = useState("")

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
                <DialogBox
                    triggerName={pageT("save")}
                    leftButtonName={pageT("cancel")}
                    rightButtonName={pageT("turnOff")}
                />
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
                            <Slider defaultValue={1} aria-label="Default" valueLabelDisplay="auto"
                                onChange={(e) => {
                                    setReservedForGridOutage(Number(e.target.value))
                                    setAvailableRegularUsage(100 - Number(e.target.value))
                                }} />
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
        <PowerOutageCard />
        <TimeOfUseCard data={clockDataset} />
        <DemandChargeCard
            data={maxDemandCapacity}
            title={pageT("maximumDemandCapacity")}
        />
    </>
}