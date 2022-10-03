import { Button, Stack } from "@mui/material"
import { useState } from "react"
import { useTranslation } from "react-multi-lang"
import moment from "moment"

import { ReactComponent as BatteryIcon } from "../assets/icons/battery.svg"
import { ReactComponent as DemandCharge }
    from "../assets/icons/demand_charge_line.svg"
import { ReactComponent as TimerIcon } from "../assets/icons/timer.svg"

export default function Settings() {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("settings." + string, params)

    return <>
        <h1 className="mb-8">{pageT("settings")}</h1>
        <div className="card">
            <div className="flex justify-between sm:col-span-2 items-center">
                <div className="flex items-center">
                    <div
                        className="bg-gray-400-opacity-20 grid h-12 w-12
                            place-items-center rounded-full">
                        <BatteryIcon className="h-8 text-gray-400 w-8" />
                    </div>
                    <h2 className="font-bold ml-4">{commonT("battery")}</h2>
                </div>
                <div>
                    <Button
                        // onClick={() => setTab(t)}
                        key={"s-b-"}
                        radius="pill"
                        variant="contained">
                        {pageT("save")}
                    </Button>
                </div>
            </div>
            <div className="lg:grid grid-cols-3-auto  lg:column-separator">
                <div className="col-span-2 border-gray-400 lg:border-0 py-4 lg:py-0">
                    <h6 className="font-bold text-white">888</h6>
                    <div className="flex justify-between items-center">
                    </div>
                </div>

            </div>
        </div>
    </>
}