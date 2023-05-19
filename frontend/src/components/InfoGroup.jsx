import { Button, DialogActions, Divider, ListItem } from "@mui/material"
import { useTranslation, setDefaultTranslations } from "react-multi-lang"
import { useMemo, useState, useEffect } from "react"

import { apiCall } from "../utils/api"

import DialogForm from "./DialogForm"

import { ReactComponent as NoticeIcon } from "../assets/icons/notice.svg"

export default function InfoGroup(props) {
    const { row, groupList, groupTypeDict, groupDictionary } = props

    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        errorT = string => t("error." + string),
        pageT = (string, params) => t("accountManagementGroup." + string, params)

    const
        [openNotice, setOpenNotice] = useState(false),
        [loading, setLoading] = useState(false),
        [infoError, setInfoError] = useState(""),
        [groupName, setGroupName] = useState(""),
        [groupType, setGroupType] = useState(null),
        [parentGroup, setParentGroup] = useState(null),
        [fieldList, setFieldList] = useState([]),
        [fullWidth, setFullWidth] = useState(true),
        [maxWidth, setMaxWidth] = useState("lg")

    const
        iconOnClick = () => {
            setOpenNotice(true)

            const groupID = row.id
            apiCall({
                onComplete: () => setLoading(false),
                onError: error => setInfoError(error),
                onStart: () => setLoading(true),
                onSuccess: rawData => {

                    if (!rawData?.data) return

                    const { data } = rawData

                    setGroupName(data.name || "")
                    setParentGroup(data.parentID || null)
                    setFieldList(data.gateways)

                    // if (Array.isArray(data.gateways)) {
                    //     setFieldList(data.gateways)
                    // }

                },
                url: `/api/account-management/groups/${groupID}`
            })
            console.log(fieldList)

        }
    return <>
        <NoticeIcon
            className="mr-5"
            onClick={iconOnClick}
        />
        <DialogForm
            dialogTitle={commonT("group")}
            fullWidth={fullWidth}
            maxWidth={maxWidth}
            open={openNotice}
            setOpen={setOpenNotice}>
            <Divider variant="middle" />
            <div className="flex flex-col m-auto mt-4 min-w-49 w-fit">
                <div className="grid grid-cols-1fr-auto">
                    <h5 className="ml-6 mt-2">{commonT("groupName")} :</h5>
                    <ListItem
                        id="name"
                        label={commonT("groupName")}>
                        {groupName || ""}
                    </ListItem>
                    <h5 className="ml-6 mt-2">{pageT("groupType")} :</h5>
                    <ListItem
                        id="group-type"
                        label={pageT("groupType")}
                    >
                        {groupTypeDict[row?.typeID] || ""}
                    </ListItem>
                    {parentGroup
                        ? <> <h5 className="ml-6 mt-2">{pageT("parentGroup")} :</h5>
                            <ListItem
                                id="parent-group-type"
                                label={pageT("parentGroup")}
                            >
                                {groupDictionary[parentGroup]}
                            </ListItem></>
                        : null}
                    {/* <div className="flex flex-wrap"> */}
                    <h5 className="ml-6 mt-2">{pageT("fieldList")} :</h5>
                    <ListItem
                        id="field-list"
                        label={pageT("fieldList")}
                        className="grid-col-2 grid"
                    >
                        {fieldList?.map((field, index) => (
                            <p key={index}>
                                {field.locationName}-{field.gatewayID}
                            </p>
                        ))}
                    </ListItem>
                    {/* </div> */}
                </div>
            </div>
            <DialogActions sx={{ margin: "1rem 0.5rem 1rem 0" }}>
                <Button onClick={() => { setOpenNotice(false) }}
                    radius="pill"
                    variant="contained"
                    color="primary">
                    {commonT("okay")}
                </Button>
            </DialogActions>
        </DialogForm>
    </>
}