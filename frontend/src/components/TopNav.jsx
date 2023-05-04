import {
    Button, Divider, FormControl, ListItemIcon, Menu, MenuItem,
    OutlinedInput, Select
} from "@mui/material"
import { NavLink } from "react-router-dom"
import { connect } from "react-redux"
import { Language as LanguageIcon, Logout as LogoutIcon }
    from "@mui/icons-material"
import React, { useState, useEffect, useMemo } from "react"
import { useTranslation } from "react-multi-lang"

import LanguageSelector from "./LanguageSelector"
import logout from "../utils/logout"

import { ReactComponent as UserCircleIcon } from "../assets/icons/profile.svg"
import { ReactComponent as UserIcon } from "../assets/icons/user.svg"

function TopNav(props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string)

    const locationMap = useMemo(() =>
        props.gatewayList.reduce((acc, cur) => {
            const
                { gatewayID, permissions } = cur,
                { name, address } = permissions[0].location
            acc[name] = {
                address: address,
                gateways: [...(acc[name]?.gateways || []), gatewayID]
            }
            return acc
        }, {}), [props.gatewayList])

    const locationOptions = useMemo(() =>
        Object.entries(locationMap)
            .map(([key, value]) => ({
                name: key,
                ...value
            }))
        , [locationMap])

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
        },
        containerClasses = "border-b border-black-main bg-gray-900 flex "
            + "flex-col h-20 items-end overflow-visible"
            + (className ? " " + className : "")
    const
        locationHandleChange = (event) => {
            setLocationName(event.target.value)
        },
        gatewayHandleChange = (event) => {
            setGatewayID(event.target.value)
        }

    useEffect(() => {
        const newLocationName = locationOptions[0]?.name
        if (newLocationName) {
            setLocationName(newLocationName)
            setGatewayID(locationMap[newLocationName].gateways[0])
        }
    }, [locationOptions])

    useEffect(() => {
        const active = props.gatewayList.find((gateway) => gateway.gatewayID === gatewayID)
        if (active) props.changeGateway(active)
    }, [gatewayID])

    return <div className={containerClasses}>
        <div className="flex flex-row-reverse h-20 items-center
                        justify-between px-12 z-10 w-full">
            <div className="flex h-20 items-center">
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
                        onChange={locationHandleChange}
                        input={<OutlinedInput />}
                        inputProps={{ "aria-label": "Without label" }}
                        displayEmpty
                    >
                        <MenuItem disabled value="">
                            <em>{commonT("locationName")}</em>
                        </MenuItem>
                        {locationOptions.map((option) => (
                            <MenuItem key={option.name} value={option.name}>
                                {option.name}
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
                        {locationMap[locationName]?.gateways.map((option) => (
                            <MenuItem key={option} value={option}>
                                {option}
                            </MenuItem>
                        )) || null}
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
    lang: state.lang.value,
    user: state.user,
    gatewayList: state.gateways.list
})
const mapDispatch = dispatch => ({
    changeGateway: value =>
        dispatch({ type: "gateways/changeGateway", payload: value })
})

export default connect(mapState, mapDispatch)(TopNav)