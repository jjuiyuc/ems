import axios from "axios"
import { API_HOST } from "../constant/env"
import store from "../store"


export const apiCall = ({
    url = "/",
    method = "post",
    data = {},
    contentType = "",
    onSuccess = () => {},
    onError = () => {}
}) => {
    url = `${API_HOST}${url}`
    const token = store.getState().user.value
    axios({ method, url, data, token })
        .then((res) => {
            if (res.status === 200) onSuccess(res)
            else console.log(res.status, res);
        })
        .catch((err) => {
            onError(err.response?.data?.code);
            console.error(err)
        })
}

