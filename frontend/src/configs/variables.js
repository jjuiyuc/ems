import palette from "./palette.json"

const {blue, gray, green, indigo, primary, yellow} = palette

const variables = {
    colors: {
        battery: blue.main,
        blue,
        gray,
        green,
        grid: palette.indigo.main,
        indigo,
        midPeak: palette.green.main,
        primary,
        onPeak: blue.main,
        offPeak: palette.purple.main,
        solar: palette.yellow.main,
        superOffPeak: palette.yellow.main,
        peakShave: palette.green.main,
        yellow
    },
    languages: {
        en: "English",
        zhtw: "中文"
    }
}

export default variables