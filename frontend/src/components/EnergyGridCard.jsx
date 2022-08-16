import { useTranslation } from "react-multi-lang"

import { ReactComponent as GridExportIcon } from "../assets/icons/grid_export.svg"
import { ReactComponent as GridImportIcon } from "../assets/icons/grid_import.svg"
import { ReactComponent as NetExportIcon } from "../assets/icons/emergy_export.svg"

const icons = {
    exportToGrid: GridExportIcon,
    importFromGrid: GridImportIcon,
    netExport: NetExportIcon
}

export default function EnergyGridCard(props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("energyResources.grid." + string, params)

    const { kwh, types } = props.data

    return <div className="card">
        <div className="flex flex-wrap items-baseline mb-8">
            <h5 className="font-bold">{props.title}</h5>
        </div>
        <div className="lg:grid grid-cols-3 three-columns gap-x-5">
            {types.map((t, i) => {
                const Icon = icons[t.type]
                return <div key={"detail-" + i} className="flex justify-between items-end">
                    <div>
                        <h6 className="font-bold text-white mb-2">{pageT(t.type)}</h6>
                        <h3 className="my-2.5">{t.kwh} {commonT("kwh")}</h3>
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