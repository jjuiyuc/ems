const Spinner = ({borderWidth="border-4", size="h-10 w-10"}) =>
    <div className={`relative ${size}`}>
        <div className={`${borderWidth} border-primary-main-opacity-30 h-full
                        rounded-full w-full`} />
        <div className={`absolute animate-spin ${borderWidth}
                        border-transparent border-r-primary-main
                        border-t-primary-main h-full inset-0 rounded-full
                        w-full`} />
    </div>
export default Spinner