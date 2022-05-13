import {createSlice} from "@reduxjs/toolkit"

const lsUser = window.localStorage.user

export const userSlice = createSlice({
    name: "user",
    initialState: {value: lsUser ? JSON.parse(lsUser) : {}},
    reducers: {
        updateUser: (state, action) => {
            localStorage.setItem("user", JSON.stringify(action.payload))
            state.value = action.payload
        }
    }
})

export const {updateUser} = userSlice.actions
export default userSlice.reducer