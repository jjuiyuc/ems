import { connect } from "react-redux"
import { Navigate, Route, Routes, useLocation } from "react-router-dom"
import React, { useEffect } from "react"
import { useTranslation } from "react-multi-lang"

import logout from "../utils/logout"

import Sidebar from "../components/Sidebar"
import TopNav from "../components/TopNav"
import Sample from "../configs/Sample"

import Dashboard from "../pages/Dashboard"
import DemandCharge from "../pages/DemandCharge"
import EnergyResoucesBattery from "../pages/EnergyResoucesBattery"
import EnergyResoucesSolar from "../pages/EnergyResoucesSolar"
import Analysis from "../pages/Analysis"
import TimeOfUse from "../pages/TimeOfUse"

function LoggedIn(props) {
    const
        location = useLocation(),
        isDashboard = location.pathname === "/dashboard",
        { sidebarStatus } = props,
        sidebarW = sidebarStatus === "expand" ? "w-60" : "w-20",
        t = useTranslation()

    useEffect(() => {
        if (new Date().getTime() > props.tokenExpiryTime) logout()
    })

    return <div className="grid grid-rows-1fr-auto min-h-screen">
        <div className="align-items-stretch flex">
            <div className={"duration-300 transition-width " + sidebarW}>
                <Sidebar />
            </div>
            <div className="flex-auto bg-gray grid grid-rows-auto-1fr">
                <TopNav className="z-10" />
                <div className={"bg-gray-700 min-w-0 shadow-main z-0 "
                    + "pl-10 pr-8 py-8 lg:pl-25 lg:pr-20 lg:py-20"
                    + (isDashboard ? " bg-image-grid" : "")}>
                    <Routes>
                        <Route element={<Dashboard />} path="/dashboard" />
                        <Route element={<Analysis />} path="/analysis" />
                        <Route element={<TimeOfUse />} path="/time-of-use" />
                        <Route element={<Sample />} path="/economics" />
                        <Route
                            element={<DemandCharge />}
                            path="/demand-charge" />
                        <Route
                            element={<EnergyResoucesBattery />}
                            path="/energy-resources/battery" />
                        <Route element={<Sample />} path="/settings" />
                        <Route
                            element={<EnergyResoucesSolar />}
                            path="/energy-resources/solar" />
                        <Route
                            element={<Navigate
                                replace
                                to="/energy-resources/solar" />}
                            path="/energy-resources" />
                        <Route
                            element={<Navigate to="/dashboard" replace />}
                            path="*" />
                    </Routes>
                </div>
            </div>
        </div>
        <footer className="bg-gray-800 flex items-center justify-between
                            text-center text-gray-300 text-sm
                            h-14 md:h-20 px-10 md:px-20">
            <span className="font-mono ml-4 text-gray-500 text-13px">
                {import.meta.env.VITE_APP_VERSION}
            </span>
            {t("common.copyright")}
        </footer>
    </div>
}

const mapState = state => ({
    sidebarStatus: state.sidebarStatus.value,
    tokenExpiryTime: state.user.tokenExpiryTime
})

export default connect(mapState)(LoggedIn)