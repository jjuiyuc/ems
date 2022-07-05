import CheckCircleIcon from "@mui/icons-material/CheckCircle"
import { connect } from "react-redux"
import { Link } from "react-router-dom"
import React, { useState } from "react"
import { Button, FormControl, TextField } from "@mui/material"
import { useTranslation } from "react-multi-lang"

import { ValidateEmail } from "../utils/utils"
import { apiCall } from "../utils/api"

import AlertBox from "../components/AlertBox"
import LanguageField from "../components/NonLoggedInLanguageField"

function ForgotPassword() {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        errorT = string => t("error." + string),
        formT = string => t("form." + string),
        pageT = (string, params) => t("forgotPassword." + string, params)

    const
        [email, setEmail] = useState(""),
        [emailError, setEmailError] = useState(null),
        [isReset, setIsReset] = useState(false)

    const
        changeEmail = e => {
            setEmail(e.target.value)
            setEmailError(null)
        },
        submit = () => {
            const isEmail = ValidateEmail(email)

            if (!isEmail) {
                setEmailError({ type: "emailFormat" })
                return
            }
            const onSuccess = () => {
                setIsReset(true)
            };
            const onError = (err) => {

            };

            const data = { username: email };
            apiCall({
                url: "/api/users/password/lost",
                method: "put",
                data: data,
                onSuccess,
                onError
            });

        }

    return <div>
        <h1 className="mb-8 md:mb-16">{pageT("forgotPassword")}</h1>
        {isReset
            ? <>
                <AlertBox
                    boxClass="mb-8"
                    content={<p>{pageT("sentResetLink", { email })}</p>}
                    icon={CheckCircleIcon} />
                <FormControl fullWidth>
                    <Button
                        color="primary"
                        href="/"
                        size="x-large"
                        variant="contained">
                        {pageT("backToLogin")}
                    </Button>
                </FormControl>
            </>
            : <>
                <FormControl fullWidth>
                    <LanguageField />
                    <TextField
                        error={emailError !== null}
                        helperText={emailError ? errorT(emailError.type) : ""}
                        label={formT("email")}
                        classes={{ root: "test" }}
                        onChange={changeEmail}
                        type="email"
                        variant="outlined"
                        value={email} />
                    <Button
                        color="primary"
                        disabled={!email}
                        onClick={submit}
                        size="x-large"
                        variant="contained">
                        {pageT("resetPassword")}
                    </Button>
                </FormControl>
                <div className="mt-8">
                    <Link to="/">{commonT("logIn")}</Link>
                </div>
            </>}
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

export default connect(mapState, mapDispatch)(ForgotPassword)