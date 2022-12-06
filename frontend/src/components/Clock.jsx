import { ArcElement, Chart, DoughnutController } from "chart.js"
import { useEffect, useState } from "react"

import "../assets/css/clock.scss"

Chart.register(ArcElement, DoughnutController)

export default function Clock(props) {
    const [chart, setChart] = useState(null)

    useEffect(() => {
        const
            ctx = document.getElementById(props.id),
            chart = new Chart(ctx, {
                data: { datasets: [props.dataset] },
                options: {
                    borderWidth: 0,
                    cutout: 90,
                    plugins: {
                        legend: { display: false },
                        tooltip: { enabled: false }
                    }
                },
                type: "doughnut"
            })

        setChart(chart)

        return () => chart.destroy()
    }, [])

    useEffect(() => {
        if (chart) {
            chart.data.datasets = [props.dataset]
            chart.update()
        }
    }, [props.dataset])

    return <div className={"clock " + (props.size ? "" : "h-48 w-48")}
        style={props.size}>
        <div className="bg">
            <div className="ticks">
                {Array.from(new Array(12).keys()).map((v, i) =>
                    <div key={"tick-" + i} />)}
            </div>
            <div className="text text-11px">
                <div>12 am</div>
                <div>6 pm</div>
                <div>6 am</div>
                <div>12 pm</div>
            </div>
        </div>
        <canvas id={props.id} style={{ height: "100%", width: "100%" }} />
    </div>
}