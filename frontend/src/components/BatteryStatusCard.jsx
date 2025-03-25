import { useTranslation } from "react-multi-lang"
import BatteryDiagram from "../components/BatteryDiagram"

export default function BatteryStatusCard({ data }) {

    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("timeOfUse." + string, params)

    return <div className="card">
        <div className="header">
            <h4>{pageT("batteryStatus")}</h4>
        </div>
        <div className="flex flex-wrap items-center justify-evenly">
            <BatteryDiagram state={data?.state} direction={data?.direction}
                target={data?.target}
            />
            <div className="column-separator grid grid-cols-3 my-6
                mw-88 gap-x-5 sm:gap-x-10">
                <div>
                    <h3 className="text-22px">{data.state}%</h3>
                    <span className="text-13px">
                        {commonT("stateOfCharge")}
                    </span>
                </div>
                <div>
                    <h3 className="text-22px">{data.power} {commonT("kw")}</h3>
                    <span className="text-13px">
                        {commonT("batteryPower")}
                    </span>
                </div>
                <div>
                    <h3 className="text-22px">{data.target ? data.target : "-"}</h3>
                    <span className="text-13px">
                        {pageT(data.direction)}
                    </span>
                </div>
            </div>
        </div>
    </div>
}