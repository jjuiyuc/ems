import { Dialog, DialogTitle } from "@mui/material"
import { useState } from "react"

export default function DialogForm({
    children = null,
    dialogTitle = "",
    open,
    setOpen,
    closeOutside = false
}) {

    const
        [fullWidth, setFullWidth] = useState(true),
        [maxWidth, setMaxWidth] = useState("sm")

    return <>
        <Dialog
            fullWidth={fullWidth}
            maxWidth={maxWidth}
            open={open}
            onClose={closeOutside ? handleClose : () => { }}
        >
            <DialogTitle id="form-dialog-title">
                {dialogTitle}
            </DialogTitle>
            {children}
        </Dialog>
    </>
}