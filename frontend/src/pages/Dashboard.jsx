import {useState}from "react"
import {useTranslation} from "react-multi-lang"
import variables from "../configs/variables"



export default function Dashboard () {
    const
    t = useTranslation(),
    commonT = string => t("common." + string)

    const
    [current,setCurrent] = useState(13),
    [threshhold,setThreshhold] = useState(15)

    return <>
    <h1>Dashboard</h1>
    <div className="card">
        <div className="flex flex-wrap items-baseline mb-6">
            <h5 className="font-bold">Peak Shave</h5>
        </div>
        <div className="flex h-2 overflow-hidden rounded-full w-full bg-gray-600">
            <div
                className="bg-positive-main h-full"
                style={{ width:'40%' }}
            />
        </div>
        <div className="text-28px font-bold">
            <span className={current > threshhold ? "text-negative-main" : "text-positive-main"}>
                {current} 
            </span>/ {threshhold} {commonT("kw")}
        </div>
    </div>
       
    </>
    

}