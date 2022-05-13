import {connect} from "react-redux"
import React from "react"

import LoggedIn from "./LoggedIn"
import NonLoggedIn from "./NonLoggedIn"

const Main = props => props.user.username ? <LoggedIn /> : <NonLoggedIn />
const mapState = state => ({user: state.user.value})

export default connect(mapState)(Main)