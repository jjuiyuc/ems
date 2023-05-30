package kafka

const (
	// ReceiveWeatherData godoc
	ReceiveWeatherData = "open-data.weather.cdc.forecast.0"
	// SendWeatherDataToLocalGW godoc
	SendWeatherDataToLocalGW = "core.weather.cdc.forecast.{gw-id}.0"
	// SendGPSLocation godoc
	SendGPSLocation = "core.location.fct.gps.0"

	// SendDeviceMappingToLocalGW godoc
	SendDeviceMappingToLocalGW = "core.meta.cmd.gw-ed-mapping.{gw-id}.0"

	// SendAIBillingParamsToLocalGW godoc
	SendAIBillingParamsToLocalGW = "core.ai.cmd.billing-params.{gw-id}.0"

	// SendLeapNotificationToLocalGW godoc
	SendLeapNotificationToLocalGW = "core.leap.cdc.notification.{gw-id}.0"

	// ReceiveLocalCCData godoc
	ReceiveLocalCCData = "iot.cc.fct.record.0"

	// ReceiveLocalAIData godoc
	ReceiveLocalAIData = "iot.ai.fct.record.0"
)
