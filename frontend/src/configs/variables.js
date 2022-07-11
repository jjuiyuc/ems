import palette from "./palette.json"

const {gray, primary} = palette

const variables = {
    colors: {
        battery: palette.blue.main,
        gray,
        grid: palette.indigo.main,
        midPeak: palette.green.main,
        primary,
        onPeak: palette.blue.main,
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