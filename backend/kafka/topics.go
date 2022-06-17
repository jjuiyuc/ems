package kafka

const (
	// Weather
	ReceiveWeatherData       = "open-data.weather.cdc.forecast.0"
	SendWeatherDatatoLocalGW = "core.weather.cdc.forecast.{gw-id}.0"
	// CC(Current Condition)
	ReceiveLocalCCData  = "iot.cc.fct.record.0"
	SendParamsToCloudCC = "core.cc.cmd.cloud-cc-params.0"
)
