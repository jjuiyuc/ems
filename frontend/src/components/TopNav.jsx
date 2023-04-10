import {
    Button, Divider, FormControl, ListItemIcon, Menu, MenuItem,
    OutlinedInput, Select
} from "@mui/material"
import { NavLink } from "react-router-dom"
import { connect } from "react-redux"
import { Language as LanguageIcon, Logout as LogoutIcon }
    from "@mui/icons-material"
import React, { useState } from "react"
import { useTranslation } from "react-multi-lang"

import LanguageSelector from "./LanguageSelector"
import logout from "../utils/logout"

import { ReactComponent as AlertIcon } from "../assets/icons/alert_default.svg"
import { ReactComponent as LocationIcon } from "../assets/icons/location.svg"
import { ReactComponent as UserCircleIcon } from "../assets/icons/profile.svg"
import { ReactComponent as UserIcon } from "../assets/icons/user.svg"

function TopNav(props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string)
    const
        locationNameData = [
            {
                value: "Serenegray",
                label: "Serenegray",
            }, {
                value: "Cht_Miaoli",
                label: "Cht_Miaoli",
            }],
        gatewayData = [
            {
                value: "0E0BA27A8175AF978C49396BDE9D7A1E",
                label: "0E0BA27A8175AF978C49396BDE9D7A1E",
            },
            {
                value: "018F1623ADD8E739F7C6CBE62A7DF3C0",
                label: "018F1623ADD8E739F7C6CBE62A7DF3C0",
            }
        ]

    // const routes = [
    //     { icon: <Account />, path: "account" }
    // ],


    const
        [menuAnchorEl, setMenuAnchorEl] = useState(null),
        [locationName, setLocationName] = useState(""),
        [gatewayID, setGatewayID] = useState("")

    const
        closeMenu = () => setMenuAnchorEl(null),
        openMenu = e => setMenuAnchorEl(e.currentTarget)

    const
        { className } = props,
        { name } = props.user,
        menuPaperProps = {
            sx: {
                borderTopLeftRadius: 0,
                borderTopRightRadius: 0,
                "&.MuiPaper-root": { marginTop: "1.125rem" }
            }
        }

    const
        locationHandleChange = (event) => {
            setLocationName(event.target.value)
        },
        gatewayHandleChange = (event) => {
            setGateway(event.target.value)
        }
    const containerClasses = "border-b border-black-main bg-gray-900 flex "
        + "flex-col h-20 items-end overflow-visible"
        + (className ? " " + className : "")

    return <div className={containerClasses}>
        <div className="flex flex-row-reverse h-20 items-center
                        justify-between px-12 z-10 w-full">
            <div className="flex h-20 items-center">
                {/* <AlertIcon className="h-8 w-8 opacity-30" /> */}
                <Button onClick={openMenu} sx={{ marginLeft: "1.5rem" }}>
                    <UserCircleIcon className="h-8 mr-2 w-8" />
                    {name}
                </Button>
            </div>
            <div className="items-center  md:flex">
                <FormControl sx={{ m: 1, minWidth: 200 }}>
                    <Select
                        id="location-name"
                        label={commonT("locationName")}
                        value={locationName}
                        size="small"
                        defaultValue={locationNameData[0]}
                        onChange={locationHandleChange}
                        input={<OutlinedInput />}
                        inputProps={{ "aria-label": "Without label" }}
                        displayEmpty
                    >
                        <MenuItem disabled value="">
                            <em>{commonT("locationName")}</em>
                        </MenuItem>
                        {locationNameData.map((option) => (
                            <MenuItem key={option.value} value={option.value}>
                                {option.label}
                            </MenuItem>
                        ))}
                    </Select>
                </FormControl>
                <FormControl sx={{ m: 1, minWidth: 300 }}>
                    <Select
                        id="location-name"
                        label={commonT("gatewayID")}
                        value={gatewayID}
                        size="small"
                        onChange={gatewayHandleChange}
                        input={<OutlinedInput />}
                        inputProps={{ "aria-label": "Without label" }}
                        displayEmpty >
                        <MenuItem disabled value="">
                            <em>{commonT("gatewayID")}</em>
                        </MenuItem>
                        {gatewayData.map((option) => (
                            <MenuItem key={option.value} value={option.value}>
                                {option.label}
                            </MenuItem>
                        ))}
                    </Select>
                </FormControl>
            </div>
        </div>
        <Menu
            anchorEl={menuAnchorEl}
            anchorOrigin={{ horizontal: "right", vertical: "bottom" }}
            onClick={closeMenu}
            onClose={closeMenu}
            open={menuAnchorEl !== null}
            PaperProps={menuPaperProps}
            transformOrigin={{ horizontal: "right", vertical: "top" }}>
            <MenuItem>
                <ListItemIcon><LanguageIcon /></ListItemIcon>
                <LanguageSelector id="lang" size="small" />
            </MenuItem>
            <Divider />
            <MenuItem className="top-menu">
                <NavLink
                    className={({ isActive }) => isActive ? " active" : ""}
                    to="/account">
                    <ListItemIcon><UserIcon /></ListItemIcon>
                    {commonT("account")}
                </NavLink>

            </MenuItem>
            <Divider />
            <MenuItem onClick={logout}>
                <ListItemIcon><LogoutIcon /></ListItemIcon>
                {commonT("logOut")}
            </MenuItem>
        </Menu>
    </div>
}

const mapState = state => ({
    address: state.gateways.active.address,
    lang: state.lang.value,
    user: state.user
})

export default connect(mapState)(TopNav)