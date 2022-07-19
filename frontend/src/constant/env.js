const
    envPort = import.meta.env.VITE_APP_API_PORT,
    windowPort = window.location.port,
    port = (envPort || windowPort) ? (":" + (envPort || windowPort)) : ""

export const API_HOST = import.meta.env.VITE_APP_API_ENDPOINT + port