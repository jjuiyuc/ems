import {
    Button, Checkbox, DialogActions, Divider, FormGroup, FormControlLabel,
    Switch, TextField
} from "@mui/material"
import { useTranslation } from "react-multi-lang"
import { useEffect, useMemo, useState } from "react"
import { ValidateNumPercent } from "../utils/utils"

import DialogForm from "../components/DialogForm"
import ExtraDeviceInfoForm from "../components/ExtraDeviceInfoForm"
import SubDeviceForm from "../components/SubDeviceForm"

import { ReactComponent as EditIcon } from "../assets/icons/edit.svg"

export default function EditField({
    children = null,
    dialogTitle = "",
    openEdit,
    setOpenEdit,
    target,
    setTarget,
    checkState,
    setCheckState,
    groupState,
    setGroupState,
    onClick,
    closeOutside = false
}) {

    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        dialogT = (string) => t("dialog." + string),
        formT = (string) => t("form." + string)

    const
        [fullWidth, setFullWidth] = useState(true),
        [maxWidth, setMaxWidth] = useState("md")

    const { AreaOwner_TW, AreaMaintainer, Serenegray } = checkState

    const
        handleChange = (event) => {
            setCheckState({
                ...checkState,
                [event.target.name]: event.target.checked,
            })
        }

    return <>
        <EditIcon
            className="mr-5"
            onClick={onClick} />
        <DialogForm
            dialogTitle={dialogT("editFieldInfo")}
            fullWidth={fullWidth}
            maxWidth={maxWidth}
            open={openEdit}
            setOpen={setOpenEdit}>
            <Divider variant="middle" />
            <div className="flex flex-col m-auto mt-4 min-w-49 w-fit">
                <div className="mb-5 flex items-baseline">
                    <p className="ml-1 mr-2">{formT("enableField")}</p>
                    <Switch />
                </div>
                <Divider variant="middle" sx={{ margin: "0 0 1rem" }} />
                <h5 className="mb-5">{commonT("group")}</h5>
                <div className="border-gray-400 border rounded-xl
                    grid grid-cols-2 gap-2 items-center mb-4 p-4">
                    <FormGroup>
                        <FormControlLabel
                            control={
                                <Checkbox
                                    checked={AreaOwner_TW}
                                    onChange={handleChange}
                                    value={target?.checkState || ""}
                                    name="AreaOwner_TW" />
                            }
                            label="AreaOwner_TW"
                        />
                        <FormControlLabel
                            control={
                                <Checkbox
                                    checked={AreaMaintainer}
                                    onChange={handleChange}
                                    value={target?.checkState || ""}
                                    name="AreaMaintainer" />
                            }
                            label="Area Maintainer"
                        />
                        <FormControlLabel
                            control={
                                <Checkbox
                                    checked={Serenegray}
                                    onChange={handleChange}
                                    value={target?.checkState || ""}
                                    name="Serenegray" />
                            }
                            label="Serenegray"
                        />
                    </FormGroup>
                </div>
            </div>
            <DialogActions sx={{ margin: "1rem 1.5rem 1rem 0" }}>
                <Button onClick={() => { setOpenEdit(false) }}
                    radius="pill"
                    variant="outlined"
                    color="gray">
                    {commonT("cancel")}
                </Button>
                <Button onClick={() => {
                    setOpenEdit(false)
                    setGroupState()
                }}
                    size="large"
                    radius="pill"
                    variant="contained"
                    color="primary">
                    {commonT("save")}
                </Button>
            </DialogActions>
        </DialogForm>
    </>
}