import { connect } from "react-redux"
import { Navigate, Route, Routes, useLocation, useNavigate } from "react-router-dom"
import React, { useEffect } from "react"
import { useTranslation } from "react-multi-lang"
import { Snackbar, Alert } from "@mui/material"

import logout from "../utils/logout"

import Sidebar from "../components/Sidebar"
import TopNav from "../components/TopNav"
// import Sample from "../configs/Sample"

import Account from "../pages/Account"
import AccountManagementGroup from "../pages/AccountManagementGroup"
import Analysis from "../pages/Analysis"
import Dashboard from "../pages/Dashboard"
import TimeOfUse from "../pages/TimeOfUse"
import AdvancedSettings from "../pages/AdvancedSettings"
import Settings from "../pages/Settings"

const routes = {
    dashboard: [
        <Route element={<Dashboard />} path="/dashboard" key="dashboard" />,
    ],
    analysis: [
        <Route element={<Analysis />} path="/analysis" key="analysis" />
    ],
    timeOfUseEnergy: [
        <Route element={<TimeOfUse />} path="/time-of-use" key="timeOfUseEnergy" />
    ],
    accountManagementGroup: [
        <Route
            element={<AccountManagementGroup />}
            path="/account-management-group"
            key="accountManagementGroup" />
    ],
    settings: [
        <Route
            element={<Settings />}
            path="/settings"
            key="settings" />
    ],
    advancedSettings: [
        <Route
            element={<AdvancedSettings />}
            path="/advanced-settings"
            key="advancedSettings" />
    ]
}
function LoggedIn(props) {
    const
        location = useLocation(),
        navigate = useNavigate(),
        isDashboard = location.pathname === "/dashboard",
        { sidebarStatus } = props,
        sidebarW = sidebarStatus === "expand" ? "w-60" : "w-20",
        t = useTranslation()

    const handleClose = () => {

        props.updateSnackbarMsg({
            msg: "", type: props.snackbarMsg.type
        })
    }
    useEffect(() => {
        if (new Date().getTime() > props.tokenExpiryTime) {
            logout()
            navigate("*")
        }
    }, [props.tokenExpiryTime, navigate])

    const authPages = props.webpages
        .map((authPage) => routes[authPage.name])
        .flat()

    return <div className="grid grid-rows-1fr-auto min-h-screen">
        <div className="align-items-stretch flex">
            <div className={"duration-300 transition-width " + sidebarW}>
                <Sidebar />
            </div>
            <div className="flex-auto bg-gray grid grid-rows-auto-1fr">
                <TopNav className="z-10" />
                <div className={"bg-gray-700 min-w-0 shadow-main z-0 "
                    + "pl-10 pr-8 py-8 lg:pl-25 lg:pr-20 lg:py-20"}>
                    <Routes>
                        <Route element={<Account />} path="/account" />
                        {authPages}
                        {authPages && <Route
                            element={<Navigate to={authPages[0].props.path} replace />}
                            path="*" />}
                    </Routes>
                </div>
            </div>
        </div>
        <Snackbar
            open={props.snackbarMsg.msg != ""}
            autoHideDuration={4000}
            onClose={handleClose}
            anchorOrigin={{ vertical: "top", horizontal: "center" }}
            sx={{ width: "500px" }}
        >
            <Alert
                severity={props.snackbarMsg.type}
                sx={{ width: "100%", fontSize: "16px" }}>
                {props.snackbarMsg.msg}
            </Alert>
        </Snackbar>
    </div >
}

const
    mapState = state => ({
        sidebarStatus: state.sidebarStatus.value,
        snackbarMsg: state.snackbarMsg,
        tokenExpiryTime: state.user.tokenExpiryTime,
        webpages: (state?.user?.webpages) || []
    }),
    mapDispatch = dispatch => ({
        updateSnackbarMsg: value =>
            dispatch({ type: "snackbarMsg/updateSnackbarMsg", payload: value })
    })
export default connect(mapState, mapDispatch)(LoggedIn)