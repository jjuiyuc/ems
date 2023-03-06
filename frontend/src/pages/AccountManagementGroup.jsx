import DataTable from "react-data-table-component"
import DatePicker from "react-datepicker"
import moment from "moment"
import { useEffect, useMemo, useState } from "react"


const AccountManagementGroup = props => {
    const sysFuncNameIndex = useMemo(() => {
        let result = {}

        Object.keys(props.systemFunctions).forEach(id => {
            const func = props.systemFunctions[id], { enable, name } = func

            result[name] = { enable, id: parseInt(id) }
        })

        return result
    }, [props.systemFunctions])

    const
        [account, setAccount] = useState({}),
        [accountAction, setAccountAction] = useState("add"),
        [accountList, setAccountList] = useState([]),
        [accountSending, setAccountSending] = useState(false),
        [detailAction, setDetailAction] = useState("detail"),
        [errorData, setErrorData] = useState(null),
        [group, setGroup] = useState({}),
        [groupList, setGroupList] = useState([]),
        [groupSending, setGroupSending] = useState(false),
        [isPasswordValid, setPasswordValid] = useState(false),
        [isSearched, setIsSearched] = useState(false),
        [isUsernameValid, setEmailValid] = useState(false),
        [keyword, setKeyword] = useState("")

    const
        search = () => {
            setIsSearched(!isSearched)
            getAccountList(
                props.apiHeader,
                setAccountList,
                isSearched ? "" : keyword
            )

            if (isSearched) setKeyword("")
        },
        submitAccount = action => {
            setAccountSending(true)

            const configs = { ...props.apiHeader }

            // URL
            if (action === "add") {
                configs.url = "api/accounts"
            }
            else {
                configs.url = "api/accounts/" + account.id
            }

            // Method
            if (action === "add") {
                configs.method = "POST"
            }
            else if (action === "delete") {
                configs.method = "DELETE"
            }
            else {
                configs.method = "PUT"
            }

            // Data
            if (action === "add" || action === "modify") {
                configs.data = {
                    email: account.email,
                    name: account.name,
                    password: account.password
                }

                if (account.group) {
                    configs.data.groups = [parseInt(account.group)]
                }

                if (account.expirationDate) {
                    configs.data.expiration_date = account.expirationDate
                }

                if (action === "add") {
                    configs.data.username = account.username
                }
                else {
                    // This param is needed, otherwise it would be emptied.
                    configs.data.password_error_try = account.password_error_try
                }
            }
            else if (action === "unlock") {
                configs.data = {
                    // This param is needed, otherwise it would be emptied.
                    groups: account.groups?.map(g => g.id) || [],
                    password_error_try: 0
                }
            }

            // Modal
            let modalID = ""

            if (action === "add" || action === "modify") {
                modalID = "accountModal"
            }
            else {
                modalID = "detailModal"
            }

            AxiosInstance(configs)
                .then(() => {
                    getAccountList(
                        props.apiHeader,
                        setAccountList,
                        isSearched ? keyword : ""
                    )
                    document.querySelector("#" + modalID + " .close").click()
                })
                .catch(error => setErrorData(error))
                .then(() => setAccountSending(false))
        },
        submitGroup = () => {
            setErrorData(null)
            setGroupSending(true)

            const
                functions = group.functions.map(f => ({
                    function_id: f.id,
                    enable: f.enable,
                    add_data: f.enable,
                    edit_data: f.enable,
                    delete_data: f.enable,
                    data_import: f.enable,
                    data_export: f.enable
                })),
                configs = {
                    ...props.apiHeader,
                    data: { functions, name: group.name, type: group.type },
                    method: "PUT",
                    url: "api/groups/" + group.id
                }

            AxiosInstance(configs)
                .then(() => {
                    getAccountList(
                        props.apiHeader,
                        setAccountList,
                        isSearched ? keyword : ""
                    )
                    getGroupList(
                        props.apiHeader,
                        props.routeNames,
                        setGroupList,
                        sysFuncNameIndex
                    )
                    document.querySelector("#modifyGroup .close").click()
                })
                .catch(error => setErrorData(error))
                .then(() => setGroupSending(false))
        },
        updateGroup = e => {
            let modifiedGroup = { ...group }

            modifiedGroup[e.target.dataset.target] = e.target.value

            setGroup(modifiedGroup)
        },
        updateGroupFunctions = i => {
            let modifiedGroup = { ...group }

            modifiedGroup.functions[i].enable
                = group.functions[i].enable ? 0 : 1

            setGroup(modifiedGroup)
        },
        updateAccount = e => {
            const { target } = e.target.dataset, { value } = e.target

            let modifiedAccount = { ...account }

            modifiedAccount[target] = value

            setAccount(modifiedAccount)

            if (target === "username") {
                setEmailValid(
                    new RegExp(EMAIL_RULE).test(modifiedAccount.username)
                )
            }
            if (target === "password" || target === "confirmPassword") {
                setPasswordValid(
                    new RegExp(PASSWORD_RULE).test(modifiedAccount.password)
                )
            }
        },
        updateAccountExpDate = (date, e) => {
            let modifiedAccount = { ...account }

            modifiedAccount.expirationDate
                = date ? moment(date).format(dateFormat) : null

            setAccount(modifiedAccount)
        }

    useEffect(() => {
        getAccountList(props.apiHeader, setAccountList)
        getGroupList(
            props.apiHeader,
            props.routeNames,
            setGroupList,
            sysFuncNameIndex
        )
    }, [props.apiHeader, props.routeNames, sysFuncNameIndex])

    const
        { t, userGroupID } = props,
        isNotAdmin = userGroupID !== groupAdminID,
        isNotSuperAdmin = userGroupID !== groupSuperAdminID,
        pageT = string => t("accountManagement." + string),
        accountColumns = (t, pageT) => [
            {
                cell: row => <u
                    className="cursor-pointer d-inline-block text-center
                                            text-nowrap text-primary"
                    data-toggle="modal"
                    data-target="#detailModal"
                    onClick={() => {
                        setAccount(row)
                        setDetailAction("detail")
                    }}>
                    {row.username}
                </u>,
                center: true,
                grow: 2,
                name: pageT("account"),
                selector: "username",
                sortable: true
            },
            {
                center: true,
                name: pageT("name"),
                selector: "name",
                sortable: true
            },
            {
                cell: (row, i) => row.groups
                    ? row.groups.map((g, j) =>
                        <span
                            className={"badge badge-primary cursor-default"
                                + (j > 0 ? " ml-2" : "")}
                            key={"g-" + i + "-" + j}>
                            {g.name}
                        </span>)
                    : "",
                center: true,
                name: pageT("group")
            },
            {
                cell: row => {
                    const isExpired = moment(row.expiration_date) < moment()

                    return <>
                        <span
                            className={"material-icons mr-2 text-"
                                + (isExpired ? "warning" : "success")}>
                            {isExpired ? "warning_amber" : "task_alt"}
                        </span>
                        {pageT(isExpired ? "expired" : "enabled")}
                    </>
                },
                center: true,
                name: pageT("accountStatus"),
                selector: "disable",
                sortable: true
            },
            {
                cell: row => <>
                    {row.password_error_try}
                    {row.password_error_try >= 5
                        ? <button
                            className="btn btn-sm btn-warning ml-2"
                            data-toggle="modal"
                            data-target="#detailModal"
                            onClick={() => {
                                setAccount(row)
                                setDetailAction("unlock")
                                setErrorData(null)
                            }}>
                            {pageT("unlock")}
                        </button>
                        : null}
                </>,
                center: true,
                name: pageT("failedLoginCount"),
                selector: "failCount",
                sortable: true
            },
            {
                cell: row => <>
                    <span
                        className="cursor-pointer material-icons text-primary"
                        data-toggle="modal"
                        data-target="#accountModal"
                        onClick={() => {
                            row.expirationDate = (row.expiration_date
                                ? moment(row.expiration_date).format(dateFormat)
                                : null)
                            row.group = row.groups
                                ? row.groups[row.groups.length - 1].id
                                : "1"
                            setAccount(row)
                            setAccountAction("modify")
                            setErrorData(null)
                        }}>
                        edit
                    </span>
                    {row.id === props.userID
                        ? <span className="material-icons ml-4 text-light-gray">
                            delete
                    </span>
                        : <span
                            className="material-icons ml-4 text-danger
                                    cursor-pointer"
                            data-toggle="modal"
                            data-target="#detailModal"
                            onClick={() => {
                                setAccount(row)
                                setDetailAction("delete")
                                setErrorData(null)
                            }}>
                            delete
                    </span>}
                </>,
                center: true,
                omit: isNotAdmin && isNotSuperAdmin
            }
        ],
        isAccountEmailValid = new RegExp(EMAIL_RULE).test(account.email),
        isPasswordMatch = account.password === account.confirmPassword,
        isSubmittable = accountAction === "add"
            ? (account.username
                && isUsernameValid
                && account.name
                && isPasswordValid
                && isPasswordMatch
                && isAccountEmailValid)
            : !(((account.password || account.confirmPassword)
                && (!isPasswordValid || !isPasswordMatch))
                || !account.name
                || !isAccountEmailValid),
        showPasswordRule = account.password && !isPasswordValid,
        showNoMatch = account.confirmPassword && !isPasswordMatch,
        accountModal = action =>
            <div
                aria-hidden="true"
                aria-labelledby="accountLabel"
                className="modal fade"
                id="accountModal"
                tabIndex="-1">
                <div className="modal-dialog">
                    <div className="modal-content">
                        <div className="modal-header">
                            <h5 className="modal-title" id="accountLabel">
                                {pageT(action + "Account")}
                            </h5>
                            <button
                                aria-label="Close"
                                className="close"
                                data-dismiss="modal"
                                type="button">
                                <span aria-hidden="true">&times;</span>
                            </button>
                        </div>
                        <div className="modal-body">
                            <div className="form-group">
                                <label
                                    className="form-label"
                                    htmlFor="account">
                                    {pageT("account")}
                                    {action === "add"
                                        ? <>
                                            <small className="ml-2 text-muted">
                                                ({pageT("usernameHint")})
                                    </small>
                                            <small className="ml-2 text-danger">
                                                ({t("common.required")})
                                    </small>
                                        </>
                                        : null}
                                </label>
                                <input
                                    className={"form-control"
                                        + (action === "add"
                                            && account.username
                                            && !isUsernameValid
                                            ? " is-invalid"
                                            : "")}
                                    disabled={action === "modify"}
                                    data-target="username"
                                    id="account"
                                    onChange={updateAccount}
                                    type="email"
                                    value={account.username || ""} />
                            </div>
                            <div className="form-group">
                                <label
                                    className="form-label"
                                    htmlFor="accountName">
                                    {pageT("name")}
                                    <small className="ml-2 text-danger">
                                        ({t("common.required")})
                                    </small>
                                </label>
                                <input
                                    className={"form-control"
                                        + (action === "modify"
                                            && !account.name
                                            ? " is-invalid"
                                            : "")}
                                    data-target="name"
                                    id="accountName"
                                    onChange={updateAccount}
                                    type="text"
                                    value={account.name || ""} />
                            </div>
                            <div className="form-group">
                                <label
                                    className="form-label"
                                    htmlFor="password">
                                    {t("password.password")}
                                    {action === "add"
                                        ? <small className="ml-2 text-danger">
                                            ({t("common.required")})
                                    </small>
                                        : null}
                                </label>
                                <input
                                    className={"form-control"
                                        + (showPasswordRule ? " is-invalid" : "")}
                                    data-target="password"
                                    id="password"
                                    onChange={updateAccount}
                                    type="password"
                                    value={account.password || ""} />
                                {showPasswordRule
                                    ? <div className="invalid-feedback">
                                        {t("password.rule")}
                                    </div>
                                    : null}
                            </div>
                            <div className="form-group">
                                <label
                                    className="form-label"
                                    htmlFor="confirmPassword">
                                    {t("password.confirmPassword")}
                                    {action === "add"
                                        ? <small className="ml-2 text-danger">
                                            ({t("common.required")})
                                    </small>
                                        : null}
                                </label>
                                <input
                                    className={"form-control"
                                        + (showNoMatch ? " is-invalid" : "")}
                                    data-target="confirmPassword"
                                    id="confirmPassword"
                                    onChange={updateAccount}
                                    type="password"
                                    value={account.confirmPassword || ""} />
                                {showNoMatch
                                    ? <div className="invalid-feedback">
                                        {t("password.confirmPWFailMsg")}
                                    </div>
                                    : null}
                            </div>
                            <div className="form-group">
                                <label
                                    className="form-label"
                                    htmlFor="group">
                                    {pageT("group")}
                                </label>
                                <select
                                    className="form-control"
                                    data-target="group"
                                    id="group"
                                    onChange={updateAccount}
                                    value={account.group}>
                                    {groupOptions}
                                </select>
                            </div>
                            <div className="form-group">
                                <label
                                    className="form-label"
                                    htmlFor="email">
                                    {t("common.email")}
                                    <small className="ml-2 text-danger">
                                        ({t("common.required")})
                                    </small>
                                </label>
                                <input
                                    className={"form-control"
                                        + ((action === "add"
                                            && account.email
                                            && !isAccountEmailValid)
                                            || (action === "modify"
                                                && !isAccountEmailValid)
                                            ? " is-invalid"
                                            : "")}
                                    data-target="email"
                                    id="email"
                                    onChange={updateAccount}
                                    type="text"
                                    value={account.email || ""} />
                            </div>
                            <div className="form-group">
                                <label
                                    className="form-label"
                                    htmlFor="accountExpirationDate">
                                    {pageT("accountExpirationDate")}
                                </label>
                                <DatePicker
                                    className="form-control"
                                    dateFormat="yyyy-MM-dd"
                                    dropdownMode="select"
                                    isClearable={!account.expiration_date}
                                    onChange={updateAccountExpDate}
                                    selected={account.expirationDate
                                        ? new Date(account.expirationDate)
                                        : null}
                                    showMonthDropdown
                                    showYearDropdown />
                                <small className="text-muted">
                                    {pageT("expDateHint")}
                                </small>
                                <ErrorMessage
                                    className="alert alert-danger font-mono
                                                mb-0 mt-4 text-center"
                                    data={errorData} />
                            </div>
                        </div>
                        <div className="modal-footer">
                            <button
                                className="btn btn-secondary"
                                data-dismiss="modal"
                                type="button">{t("common.cancel")}</button>
                            <button
                                className="btn btn-primary"
                                disabled={!isSubmittable || accountSending}
                                onClick={submitAccount.bind(this, action)}
                                type="button">
                                {accountSending
                                    ? <>
                                        <div className="align-middle mr-2 spinner-border
                                                    spinner-border-sm" />
                                        {t("common.sending")}
                                    </>
                                    : t("common." + action)}
                            </button>
                        </div>
                    </div>
                </div>
            </div>,
        detailModal = action =>
            <div
                aria-hidden="true"
                aria-labelledby="detailLabel"
                className="modal fade"
                id="detailModal"
                tabIndex="-1">
                <div className="modal-dialog">
                    <div className="modal-content">
                        <div className="modal-header">
                            <h5 className="modal-title" id="detailLabel">
                                {pageT(action === "detail"
                                    ? "accountDetail"
                                    : action + "Account")}
                            </h5>
                            <button
                                aria-label="Close"
                                className="close"
                                data-dismiss="modal"
                                type="button">
                                <span aria-hidden="true">&times;</span>
                            </button>
                        </div>
                        <div className="modal-body">
                            {action !== "detail"
                                ? <p>{pageT(action + "Confirm")}</p>
                                : null}
                            {account.id
                                ? <table className="table table-bordered
                                                table-striped">
                                    <tbody>
                                        <tr>
                                            <th>{t("common.id")}</th>
                                            <td>{account.id}</td>
                                        </tr>
                                        <tr>
                                            <th>{pageT("account")}</th>
                                            <td>{account.username}</td>
                                        </tr>
                                        <tr>
                                            <th>{pageT("name")}</th>
                                            <td>{account.name}</td>
                                        </tr>
                                        <tr>
                                            <th>{pageT("group")}</th>
                                            <td>{account.groups
                                                ? account.groups.map((g, i) =>
                                                    <span
                                                        className={"badge badge-primary cursor-default"
                                                            + (i > 0 ? " ml-2" : "")}
                                                        key={"dg-" + i}>
                                                        {g.name}
                                                    </span>)
                                                : ""}</td>
                                        </tr>
                                        <tr>
                                            <th>{t("common.email")}</th>
                                            <td>{account.email}</td>
                                        </tr>
                                        <tr>
                                            <th>{pageT("accountStatus")}</th>
                                            <td>
                                                {pageT(account.disable
                                                    ? "disabled"
                                                    : "enabled")}
                                            </td>
                                        </tr>
                                        <tr>
                                            <th>{pageT("failedLoginCount")}</th>
                                            <td>{account.password_error_try}</td>
                                        </tr>
                                        <tr>
                                            <th>
                                                {pageT("accountExpirationDate")}
                                            </th>
                                            <td>
                                                {account.expiration_date
                                                    ? moment(account.expiration_date)
                                                        .format("YYYY-MM-DD")
                                                    : "-"}
                                            </td>
                                        </tr>
                                    </tbody>
                                </table>
                                : null}
                        </div>
                        <div className="modal-footer">
                            <button
                                className="btn btn-secondary"
                                data-dismiss="modal"
                                type="button">
                                {t("common." + (action === "detail"
                                    ? "close"
                                    : "cancel"))}
                            </button>
                            {action !== "detail"
                                ? <button
                                    className={"btn btn-"
                                        + (action === "delete"
                                            ? "danger"
                                            : "primary")}
                                    disabled={accountSending}
                                    onClick={submitAccount.bind(this, action)}
                                    type="button">
                                    {accountSending
                                        ? <>
                                            <div className="align-middle mr-2
                                                        spinner-border
                                                        spinner-border-sm" />
                                            {t("common.sending")}
                                        </>
                                        : action === "delete"
                                            ? t("common.delete")
                                            : pageT(action)}
                                </button>
                                : null}
                        </div>
                    </div>
                </div>
            </div>,
        groupColumns = (t, pageT) => [
            {
                cell: row => <u
                    className="cursor-pointer d-inline-block
                                            text-center text-nowrap
                                            text-primary"
                    data-toggle="modal"
                    data-target="#groupDetailModal"
                    onClick={setGroup.bind(this, row)}>
                    {row.name}
                </u>,
                center: true,
                name: pageT("groupName"),
                selector: "name"
            },
            {
                center: true,
                name: pageT("groupType"),
                selector: "type"
            },
            {
                center: true,
                cell: row => row.type === "Super-Admin" && isNotSuperAdmin
                    ? "-"
                    : <span
                        className="cursor-pointer material-icons text-primary"
                        data-toggle="modal"
                        data-target="#modifyGroup"
                        onClick={() => {
                            setErrorData(null)
                            setGroup(JSON.parse(JSON.stringify(row)))
                        }}>
                        edit
                    </span>,
                name: t("common.modify"),
                omit: isNotAdmin && isNotSuperAdmin
            }
        ],
        groupOptions = groupList.map((group, i) =>
            <option
                key={"group-" + i}
                value={group.id}>
                {group.name}
            </option>),
        groupDetailFuncList = group => group.functions.map((f, i) => {
            const isEnabled = f.enable

            return <li
                className={"py-1 rounded-sm"
                    + (isEnabled ? "" : " bg-light-gray")}
                key={"g-d-f-" + i}
                style={{ color: isEnabled ? "inherit" : "#aaa" }}>
                <i
                    aria-label={pageT(isEnabled ? "enabled" : "disabled")}
                    className="align-middle material-icons mr-1">
                    {isEnabled ? "task_alt" : "remove"}
                </i>
                {t("menuItems." + f.name)}
            </li>
        }),
        groupDetailContent = group =>
            <table className="table table-bordered table-striped">
                <tbody>
                    <tr><th>{t("common.id")}</th><td>{group.id}</td></tr>
                    <tr><th>{pageT("groupName")}</th><td>{group.name}</td></tr>
                    <tr><th>{pageT("groupType")}</th><td>{group.type}</td></tr>
                    <tr>
                        <th>{pageT("functions")}</th>
                        <td>
                            <ul
                                className="list-unstyled mb-0"
                                style={{
                                    display: "grid",
                                    gridGap: ".5rem",
                                    gridTemplateColumns: "auto auto"
                                }}>
                                {groupDetailFuncList(group)}
                            </ul>
                        </td>
                    </tr>
                </tbody>
            </table>

    return <>
        {props.availableFuncs.includes("accountGroup")
            ? <>
                <h5>{pageT("group")}</h5>
                <div className="react-data-table-component mb-5">
                    <DataTable
                        columns={groupColumns(t, pageT)}
                        data={groupList}
                        fixedHeader={false}
                        noDataComponent={t("dataTable.noDataMsg")}
                        pagination={false}
                        striped={true} />
                </div>
                <div
                    aria-hidden="true"
                    aria-labelledby="modifyGroupLabel"
                    className="modal fade"
                    id="modifyGroup"
                    tabIndex="-1">
                    <div className="modal-dialog">
                        <div className="modal-content">
                            <div className="modal-header">
                                <h5 className="modal-title" id="modifyGroupLabel">
                                    {pageT("modifyGroup")}
                                </h5>
                                <button
                                    aria-label="Close"
                                    className="close"
                                    data-dismiss="modal"
                                    type="button">
                                    <span aria-hidden="true">&times;</span>
                                </button>
                            </div>
                            <div className="modal-body">
                                {group.id
                                    ? <>
                                        <div className="form-group">
                                            <label
                                                className="form-label"
                                                htmlFor="groupName">
                                                {pageT("groupName")}
                                                <small className="ml-2 text-danger">
                                                    ({t("common.required")})
                                            </small>
                                            </label>
                                            <input
                                                className="form-control"
                                                data-target="name"
                                                id="groupName"
                                                onChange={updateGroup}
                                                type="text"
                                                value={group.name} />
                                        </div>
                                        <div className="form-group">
                                            <label
                                                className="form-label"
                                                htmlFor="groupType">
                                                {pageT("groupType")}
                                            </label>
                                            <select
                                                className="form-control"
                                                data-target="type"
                                                id="groupType"
                                                onChange={updateGroup}
                                                value={group.type}>
                                                {groupTypeOptions}
                                            </select>
                                        </div>
                                        <label className="form-label mb-3">
                                            {pageT("functions")}
                                        </label>
                                        <div style={{
                                            display: "grid",
                                            gridGap: ".5rem",
                                            gridTemplateColumns: "auto auto"
                                        }}>
                                            {group.functions.map((func, i) =>
                                                <div key={"func-" + i}>
                                                    <input
                                                        checked={func.enable}
                                                        className={"align-middle mr-2"
                                                            + (func.editable
                                                                ? ""
                                                                : " cursor-default")}
                                                        disabled={!func.editable}
                                                        id={"func-" + func.name}
                                                        onChange={() =>
                                                            updateGroupFunctions(i)}
                                                        type="checkbox" />
                                                    <label
                                                        className={func.editable
                                                            ? ""
                                                            : "cursor-default text-muted"}
                                                        htmlFor={"func-" + func.name}>
                                                        {t("menuItems." + func.name)}
                                                    </label>
                                                </div>)}
                                        </div>
                                        <ErrorMessage
                                            className="alert alert-danger font-mono
                                                    mb-0 mt-4 text-center"
                                            data={errorData} />
                                    </>
                                    : null}
                            </div>
                            <div className="modal-footer">
                                <button
                                    className="btn btn-secondary"
                                    data-dismiss="modal"
                                    type="button">{t("common.cancel")}</button>
                                <button
                                    className="btn btn-primary"
                                    disabled={!group.name || groupSending}
                                    onClick={submitGroup}
                                    type="button">
                                    {groupSending
                                        ? <>
                                            <div className="align-middle mr-2
                                                    spinner-border
                                                    spinner-border-sm" />
                                            {t("common.sending")}
                                        </>
                                        : t("common.save")}
                                </button>
                            </div>
                        </div>
                    </div>
                </div>
                <div
                    aria-hidden="true"
                    aria-labelledby="groupDetailModalLabel"
                    className="modal fade"
                    id="groupDetailModal"
                    tabIndex="-1">
                    <div className="modal-dialog">
                        <div className="modal-content">
                            <div className="modal-header">
                                <h5
                                    className="modal-title"
                                    id="groupDetailModalLabel">
                                    {pageT("groupDetail")}
                                </h5>
                                <button
                                    aria-label="Close"
                                    className="close"
                                    data-dismiss="modal"
                                    type="button">
                                    <span aria-hidden="true">&times;</span>
                                </button>
                            </div>
                            <div className="modal-body">
                                {group.id ? groupDetailContent(group) : null}
                            </div>
                            <div className="modal-footer">
                                <button
                                    className="btn btn-secondary"
                                    data-dismiss="modal"
                                    type="button">{t("common.close")}</button>
                            </div>
                        </div>
                    </div>
                </div>
            </>
            : null}
        {props.availableFuncs.includes("account")
            ? <>
                <div className="align-items-center d-flex justify-content-between mb-2">
                    <h5 className="mb-0">{pageT("account")}</h5>
                    <div className="input-group w-auto">
                        <div className="input-group-prepend">
                            <div className="input-group-text">
                                {pageT("accountOrName")}
                            </div>
                        </div>
                        <input
                            className="form-control"
                            disabled={isSearched}
                            onChange={e => setKeyword(e.target.value)}
                            onKeyUp={e => keyword && e.key === "Enter"
                                ? search()
                                : null}
                            value={keyword}
                            type="text" />
                        <div className="input-group-append">
                            <button
                                className="btn btn-primary py-0"
                                disabled={!keyword}
                                onClick={search}>
                                <span className="align-middle material-icons">
                                    {isSearched ? "close" : "search"}
                                </span>
                            </button>
                        </div>
                    </div>
                    {userGroupID === groupAdminID || userGroupID === groupSuperAdminID
                        ? <button
                            className="btn btn-primary"
                            data-toggle="modal"
                            data-target="#accountModal"
                            onClick={() => {
                                setAccount({})
                                setAccountAction("add")
                                setErrorData(null)
                            }}>
                            {t("common.add")}
                        </button>
                        : null}
                </div>
                <div className="react-data-table-component mb-4">
                    <DataTable
                        columns={accountColumns(t, pageT)}
                        data={accountList}
                        noDataComponent={t("dataTable.noDataMsg")}
                        pagination={true}
                        paginationComponentOptions={{
                            rowsPerPageText: t("dataTable.rowsPerPage")
                        }}
                        striped={true} />
                </div>
                {userGroupID === groupAdminID || userGroupID === groupSuperAdminID
                    ? accountModal(accountAction)
                    : null}
                {detailModal(detailAction)}
            </>
            : null}
    </>
}

export default AccountManagementGroup