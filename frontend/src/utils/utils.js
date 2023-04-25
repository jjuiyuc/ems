import moment from "moment"
import variables from "../configs/variables"

const { colors } = variables

const convertTimeToNumber = (string, timezone) => {
    const
        today = moment().format("YYYY-MM-DD"),
        time = moment(today + " " + string + timezone),
        minute = time.minute()

    let hour = time.hour()
    if (hour == 0 && time.date() !== moment().date()) {
        hour = 24
    }
    return hour + Number((minute / 60).toFixed(1))
}
const drawHighPeak = (onPeak) => chart => {
    if (chart.scales.x._gridLineItems && Array.isArray(onPeak)) {
        onPeak.forEach(item => {
            const { start, end } = item
            if (!start || !end) return
            const
                ctx = chart.ctx,
                xLines = chart.scales.x._gridLineItems,
                xLineFirst = xLines[0],
                yFirstLine = chart.scales.y._gridLineItems[0],
                xLeft = yFirstLine.x1,
                xFullWidth = yFirstLine.x2 - xLeft,
                xWidth = (end - start) / 24 * xFullWidth,
                xStart = start / 24 * xFullWidth + xLeft,
                yTop = xLineFirst.y1,
                yFullHeight = xLineFirst.y2 - yTop

            ctx.beginPath()
            ctx.fillStyle = "#ffffff10"
            ctx.strokeStyle = colors.gray[400]
            ctx.rect(xStart, yTop, xWidth, yFullHeight)
            ctx.fill()
            ctx.stroke()
        }
        )
    }
}
const validateEmail = email =>
    /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(email)

const validatePassword = password =>
    /^(?=.{8,})((?=.*[^a-zA-Z]))|(?=.*\d)(?=.*[a-zA-Z])|(?=.*[^a-zA-Z0-9]).*$/.test(password)

const validateNum = num =>
    /^(0?|[1-9][0-9]*)$/.test(num)

const validateNumPercent = num =>
    /^(0?|[1-9]\d?|100)$/.test(num)

const validateNumTwoDecimalPlaces = num =>
    /^((0?|[1-9][0-9]*)|([0-9]*\.)|([0]\.\d{1,2}|[1-9][0-9]*\.\d{1,2}))$/.test(num)

export { convertTimeToNumber, drawHighPeak, validateEmail, validatePassword, validateNum, validateNumPercent, validateNumTwoDecimalPlaces }