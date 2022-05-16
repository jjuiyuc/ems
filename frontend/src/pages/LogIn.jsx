import {connect} from "react-redux"
import {Link} from "react-router-dom"
import React, {useState} from "react"
import {Button, FormControl, TextField} from "@mui/material"
import {useTranslation} from "react-multi-lang"

import {ValidateEmail} from "../utils/utils"

import LanguageSelector from "../components/LanguageSelector"

function LogIn (props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        errorT = string => t("error." + string),
        formT = string => t("form." + string),
        pageT = string => t("login." + string)

    const
        [email, setEmail] = useState(""),
        [emailError, setEmailError] = useState(null),
        [password, setPassword] = useState(""),
        [passwordError, setPasswordError] = useState(false)

    const
        changeEmail = e => {
            setEmail(e.target.value)
            setEmailError(null)
        },
        changePassword = e => {
            setPassword(e.target.value)
            setPasswordError(false)
        },
        submit = () => {
            const isEmail = ValidateEmail(email)

            if (!isEmail) {
                setEmailError({type: "emailFormat"})

                return
            }

            props.updateUser({username: email})
        }

    return <div>
        <h1 className="mb-8 md:mb-16">{commonT("login")}</h1>
        <FormControl fullWidth>
            <LanguageSelector />
            <TextField
                error={emailError !== null}
                helperText={emailError ? errorT(emailError.type) : ""}
                label={formT("email")}
                onChange={changeEmail}
                type="email"
                variant="outlined"
                value={email} />
            <TextField
                error={passwordError}
                helperText={passwordError ? errorT("passwordIncorrect") : ""}
                label={formT("password")}
                onChange={changePassword}
                type="password"
                variant="outlined"
                value={password} />
            <Button
                color="primary"
                disabled={!email || !password}
                onClick={submit}
                size="x-large"
                variant="contained">
                {commonT("login")}
            </Button>
        </FormControl>
        <div className="mt-8">
            <Link to="/forgotPassword">{pageT("forgotPassword")}</Link>
        </div>
    </div>
}

const
    mapState = state => ({
        lang: state.lang.value,
        user: state.user.value
    }),
    mapDispatch = dispatch => ({
        updateLang: value => dispatch({
            type: "lang/updateLang", payload: value
        }),
        updateUser: value => dispatch({
            type: "user/updateUser", payload: value
        })
    })

export default connect(mapState, mapDispatch)(LogIn)