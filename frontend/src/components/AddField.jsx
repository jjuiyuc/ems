import {
    Button, DialogActions, Divider, FormControl, ListItem,
    MenuItem, TextField
} from "@mui/material"
import AddIcon from "@mui/icons-material/Add"
import { useTranslation } from "react-multi-lang"
import { useEffect, useMemo, useState } from "react"

import DialogForm from "../components/DialogForm"

export default function AddField({
    children = null,
    dialogTitle = "",
    locationTitle,
    open,
    setOpen,
    closeOutside = false
}) {

    const
        powerCompany = [
            {
                value: "TPC",
                label: "TPC",
            }
        ],
        voltageType = [
            {
                value: "lowVoltage",
                label: "lowVoltage",
            },
            {
                value: "highVoltage",
                label: "highVoltage",
            }
        ],
        touType = [
            {
                value: "twoSection",
                label: "twoSection",
            }
        ]


    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        dialogT = (string) => t("dialog." + string),
        formT = (string) => t("form." + string)

    const
        [fullWidth, setFullWidth] = useState(true),
        [maxWidth, setMaxWidth] = useState("xl"),
        [openAdd, setOpenAdd] = useState(false)

    return <>
        <Button
            onClick={() => { setOpenAdd(true) }}
            size="x-large"
            variant="outlined"
            radius="pill"
            fontSize="large"
            color="brand"
            startIcon={<AddIcon />}>
            {commonT("add")}
        </Button>
        <DialogForm
            dialogTitle={dialogT("addField")}
            open={openAdd}
            setOpen={setOpenAdd}
            fullWidth={fullWidth}
            maxWidth={maxWidth}>
            <Divider variant="middle" />
            <FormControl sx={{
                display: "flex",
                flexDirection: "column",
                margin: "auto",
                width: "fit-content",
                mt: 2,
                minWidth: 120
            }}
                fullWidth={true}>
                <div className="flex-nowrap">
                    <TextField
                        id="gatewayID"
                        label={commonT("gatewayID")}
                    // value={gatewayID}
                    // focused
                    />
                    <Button
                        // onClick={}
                        sx={{ margin: "0.5rem 0.5rem 0.5rem 1rem" }}
                        radius="pill"
                        variant="contained"
                        color="primary">
                        {commonT("verify")}
                    </Button>
                </div>
                <h5 className="mb-4 ml-2">{locationTitle}</h5>
                <TextField
                    id="address"
                    label={formT("address")}
                // value={address}
                />
                <div className="flex-nowrap">
                    <TextField
                        id="lat"
                        label={formT("lat")}
                    // value={lat}
                    />
                    <TextField
                        id="lng"
                        label={formT("lng")}
                        sx={{ marginLeft: "1rem" }}
                    // value={lng}
                    />
                </div>
                <TextField
                    id="powerCompany"
                    select
                    label={formT("powerCompany")}
                    defaultValue=""
                >
                    {powerCompany.map((option) => (
                        <MenuItem key={option.value} value={option.value}>
                            {option.label}
                        </MenuItem>
                    ))}
                </TextField>
                <TextField
                    id="voltageType"
                    select
                    label={formT("voltageType")}
                    defaultValue=""
                >
                    {voltageType.map((option) => (
                        <MenuItem key={option.value} value={formT(`${option.value}`)}>
                            {formT(`${option.label}`)}
                        </MenuItem>
                    ))}
                </TextField>
                <TextField
                    id="touType"
                    select
                    label={formT("touType")}
                    defaultValue=""
                >
                    {touType.map((option) => (
                        <MenuItem key={option.value} value={formT(`${option.value}`)}>
                            {formT(`${option.label}`)}
                        </MenuItem>
                    ))}
                </TextField>

            </FormControl>
            <Divider variant="middle" />
            <DialogActions sx={{ margin: "1rem 0.5rem 1rem 0" }}>
                <Button onClick={() => { setOpenAdd(false) }}
                    radius="pill"
                    variant="outlined"
                    color="gray">
                    {commonT("cancel")}
                </Button>
                <Button onClick={() => { setOpenAdd(false) }}
                    radius="pill"
                    variant="contained"
                    color="primary">
                    {commonT("add")}
                </Button>
            </DialogActions>
        </DialogForm>
    </>
}