import { connect } from "react-redux"
import React from "react"

import LoggedIn from "./LoggedIn"
import NonLoggedIn from "./NonLoggedIn"

const Main = props => props.webpages ? <LoggedIn /> : <NonLoggedIn />
const mapState = state => ({ webpages: state.user?.webpages })

export default connect(mapState)(Main)