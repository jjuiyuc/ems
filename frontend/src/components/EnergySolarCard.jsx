import { useTranslation } from "react-multi-lang"

import { ReactComponent as DirectUseIcon } from "../assets/icons/direct_use.svg"
import { ReactComponent as ChargedIcon } from "../assets/icons/battery_charged.svg"
import { ReactComponent as GridImportIcon } from "../assets/icons/grid_import.svg"

const colors = {
    directUsage: "bg-green-main",
    chargeToBattery: "bg-blue-main",
    exportToGrid: "bg-indigo-main"
}

const icons = {
    directUsage: DirectUseIcon,
    chargeToBattery: ChargedIcon,
    exportToGrid: GridImportIcon
}

export default function EnergySolarCard(props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("energyResources.solar." + string, params)

    const { kwh, types } = props.data

    return <div className="card">
        <div className="flex flex-wrap items-baseline mb-6">
            <h2 className="mr-2 whitespace-nowrap">{kwh} {commonT("kwh")}</h2>
            <h5 className="font-bold">{props.title}</h5>
        </div>
        <div className="flex h-2 overflow-hidden rounded-full w-full">
            {types.map((t, i) =>
                <div
                    className={colors[t.type] + " h-full"}
                    key={"bar-" + i}
                    style={{ width: t.percentage + "%" }} />)}
        </div>
        <div className="flex justify-center mb-12 mt-4">
            {types.map((t, i) =>
                <div
                    className="flex items-center mx-2.5 text-white text-xs"
                    key={"legend-" + i}>
                    <div className={colors[t.type] + " h-3 mr-2 rounded-full w-3"} />
                    {pageT(t.type)}
                </div>)}
        </div>
        <div className="grid-cols-3 three-columns gap-x-5 sm:gap-x-10 lg:grid">
            {types.map((t, i) => {
                const Icon = icons[t.type]
                return <div key={"detail-" + i} className="flex justify-between items-end">
                    <div>
                        <h6 className="font-bold text-white">{pageT(t.type)}</h6>
                        <h3 className="my-1">{t.percentage}%</h3>
                        <p className="lg:test text-13px text-white">
                            {t.kwh} {commonT("kwh")}
                        </p>
                    </div>
                    <div
                        className="bg-gray-400-opacity-20 grid h-12 w-12
                            place-items-center rounded-full">
                        <Icon className="h-8 text-gray-400 w-8" />
                    </div>
                </div>
            }

            )}
        </div>
    </div>
}