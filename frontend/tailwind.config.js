const palette = require("./src/configs/palette.json")
const plugin = require("tailwindcss/plugin")

module.exports = {
    content: [
        "./index.html",
        "./src/**/*.{vue,js,ts,jsx,tsx}",
    ],
    theme: {
        extend: {
            borderRadius: {
                "2.5xl": "1.25rem"
            },
            borderWidth: {
                "6": "6px"
            },
            boxShadow: {
                main: "4px 0px 4px rgba(0, 0, 0, 0.05) inset"
            },
            colors: { ...palette },
            fontSize: {
                "13px": ".8125rem",
                "22px": "1.375rem",
                "28px": "1.75rem",
                "3.5xl": "2rem"
            },
            gridTemplateColumns: {
                "auto": "auto",
                "auto-1fr": "auto 1fr",
                "auto-19rem": "auto 19rem",
                "1fr-1px-1fr": "1fr 1px 1fr",
                "1fr-auto": "1fr auto",
                "3-auto": "auto auto auto",
                "5rem-1fr": "5rem 1fr",
                "8rem-auto": "8rem auto",
                "15rem-1fr": "15rem 1fr",
                "1fr-auto-1fr": "1fr auto 1fr",
                "settings-col4": "auto 50px auto auto",
                "settings-input": "auto 1fr auto auto auto 30px",
                "settings-input-col5": "auto 1fr auto auto 30px"
            },
            gridTemplateRows: {
                "1fr-auto": "1fr auto",
                "auto-1fr": "auto 1fr"
            },
            height: {
                "15": "3.75rem",
                "38": "9.5rem",
                "45": "11.25rem",
                "160": "40rem"
            },
            maxHeight: {
                "4xl": "56rem",
                "80vh": "80vh"
            },
            margin: {
                "0.7": "0.188rem",
                "18": "4.5rem",
                "auto": "auto"
            },
            maxWidth: {
                "2/3": "66.6%"
            },
            minWidth: {
                "16": "4rem",
                "20": "5rem",
                "42": "10.5rem",
                "49": "12.25rem"
            },
            padding: {
                "7": "1.75rem",
                "25": "6.25rem"
            },
            transitionProperty: {
                "opacity": "opacity",
                "width": "width"
            },
            width: {
                "3/8": "37.5%",
                "5/8": "65.5%",
                "41": "10.25rem",
                "42": "10.5rem",
                "45": "11.25rem",
                "38": "9.5rem",
                "60": "15rem",
                "88": "22rem"
            }
        },
        screens: {
            sm: "600px",
            md: "900px",
            lg: "1200px",
            xl: "1536px",
            "h-sm": { raw: "(min-height: 600px)" },
            "h-md": { raw: "(min-height: 900px)" },
            "h-lg": { raw: "(min-height: 1200px)" },
            "h-xl": { raw: "(min-height: 1536px)" },
            "sm-sm": { raw: "(min-width: 600px) and (min-height: 600px)" },
            "sm-md": { raw: "(min-width: 600px) and (min-height: 900px)" },
            "sm-lg": { raw: "(min-width: 600px) and (min-height: 1200px)" },
            "sm-xl": { raw: "(min-width: 600px) and (min-height: 1536px)" },
        }
    },
    plugins: [
        plugin(function ({ addUtilities }) {
            const bgImage = gap => {
                const
                    gray = palette.gray[400],
                    halfGap = gap / 2 + "rem",
                    twiceGap = gap * 2 + "rem",
                    pos1 = `calc((100% - ${twiceGap}) / 3 + ${halfGap})`,
                    pos1_1 = `calc((100% - ${twiceGap}) / 3 + ${halfGap} + 1px)`,
                    pos2 = `calc((100% - ${twiceGap}) / 3 * 2 + ${gap}rem + ${halfGap})`,
                    pos2_1 = `calc((100% - ${twiceGap}) / 3 * 2 + ${gap}rem + ${halfGap} + 1px)`,
                    trans = "transparent"

                return ({
                    "background-image": "linear-gradient(to right, "
                        + `${trans} ${pos1}, ${gray} ${pos1_1}, ${trans} ${pos1_1}, `
                        + `${trans} ${pos2}, ${gray} ${pos2}, ${trans} ${pos2_1})`
                })
            }
            addUtilities({
                "@media (max-width: 599px)": {
                    ".column-separator": bgImage(1.25)
                },
                "@media (min-width: 600px)": {
                    ".column-separator": bgImage(2.5)
                }
            })
        })
    ],
    safelist: [
        "bg-blue-main-opacity-20",
        "bg-green-main-opacity-20",
        "bg-indigo-main-opacity-20",
        "bg-negative-main",
        "bg-positive-main",
        "bg-yellow-main-opacity-20",
        "h-10",
        "w-10",
        "text-blue-main",
        "text-green-main",
        "text-indigo-main",
        "text-negative-main",
        "text-positive-main",
        "text-yellow-main"
    ]
}