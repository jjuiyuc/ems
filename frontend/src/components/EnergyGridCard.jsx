import { useTranslation } from "react-multi-lang"

import { ReactComponent as GridExportIcon } from "../assets/icons/grid_export.svg"
import { ReactComponent as GridImportIcon } from "../assets/icons/grid_import.svg"
import { ReactComponent as NetExportIcon } from "../assets/icons/emergy_export.svg"

const icons = {
    importFromGrid: GridImportIcon,
    exportToGrid: GridExportIcon,
    netImport: NetExportIcon
}

export default function EnergyGridCard(props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("energyResources.grid." + string, params)

    return <div className="card">
        <h5 className="font-bold mb-8">{props.title}</h5>
        <div className="xl:grid grid-cols-3 gap-x-10 xl:column-separator">
            {props.data.map((t, i) => {
                const Icon = icons[t.type]
                return <div className={"border-gray-400 xl:border-0 py-4 xl:py-0"
                    + (i > 0 ? " border-t" : "")}
                    key={"detail-" + i} >
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