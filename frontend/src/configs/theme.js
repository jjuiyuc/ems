import {createTheme} from "@mui/material/styles"
import palette from "./palette"

const theme = createTheme({
    components: {
        MuiButton: {
            styleOverrides: {
                root: {
                    color: "white",
                    fontSize: "1rem",
                    textTransform: "none",
                    transitionDuration: ".3s",
                    transitionProperty: "background-color, opacity"
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
        MuiFormHelperText: {
            styleOverrides: {
                root: {
                   fontSize: ".8125rem",
                   lineHeight: "1.5374",
                   marginTop: ".5rem"
                }
            }
        },
        MuiInputBase: {
            styleOverrides: {
                root: {
                    borderRadius: "16px !important",
                    color: "white",
                    "fieldset": {transition: "border-color .3s"}
                },
                input: ({ownerState, theme}) => ({
                    "&:-webkit-autofill": {
                        boxShadow: `0 0 0 100.09px ${theme.palette.gray[600]} inset !important`
                    }
                })
            }
        },
        MuiInputLabel: {
            styleOverrides: {
                root: ({ownerState, theme}) => ({
                    color: theme.palette.gray[300]
                })
            }
        },
        MuiMenuItem: {
            styleOverrides: {
                root: {
                    transitionDuration: ".3s",
                    transitionProperty: "background-color"
                }
            }
        },
        MuiTextField: {
            styleOverrides: {
                root: {
                    marginBottom: "2rem"
                }
            },
            variants: [
                {
                    props: {font: "mono"},
                    style: {
                        "input": {
                            fontFamily: `Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace`
                        }
                    }
                },
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
        ...palette,
        background: {default: "#1C1C1E"},
        mode: "dark",
        text: {"primary": "#E0E0E0"}
    },
    typography: {
        "fontFamily": `"Roboto", "Noto Sans TC", "Helvetica", "Arial", sans-serif`
    }
})

export default theme