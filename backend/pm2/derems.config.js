const path = require("path");
const binPath = "/opt/derems/sbin";
const args = "-d /opt/derems/etc -e derems.yaml";


module.exports = {
  apps: [
    /***** core-derems *****/
    {
      name: "core-weather-worker",
      script: path.join(binPath, "weather-worker"),
      args: args,
      namespace: "core-derems",
    },
    {
      name: "core-api-worker",
      script: path.join(binPath, "api-worker"),
      args: args,
      namespace: "core-derems",
    },
    {
      name: "core-local-cc-worker",
      script: path.join(binPath, "local-cc-worker"),
      args: args,
      namespace: "core-derems",
    },
    {
      name: "core-billing-worker",
      script: path.join(binPath, "billing-worker"),
      args: args,
      namespace: "core-derems",
    },
    {
      name: "core-local-ai-worker",
      script: path.join(binPath, "local-ai-worker"),
      args: args,
      namespace: "core-derems",
    },
    /***** test-derems *****/
    {
      name: "test-local-cc-uplink-worker",
      script: path.join(binPath, "local-cc-uplink-worker.sh"),
      namespace: "test-derems",
    },
  ]
}
