import { connect } from "react-redux"
import { Link, NavLink } from "react-router-dom"
import React from "react"
import { Fade, Tooltip } from "@mui/material"
import { useTranslation } from "react-multi-lang"
import AdvancedSettings from "@mui/icons-material/SettingsSuggest"

import { ReactComponent as AccountGroup } from "../assets/icons/group.svg"
import { ReactComponent as Analysis } from "../assets/icons/analysis.svg"
import { ReactComponent as Dashboard } from "../assets/icons/dashboard.svg"
import { ReactComponent as Menu } from "../assets/icons/menu.svg"
import { ReactComponent as Settings } from "../assets/icons/settings.svg"
import { ReactComponent as Timer } from "../assets/icons/timer.svg"

const navs = {
    dashboard:
        { icon: <Dashboard />, path: "dashboard" },
    analysis:
        { icon: <Analysis />, path: "analysis" },
    timeOfUseEnergy:
        { icon: <Timer />, path: "time-of-use", text: "timeOfUseEnergy" },
    accountManagementGroup:
        { icon: <AccountGroup />, path: "account-management-group", text: "accountManagementGroup" },
    settings:
        { icon: <Settings />, path: "settings", text: "settings" },
    advancedSettings:
        { icon: <AdvancedSettings fontSize="medium" />, path: "advanced-settings", text: "advancedSettings" }
}

function Sidebar(props) {
    const toggle = () =>
        props.updateSidebarStatus(isExpanded ? "collapse" : "expand")

    const
        { status } = props,
        isExpanded = status === "expand",
        transitionClasses = "duration-300 transition "
            + (isExpanded ? "opacity-100" : "opacity-0")

    const MenuIcon = isExpanded ? Menu : Menu
    const
        t = useTranslation(),
        navT = string => t("navigator." + string)

    const ListItem = ({ icon, path, text }, index) => {
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
    }
    const authNavList = props.webpages.map((authPage, i) =>
        <ListItem key={"sb-n-" + i} {...navs[authPage.name]} />)

    return <aside className="h-full overflow-x-hidden">
        <div className="w-60">
            <div className="border-b border-black-main flex h-20 items-center
                            justify-start px-7">
                <MenuIcon
                    className="cursor-pointer h-6 text-gray-200 w-6"
                    onClick={toggle} />
                <Link className={"ml-2 " + transitionClasses} to="/">
                    {/* <LogoWithName className="logo" /> */}
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