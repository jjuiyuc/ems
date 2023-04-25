import axios from "axios"
import { API_HOST } from "../constant/env"
import store from "../store"

export const apiCall = async ({
    contentType = "",
    data = {},
    method = "GET",
    onComplete = () => { },
    onError = () => { },
    onStart = () => { },
    onSuccess = () => { },
    url = "/"
}) => {
    const { protocol } = window.location, { token } = store.getState().user

    url = `${protocol}//${API_HOST}${url}`

    onStart()

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
            console.log(onSuccess)

            onSuccess(res.data)
            return res.data
        }
    }
    catch (err) {
        let result = err.code

        if (err.response) {
            result = err.response.status

            if (typeof (err.response.data) === "object") {
                result = err.response.data.code
            }
        }

        onError(result)
    }
    finally {
        onComplete()
    }
}
