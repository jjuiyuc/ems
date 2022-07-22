import {Button, Divider, ListItemIcon, Menu, MenuItem} from "@mui/material"
import {connect} from "react-redux"
import {Language as LanguageIcon, Logout as LogoutIcon}
    from "@mui/icons-material"
import React, {useState} from "react"
import {useTranslation} from "react-multi-lang"

import {ReactComponent as AlertIcon} from "../assets/icons/alert_ring.svg"
import {ReactComponent as LocationIcon} from "../assets/icons/location.svg"
import {ReactComponent as UserIcon} from "../assets/icons/profile.svg"

import LanguageSelector from "./LanguageSelector"

function TopNav (props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string)

    const [menuAnchorEl, setMenuAnchorEl] = useState(null)

    const
        closeMenu = () => setMenuAnchorEl(null),
        openMenu = e => setMenuAnchorEl(e.currentTarget)

    const
        {className} = props,
        {gateways, name} = props.user,
        gateway = gateways.length > 0
            ? gateways.filter(g => g.active)[0]
            : {address: ""},
        menuPaperProps = {
            sx: {
                borderTopLeftRadius: 0,
                borderTopRightRadius: 0,
                "&.MuiPaper-root": {marginTop: "1.125rem"}
            }
        }

    const containerClasses = "border-b border-black-main bg-gray-900 flex "
                                + "flex-col h-20 items-end overflow-visible"
                                + (className ? " " + className : "")

    return <div className={containerClasses}>
        <div className="flex flex-row-reverse h-20 items-center
                        justify-between px-12 z-10 w-full">
            <div className="flex h-20 items-center">
                <AlertIcon className="h-8 w-8 opacity-30" />
                <Button onClick={openMenu} sx={{marginLeft: "1.5rem"}}>
                    <UserIcon className="h-8 mr-2 w-8" />
                    {name}
                </Button>
            </div>
            <div className="items-center hidden md:flex">
                <LocationIcon className="h-8 mr-1 w-8" />
                {gateway.address}
            </div>
        </div>
        <Menu
            anchorEl={menuAnchorEl}
            anchorOrigin={{horizontal: "right", vertical: "bottom"}}
            onClick={closeMenu}
            onClose={closeMenu}
            open={menuAnchorEl !== null}
            PaperProps={menuPaperProps}
            transformOrigin={{horizontal: "right", vertical: "top"}}>
            <MenuItem>
                <ListItemIcon><LanguageIcon /></ListItemIcon>
                <LanguageSelector id="lang" size="small" />
            </MenuItem>
            <Divider />
            <MenuItem onClick={props.logout}>
                <ListItemIcon><LogoutIcon /></ListItemIcon>
                {commonT("logOut")}
            </MenuItem>
      </Menu>
    </div>
}

const
    mapState = state => ({lang: state.lang.value, user: state.user}),
    mapDispatch = dispatch => ({logout: () => dispatch({type: "user/logout"})})

export default connect(mapState, mapDispatch)(TopNav)