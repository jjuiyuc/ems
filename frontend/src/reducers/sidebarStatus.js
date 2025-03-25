import {createSlice} from "@reduxjs/toolkit"

export const sidebarStatusSlice = createSlice({
    name: "sidebarStatus",
    initialState: {value: window.localStorage.sidebarStatus || "expand"},
    reducers: {
        updateSidebarStatus: (state, action) => {
            localStorage.setItem("sidebarStatus", action.payload)
            state.value = action.payload
        }
    }
})

export default sidebarStatusSlice.reducer