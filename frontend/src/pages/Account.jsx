import { Button, Divider, InputAdornment, IconButton, OutlinedInput } from "@mui/material"
import Visibility from "@mui/icons-material/Visibility"
import VisibilityOff from "@mui/icons-material/VisibilityOff"
import { useTranslation } from "react-multi-lang"
import { useState } from "react"

import { ValidatePassword } from "../utils/utils"
import { apiCall } from "../utils/api"

import AccountInfoModify from "../components/AccountInfoModify"

export default function Account() {

    const t = useTranslation(),
        commonT = string => t("common." + string),
        dialogT = (string) => t("dialog." + string),
        pageT = (string, params) => t("account." + string, params),
        errorT = (string) => t("error." + string)

    const
        [data, setData] = useState(
            {
                id: 3,
                account: "serenegray@ubiik.com",
                password: "9977xxxxll",
                name: "Serenegray"
            }
        ),
        [curPassword, setCurPassword] = useState(""),
        [curPasswordError, setCurPasswordError] = useState(false),
        [newPassword, setNewPassword] = useState(""),
        [newPasswordError, setNewPasswordError] = useState(false),
        [showPassword, setShowPassword] = useState(false),
        [loading, setLoading] = useState(false)

    const
        handleClickOpen = () => {
            setOpen(true)
        },
        handleClose = () => {
            setOpen(false)
        },
        handleClickShowPassword = () => setShowPassword((show) => !show),
        handleMouseDownPassword = (event) => {
            event.preventDefault()
        },
        handleChange = (e) => {
            setDefaultValue(e.target.value)
        }
    const
        changePassword = (e) => {
            setCurPassword(e.target.value)
            setCurPasswordError(false)
            setOtherError("")
        },
        changeNewPassword = (e) => {
            setNewPassword(e.target.value)
            setNewPasswordError(false)
            setOtherError("")
        }

    return <>
        <h1 className="mb-9">{commonT("account")}</h1>
        {/* <div className="card flex flex-col m-auto mt-4 min-w-49 w-fit"> */}
        <div className="gap-y-5 flex flex-wrap lg:gap-x-5">
            <AccountInfoModify />
            <div className="card w-fit">
                <h4 className="mb-6">
                    {pageT("changePassword")}
                </h4>
                <Divider variant="fullWidth" sx={{ marginBottom: "1.5rem" }} />
                <form className="grid grid-cols-1fr-auto gap-x-5 gap-y-6">
                    <label className="pt-2">{pageT("currentPassword")}</label>
                    <OutlinedInput
                        id="cur-password"
                        size="small"
                        type={showPassword ? "text" : "password"}
                        value={curPassword || ""}
                        onChange={changePassword}
                        autoComplete="current-password"
                        endAdornment={
                            <InputAdornment position="end">
                                <IconButton
                                    aria-label="toggle password visibility"
                                    onClick={handleClickShowPassword}
                                    onMouseDown={handleMouseDownPassword}
                                    edge="end"
                                >
                                    {showPassword
                                        ? <Visibility />
                                        : <VisibilityOff />
                                    }
                                </IconButton>
                            </InputAdornment>
                        }
                    />
                    <label className="pt-2">{pageT("newPassword")}</label>
                    <OutlinedInput
                        id="new-password"
                        size="small"
                        type={showPassword ? "text" : "password"}
                        value={newPassword || ""}
                        onChange={changeNewPassword}
                        autoComplete="current-password"
                        endAdornment={
                            <InputAdornment position="end">
                                <IconButton
                                    aria-label="toggle password visibility"
                                    onClick={handleClickShowPassword}
                                    onMouseDown={handleMouseDownPassword}
                                    edge="end"
                                >
                                    {showPassword
                                        ? <Visibility />
                                        : <VisibilityOff />
                                    }
                                </IconButton>
                            </InputAdornment>
                        }
                    />

                </form>
                <Divider variant="fullWidth" sx={{ marginTop: "1.5rem" }} />
                <div className="flex flex-row-reverse mt-6">
                    <Button
                        sx={{ marginLeft: "0.5rem" }}
                        // onClick={() => {

                        // }}
                        radius="pill"
                        variant="contained"
                        color="primary">
                        {commonT("save")}
                    </Button>
                    <Button
                        // onClick={}
                        variant="outlined"
                        radius="pill"
                        color="gray">
                        {commonT("cancel")}
                    </Button>
                </div>
            </div>
        </div>
    </>
}