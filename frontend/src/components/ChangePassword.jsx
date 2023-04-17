import { connect } from "react-redux"
import { Button, Divider, IconButton, InputAdornment, OutlinedInput } from "@mui/material"
import Visibility from "@mui/icons-material/Visibility"
import VisibilityOff from "@mui/icons-material/VisibilityOff"
import { useTranslation } from "react-multi-lang"
import { useEffect, useState } from "react"

import { apiCall } from "../utils/api"
import { ValidatePassword } from "../utils/utils"


export default function ChangePassword(props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        formT = (string) => t("form." + string),
        pageT = (string, params) => t("account." + string, params),
        errorT = (string) => t("error." + string)

    const
        [curPassword, setCurPassword] = useState(""),
        [curPasswordError, setCurPasswordError] = useState(false),
        [newPassword, setNewPassword] = useState(""),
        [newPasswordError, setNewPasswordError] = useState(false),
        [showPassword, setShowPassword] = useState(false),
        [loading, setLoading] = useState(false)

    const
        handleClickShowPassword = () => setShowPassword((show) => !show),
        handleMouseDownPassword = (event) => {
            event.preventDefault()
        },
        changePassword = (e) => {
            setCurPassword(e.target.value)
            setCurPasswordError(false)
            // setOtherError("")
        },
        changeNewPassword = (e) => {
            setNewPassword(e.target.value)
            setNewPasswordError(false)
            // setOtherError("")
        },
        submit = () => {
            apiCall({
                data: {
                    currentPassword: curPassword,
                    newPassword: newPassword
                },
                method: "put",
                onSuccess: () => console.log("okk"),
                url: "/api/users/password"
            })
        }

    return <>
        <div className="card w-fit">
            <h4 className="mb-6">
                {pageT("changePassword")}
            </h4>
            <Divider variant="fullWidth" sx={{ marginBottom: "1.5rem" }} />
            <form className="grid grid-cols-1fr-auto gap-x-5 gap-y-6">
                <label className="pt-2">{pageT("currentPassword")}</label>
                <OutlinedInput
                    id="cur-password"
                    size="medium"
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
                    size="medium"
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
            <Divider variant="fullWidth" sx={{ marginTop: "2.8rem" }} />
            <div className="flex flex-row-reverse mt-6">
                <Button
                    sx={{ marginLeft: "0.5rem" }}
                    onClick={submit}
                    radius="pill"
                    variant="contained"
                    color="primary">
                    {commonT("save")}
                </Button>
            </div>
        </div>
    </>
}