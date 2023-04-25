import { useState } from "react"
import { useTranslation } from "react-multi-lang"

import DialogBox from "../components/DialogBox"
import TimeOfUseCard from "../components/TimeOfUseCard"

export default function AdvancedSettings(props) {
    const
        t = useTranslation()

    return <>
        <h1 className="mb-8">{t("navigator.advancedSettings")}</h1>
        <TimeOfUseCard />
    </>
}