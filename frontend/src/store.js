import {configureStore} from "@reduxjs/toolkit"

import lang from "./reducers/lang"
import user from "./reducers/user"

export default configureStore ({reducer: {lang, user}})