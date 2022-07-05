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
        [password, setPassword] = useState(""),
        [passwordError, setPasswordError] = useState(false)

    const
        changeEmail = (e) => {
            setEmail(e.target.value)
            setEmailError(null)
        },
        changePassword = (e) => {
            setPassword(e.target.value)
            setPasswordError(false)
        },
        submit = async () => {
            const isEmail = ValidateEmail(email)

            if (!isEmail) {
                setEmailError({ type: "emailFormat" })
                return
            }

            const onError = (err) => {
                if (err == 20004) {
                    setEmailError({ type: "emailNotExist" })
                }
                else if (err == 20006) {
                    setEmailError({ type: "userLocked" })
                }
                else if (err == 20007) {
                    setPasswordError(true)
                }
                else if (err = 40000) {
                    console.log('error')
                }
            }

            const data = { username: email, password };
            const loginStatus = await apiCall({
                url: "/api/auth",
                method: "post",
                data,
                onError
            })
            const token = loginStatus.data.token
            props.updateUser({
                username: email,
                token
            })
            const userProfile = await apiCall({
                url: "/api/users/profile",
                onError
            })
            props.updateUserProfile(userProfile.data)
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
            }),
        updateUserProfile: (value) =>
            dispatch({
                type: "user/updateUserProfile",
                payload: value,
            })

    })

export default connect(mapState, mapDispatch)(LogIn)
