package kafka

const (
	// ReceiveWeatherData godoc
	ReceiveWeatherData = "open-data.weather.cdc.forecast.0"
	// SendWeatherDatatoLocalGW godoc
	SendWeatherDatatoLocalGW = "core.weather.cdc.forecast.{gw-id}.0"

	// SendAIBillingParamsToLocalGW godoc
	SendAIBillingParamsToLocalGW = "core.ai.cmd.billing-params.{gw-id}.0"

	// ReceiveLocalCCData godoc
	ReceiveLocalCCData = "iot.cc.fct.record.0"
)
