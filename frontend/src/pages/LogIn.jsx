import { connect } from "react-redux"
import { Link } from "react-router-dom"
import { useState } from "react"
import { Button, FormControl, TextField } from "@mui/material"
import { useTranslation } from "react-multi-lang"

import { validateEmail } from "../utils/utils"
import { apiCall } from "../utils/api"

import LanguageField from "../components/NonLoggedInLanguageField"

function LogIn(props) {
    const t = useTranslation(),
        commonT = (string) => t("common." + string),
        errorT = (string) => t("error." + string),
        formT = (string) => t("form." + string),
        pageT = (string) => t("logIn." + string);

    const
        [email, setEmail] = useState(""),
        [emailError, setEmailError] = useState(null),
        [otherError, setOtherError] = useState(""),
        [profileError, setProfileError] = useState(""),
        [password, setPassword] = useState(""),
        [passwordError, setPasswordError] = useState(false)

    const
        changeEmail = (e) => {
            setEmail(e.target.value)
            setEmailError(null)
            setOtherError("")
        },
        changePassword = (e) => {
            setPassword(e.target.value)
            setPasswordError(false)
            setOtherError("")
        },
        submit = async () => {
            const isEmail = validateEmail(email)

            if (!isEmail) {
                setEmailError({ type: "emailFormat" })
                return
            }

            const onError = (err) => {
                switch (err) {
                    case 20004:
                        setEmailError({ type: "emailNotExist" })
                        break
                    case 20006:
                        setEmailError({ type: "userLocked" })
                        break
                    case 20007:
                        setPasswordError(true)
                        break
                    default: setOtherError(err)
                }
            }

            const data = { username: email, password }
            const loginStatus = await apiCall({
                url: "/api/auth",
                method: "post",
                data,
                onError
            })

            if (!loginStatus) return

            const token = loginStatus.data.token

            props.updateUser({ token, username: email })


            const profileOnError = (err) => {
                setProfileError(err)
            }
            const userProfile = await apiCall({
                url: "/api/users/profile",
                profileOnError
            })
            if (!userProfile) return
            const
                { gateways, group, id, name, username } = userProfile.data,
                webpages = group.webpages.filter(webpage => webpage?.permissions?.read)

            // Tokens will expire in 3 hours
            const tokenExpiryTime = new Date().getTime() + 1000 * 60 * 60 * 3

            if (gateways && gateways.length > 0) {
                props.setGateway(gateways[0])
                props.setGatewayList(gateways)
            }
            props.updateUserProfile({ group, id, name, username, tokenExpiryTime, webpages })

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
                {otherError
                    ? <div className="box mb-8 negative text-center text-red-400">
                        {otherError}
                    </div>
                    : null}
                {profileError
                    ? <div className="box mb-8 negative text-center text-red-400">
                        {profileError ? errorT("userProfileError") : ""}
                    </div>
                    : null}
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

const mapDispatch = dispatch => {
    const updateStore = (payload, type) => dispatch({ payload, type })

    return ({
        setGateway: v => updateStore(v, "gateways/changeGateway"),
        setGatewayList: v => updateStore(v, "gateways/updateList"),
        updateUser: v => updateStore(v, "user/updateUser"),
        updateUserProfile: v => updateStore(v, "user/updateUserProfile")
    })
}

export default connect(null, mapDispatch)(LogIn)