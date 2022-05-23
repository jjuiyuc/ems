import {Navigate, Route, Routes, useLocation} from "react-router-dom"
import {connect} from "react-redux"
import React, { useEffect } from "react"

import Sidebar from "../components/Sidebar"
import TopNav from "../components/TopNav"
import Sample from "../configs/Sample"

import Dashboard from "../pages/Dashboard"

function LoggedIn (props) {
    const
        location = useLocation(),
        isDashboard = location.pathname === "/dashboard",
        {sidebarStatus} = props,
        sidebarW = sidebarStatus === "expand" ? "w-60" : "w-20"

    return <div className="align-items-stretch flex min-h-screen">
        <div className={"duration-300 transition-width " + sidebarW}>
            <Sidebar />
        </div>
        <div className="flex-auto bg-gray grid grid-rows-auto-1fr">
            <TopNav className="z-10" />
            <div className={"bg-gray-700 shadow-main pl-25 pr-20 py-20 z-0"
                            + (isDashboard ? " bg-image-grid" : "")}>
                <Routes>
                    <Route element={<Dashboard />} path="/dashboard" />
                    <Route element={<Sample />} path="/analysis" />
                    <Route element={<Sample />} path="/time-of-use" />
                    <Route element={<Sample />} path="/economics" />
                    <Route element={<Sample />} path="/demand-charge" />
                    <Route element={<Sample />} path="/energy-resources" />
                    <Route element={<Sample />} path="/settings" />
                    <Route
                        path="*"
                        element={<Navigate to="/dashboard" replace />}/>
                </Routes>
            </div>
        </div>
    </div>
}

const mapState = state => ({sidebarStatus: state.sidebarStatus.value})

export default connect(mapState)(LoggedIn)