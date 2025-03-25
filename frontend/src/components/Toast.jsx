import Alert from "@mui/material"

export default function Toast() {
    return (
        <>
            <Alert variant="filled" severity="error">
                This is an error alert — check it out!
      </Alert>

            <Alert variant="filled" severity="success">
                This is a success alert — check it out!
      </Alert>
        </>
    )
}