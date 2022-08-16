import palette from "./palette.json"

const { blue, gray, green, indigo, primary, yellow, purple } = palette

const variables = {
    colors: {
        battery: blue.main,
        blue,
        gray,
        indigo,
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
        yellow,
        purple
    },
    languages: {
        en: "English",
        zhtw: "中文"
    }
}

export default variables