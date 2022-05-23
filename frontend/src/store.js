import {configureStore} from "@reduxjs/toolkit"

import lang from "./reducers/lang"
import sidebarStatus from "./reducers/sidebarStatus"
import user from "./reducers/user"

export default configureStore ({reducer: {lang, sidebarStatus, user}})