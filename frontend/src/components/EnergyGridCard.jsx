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

    return <div className="card">
        <h5 className="font-bold mb-8">{props.title}</h5>
        <div className="lg:grid grid-cols-3 three-columns gap-x-10">
            {props.data.map((t, i) => {
                const Icon = icons[t.type]
                return <div key={"detail-" + i} className="mb-4 lg:m-0">
                    <h6 className="font-bold">{pageT(t.type)}</h6>
                    <div className="flex justify-between items-center mt-3.5">
                        <h3>{t.kwh} {commonT("kwh")}</h3>
                        <div
                            className="bg-gray-400-opacity-20 h-12 w-12 grid
                            place-items-center rounded-full">
                            <Icon className="text-gray-400 w-8 h-8" />
                        </div>
                    </div>
                </div>
            }
            )}
        </div>
    </div>
}