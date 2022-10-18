import { useTranslation } from "react-multi-lang"

const colors = {
    battery: "bg-blue-main",
    grid: "bg-indigo-main",
    solar: "bg-yellow-main"
}

export default function EnergyCard(props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string)

    const { kwh, types } = props.data

    return <div className="card energyCard">
        <div className="flex flex-wrap items-baseline mb-8">
            <h2 className="mr-2 whitespace-nowrap">{kwh} {commonT("kwh")}</h2>
            <h5 className="font-bold">{props.title}</h5>
        </div>
        <div className="flex h-2 overflow-hidden rounded-full w-full bg-gray-600">
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
                    {commonT(t.type)}
                </div>)}
        </div>
        <div className="grid grid-cols-3 column-separator gap-x-5 sm:gap-x-10">
            {types.map((t, i) =>
                <div key={"detail-" + i}>
                    <h6 className="font-bold text-white">{commonT(t.type)}</h6>
                    <h3 className="my-1">{t.percentage}%</h3>
                    <p className="lg:test text-13px text-white">
                        {t.kwh} {commonT("kwh")}
                    </p>
                </div>)}
        </div>
    </div>
}