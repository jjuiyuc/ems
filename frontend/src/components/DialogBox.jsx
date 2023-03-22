import { Button, Dialog, DialogTitle, DialogContent, DialogContentText, DialogActions } from "@mui/material"
import { useState } from "react"
import { useTranslation } from "react-multi-lang"

export default function DialogBox({
    type = "",
    triggerName = "",
    leftButtonName = "",
    rightButtonName = "",
    closeOutside = false
}) {
    const t = useTranslation(),
        dialogT = (string) => t("dialog." + string)

    const [open, setOpen] = useState(false),
        [fullWidth, setFullWidth] = useState(true),
        [maxWidth, setMaxWidth] = useState("sm")

    const
        handleClickOpen = () => {
            setOpen(true)
        },
        handleClose = () => {
            setOpen(false)
        }
    return <>
        <div>
            <Button
                key={"s-b-"}
                radius="pill"
                variant="contained"
                onClick={handleClickOpen}>
                {triggerName}
            </Button>
            <Dialog
                open={open}
                onClose={closeOutside ? handleClose : () => { }}
                aria-labelledby="alert-dialog-title"
                aria-describedby="alert-dialog-description"
            >
                <DialogTitle id="alert-dialog-title">
                    {dialogT("confirmedMsg")}
                </DialogTitle>
                <DialogContent>
                    <DialogContentText id="alert-dialog-description">
                        {dialogT("promptMsg")}
                    </DialogContentText>
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleClose}
                        radius="pill"
                        variant="outlined"
                        color="gray">
                        {leftButtonName}
                    </Button>
                    <Button onClick={handleClose} autoFocus
                        radius="pill"
                        variant="outlined"
                        color="negative">
                        {rightButtonName}
                    </Button>
                </DialogActions>
            </Dialog>
        </div>
    </>

}