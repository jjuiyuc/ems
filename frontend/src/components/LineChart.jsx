
import {
    CategoryScale,
    Chart,
    Filler,
    LinearScale,
    LineController,
    LineElement,
    PointElement,
    Tooltip
} from "chart.js"
import {useEffect, useState} from "react"

Chart.register(
    CategoryScale,
    Filler,
    LinearScale,
    LineController,
    LineElement,
    PointElement,
    Tooltip
)

import variables from "../configs/variables"

export default function LineChart (props) {
    const [chart, setChart] = useState(null)

    const {colors} = variables

    useEffect(() => {
        const
            ctx = document.getElementById(props.id),
            chart = new Chart(ctx, {
                data: {
                    datasets: props.data.datasets,
                    labels: props.data.labels
                },
                options: {
                    interaction: {
                        intersect: false,
                        mode: "index"
                    },
                    plugins: {
                        legend: {display: false},
                        tooltip: {
                            backgroundColor: colors.gray[600],
                            boxPadding: 4,
                            bodyFont: {size: 13},
                            callbacks: {
                                ...props.data.tooltipCallbacks
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
                                lineWidth: 0
                            },
                            ticks: {
                                autoSkip: true,
                                autoSkipPadding: 60,
                                color: colors.gray[200],
                                font: {
                                    size: 11
                                },
                                padding: 0
                            }
                        },
                        y: {
                            grid: {
                                borderDash: [1, 2],
                                color: colors.gray[400],
                                drawTicks: false
                            },
                            max: 80,
                            min: 0,
                            ticks: {
                                color: colors.gray[200],
                                font: {
                                    size: 11
                                },
                                padding: 10
                            }
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
                        }
                    }
                ],
                type: "line"
        })

        setChart(chart)

        return () => chart.destroy()
    }, [])

    // useEffect(() => {
    //     if (chart) {
    //         chart.data.datasets = [props.dataset]
    //         chart.update()
    //     }
    // }, [props.dataset])

    return <canvas className="h-full relative w-full" id={props.id} />
}