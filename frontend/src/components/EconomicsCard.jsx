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
                            <h5 className="mr-1">{props.title}</h5>{subTitle}
                        </div>
                        <h2 className="flex">{props.value}</h2>
                    </div>
                    <div className="bg-gray-400-opacity-20 grid h-16 min-w-16
                            place-items-center rounded-full w-16">
                        <Icon className="h-9 text-gray-400 w-9" />
                    </div>
                </div>
                {props.tabs.includes("lastMonth")
                    ? <div className="light-card font-bold mt-6">
                        <p className="mb-4">{props.leftTitle}</p>
                        <h2>{props.leftValue}</h2>
                    </div>
                    : null}
                {props.tabs.includes("sameMonthLastYear")
                    ? <div className="light-card font-bold mt-6">
                        <p className="mb-4">{props.rightTitle}</p>
                        <h2>{props.rightValue}</h2>
                    </div>
                    : null}
            </div>
        </div>
    </>
}