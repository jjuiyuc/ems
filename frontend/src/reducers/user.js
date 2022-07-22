import { createSlice } from "@reduxjs/toolkit"

const lsUser = window.localStorage.user

export const userSlice = createSlice({
    name: "user",
    initialState: lsUser ? JSON.parse(lsUser) : {},
    reducers: {
        changeGateway: (state, action) => {
            state.gateways.forEach((g, i) => g.active = i === action.payload)
            localStorage.setItem("user", JSON.stringify(state))

            return state
        },
        logout: state => {
            localStorage.removeItem("user")
            return {}
        },
        updateUser: (state, action) => {
            localStorage.setItem("user", JSON.stringify(action.payload))
            return action.payload
        },
        updateUserProfile: (state, action) => {
            const userValue = { ...state, ...action.payload }
            localStorage.setItem("user", JSON.stringify(userValue))
            return userValue
        }
    }
})

export const { updateUser } = userSlice.actions
export default userSlice.reducer