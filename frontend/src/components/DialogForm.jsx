import { Dialog, DialogTitle } from "@mui/material"
import { useState } from "react"

export default function DialogForm({
    children = null,
    dialogTitle = "",
    fullWidth,
    maxWidth,
    open,
    setOpen,
    closeOutside = false
}) {

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