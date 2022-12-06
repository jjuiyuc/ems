const checkbox = new Image(18, 18)

checkbox.src = "data:image/svg+xml,%3Csvg "
    + "xmlns='http://www.w3.org/2000/svg' height='18' width='18'"
    + " fill='none'%3E"
    + "%3Crect x='1' y='1' width='16' height='16' rx='2' stroke='%23606060'"
    + " stroke-width='1.8' /%3E%3C/svg%3E"

const checkedBox = color => {
    const img = new Image(18, 18)

    img.src = "data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg'"
        + " height='18' width='18' fill='none'%3E %3Cpath"
        + " fill-rule='evenodd' clip-rule='evenodd'"
        + " d='M3 0C1.34315 0 0 1.34315 0 3V15C0 16.6569 1.34315 18 3 18H15C16.6569 18 18 16.6569 18 15V3C18 1.34315 16.6569 0 15 0H3ZM8.58666 12.326L13.6123 7.65936L12.3875 6.34033L7.96361 10.4482L5.60069 8.32973L4.39911 9.66996L7.37347 12.3366C7.71983 12.6472 8.24578 12.6426 8.58666 12.326Z'"
        + " fill='" + color + "' /%3E%3C/svg%3E"

    return img
}

const square = color => {
    const img = new Image(18, 18)

    img.src = "data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' "
        + `height='18' width='18' fill='${color}'%3E`
        + "%3Crect x='1' y='1' width='16' height='16' rx='2' /%3E"
        + "%3C/svg%3E"

    return img
}

import { Chart } from "chart.js"

const legendLabels = legendStyle => chart => {
    const labels = Chart.defaults.plugins.legend.labels.generateLabels(chart)

    for (var key in labels) {
        const label = labels[key], color = label.fillStyle.replace("#", "%23")

        label.pointStyle = legendStyle === "square"
            ? square(color)
            : (label.pointStyle = label.hidden ? checkbox : checkedBox(color))
        label.hidden = false
    }

    return labels
}

const square = color => {
    const img = new Image(18, 18)

    img.src = "data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' "
        + `height='18' width='18' fill='${color}'%3E`
        + "%3Crect x='1' y='1' width='16' height='16' rx='2' /%3E"
        + "%3C/svg%3E"

    return img
}

import { Chart } from "chart.js"

const legendLabels = legendStyle => chart => {
    const labels = Chart.defaults.plugins.legend.labels.generateLabels(chart)

    for (var key in labels) {
        const label = labels[key], color = label.fillStyle.replace("#", "%23")

        label.pointStyle = legendStyle === "square"
            ? square(color)
            : (label.pointStyle = label.hidden ? checkbox : checkedBox(color))
        label.hidden = false
    }

    return labels
}

const tooltipLabelPoint = color => {
    const img = new Image(8, 8)

    img.src = "data:image/svg+xml,%3Csvg "
        + "xmlns='http://www.w3.org/2000/svg' height='8' width='8'%3E"
        + "%3Ccircle cx='4' cy='4' r='4' fill='" + color
        + "' /%3E%3C/svg%3E"

    return img
}

export { legendLabels, tooltipLabelPoint }