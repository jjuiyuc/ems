import { Button, Dialog, DialogTitle, DialogContent, DialogContentText, DialogActions } from "@mui/material"
import { useEffect, useState } from "react"

export default function DialogBox({
    triggerName = "",
    leftButtonName = "",
    rightButtonName = "",
    closeOutside = false
}) {
    const [open, setOpen] = useState(false);

    const handleClickOpen = () => {
        setOpen(true)
    }

    const handleClose = () => {
        setOpen(false)
    }

    return (
        <>
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
                        {"Use Google's location service?"}
                    </DialogTitle>
                    <DialogContent>
                        <DialogContentText id="alert-dialog-description">
                            Let Google help apps determine location. This means sending anonymous
                            location data to Google, even when no apps are running.
                        </DialogContentText>
                    </DialogContent>
                    <DialogActions>
                        <Button onClick={handleClose}>
                            {leftButtonName}
                        </Button>
                        <Button onClick={handleClose} autoFocus>
                            {rightButtonName}
                        </Button>
                    </DialogActions>
                </Dialog>
            </div>
        </>
    )
}

