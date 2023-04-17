import { Button, Slider, Switch } from "@mui/material"
import { useState } from "react"
import { useTranslation } from "react-multi-lang"

import DialogBox from "../components/DialogBox"
import TimeOfUseCard from "../components/TimeOfUseCard"

import variables from "../configs/variables"

import { ReactComponent as BatteryIcon } from "../assets/icons/battery.svg"

const { colors } = variables

export default function AdvancedSettings(props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("settings." + string, params)

    return <>
        <h1 className="mb-8">{pageT("settings")}</h1>
        <TimeOfUseCard />

    </>
}