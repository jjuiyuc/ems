import {
    Chart,
    Filler,
    LinearScale,
    LineController,
    LineElement,
    PointElement,
    TimeSeriesScale,
    Tooltip
} from "chart.js"
import "chartjs-adapter-moment"
import { useEffect, useState } from "react"

Chart.register(
    Filler,
    LinearScale,
    LineController,
    LineElement,
    PointElement,
    TimeSeriesScale,
    Tooltip
)

import { checkbox, checkboxChecked, tooltipLabelPoint } from "../utils/chart"
import variables from "../configs/variables"

const
    { colors } = variables,
    datasetCommonOpts = {
        borderWidth: 1,
        hoverRadius: 3,
        pointHoverBorderWidth: 6,
        pointBorderWidth: 0,
        radius: 3,
        tension: 0
    }

const createDatasets = datasets => datasets.map(item => ({
    ...datasetCommonOpts,
    ...item
}))

export default function LineChart(props) {
    const [chart, setChart] = useState(null)

    useEffect(() => {
        const chart = new Chart(document.getElementById(props.id), {
            data: {
                datasets: createDatasets(props.data.datasets),
                labels: props.data.labels
            },
            options: {
                interaction: {
                    intersect: false,
                    mode: "index"
                },
                plugins: {
                    legend: {
                        align: "end",
                        display: props.data.legend || false,
                        labels: {
                            boxHeight: 18,
                            boxWidth: 18,
                            color: "white",
                            font: {
                                size: 16
                            },
                            generateLabels: function (chart) {
                                const labels = Chart.defaults.plugins.legend
                                    .labels.generateLabels(chart)

                                for (var key in labels) {
                                    const
                                        label = labels[key],
                                        color = label.fillStyle
                                            .replace("#", "%23")

                                    label.pointStyle = label.hidden
                                        ? checkbox
                                        : checkboxChecked(color)
                                    label.hidden = false
                                }
                                return labels
                            },
                            padding: 20,
                            usePointStyle: true
                        },
                        position: "top"
                    },
                    tooltip: {
                        backgroundColor: colors.gray[600],
                        boxPadding: 4,
                        bodyFont: { size: 13 },
                        callbacks: {
                            afterBody: props.data.tooltipAfterBody,
                            label: props.data.tooltipLabel,
                            labelPointStyle: context => {
                                const color = context.dataset.backgroundColor
                                    .replace("#", "%23")
                                return { pointStyle: tooltipLabelPoint(color) }
                            }
                        },
                        caretPadding: 12,
                        caretSize: 8,
                        cornerRadius: 20,
                        padding: {
                            bottom: 16,
                            left: 24,
                            right: 24,
                            top: 16
                        },
                        titleFont: {
                            size: 13,
                            weight: "bold"
                        },
                        titleMarginBottom: 8,
                        usePointStyle: true
                    }
                },
                maintainAspectRatio: false,
                scales: {
                    x: {
                        grid: {
                            borderDash: [1, 2],
                            color: colors.gray[400],
                            drawTicks: false
                        },
                        ticks: {
                            autoSkip: true,
                            autoSkipPadding: 60,
                            color: colors.gray[200],
                            font: {
                                size: 11
                            },
                            padding: 10,
                            source: props.data?.xTickSource || "auto"
                        },
                        time: {
                            displayFormats: {
                                hour: "h A"
                            },
                            tooltipFormat: "hh:mm:ss A"
                        },
                        type: "time",
                        ...props.data?.x
                    },
                    y: {
                        grid: {
                            borderDash: [1, 2],
                            color: colors.gray[400],
                            drawTicks: false
                        },
                        ticks: {
                            color: colors.gray[200],
                            font: {
                                size: 11
                            },
                            padding: 10,
                            callback: props.data.tickCallback
                        },
                        position: "left",
                        ...props.data?.y
                    },
                    y1: {
                        grid: {
                            borderDash: [1, 2],
                            color: colors.gray[400],
                            drawTicks: false
                        },
                        ticks: {
                            color: colors.gray[200],
                            callback: props.data.y1TickCallback,
                            font: {
                                size: 11
                            },
                            padding: 10,
                        },
                        display: props.data.y1 || false,
                        position: "right",
                        ...props.data.y1
                    }
                }
            },
            plugins: [
                {
                    beforeDraw: chart => {
                        if (chart.tooltip?._active?.length) {
                            let x = chart.tooltip._active[0].element.x
                            let yAxis = chart.scales.y
                            let ctx = chart.ctx
                            ctx.save()
                            ctx.beginPath()
                            ctx.moveTo(x, yAxis.top)
                            ctx.lineTo(x, yAxis.bottom)
                            ctx.lineWidth = 2
                            ctx.strokeStyle = colors.gray[400]
                            ctx.stroke()
                            ctx.restore()
                        }
                        if (props.data.beforeDraw) {
                            props.data.beforeDraw(chart)
                        }

                    }
                }
            ],
            type: "line"
        })

        setChart(chart)

        return () => chart.destroy()
    }, [])

    useEffect(() => {
        if (chart && chart.canvas) {
            chart.data.datasets = createDatasets(props.data.datasets)
            chart.data.labels = props.data.labels
            if (props.data?.x?.time) {
                chart.options.scales.x.time = props.data.x.time
            }
            chart.update()
        }
    }, [chart, props.data])

    return <canvas className="h-full relative w-full" id={props.id} />
}