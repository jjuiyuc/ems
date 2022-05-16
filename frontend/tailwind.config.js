const palette = require("./src/configs/palette.json")

module.exports = {
    content: [
        "./index.html",
        "./src/**/*.{vue,js,ts,jsx,tsx}",
    ],
    theme: {
        extend: {
            borderRadius: {
                "4xl": "20px"
            },
            colors: {...palette},
            fontSize: {
                "13px": ".8125rem", // 13 px
                "3.5xl": "2rem"
            },
            gridTemplateRows: {
                "1fr-auto": "1fr auto"
            },
            height: {
                "15": "3.75rem"
            },
            maxHeight: {
                "4xl": "56rem"
            },
            maxWidth: {
                "2/3": "66.6%"
            },
            width: {
                "3/8": "37.5%",
                "5/8": "65.5%"
            }
        },
        screens: {
            sm: "600px",
            md: "900px",
            lg: "1200px",
            xl: "1536px",
            "h-sm": {raw: "(min-height: 600px)"},
            "h-md": {raw: "(min-height: 900px)"},
            "h-lg": {raw: "(min-height: 1200px)"},
            "h-xl": {raw: "(min-height: 1536px)"},
            "sm-sm": {raw: "(min-width: 600px) and (min-height: 600px)"},
            "sm-md": {raw: "(min-width: 600px) and (min-height: 900px)"},
            "sm-lg": {raw: "(min-width: 600px) and (min-height: 1200px)"},
            "sm-xl": {raw: "(min-width: 600px) and (min-height: 1536px)"},
        }
    },
    plugins: []
}