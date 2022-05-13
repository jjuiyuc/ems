import {createSlice} from "@reduxjs/toolkit"
import {setLanguage} from "react-multi-lang"

export const langSlice = createSlice({
    name: "lang",
    initialState: {value: window.localStorage.lang || "zhtw"},
    reducers: {
        updateLang: (state, action) => {
            localStorage.setItem("lang", action.payload)
            state.value = action.payload

            setLanguage(action.payload)
        }
    }
})

export const {updateLang} = langSlice.actions
export default langSlice.reducer