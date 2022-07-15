import { createSlice } from "@reduxjs/toolkit"

const lsUser = window.localStorage.user

export const userSlice = createSlice({
    name: "user",
    initialState: { value: lsUser ? JSON.parse(lsUser) : {} },
    reducers: {
        logout: state => {
            localStorage.removeItem("user")
            state.value = {}
        },
        updateUser: (state, action) => {
            localStorage.setItem("user", JSON.stringify(action.payload))
            state.value = action.payload
        },
        updateUserProfile: (state, action) => {
            const userValue = { ...state.value, ...action.payload }
            localStorage.setItem("user", JSON.stringify(userValue))
            state.value = userValue
        }
    }
})

export const { updateUser } = userSlice.actions
export default userSlice.reducer