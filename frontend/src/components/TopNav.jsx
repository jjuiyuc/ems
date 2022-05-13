import {Button, Grid} from "@mui/material"
import {connect} from "react-redux"
import React from "react"

function TopNav (props) {
    const logout = () => props.updateUser({})

    return <Grid container justifyContent="flex-end">
        <Button color="primary" onClick={logout} variant="contained">
            Log Out
        </Button>
    </Grid>
}

const
    mapState = state => ({user: state.user.value}),
    mapDispatch = dispatch => ({
        updateUser: value => dispatch({
            type: "user/updateUser", payload: value
        })
    })

export default connect(mapState, mapDispatch)(TopNav)