import {configureStore} from "@reduxjs/toolkit"

import gateways from "./reducers/gateways"
import lang from "./reducers/lang"
import sidebarStatus from "./reducers/sidebarStatus"
import user from "./reducers/user"

export default configureStore({reducer: {gateways, lang, sidebarStatus, user}})