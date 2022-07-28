import store from "../store"

const logout = () => {
    store.dispatch({type: "gateways/clear"})
    store.dispatch({type: "user/clear"})
}

export default logout