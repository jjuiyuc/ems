import { createTheme } from "@mui/material/styles"
import palette from "./palette"

const theme = createTheme({
    components: {
        MuiButton: {
            styleOverrides: {
                root: ({ ownerState, theme }) => {
                    let styles = {
                        color: ownerState.color === "primary"
                            ? "white"
                            : theme.palette[ownerState.color].main,
                        fontSize: "1rem",
                        fontWeight: 400,
                        textTransform: "none",
                        transitionDuration: ".3s",
                        transitionProperty: "background-color, opacity"
                    }

                    if ("filter" in ownerState) {
                        styles = {
                            ...styles,
                            background: theme.palette.gray[600],
                            color: theme.palette.gray[200],
                            boxShadow: "none",
                            minWidth: "6.75rem",
                            paddingLeft: "1rem",
                            paddingRight: "1rem",
                            "&:hover": {
                                background: theme.palette.gray[500],
                                boxShadow: "none"
                            }
                        }
                    }

                    return styles
                }
            },
            variants: [
                {
                    props: { disabled: true },
                    style: props => ({
                        background: `${props.theme.palette[props.ownerState.color].main} !important`,
                        color: "white !important",
                        opacity: .3
                    })
                },
                {
                    props: { filter: "selected" },
                    style: props => ({
                        fontWeight: "bold",
                        background: theme.palette[ownerState.color].main,
                        color: theme.palette.gray[900],
                        cursor: "default",
                        "&:hover": {
                            background: theme.palette[ownerState.color].main
                        }
                    }),
                },

                {
                    props: { radius: "pill" },
                    style: { borderRadius: "100vh" },
                },
                {
                    props: { size: "x-large" },
                    style: {
                        borderRadius: "20px",
                        fontSize: "1rem",
                        height: "60px"
                    }
                },
                {
                    props: { size: "small", variant: "text" },
                    style: {
                        fontWeight: "bold",
                        lineHeight: "1.17",
                        padding: ".5rem 1rem"
                    },
                },
            ]
        },
        MuiButtonGroup: {
            styleOverrides: {
                root: ({ ownerState, theme }) => {
                    if (!ownerState.variant.includes("ubiik")) return null

                    let styles = {
                        boxShadow: "none",
                        "button,a": {
                            background: theme.palette.gray[900],
                            border: "none !important",
                            borderRadius: "1.25em",
                            color: theme.palette.gray[200],
                            fontWeight: "700",
                            lineHeight: "1.171875",
                            padding: "1em 2em",
                            "&:hover": {
                                background: theme.palette.gray[800],
                            },
                            "&[aria-current=true]": {
                                background: theme.palette.gray[600],
                            }
                        }
                    }

                    if ("amount" in ownerState) {
                        styles.display = "inline-grid"
                        styles.gridTemplateColumns
                            = `repeat(${ownerState.amount}, 1fr)`
                    }

                    return styles
                }
            }
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
                    "fieldset": { transition: "border-color .3s" }
                },
                input: ({ theme }) => ({
                    "&:-webkit-autofill": {
                        boxShadow: `0 0 0 100.09px ${theme.palette.gray[600]} inset !important`
                    }
                })
            }
        },
        MuiInputLabel: {
            styleOverrides: {
                root: ({ theme }) => ({
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
                    props: { font: "mono" },
                    style: {
                        "input": {
                            fontFamily: `Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace`
                        }
                    }
                },
                {
                    props: { variant: "outlined" },
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
        background: { default: "#1C1C1E" },
        mode: "dark",
        text: { "primary": "#E0E0E0" }
    },
    typography: {
        "fontFamily": `"Roboto", "Noto Sans TC", "Helvetica", "Arial", sans-serif`
    }
})

export default theme