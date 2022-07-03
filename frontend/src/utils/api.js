import axios from "axios"
import { API_HOST } from "../constant/env"
import store from "../store"

export const apiCall = async ({
    url = "/",
    method = "get",
    data = {},
    contentType = "",
    onSuccess = () => { },
    onError = () => { }
}) => {
    url = `${API_HOST}${url}`
    const { token } = store.getState().user.value

    try {
        const res = await axios({
            url,
            method,
            data,
            headers: {
                'Authorization': `Bearer ${token}`
            }
        })
        if (res.status === 200) {
            onSuccess(res.data)
            return res.data
        }
        else console.log(res.status, res)

    } catch (err) {
        onError(err.response?.data?.code);
        console.error(err)
    }
}
