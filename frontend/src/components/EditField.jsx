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

const groups = [
    { label: "AreaOwner_TW", value: "AreaOwnerTW" },
    { label: "Area Maintainer", value: "AreaMaintainer" },
    { label: "Serenegray", value: "Serenegray" },
    { label: "Cht_Miaoli", value: "ChtMiaoli" },
]

const defaultChecked = Object.fromEntries(
    groups.map(({ value }) => [value, false])
)

export default function EditField({
    row,
    onSave = () => { }
}) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        dialogT = (string) => t("dialog." + string),
        formT = (string) => t("form." + string)

    const
        [openEdit, setOpenEdit] = useState(false),
        [checked, setChecked] = useState(defaultChecked),
        [enableField, setEnableField] = useState(false)
    const
        handleSwitch = () => {
            setEnableField(!enableField)
        },
        handleChange = (e) => {
            setChecked({
                ...checked,
                [e.target.name]: e.target.checked
            })
        },
        handleClick = () => {
            setOpenEdit(true)
        },
        handleSave = () => {
            const newGroup = Object
                .entries(checked)
                .filter(([value, boo]) => boo)
                .map(([value]) => value)
            onSave({
                ...row,
                group: newGroup,
                enableField
            })
            setOpenEdit(false)
        }

    useEffect(() => {
        if (openEdit) {
            const newChecked = row.group.reduce((acc, cur) => {
                acc[cur] = true
                return acc
            }, { ...defaultChecked })
            setChecked(newChecked)
        }
    }, [row.group, openEdit])

    useEffect(() => {
        if (openEdit) setEnableField(row.enableField)
    }, [row.enableField, openEdit])

    return <>
        <EditIcon
            className="mr-5"
            onClick={handleClick} />
        <DialogForm
            dialogTitle={dialogT("editFieldInfo")}
            fullWidth={true}
            maxWidth="md"
            open={openEdit}
            setOpen={setOpenEdit}>
            <Divider variant="middle" />
            <div className="flex flex-col m-auto mt-4 min-w-49 w-fit">
                <div className="mb-5 flex items-baseline">
                    <p className="ml-1 mr-2">{formT("enableField")}</p>
                    <Switch
                        checked={enableField}
                        onChange={handleSwitch}
                    />
                </div>
                <Divider variant="middle" sx={{ margin: "0 0 1rem" }} />
                <h5 className="mb-5">{commonT("group")}</h5>
                <div className="border-gray-400 border rounded-xl
                    grid grid-cols-2 gap-2 items-center mb-4 p-4">
                    <FormGroup>
                        {groups.map((option) =>
                            <FormControlLabel
                                control={
                                    <Checkbox
                                        checked={checked[option.value]}
                                        onChange={handleChange}
                                        name={option.value} />
                                }
                                label={option.label}
                            />
                        )}
                    </FormGroup>
                </div>
            </div>
            <DialogActions sx={{ margin: "1rem 1.5rem 1rem 0" }}>
                <Button
                    onClick={() => { setOpenEdit(false) }}
                    radius="pill"
                    variant="outlined"
                    color="gray">
                    {commonT("cancel")}
                </Button>
                <Button
                    onClick={handleSave}
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