import {connect} from "react-redux"
import CssBaseline from "@mui/material/CssBaseline"
import React from "react"
import {setDefaultLanguage, setTranslations, withTranslation}
    from "react-multi-lang"
import {ThemeProvider} from "@mui/material/styles"

import theme from "./configs/theme"

import en from "./translations/en.json"
import zhtw from "./translations/zhtw.json"

import Main from "./containers/Main"

setDefaultLanguage(window.localStorage.lang || "zhtw")
setTranslations({en, zhtw})

function App () {
    return <ThemeProvider theme={theme}>
        <CssBaseline />
        <Main />
    </ThemeProvider>
}

const mapState = state => ({lang: state.lang.value})

export default connect(mapState)(withTranslation(App))