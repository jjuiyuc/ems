import { connect } from "react-redux"
import { useEffect, useState } from "react"
import { useTranslation } from "react-multi-lang"

import { API_HOST } from "../constant/env"
import PriceCard from "../components/PriceCard"

import { ReactComponent as EconomicsIcon } from "../assets/icons/economics.svg"

export default function Economics(props) {

    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("economics." + string, params)
    const
        [total, setTotal] = useState(630),
        [ancillaryServices, setAncillaryServices] = useState(15),
        [demandCharge, setDemandCharge] = useState(150),
        [timeOfUseArbitrage, setTimeOfUseArbitrage] = useState(130),
        [renewableEnergyCertificate, setRenewableEnergyCertificate] = useState(150),
        [solarLocalUsage, setSolarLocalUsage] = useState(160),
        [exportToGrid, setExportToGrid] = useState(130)

    return <>
        <h1 className="mb-9">{pageT("economics")}</h1>
        <div className="flex max-w-7xl">
            <div className="card w-3/12 mb-8">
                <h5 className="font-bold">{pageT("total")}</h5>
                <h2 className="font-bold mb-1 pt-4">{pageT("february")} 2022</h2>
                <h2 className="font-bold mb-1 pt-4">${total}</h2>
                <div className="bg-primary-main-opacity-20 w-20 h-20 rounded-full relative">
                    <EconomicsIcon className="text-brand-main w-12 h-12 ml-3 absolute" />
                </div>
            </div>
            <div className="flex flex-wrap ml-5">
                <div className="lg:grid grid-cols-3 auto-cols-max">
                    <PriceCard
                        price={ancillaryServices}
                        title={pageT("ancillaryServices")} />
                    <PriceCard
                        price={demandCharge}
                        title={pageT("demandCharge")} />
                    <PriceCard
                        price={timeOfUseArbitrage}
                        title={pageT("timeOfUseArbitrage")} />
                    <PriceCard
                        price={renewableEnergyCertificate}
                        title={pageT("renewableEnergyCertificate")} />
                    <PriceCard
                        price={solarLocalUsage}
                        title={pageT("solarLocalUsage")} />
                    <PriceCard
                        price={exportToGrid}
                        title={pageT("exportToGrid")} />
                </div>
            </div>


        </div>

    </>
}