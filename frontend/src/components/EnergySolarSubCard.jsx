import { Button } from "@mui/material"
import { useTranslation } from "react-multi-lang"

export default function EnergySolarSubCard(props) {
    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("energyResources.solar." + string, params)

    const Icon = props.icon

    return <>
        <div className="card short flex justify-between items-center">
            <div>
                <div className="flex items-center mb-4">
                    <h5>{props.data.title}</h5>
                    {props.data.subTitle}
                </div>
                {props.data.value}
            </div>
            <div className="bg-gray-400-opacity-20 grid h-12 w-12
                            place-items-center rounded-full">
                <Icon className="h-8 text-gray-400 w-8" />
            </div>
        </div>
    </>
}