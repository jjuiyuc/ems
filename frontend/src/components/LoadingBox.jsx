import Spinner from "../components/Spinner"

const LoadingBox = ({ loading }) => loading
    ? <div className="absolute bg-black-main-opacity-95 grid inset-0
    place-items-center rounded-3xl">
        <Spinner />
    </div>
    : null

export default LoadingBox