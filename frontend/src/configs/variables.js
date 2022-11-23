import palette from "./palette.json"

const { blue, gray, green, indigo, primary, yellow, purple, negative, red } = palette

const variables = {
    colors: {
        battery: blue.main,
        blue,
        gray,
        indigo,
        green,
        grid: palette.indigo.main,
        indigo,
        midPeak: palette.yellow.main,
        primary,
        onPeak: palette.negative.main,
        offPeak: palette.green.main,
        solar: palette.yellow.main,
        superOffPeak: palette.purple.main,
        peakShave: palette.green.main,
        yellow,
        purple,
        negative,
        red

    },
    languages: {
        en: "English",
        zhtw: "中文"
    }
}

export default variables