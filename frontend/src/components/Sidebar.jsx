import { connect } from "react-redux"
import { Link, NavLink } from "react-router-dom"
import React from "react"
import { Fade, Tooltip } from "@mui/material"
import { useTranslation } from "react-multi-lang"

import { ReactComponent as AccountGroup } from "../assets/icons/group.svg"
import { ReactComponent as AccountUser } from "../assets/icons/user.svg"
import { ReactComponent as Analysis } from "../assets/icons/analysis.svg"
import { ReactComponent as Dashboard } from "../assets/icons/dashboard.svg"
import { ReactComponent as Demand } from "../assets/icons/demand_charge.svg"
import { ReactComponent as Economics } from "../assets/icons/economics.svg"
import { ReactComponent as Field } from "../assets/icons/field_management.svg"
import { ReactComponent as Menu } from "../assets/icons/menu.svg"
import { ReactComponent as Logo } from "../assets/images/logo.svg"
import { ReactComponent as LogoWithName } from "../assets/images/logoWithName.svg"
import { ReactComponent as Resource } from "../assets/icons/resource.svg"
import { ReactComponent as Settings } from "../assets/icons/settings.svg"
import { ReactComponent as AdvancedSettings } from "../assets/icons/settings.svg"
import { ReactComponent as Timer } from "../assets/icons/timer.svg"

const navs = {
    dashboard:
        { icon: <Dashboard />, path: "dashboard" },
    analysis:
        { icon: <Analysis />, path: "analysis" },
    timeOfUseEnergy:
        { icon: <Timer />, path: "time-of-use", text: "timeOfUseEnergy" },
    economics:
        { icon: <Economics />, path: "economics" },
    demandCharge:
        { icon: <Demand />, path: "demand-charge", text: "demandCharge" },
    energyResources:
        { icon: <Resource />, path: "energy-resources", text: "energyResources" },
    fieldManagement:
        { icon: <Field />, path: "field-management", text: "fieldManagement" },
    accountManagementGroup:
        { icon: <AccountGroup />, path: "account-management-group", text: "accountManagementGroup" },
    accountManagementUser:
        { icon: <AccountUser />, path: "account-management-user", text: "accountManagementUser" },
    settings:
        { icon: <Settings />, path: "settings", text: "settings" },
    advancedSettings:
        { icon: <AdvancedSettings />, path: "advanced-settings", text: "advancedSettings" }
}

function Sidebar(props) {
    const toggle = () =>
        props.updateSidebarStatus(isExpanded ? "collapse" : "expand")

    const
        { status } = props,
        isExpanded = status === "expand",
        transitionClasses = "duration-300 transition "
            + (isExpanded ? "opacity-100" : "opacity-0")

    const MenuIcon = isExpanded ? Menu : Logo
    const
        t = useTranslation(),
        navT = string => t("navigator." + string),
        navLink = ({ icon, path, text }, index) => {
            const name = navT(text || path)

            return <li key={"sb-n-" + index}>
                <NavLink
                    className={({ isActive }) => isActive ? " active" : ""}
                    to={"/" + path}>
                    <Tooltip
                        arrow
                        disableHoverListener={isExpanded}
                        placement="left"
                        TransitionComponent={Fade}
                        TransitionProps={{ timeout: 300 }}
                        title={name}>
                        <div>{icon}</div>
                    </Tooltip>
                    <span className={transitionClasses}>{name}</span>
                </NavLink>
            </li>
        },
        authNavList = props.webpages
            .map((authPage, i) => navLink(navs[authPage.name], i))

    return <aside className="h-full overflow-x-hidden">
        <div className="w-60">
            <div className="border-b border-black-main flex h-20 items-center
                            justify-start px-7">
                <MenuIcon
                    className="cursor-pointer h-6 text-gray-200 w-6"
                    onClick={toggle} />
                <Link className={"ml-2 " + transitionClasses} to="/">
                    <LogoWithName className="logo" />
                </Link>
            </div>
            <ul className="mt-6 sidebar-menu">{authNavList}</ul>
        </div>
    </aside>
}
const
    mapState = state => ({
        status: state.sidebarStatus.value,
        webpages: (state.user.webpages || [])
    }),
    mapDispatch = dispatch => ({
        updateSidebarStatus: value => dispatch({
            type: "sidebarStatus/updateSidebarStatus", payload: value
        })
    })

export default connect(mapState, mapDispatch)(Sidebar)