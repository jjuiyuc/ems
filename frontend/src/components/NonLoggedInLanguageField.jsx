import {InputLabel} from "@mui/material"
import LanguageIcon from "@mui/icons-material/Language"
import React from "react"
import {useTranslation} from "react-multi-lang"

import LanguageSelector from "./LanguageSelector"

function LanguageField () {
    const
        t = useTranslation(),
        langLabel = <><LanguageIcon /> {t("form.language")}</>

    return <>
        <InputLabel id="lang-label">{langLabel}</InputLabel>
        <LanguageSelector
            className="mb-8 text-left"
            id="lang"
            label={langLabel}
            labelId="lang-label" />
    </>
}

export default LanguageField