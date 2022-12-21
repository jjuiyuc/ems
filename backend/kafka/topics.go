package kafka

const (
	// ReceiveWeatherData godoc
	ReceiveWeatherData = "open-data.weather.cdc.forecast.0"
	// SendWeatherDataToLocalGW godoc
	SendWeatherDataToLocalGW = "core.weather.cdc.forecast.{gw-id}.0"

	// SendAIBillingParamsToLocalGW godoc
	SendAIBillingParamsToLocalGW = "core.ai.cmd.billing-params.{gw-id}.0"

	// SendLeapBiddingDispatchToLocalGW godoc
	SendLeapBiddingDispatchToLocalGW = "core.ai.fct.leap-bidding-dispatch.{gw-id}.0"

	// ReceiveLocalCCData godoc
	ReceiveLocalCCData = "iot.cc.fct.record.0"

	// ReceiveLocalAIData godoc
	ReceiveLocalAIData = "iot.ai.fct.record.0"
)
