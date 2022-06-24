import { connect } from "react-redux"
import { Link } from "react-router-dom"
import { useState } from "react"
import { Button, FormControl, TextField } from "@mui/material"
import { useTranslation } from "react-multi-lang"

import { ValidateEmail } from "../utils/utils"
import { apiCall } from "../utils/api"

import LanguageField from "../components/NonLoggedInLanguageField"

function LogIn(props) {
    const t = useTranslation(),
        commonT = (string) => t("common." + string),
        errorT = (string) => t("error." + string),
        formT = (string) => t("form." + string),
        pageT = (string) => t("logIn." + string);

    const [email, setEmail] = useState(""),
        [emailError, setEmailError] = useState(null),
        savedEmail = localStorage.getItem("email") || "",
        [password, setPassword] = useState(""),
        [passwordError, setPasswordError] = useState(false);

    const changeEmail = (e) => {
            setEmail(e.target.value)
            setEmailError(null)
        },
        changePassword = (e) => {
            setPassword(e.target.value)
            setPasswordError(false)
        },
        submit = () => {
            const isEmail = ValidateEmail(email)

            if (!isEmail) {
                setEmailError({ type: "emailFormat" })

                return;
            }
            const onSuccess = (token) => {
                props.updateUser({
                    // address: "1915 11th Ave. San Francisco, CA",
                    // name: "Suncat",
                    username: email,
                    token
                })
            }

            const data = { username: email, password };
            apiCall({
                url: "api/auth",
                method: "post",
                data: data,
                onSuccess
            })
        }

   

    return (
        <div>
            <h1 className="mb-8 md:mb-16">{commonT("logIn")}</h1>
            <FormControl fullWidth>
                <LanguageField />
                <TextField
                    error={emailError !== null}
                    helperText={emailError ? errorT(emailError.type) : ""}
                    label={formT("email")}
                    onChange={changeEmail}
                    type="email"
                    variant="outlined"
                    value={email}
                />
                <TextField
                    error={passwordError}
                    helperText={
                        passwordError ? errorT("passwordIncorrect") : ""
                    }
                    label={formT("password")}
                    onChange={changePassword}
                    type="password"
                    variant="outlined"
                    value={password}
                />
                <Button
                    color="primary"
                    disabled={!email || !password}
                    onClick={submit}
                    size="x-large"
                    variant="contained"
                >
                    {commonT("logIn")}
                </Button>
            </FormControl>
            <div className="mt-8">
                <Link to="/forgotPassword">{pageT("forgotPassword")}</Link>
            </div>
        </div>
    );
}

const mapState = (state) => ({
        lang: state.lang.value,
        user: state.user.value,
    }),
    mapDispatch = (dispatch) => ({
        updateLang: (value) =>
            dispatch({
                type: "lang/updateLang",
                payload: value,
            }),
        updateUser: (value) =>
            dispatch({
                type: "user/updateUser",
                payload: value,
            })
    })

export default connect(mapState, mapDispatch)(LogIn)
