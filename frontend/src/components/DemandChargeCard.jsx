import { Button, TextField, InputAdornment } from "@mui/material"
import { useState } from "react"
import { useTranslation } from "react-multi-lang"

import { validateNum } from "../utils/utils"

import { ReactComponent as DemandChargeIcon } from
    "../assets/icons/demand_charge_line.svg"

export default function DemandChargeCard(props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("settings." + string, params)
    const
        [inputValue, setInputValue] = useState("")
    const
        inputNum = (e) => {
            const num = e.target.value
            const isNum = validateNum(num)
            if (!isNum) return
            setInputValue(num)
        }
    return <div className="card mt-8">
        <div className="flex justify-between sm:col-span-2 items-center">
            <div className="flex items-center">
                <div
                    className="bg-gray-400-opacity-20 grid h-12 w-12
                place-items-center rounded-full">
                    <DemandChargeIcon className="h-8 text-gray-400 w-8" />
                </div>
                <h2 className="font-bold ml-4">{commonT("demandCharge")}</h2>
            </div>
            <Button
                // onClick={() => setTab(t)}
                key={"s-b-"}
                radius="pill"
                variant="contained">
                {commonT("save")}
            </Button>
        </div>
        <div className="flex items-center mt-12">
            <h5 className="mr-8">{props.title}</h5>
            <div className="mt-6">
                <TextField
                    label="Maximum Demand"
                    id="outlined-end-adornment"
                    value={inputValue}
                    onChange={inputNum}
                    InputProps={{
                        endAdornment:
                            <InputAdornment position="end">
                                {commonT("kw")}
                            </InputAdornment>
                    }}
                    focused
                />
            </div>
        </div>
    </div>
}