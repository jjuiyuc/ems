export default function PriceCard(props) {

    const { price } = props

    return <div className="card mr-5 md:mb-8">
        <div>
            <p>{props.title}</p>
            <h2 className="font-bold mb-1">${price}</h2>
        </div>
    </div>
}