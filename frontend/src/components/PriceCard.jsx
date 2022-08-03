const PriceCard = props =>
    <div className="card font-bold mb-8">
        <p className="mb-4">{props.title}</p>
        <h2>$ {props.price}</h2>
    </div>

export default PriceCard