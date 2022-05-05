import { createTheme } from "@mui/material/styles"

const theme = createTheme({
    components: {
        MuiButton: {
            styleOverrides: {
                root: {
                    color: "white",
                    fontSize: "1rem",
                    textTransform: "none"
                }
            },
            variants: [
                {
                    props: {radius: "pill"},
                    style: {borderRadius: "100vh"},
                },
                {
                    props: {size: "x-large"},
                    style: {
                        borderRadius: "20px",
                        fontSize: "1rem",
                        height: "60px"
                    }
                },
                {
                    props: {disabled: true},
                    style: props => ({
                        background: `${props.theme.palette[props.ownerState.color].main} !important`,
                        color: "white !important",
                        opacity: .3
                    })
                }
            ]
        },
        MuiInputBase: {
            styleOverrides: {
                root: {
                    borderRadius: "16px !important",
                }
            }
        },
        MuiTextField: {
            styleOverrides: {
                root: {
                    "& fieldset": {
                        transition: "border-color .2s"
                    }
                }
            },
            variants: [
                {
                    props: {variant: "outlined"},
                    style: props => ({
                        "& fieldset": {
                            borderColor: props.theme.palette.gray[400]
                        }
                    })
                }
            ]
        }
    },
    palette: {
        background: {default: "#1C1C1E"},
        black: {main: "#020202"},
        blue: {main: "#43B0FF"},
        gray: {
            200: "#E0E0E0",
            300: "#BDBDBD",
            400: "#606060",
            500: "#404040",
            600: "#303034",
            700: "#262628",
            800: "#232325",
            900: "#1C1C1E"
        },
        green: {main: "#14EEC7"},
        indigo: {main: "#7357FF"},
        error: {main: "#FF4E3E"},
        mode: "dark",
        primary: {main: "#12C9C9"},
        purple: {main: "#EF7BE3"},
        success: {main: "#0FD76B"},
        yellow: {main: "#FFDA15"}
    }
})

export default theme