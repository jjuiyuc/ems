import React, { useState } from "react"
import { Link, useSearchParams } from "react-router-dom"
import { useTranslation } from "react-multi-lang"
import CheckCircleIcon from "@mui/icons-material/CheckCircle"
import HighlightOffIcon from '@mui/icons-material/HighlightOff';
import { Button, FormControl, TextField } from "@mui/material"

import { ValidatePassword } from "../utils/utils"
import { apiCall } from "../utils/api"

import AlertBox from "../components/AlertBox"
import LanguageField from "../components/NonLoggedInLanguageField"

function ResetPassword() {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        errorT = string => t("error." + string),
        formT = string => t("form." + string),
        pageT = string => t("resetPassword." + string)

    const
        [newPassword, setNewPassword] = useState(""),
        [newPasswordError, setNewPasswordError] = useState(false),
        [confirmPassword, setConfirmPassword] = useState(""),
        [confirmPasswordError, setConfirmPasswordError] = useState(false),
        [isReset, setIsReset] = useState(false),
        [isError, setIsError] = useState(false)



    const
        [searchParams] = useSearchParams()


    const
        changeConfirmPassword = e => {
            setConfirmPassword(e.target.value)
            setConfirmPasswordError(false)
        },
        changeNewPassword = e => {
            setNewPassword(e.target.value)
            setNewPasswordError(false)
        },
        validateNewPassword = () => setNewPasswordError(
            newPassword.length > 0 && !ValidatePassword(newPassword)
        ),
        validateMatchPassword = () => setConfirmPasswordError(
            newPassword.length > 0
            && confirmPassword.length > 0
            && (confirmPassword !== newPassword)
        ),



        submit = () => {

            const data = { token: searchParams.get('token'), password: newPassword }

            const onSuccess = (res) => {
                setIsReset(true)
            }

            const onError = (res) => {
                setIsError(true)
            }
            apiCall({
                url: "/api/users/password/reset-by-token",
                method: "put",
                data,
                onSuccess,
                onError
            })
        }

    const
        cpHelperText = confirmPasswordError ? errorT("passwordNotMatch") : "",
        isSubmittable = confirmPassword
            && !confirmPasswordError
            && newPassword
            && !newPasswordError,
        resetMsg = <>
            <p>{pageT("hasReset")}</p>
            <p>{pageT("logInWithNewPassword")}</p>
        </>,
        errorMsg = <>
            <p>{errorT("passwordTokenError")}</p>
        </>

    return <div>
        <h1 className={"mb-8" + (isReset ? " md:mb-16" : "")}>
            {pageT("resetPassword")}
        </h1>
        {isReset
            ? <>
                <AlertBox
                    boxClass="mb-8"
                    content={resetMsg}
                    icon={CheckCircleIcon} />
                <FormControl fullWidth>
                    <Button
                        color="primary"
                        href="/"
                        size="x-large"
                        variant="contained">
                        {commonT("logIn")}
                    </Button>
                </FormControl>
            </>
            : <>
                <h6 className="text-gray-300 mb-8 md:mb-16">
                    {commonT("passwordRule")}
                </h6>
                {isError
                    ? <AlertBox
                        boxClass="mb-8"
                        content={errorMsg}
                        icon={HighlightOffIcon} />
                    : ""
                }
                <FormControl fullWidth>
                    <LanguageField />
                    <TextField
                        error={newPasswordError}
                        font="mono"
                        helperText={newPasswordError ? errorT("passwordFormat") : ""}
                        label={formT("newPassword")}
                        onBlur={validateNewPassword}
                        onChange={changeNewPassword}
                        type="password"
                        variant="outlined"
                        value={newPassword} />
                    <TextField
                        error={confirmPasswordError}
                        font="mono"
                        helperText={cpHelperText}
                        label={formT("confirmPassword")}
                        onBlur={validateMatchPassword}
                        onChange={changeConfirmPassword}
                        type="password"
                        variant="outlined"
                        value={confirmPassword} />
                    <Button
                        color="primary"
                        disabled={!isSubmittable}
                        onClick={submit}
                        size="x-large"
                        variant="contained">
                        {pageT("reset")}
                    </Button>
                </FormControl>
                <div className="mt-8">
                    <Link to="/">{commonT("logIn")}</Link>
                </div>
            </>}
    </div>
}

export default ResetPassword