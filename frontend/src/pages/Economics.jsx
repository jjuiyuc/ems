import { connect } from "react-redux"
import { useEffect, useState } from "react"
import { useTranslation } from "react-multi-lang"

import { API_HOST } from "../constant/env"
// import PriceCard from "../components/PriceCard"




export default function Economics(props) {

    const
        t = useTranslation(),
        commonT = string => t("common." + string),
        pageT = (string, params) => t("economics." + string, params)




    return <>
        <h1 className="mb-9">{pageT("economics")}</h1>
        <div className="flex lg:max-w-7xl">
            <div className="card w-3/12">
                <h6>{pageT("total")}</h6>
            </div>
            <div className="flex flex-wrap">
                <div className="row flex">
                    <div className="card">
                        <h6>{pageT("ancillaryServices")}</h6>
                        <h2>$15</h2>
                    </div>
                    <div className="card">
                        <h6>{pageT("demandCharge")}</h6>
                        <h2>$150</h2>
                    </div>
                    <div className="card">
                        <h6>{pageT("timeOfUseArbitrage")}</h6>
                        <h2>$150</h2>
                    </div>
                </div>
                <div className="row flex">
                    <div className="card ">
                        <h6>{pageT("renewableEnergyCertificate")}</h6>
                        <h2>$150</h2>
                    </div>
                    <div className="card">
                        <h6>{pageT("solarLocalUsage")}</h6>
                        <h2>$150</h2>
                    </div>
                    <div className="card">
                        <h6>{pageT("exportToGrid")}</h6>
                        <h2>$150</h2>
                    </div>
                </div>
            </div>


        </div>

    </>
}