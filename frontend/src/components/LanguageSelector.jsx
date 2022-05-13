import {connect} from "react-redux"
import {InputLabel, MenuItem, Select} from "@mui/material"
import LanguageIcon from "@mui/icons-material/Language"
import React from "react"
import {useTranslation} from "react-multi-lang"

import variables from "../configs/Variables"

function LanguageSelector (props) {
    const changeLang = e => props.updateLang(e.target.value)

    const
        t = useTranslation(),
        formT = string => t("form." + string)

    const
        {languages} = variables,
        langLabel = <><LanguageIcon /> {formT("language")}</>,
        langOpts = Object.keys(languages).map((key, i) =>
            <MenuItem key={"l-l-" + i} value={key}>{languages[key]}</MenuItem>)

    return <>
        <InputLabel id="lang-label">{langLabel}</InputLabel>
        <Select
            className="mb-8 text-left"
            id="lang"
            label={langLabel}
            labelId="lang-label"
            onChange={changeLang}
            value={props.lang}>
            {langOpts}
        </Select>
    </>
}

const
    mapState = state => ({lang: state.lang.value}),
    mapDispatch = dispatch => ({
        updateLang: value => dispatch({type: "lang/updateLang", payload: value})
    })

export default connect(mapState, mapDispatch)(LanguageSelector)