import { createSlice } from "@reduxjs/toolkit"

export const snackbarMsgSlice = createSlice({
    name: "snackbarMsg",
    initialState: {
        msg: "", type: "success"
    },
    reducers: {
        updateSnackbarMsg: (state, action) => {
            localStorage.setItem("snackbarMsg", action.payload)
            state.msg = action.payload.msg
            state.type = action.payload.type
        }
    }
})

export default snackbarMsgSlice.reducer