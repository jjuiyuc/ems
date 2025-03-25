import {BrowserRouter} from "react-router-dom"
import {Provider} from "react-redux"
import React from "react"
import ReactDOM from "react-dom/client"

import store from "./store"
import "./assets/css/main.scss"

import App from "./App"

ReactDOM.createRoot(document.getElementById("root")).render(
    <React.StrictMode>
        <Provider store={store}>
            <BrowserRouter>
                <App />
            </BrowserRouter>
        </Provider>
    </React.StrictMode>
)