import { Button, DialogActions, Divider, ListItem } from "@mui/material"
import ReportProblemIcon from "@mui/icons-material/ReportProblem"
import { useTranslation } from "react-multi-lang"
import { useState } from "react"

import { apiCall } from "../utils/api"

import AlertBox from "../components/AlertBox"
import DialogForm from "./DialogForm"
import LoadingBox from "../components/LoadingBox"
import { ReactComponent as NoticeIcon } from "../assets/icons/notice.svg"

const ErrorBox = ({ error, margin = "", message }) => error
    ? <AlertBox
        boxClass={`${margin} negative`}
        content={<>
            <span className="font-mono ml-2">{error}</span>
            <span className="ml-2">{message}</span>
        </>}
        icon={ReportProblemIcon}
        iconColor="negative-main" />
    : null

export default function InfoGroup(props) {
    const { row, groupTypeDict, groupDictionary } = props

    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("accountManagementGroup." + string, params)

    const
        [openNotice, setOpenNotice] = useState(false),
        [loading, setLoading] = useState(false),
        [infoError, setInfoError] = useState(""),
        [groupName, setGroupName] = useState(""),
        [parentGroup, setParentGroup] = useState(null),
        [fieldList, setFieldList] = useState([])

    const iconOnClick = () => {
        setOpenNotice(true)
        setInfoError("")
        setLoading(true)

        setTimeout(() => {
            setGroupName(row.name)
            setParentGroup(row.parentID)

            const mockGateways = [
                { locationName: "PLACE 0", gatewayID: "MOCK-GW-000" },
                { locationName: "PLACE 1", gatewayID: "MOCK-GW-001" },
                { locationName: "PLACE 2", gatewayID: "MOCK-GW-002" },
            ]
            setFieldList(mockGateways)
            setLoading(false)
        }, 300)
    }

    return <>
        <NoticeIcon
            className="mr-5"
            onClick={iconOnClick}
        />
        <DialogForm
            dialogTitle={commonT("group")}
            fullWidth={true}
            maxWidth="md"
            open={openNotice}
            setOpen={setOpenNotice}>
            <Divider variant="middle" />
            <div className="flex flex-col m-auto mt-4 min-w-49 w-fit">
                <div className="grid grid-cols-8rem-auto">
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
                            </ListItem>
                        </>
                        : null}
                    <h5 className="ml-6 mt-2">{pageT("fieldList")} :</h5>
                    <ListItem
                        id="field-list"
                        label={pageT("fieldList")}
                    >
                        <div className="grid gap-y-4">
                            {fieldList?.map((field, index) => (
                                <p key={index} className="w-full" >
                                    {field.locationName}-{field.gatewayID}
                                </p>
                            ))}
                        </div>
                    </ListItem>
                </div>
            </div>
            <ErrorBox
                error={infoError}
                message={t("error.noDataMsg")} />
            <LoadingBox loading={loading} />
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