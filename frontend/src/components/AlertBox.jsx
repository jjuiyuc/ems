function AlertBox (props) {
    const
        {boxClass, content, icon, iconColor} = props,
        color = iconColor || "primary-main",
        Icon = icon

    return <div className={`box text-center${boxClass ? " " + boxClass : ""}`}>
        <Icon className={`mb-2 text-${color}`} fontSize="large" />
        {content}
    </div>
}

export default AlertBox