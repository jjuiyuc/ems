export default function EconomicsCard(props) {
    const
        Icon = props.icon,
        subTitle = props.subTitle
            ? <label className="bg-gray-600 font-normal px-2 py-1 rounded-3xl
                                text-11px">
                {props.subTitle}
            </label>
            : null

    return <>
        <div className="">
            <div className="card short grid">
                <div className="flex justify-between items-center">
                    <div>
                        <div className="flex flex-wrap items-center mb-4">
                            <h5 className="mr-1">{props.data[0].title}</h5>{subTitle}
                        </div>
                        <h2 className="flex">{props.data[0].value}</h2>
                    </div>
                    <div className="bg-gray-400-opacity-20 grid h-16 min-w-16
                            place-items-center rounded-full w-16">
                        <Icon className="h-9 text-gray-400 w-9" />
                    </div>
                </div>
                {props.tabs.includes("lastMonth")
                    ? <div className="light-card font-bold mt-6">
                        <p className="mb-4">{props.data[1].title}</p>
                        <h2>{props.data[1].value}</h2>
                    </div>
                    : null}
                {props.tabs.includes("sameMonthLastYear")
                    ? <div className="light-card font-bold mt-6">
                        <p className="mb-4">{props.data[2].title}</p>
                        <h2>{props.data[2].value}</h2>
                    </div>
                    : null}
            </div>
        </div>
    </>
}