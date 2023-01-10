import { ReactComponent as EconomicsIcon } from "../assets/icons/economics.svg"
import { useTranslation } from "react-multi-lang"

export default function EconomicsCard(props) {
    const
        t = useTranslation(),
        pageT = (string, params) => t("economics." + string, params)

    const { thisMonth, type } = props.data

    return <>
        <div className="">
            <div className="card short grid">
                <div className="flex justify-between items-center">
                    <div>
                        <div className="flex flex-wrap items-center mb-4">
                            <h5 className="mr-1">{pageT(type + "Ubiik")}</h5>
                            <label className="bg-gray-600 font-normal px-2 py-1
                                rounded-3xl text-11px">
                                {pageT("thisMonth")}
                            </label>
                        </div>
                        <h2 className="flex">${thisMonth}</h2>
                    </div>
                    <div className="bg-gray-400-opacity-20 grid h-16 min-w-16
                            place-items-center rounded-full w-16">
                        <EconomicsIcon className="h-9 text-gray-400 w-9" />
                    </div>
                </div>
                {["lastMonth", "sameDayLastYear"].map(t =>
                    props.tabs.includes(t)
                        ? <div className="light-card font-bold mt-6">
                            <p className="mb-4">{pageT(t)}</p>
                            <h2>${props.data[t]}</h2>
                        </div>
                        : null)}
            </div>
        </div>
    </>
}