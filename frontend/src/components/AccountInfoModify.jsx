import { connect } from "react-redux"
import { Button, Divider, OutlinedInput } from "@mui/material"
import { useTranslation } from "react-multi-lang"
import { useEffect, useState } from "react"
import { Link, useSearchParams } from "react-router-dom"

import { apiCall } from "../utils/api"
import { PropaneSharp } from "@mui/icons-material"

const mapState = state => ({

    name: state.user.name,
    username: state.user.username
})
export default connect(mapState)(function AccountInfoModify(props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        formT = (string) => t("form." + string),
        pageT = (string, params) => t("account." + string, params)

    const
        [account, setAccount] = useState(props.username),
        [accountError, setAccountError] = useState(null),
        [name, setName] = useState(props.name),
        [nameError, setNameError] = useState(null),
        [loading, setLoading] = useState(false)

    const
        changeName = (e) => {
            setName(e.target.value)
            setNameError(null)
        },
        submit = () => {
            apiCall({
                data: { name: name },
                method: "put",
                onSuccess: () => console.log("ok"),
                url: "/users/name"
            })
        }


    return <>
        <div className="card w-fit">
            <h4 className="mb-6">
                {pageT("accountInformationModification")}
            </h4>
            <Divider variant="fullWidth" sx={{ marginBottom: "1.5rem" }} />
            <form className="grid grid-cols-1fr-auto gap-x-5 gap-y-6">
                <label>{commonT("account")}</label>
                <span className="pl-1">{account}</span>
                <label className="pt-2">{pageT("name")}</label>
                <OutlinedInput
                    id="edit-name"
                    size="small"
                    value={name || ""}
                    onChange={changeName}
                />
            </form>
            <Divider variant="fullWidth" sx={{ marginTop: "1.5rem" }} />
            <div className="flex flex-row-reverse mt-6">
                <Button
                    sx={{ marginLeft: "0.5rem" }}
                    onClick={
                        submit
                    }
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
    </>
})