import { createSlice } from "@reduxjs/toolkit"

const mockGateways = { "active": { "gatewayID": "MOCK-GW-000", "permissions": [{ "enabledAt": "2022-08-04T00:00:00Z", "enabledBy": null, "disabledAt": null, "disabledBy": null, "location": { "name": "PLACE 0", "address": "新竹縣XX鄉ＯＯＯ路" } }] }, "list": [{ "gatewayID": "MOCK-GW-000", "permissions": [{ "enabledAt": "2022-08-04T00:00:00Z", "enabledBy": null, "disabledAt": null, "disabledBy": null, "location": { "name": "PLACE 0", "address": "新竹縣XX鄉ＯＯＯ路" } }] }, { "gatewayID": "1E0BA27A8175AF978C49396BDE9D7A1E", "permissions": [{ "enabledAt": "2022-10-24T00:00:00Z", "enabledBy": null, "disabledAt": null, "disabledBy": null, "location": { "name": "PLACE 1", "address": "宜蘭縣ＸＸ鄉ＯＯＯ路" } }] }, { "gatewayID": "218F1623ADD8E739F7C6CBE62A7DF3C0", "permissions": [{ "enabledAt": "2023-01-19T00:00:00Z", "enabledBy": null, "disabledAt": null, "disabledBy": null, "location": { "name": "PLACE 2", "address": "台北市ＸＸ區ＯＯＯ路" } }] }, { "gatewayID": "3RT00000999000000001RUK", "permissions": [{ "enabledAt": "2023-06-20T16:00:00Z", "enabledBy": null, "disabledAt": null, "disabledBy": null, "location": { "name": "PLACE 3", "address": "台南市ＸＸ區ＯＯ路" } }] }] }


const setLocalStorage = state =>
    localStorage.setItem("gateways", JSON.stringify(state))

export const slice = createSlice({
    name: "gateways",
    initialState: mockGateways
        ? mockGateways
        : { active: { gatewayID: "" }, list: [] },
    reducers: {
        clear: state => {
            state = { active: { gatewayID: "" }, list: [] }
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

