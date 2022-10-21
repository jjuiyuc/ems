import { connect } from "react-redux"
import ReportProblemIcon from "@mui/icons-material/ReportProblem"
import { useEffect, useMemo, useState } from "react"
import { useTranslation } from "react-multi-lang"

import AlertBox from "../components/AlertBox"
import { API_HOST } from "../constant/env"
import ElectricalGridDiagram from "../components/ElectricalGridDiagram"

import { ReactComponent as HomeIcon } from "../assets/icons/home.svg"
import { ReactComponent as BatteryIcon } from "../assets/icons/battery.svg"
import { ReactComponent as GridIcon } from "../assets/icons/grid.svg"
import { ReactComponent as SolarIcon } from "../assets/icons/sunny.svg"

import "../assets/css/dashboard.scss"

const iconMap = {
    battery: { color: "blue", icon: BatteryIcon },
    grid: { color: "indigo", icon: GridIcon },
    home: { color: "green", icon: HomeIcon },
    solar: { color: "yellow", icon: SolarIcon }
}

const IconSet = props => {
    const { size } = props, { color, icon } = iconMap[props.icon], Icon = icon

    return <div className={`h-${size} bg-${color}-main-opacity-20 grid
                            place-items-center rounded-full w-${size}`}>
        <Icon className={`h-3/5 w-3/5 text-${color}-main`} />
    </div>
}
const Card = props => {
    const
        cardClasses = "card narrow px-6"
            + (props.className ? " " + props.className : ""),
        { data, title } = props

    return <div className={cardClasses}>
        <div className="grid sm:grid-cols-2 text-center
                        gap-x-4 sm:gap-x-6
                        gap-y-6 sm:gap-y-8">
            <div className="flex flex-wrap sm:col-span-2 items-center">
                <IconSet icon={props.icon} size="10" />
                <h5 className="font-bold ml-3">{title}</h5>
            </div>
            {data.map((item, i) =>
                <div key={"d-c-d-" + i}>
                    <h2 className="font-bold text-2xl mb-1">{item.value}</h2>
                    <p className="lg:test text-sm">{item.name}</p>
                </div>)}
        </div>
    </div>
}
const CardOnDiagram = props =>
    <div className={`arrow-${props.arrow} card-diagram bg-gray-900 flex
                    items-start px-4 py-3 rounded-lg`}>
        <IconSet icon={props.icon} size="6" />
        <div className="ml-2 mt-0.5">
            <h6 className="font-bold text-gray-200">{props.value}</h6>
            <span className="text-11px text-gray-300">{props.title}</span>
        </div>
    </div>

const mapState = state => ({
    gatewayID: state.gateways.active.gatewayID,
    token: state.user.token
})

export default connect(mapState)(function Dashboard(props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("dashboard." + string, params)

    const
        [battery, setBattery] = useState({
            direction: "",
            target: "",
            import: 0,
            power: 0,
            state: 0
        }),
        [diagram, setDiagram] = useState({
            battery: 0,
            grid: 0,
            load: 0,
            solar: 0
        }),
        [error, setError] = useState(null),
        [lines, setLines] = useState({
            battery: { grid: 0, load: 0, pv: 0 },
            grid: { battery: 0, load: 0, pv: 0 },
            load: { battery: 0, grid: 0, pv: 0 },
            pv: { battery: 0, grid: 0, load: 0 },
        }),
        [load, setLoad] = useState({ discharge: 0, import: 0, solar: 0 }),
        [peak, setPeak] = useState({ active: 0, current: 0, threshhold: 0 }),
        [solar, setSolar] = useState({ charge: 0, consume: 0, export: 0 }),
        [websocketSupport, setWebSocketSupport] = useState(true)

    const updateData = data => {
        const
            batteryPower = data.batteryProducedAveragePowerAC
                + data.batteryConsumedAveragePowerAC,
            solarPowerExport = Math.abs(data.gridProducedAveragePowerAC
                + data.gridConsumedAveragePowerAC)

        setBattery({
            direction:
                data.batteryChargingFrom ? "chargingFrom" : "dischargingTo",
            target: (data.batteryChargingFrom || data.batteryDischargingTo)
                .toLocaleLowerCase(),
            import: data.batteryGridAveragePowerAC,
            power: batteryPower,
            state: data.batterySoC
        })
        setDiagram({
            battery: batteryPower,
            grid: solarPowerExport,
            load: data.loadAveragePowerAC,
            solar: data.pvAveragePowerAC
        })
        setLines({
            battery: data.batteryLinks,
            grid: data.gridLinks,
            load: data.loadLinks,
            pv: data.pvLinks
        })
        setLoad({
            discharge: data.loadBatteryAveragePowerAC,
            import: data.loadGridAveragePowerAC,
            solar: data.loadPvAveragePowerAC
        })
        setPeak({
            active: data.gridIsPeakShaving,
            current: data.gridProducedAveragePowerAC,
            threshhold: data.gridContractPowerAC
        })
        setSolar({
            charge: data.batteryPvAveragePowerAC,
            consume: data.loadPvAveragePowerAC,
            export: data.gridPvAveragePowerAC
        })
    }

    useEffect(() => {
        if (!props.gatewayID) return

        const
            windowProtocol = window.location.protocol,
            wsProtocol = windowProtocol.replace("http", "ws")

        let wsConnection = null

        if (window["WebSocket"]) {
            const url = `${wsProtocol}//${API_HOST}/ws/${props.gatewayID}`
                + "/devices/energy-info"

            wsConnection = new WebSocket(url, props.token)
            wsConnection.onerror = () => setError({ url })
            wsConnection.onmessage = e => updateData(JSON.parse(e.data).data)
            wsConnection.onopen = () => setError(null)
        }
        else {
            setWebSocketSupport(false)
        }
        return () => {
            if (wsConnection) wsConnection.close()
        }
    }, [props.gatewayID])

    const peakShaveRate = useMemo(() => {
        if (!peak.current || !peak.threshhold) return 0

        const rawRate = peak.current / peak.threshhold

        return rawRate <= 1 ? rawRate : 1
    }, [peak.current, peak.threshhold])

    const
        batteryData = [
            { name: commonT("stateOfCharge"), value: `${battery.state}%` },
            {
                name: commonT("batteryPower"),
                value: `${battery.power} ${commonT("kw")}`
            },
            {
                name: commonT("importFromGrid"),
                value: `${battery.import} ${commonT("kw")}`
            },
            { name: pageT(battery.direction), value: battery.target || "-" }
        ],
        loadData = [
            { name: pageT("solar"), value: `${load.solar} ${commonT("kw")}` },
            { name: pageT("batteryDischarge"), value: `${load.discharge} ${commonT("kw")}` },
            {
                name: pageT("importFromGrid"),
                value: `${load.import} ${commonT("kw")}`
            }
        ],
        solarData = [
            {
                name: pageT("directConsumption"),
                value: `${solar.consume} ${commonT("kw")}`
            },
            {
                name: pageT("chargeToBattery"),
                value: `${solar.charge} ${commonT("kw")}`
            },
            { name: commonT("exportToGrid"), value: `${solar.export} ${commonT("kw")}` },
        ]

    const
        BatteryCard = props =>
            <CardOnDiagram
                arrow={props.arrow}
                icon="battery"
                title={commonT("battery")}
                value={`${diagram.battery} ${commonT("kw")}`} />,
        GridCard = props =>
            <CardOnDiagram
                arrow={props.arrow}
                icon="grid"
                title={commonT("grid")}
                value={`${diagram.grid} ${commonT("kw")}`} />,
        LoadCard = props =>
            <CardOnDiagram
                arrow={props.arrow}
                icon="home"
                title={pageT("load")}
                value={`${diagram.load} ${commonT("kw")}`} />,
        SolarCard = props =>
            <CardOnDiagram
                arrow={props.arrow}
                icon="solar"
                title={commonT("solar")}
                value={`${diagram.solar} ${commonT("kw")}`} />

    const activeIndicator = peak.active
        ? <>
            <div className="grid h-6 ml-6 mr-1 place-items-center relative w-6">
                <div className="absolute animate-ping bg-negative-main h-3
                            rounded-full w-3" />
                <div className="bg-negative-main h-2.5 rounded-full w-2.5" />
            </div>
            <h5 className="text-negative-main">{commonT("active")}</h5>
        </>
        : <>
            <div className="grid h-6 ml-6 mr-1 place-items-center relative w-6">
                <div className="absolute bg-gray-300 h-3
                    rounded-full w-3" />
                <div className="bg-gray-300 h-2.5 rounded-full w-2.5" />
            </div>
            <h5 className="">{commonT("inactive")}</h5>
        </>

    const peakShaveColor = peak.active ? "negative" : "positive"

    return <>
        <h1 className="mb-9">{pageT("dashboard")}</h1>
        {websocketSupport
            ? null
            : <div className="box mb-8 negative text-center">
                {pageT("noWebsocketSupport")}
            </div>}
        {error
            ? <AlertBox
                boxClass="mb-8 negative"
                content={<div>
                    {pageT("unableToConnect")}
                    <p className="break-all font-mono mt-2 pt-2 border-t border-red-400">
                        {error.url}
                    </p>
                </div>}
                icon={ReportProblemIcon}
                iconColor="negative-main" />
            : null}
        <div className="lg:flex items-start">
            <div className="flex-1">
                <div className="card">
                    <div className="flex flex-wrap items-baseline mb-6">
                        <h5 className="font-bold">{commonT("peakShave")}</h5>
                        {activeIndicator}
                    </div>
                    <div className="bg-gray-600 flex h-2 overflow-hidden
                                    rounded-full w-full ">
                        <div
                            className={`bg-${peakShaveColor}-main rounded-full`}
                            style={{ width: `${peakShaveRate * 100}%` }}
                        />
                    </div>
                    <div className="flex font-bold items-center mt-2">
                        <div className="text-28px">
                            <span className={`text-${peakShaveColor}-main`}>
                                {peak.current}
                            </span> / {peak.threshhold} {commonT("kw")}
                        </div>
                    </div>
                </div>
                <div className="grid-cols-3-auto items-stretch
                                block md:grid lg:block xl:grid
                                mx-5 lg:mx-9 my-10 lg:my-18">
                    <div className="grid-cols-2 mb-4
                                    grid md:hidden lg:grid xl:hidden">
                        <div className="flex"><LoadCard arrow="bottom" /></div>
                        <div className="flex justify-end">
                            <GridCard arrow="bottom" />
                        </div>
                    </div>
                    <div className="flex-col justify-between
                                    hidden md:flex lg:hidden xl:flex">
                        <div><LoadCard arrow="right" /></div>
                        <div><BatteryCard arrow="right" /></div>
                    </div>
                    <ElectricalGridDiagram
                        className="h-auto w-full
                                    md:py-8 lg:py-0 xl:py-8"
                        data={lines} />
                    <div className="grid-cols-2 mt-4
                                    grid md:hidden lg:grid xl:hidden">
                        <div className="flex"><BatteryCard arrow="top" /></div>
                        <div className="flex justify-end">
                            <SolarCard arrow="top" />
                        </div>
                    </div>
                    <div className="flex-col justify-between
                                    hidden md:flex lg:hidden xl:flex">
                        <div><GridCard arrow="left" /></div>
                        <div><SolarCard arrow="left" /></div>
                    </div>
                </div>
            </div>
            <div className="lg:ml-5 gap-5 grid md:grid-cols-2 lg:block lg:w-88">
                <Card data={loadData} icon="home" title={pageT("load")} />
                <Card
                    className="lg:mt-5"
                    data={batteryData}
                    icon="battery"
                    title={commonT("battery")} />
                <Card
                    className="lg:mt-5"
                    data={solarData}
                    icon="solar"
                    title={commonT("solar")} />
            </div>
        </div>
    </>
})