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
    try {
        url = `${API_HOST}/${url}`
        const token = store.getState().user.value
        const _token = token ? { Authorization: `Bearer ${token}` } : {}
        const _contentType = contentType ? { "Content-Type": contentType } : {}
        axios({ method, url, data, token })
            .then((res) => {
                if (res.status === 200) onSuccess(res.data.data.token)
                else console.log(res.status, res);
            })
            .catch((err) => {
                onError(err.response.data.code);
                console.error(err)
            });
    } catch (err) {
        console.error(err)
    }
};

