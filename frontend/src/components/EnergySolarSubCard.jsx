export default function EnergySolarSubCard(props) {

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
            <div className="bg-gray-400-opacity-20 grid h-16 w-16
                            place-items-center rounded-full">
                <Icon className="h-9 text-gray-400 w-9" />
            </div>
        </div>
    </>
}