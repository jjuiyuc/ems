import palette from "./palette.json"

const {blue, gray, primary} = palette

const variables = {
    colors: {
        battery: blue.main,
        blue,
        gray,
        grid: palette.indigo.main,
        midPeak: palette.green.main,
        primary,
        onPeak: blue.main,
        offPeak: palette.purple.main,
        solar: palette.yellow.main,
        superOffPeak: palette.yellow.main,
        peakShave: palette.green.main
    },
    languages: {
        en: "English",
        zhtw: "中文"
    }
}

export default variables