import { Dialog, DialogTitle } from "@mui/material"
import { styled } from "@mui/material/styles"

import palette from "../configs/palette.json"

const CustomDialog = styled(Dialog)(({ theme }) => ({
    "& .MuiPaper-root": {
        color: theme.palette.gray[200],
        backgroundColor: theme.palette.gray[900],
        backgroundImage: "none",
        border: "2px solid" + theme.palette.gray[400],
        boxShadow: "0px 2px 6px 2px rgba(0,0,0,0.14), 0px 2px 8px 2px rgba(96,96,96,0.4)"
    },
}))

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
        <CustomDialog
            fullWidth={fullWidth}
            maxWidth={maxWidth}
            open={open}
            onClose={closeOutside ? handleClose : () => { }}
        >
            <DialogTitle id="form-dialog-title">
                {dialogTitle}
            </DialogTitle>
            {children}
        </CustomDialog>
    </>
}