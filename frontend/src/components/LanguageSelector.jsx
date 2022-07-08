import {connect} from "react-redux"
import {MenuItem, Select} from "@mui/material"
import React from "react"

import variables from "../configs/variables"

function LanguageSelector (props) {
    const changeLang = e => props.updateLang(e.target.value)

    const
        {languages} = variables,
        langOpts = Object.keys(languages).map((key, i) =>
            <MenuItem key={"l-l-" + i} value={key}>{languages[key]}</MenuItem>)

    return <Select
        className={props.className}
        id={props.id}
        label={props.label}
        labelId={props.labelId}
        onChange={changeLang}
        size={props.size}
        value={props.lang}>
        {langOpts}
    </Select>
}

const
    mapState = state => ({lang: state.lang.value}),
    mapDispatch = dispatch => ({
        updateLang: value => dispatch({type: "lang/updateLang", payload: value})
    })

export default connect(mapState, mapDispatch)(LanguageSelector)