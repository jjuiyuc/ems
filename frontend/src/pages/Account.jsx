import { useTranslation } from "react-multi-lang"

import AccountInfoModify from "../components/AccountInfoModify"
import ChangePassword from "../components/ChangePassword"

export default function Account() {

    const t = useTranslation(),
        commonT = string => t("common." + string)

    return <>
        <h1 className="mb-9">{commonT("account")}</h1>
        {/* <div className="card flex flex-col m-auto mt-4 min-w-49 w-fit"> */}
        <div className="gap-y-5 flex flex-wrap md:gap-x-5">
            <AccountInfoModify />
            <ChangePassword />
        </div>
    </>
}