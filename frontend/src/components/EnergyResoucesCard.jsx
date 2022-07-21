import { Fragment } from "react"

export default function EnergyResoucesCard (props) {
    const Icon = props.icon

    return <div className="card short">
        <div className="flex flex-wrap items-center mb-8">
            <div className="bg-gray-400-opacity-20 grid h-12 mr-4
                            place-items-center rounded-full w-12">
                <Icon className="h-8 text-gray-400 w-8" />
            </div>
            <h5>{props.title}</h5>
        </div>
        <div className="grid grid-cols-1fr-1px-1fr gap-5">
        {props.data.map((item, i) => <Fragment key={"erg-d-" + i}>
            {i > 0 ? <div className="bg-gray-400"></div> : null}
            <div>
                <h6 className="mb-2">{item.title}</h6>
                <h3>{item.value}</h3>
            </div>
        </Fragment>)}
        </div>
    </div>
}