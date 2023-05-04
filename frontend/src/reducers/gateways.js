import { createSlice } from "@reduxjs/toolkit"

const lsGateways = window.localStorage.gateways

const setLocalStorage = state =>
    localStorage.setItem("gateways", JSON.stringify(state))

export const slice = createSlice({
    name: "gateways",
    initialState: lsGateways
        ? JSON.parse(lsGateways)
        : { active: { address: "", gatewayID: "" }, list: [] },
    reducers: {
        clear: state => {
            state = { active: { address: "", gatewayID: "" }, list: [] }
            localStorage.removeItem("gateways")
        },
        changeGateway: (state, action) => {
            state.active = action.payload
            setLocalStorage(state)
        },
        updateList: (state, action) => {
            state.list = action.payload
            setLocalStorage(state)
        }
    }
})

export default slice.reducer