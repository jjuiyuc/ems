import CheckCircleIcon from "@mui/icons-material/CheckCircle"
import {Link} from "react-router-dom"
import React, {useState} from "react"
import {Button, FormControl, TextField} from "@mui/material"
import {useTranslation} from "react-multi-lang"

import {ValidatePassword} from "../utils/utils"

import AlertBox from "../components/AlertBox"
import LanguageField from "../components/NonLoggedInLanguageField"

function ResetPassword () {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        errorT = string => t("error." + string),
        formT = string => t("form." + string),
        pageT = string => t("resetPassword." + string)

    const
        [confirmPassword, setConfirmPassword] = useState(""),
        [confirmPasswordError, setConfirmPasswordError] = useState(false),
        [isReset, setIsReset] = useState(false),
        [newPassword, setNewPassword] = useState(""),
        [newPasswordError, setNewPasswordError] = useState(false)

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
        submit = () => setIsReset(true)

    const
        cpHelperText = confirmPasswordError ? errorT("passwordNotMatch") : "",
        isSubmittable = confirmPassword
                        && !confirmPasswordError
                        && newPassword
                        && !newPasswordError,
        resetMsg = <>
            <p>{pageT("hasReset")}</p>
            <p>{pageT("logInWithNewPassword")}</p>
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
        <h6 className="mb-8 md:mb-16">{commonT("passwordRule")}</h6>
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