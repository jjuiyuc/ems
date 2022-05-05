import {AccessAlarm} from "@mui/icons-material"
import Box from "@mui/material/Box"
import Button from "@mui/material/Button"
import React from "react"
import TextField from "@mui/material/TextField"
import Typography from "@mui/material/Typography"

export default function Sample () {
    return <>
        <Box m={3}>
            <Button color="purple" variant="contained">Save</Button>
            <Button variant="contained">Save</Button>
            <Button radius="pill" variant="contained">Save</Button>
            <Button size="x-large" variant="contained">Save</Button>
            <TextField
                InputProps={{endAdornment: <AccessAlarm />}}
                label="Time"
                variant="outlined" />
        </Box>
        <Typography color="purple.main">text</Typography>
        {process.env.npm_package_version} | {import.meta.env.VITE_APP_TEST}
    </>
}