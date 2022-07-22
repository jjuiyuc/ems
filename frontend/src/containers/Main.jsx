import {connect} from "react-redux"
import React from "react"

import LoggedIn from "./LoggedIn"
import NonLoggedIn from "./NonLoggedIn"

const Main = props => props.username ? <LoggedIn /> : <NonLoggedIn />
const mapState = state => ({username: state.user.username})

export default connect(mapState)(Main)