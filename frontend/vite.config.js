import { defineConfig } from "vite"
import react from "@vitejs/plugin-react"
import svgr from "vite-plugin-svgr"

// https://vitejs.dev/config/
export default defineConfig({
    base: "/ems/",
    plugins: [react(), svgr()],
})
// export default defineConfig({
//     define: {
//         "global": {},
//         "process.env": process.env
//     },
//     plugins: [react(), svgr()]
// })