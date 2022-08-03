import { Fragment } from "react"
import { ReactComponent as UpIcon } from "../assets/icons/up.svg"

export default function EnergySolarSubCard(props) {
    const Icon = props.icon

    return <div className="card short flex justify-between">
        <div className="grid grid-cols-1fr-1px-1fr gap-5">
            {props.data.map((item, i) => <Fragment key={"erg-s-" + i}>
                {i > 0 ? <div className="bg-gray-400"></div> : null}
                <div>
                    <h5 className="mb-4">{item.title}</h5>
                    <h3>{item.value}</h3>
                </div>
            </Fragment>)}
        </div>
        <div className="flex flex-wrap items-center">
            <div className="items-center bg-gray-400-opacity-20 grid h-12 w-12
                            place-items-center rounded-full">
                <Icon className="h-8 text-gray-400 w-8" />
            </div>
        </div>
    </div>
}