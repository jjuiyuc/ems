import { Route, Routes } from "react-router-dom"
import React from "react"
import { useTranslation } from "react-multi-lang"

import smartGrid from "../assets/images/smartGrids.svg"

import ForgotPassword from "../pages/ForgotPassword"
import LogIn from "../pages/LogIn"
import ResetPassword from "../pages/ResetPassword"

function NonLoggedIn() {
    const t = useTranslation()

    const Version = props =>
        <div className={"font-mono p-2 text-center text-gray-400 text-xs"
            + (props.className ? " " + props.className : "")}>
            {import.meta.env.VITE_APP_VERSION}
        </div>

    const brand = <>
        <h3 className="mt-4 md:mt-8 text-xl xl:text-2xl">
            <span className="mr-2">
                Enjoy Life &
            </span>
            <span className="text-primary-main">
                The Potential of Smart Grids
            </span>
        </h3>
    </>

    return <div className="flex min-h-screen items-stretch">
        <div className="bg-gray-800 flex-col grid-rows-1fr-auto items-center
                        justify-center py-4 w-3/8
                        hidden md:grid
                        px-8 xl:px-16">
            <div>
                <h2 className="mb-12 ml-5">Energy Management</h2>
                <div className="mb-32 ml-5">{brand}</div>
                {/* <img src={smartGrid} /> */}
            </div>
            <Version />
        </div>
        <div className="flex flex-col items-center justify-center px-6
                        py-6 h-sm:py-12 h-md:py-18 h-lg:py-24
                        w-full md:w-5/8">
            <div className="flex flex-col items-center mb-5 md:hidden">
                {brand}
            </div>
            <div className="bg-gray-800 gap-10 grid grid-rows-1fr-auto h-full
                            items-start max-h-4xl max-w-xl pb-5 rounded-2.5xl
                            text-center w-full
                            pt-6 sm-sm:pt-24 sm-md:pt-32 sm-lg:pt-40
                            px-6 md:px-12">
                <Routes>
                    <Route
                        element={<ForgotPassword />}
                        path="/forgotPassword" />
                    <Route
                        element={<ResetPassword />}
                        path="/handle-reset-link" />
                    <Route element={<LogIn />} path="*" />
                </Routes>
            </div>
            <Version className="md:hidden" />
        </div>
    </div>
}

export default NonLoggedIn