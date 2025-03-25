import { connect } from "react-redux"
import { Link } from "react-router-dom"
import { useState } from "react"
import { Button, FormControl, TextField } from "@mui/material"
import { useTranslation } from "react-multi-lang"

import { validateEmail } from "../utils/utils"
import { apiCall } from "../utils/api"

import LanguageField from "../components/NonLoggedInLanguageField"

const mockWebpage = [
    {
        "id": 1,
        "name": "dashboard",
        "permissions": {
            "create": false,
            "read": true,
            "update": false,
            "delete": false
        }
    },
    {
        "id": 2,
        "name": "analysis",
        "permissions": {
            "create": false,
            "read": true,
            "update": false,
            "delete": false
        }
    },
    {
        "id": 3,
        "name": "timeOfUseEnergy",
        "permissions": {
            "create": false,
            "read": true,
            "update": false,
            "delete": false
        }
    },
    {
        "id": 8,
        "name": "accountManagementGroup",
        "permissions": {
            "create": true,
            "read": true,
            "update": true,
            "delete": true
        }
    },
    // {
    //     "id": 10,
    //     "name": "settings",
    //     "permissions": {
    //         "create": true,
    //         "read": true,
    //         "update": true,
    //         "delete": true
    //     }
    // },
    {
        "id": 11,
        "name": "advancedSettings",
        "permissions": {
            "create": true,
            "read": true,
            "update": true,
            "delete": false
        }
    }
]

const mockGateways = { "active": { "gatewayID": "MOCK-GW-000", "permissions": [{ "enabledAt": "2022-08-04T00:00:00Z", "enabledBy": null, "disabledAt": null, "disabledBy": null, "location": { "name": "PLACE 0", "address": "新竹縣XX鄉ＯＯＯ路" } }] }, "list": [{ "gatewayID": "MOCK-GW-000", "permissions": [{ "enabledAt": "2022-08-04T00:00:00Z", "enabledBy": null, "disabledAt": null, "disabledBy": null, "location": { "name": "PLACE 0", "address": "新竹縣XX鄉ＯＯＯ路" } }] }, { "gatewayID": "1E0BA27A8175AF978C49396BDE9D7A1E", "permissions": [{ "enabledAt": "2022-10-24T00:00:00Z", "enabledBy": null, "disabledAt": null, "disabledBy": null, "location": { "name": "PLACE 1", "address": "宜蘭縣ＸＸ鄉ＯＯＯ路" } }] }, { "gatewayID": "218F1623ADD8E739F7C6CBE62A7DF3C0", "permissions": [{ "enabledAt": "2023-01-19T00:00:00Z", "enabledBy": null, "disabledAt": null, "disabledBy": null, "location": { "name": "PLACE 2", "address": "台北市ＸＸ區ＯＯＯ路" } }] }, { "gatewayID": "3RT00000999000000001RUK", "permissions": [{ "enabledAt": "2023-06-20T16:00:00Z", "enabledBy": null, "disabledAt": null, "disabledBy": null, "location": { "name": "PLACE 3", "address": "台南市ＸＸ區ＯＯ路" } }] }] }

function LogIn(props) {
    const t = useTranslation(),
        commonT = (string) => t("common." + string),
        errorT = (string) => t("error." + string),
        formT = (string) => t("form." + string),
        pageT = (string) => t("logIn." + string)

    const
        [email, setEmail] = useState(""),
        [emailError, setEmailError] = useState(null),
        [otherError, setOtherError] = useState(""),
        [showProfileError, setShowProfileError] = useState(false),
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

            const MOCK_MODE = true

            let loginStatus
            if (MOCK_MODE) {
                loginStatus = {
                    data: {
                        token: "mock-token-123456"
                    }
                }
            } else {
                const data = { username: email, password }
                loginStatus = await apiCall({
                    url: "/api/auth",
                    method: "post",
                    data,
                    onError
                })
                if (!loginStatus) return
            }

            const token = loginStatus.data.token
            props.updateUser({ token, username: email })

            let userProfile
            if (MOCK_MODE) {
                userProfile = {
                    data: {
                        id: "mock-user-id",
                        name: "Mock User",
                        username: email,
                        group: {
                            webpages: mockWebpage,
                            gateways: mockGateways
                        }
                    }
                }
            } else {
                const profileOnError = () => setShowProfileError(true)
                userProfile = await apiCall({
                    url: "/api/users/profile",
                    profileOnError
                })
                if (!userProfile) return
            }

            const {
                group,
                id,
                name,
                username
            } = userProfile.data

            const webpages = group.webpages.filter(webpage => webpage?.permissions?.read)
            const gateways = group.gateways

            const tokenExpiryTime = new Date().getTime() + 1000 * 60 * 60 * 3

            if (gateways && gateways.length > 0) {
                props.setGateway(gateways[0])
                props.setGatewayList(gateways)
            }

            props.updateUserProfile({ group, id, name, username, tokenExpiryTime, webpages })
            //
            // const data = { username: email, password }
            // const loginStatus = await apiCall({
            //     url: "/api/auth",
            //     method: "post",
            //     data,
            //     onError
            // })

            // if (!loginStatus) return

            // const token = loginStatus.data.token

            // props.updateUser({ token, username: email })


            // const profileOnError = () => {
            //     setShowProfileError(true)
            // }
            // const userProfile = await apiCall({
            //     url: "/api/users/profile",
            //     profileOnError
            // })
            // if (!userProfile) return
            // const
            //     { group, id, name, username } = userProfile.data,
            //     webpages = group.webpages.filter(webpage => webpage?.permissions?.read),
            //     gateways = group.gateways

            // // Tokens will expire in 3 hours
            // const tokenExpiryTime = new Date().getTime() + 1000 * 60 * 60 * 3

            // if (gateways && gateways.length > 0) {
            //     props.setGateway(gateways[0])
            //     props.setGatewayList(gateways)
            // }
            // props.updateUserProfile({ group, id, name, username, tokenExpiryTime, webpages })

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
                {otherError && (
                    <div className="box mb-8 negative text-center text-red-400">
                        {otherError}
                    </div>
                )}
                {showProfileError && (
                    <div className="box mb-8 negative text-center text-red-400">
                        {showProfileError ? errorT("userProfileError") : ""}
                    </div>
                )}
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
    )
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