import { useTranslation } from "react-multi-lang"

import AccountInfoModify from "../components/AccountInfoModify"
import ChangePassword from "../components/ChangePassword"

export default function Account() {

    const t = useTranslation(),
        commonT = string => t("common." + string),
        dialogT = (string) => t("dialog." + string),
        pageT = (string, params) => t("account." + string, params),
        errorT = (string) => t("error." + string)


    return <>
        <h1 className="mb-9">{commonT("account")}</h1>
        {/* <div className="card flex flex-col m-auto mt-4 min-w-49 w-fit"> */}
        <div className="gap-y-5 flex flex-wrap lg:gap-x-5">
            <AccountInfoModify />
            <ChangePassword />
        </div>
    </>
}