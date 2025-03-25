const
    envPort = import.meta.env.VITE_APP_API_PORT,
    host = import.meta.env.VITE_APP_API_ENDPOINT || window.location.hostname,
    windowPort = window.location.port,
    port = (envPort || windowPort) ? (":" + (envPort || windowPort)) : ""

export const API_HOST = host + port