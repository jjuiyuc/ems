import { connect } from "react-redux"
import { Link, NavLink } from "react-router-dom"
import React from "react"
import { Fade, Tooltip } from "@mui/material"
import { useTranslation } from "react-multi-lang"

import { ReactComponent as Analysis } from "../assets/icons/analysis.svg"
import { ReactComponent as Dashboard } from "../assets/icons/dashboard.svg"
import { ReactComponent as Demand } from "../assets/icons/demand_charge.svg"
import { ReactComponent as Economics } from "../assets/icons/economics.svg"
import { ReactComponent as Menu } from "../assets/icons/menu.svg"
import { ReactComponent as Logo } from "../assets/images/logo.svg"
import { ReactComponent as Resource } from "../assets/icons/resource.svg"
import { ReactComponent as Settings } from "../assets/icons/settings.svg"
import { ReactComponent as Timer } from "../assets/icons/timer.svg"

function Sidebar(props) {
    const toggle = () =>
        props.updateSidebarStatus(isExpanded ? "collapse" : "expand")

    const
        { status } = props,
        isExpanded = status === "expand",
        transitionClasses = "duration-300 transition "
            + (isExpanded ? "opacity-100" : "opacity-0")

    const
        navs = [
            { icon: <Dashboard />, path: "dashboard" },
            { icon: <Analysis />, path: "analysis" },
            { icon: <Timer />, path: "time-of-use", text: "timeOfUseEnergy" },
            // { icon: <Economics />, path: "economics" },
            { icon: <Demand />, path: "demand-charge", text: "demandCharge" },
            { icon: <Resource />, path: "energy-resources", text: "energyResources" },
            // { icon: <Settings />, path: "settings" }
        ],
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
        navList = navs.map((item, i) => navLink(item, i))

    return <aside className="h-full overflow-x-hidden">
        <div className="w-60">
            <div className="border-b border-black-main flex h-20 items-center
                            justify-start px-7">
                <Menu
                    className="cursor-pointer h-6 text-gray-200 w-6"
                    onClick={toggle} />
                <Link className={"ml-2 " + transitionClasses} to="/">
                    <Logo className="logo" />
                </Link>
            </div>
            <ul className="mt-6 sidebar-menu">{navList}</ul>
        </div>
    </aside>
}

const
    mapState = state => ({ status: state.sidebarStatus.value }),
    mapDispatch = dispatch => ({
        updateSidebarStatus: value => dispatch({
            type: "sidebarStatus/updateSidebarStatus", payload: value
        })
    })

export default connect(mapState, mapDispatch)(Sidebar)