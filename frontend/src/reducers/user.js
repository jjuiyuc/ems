import { createSlice } from "@reduxjs/toolkit"

const lsUser = window.localStorage.user

const setLocalStorage = state =>
    localStorage.setItem("user", JSON.stringify(state))

export const userSlice = createSlice({
    name: "user",
    initialState: lsUser ? JSON.parse(lsUser) : {},
    reducers: {
        clear: () => {
            localStorage.removeItem("user")
            return {}
        },
        updateUser: (state, action) => {
            setLocalStorage(action.payload)
            return action.payload
        },
        updateUserProfile: (state, action) => {
            const user = { ...state, ...action.payload }
            setLocalStorage(user)
            return user
        }
    }
})

export default userSlice.reducer