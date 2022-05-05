import React from "react"
import CssBaseline from "@mui/material/CssBaseline"
import {ThemeProvider} from "@mui/material/styles"

import Sample from "./configs/Sample"
import theme from "./configs/theme"

function App() {
    return <ThemeProvider theme={theme}>
        <CssBaseline />
        <Sample />
    </ThemeProvider>
}

export default App