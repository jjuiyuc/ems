package main

import (
	"flag"

	"der-ems/config"
	"der-ems/kafka"
	"der-ems/models"
)

func main() {
	dir := flag.String("d", "./config", "")
	env := flag.String("e", "template", "")
	flag.Parse()

	config.Init(*dir, *env)
	models.Init()

	kafka.WeatherConsumerWorker(
		[]string{kafka.ReceiveWeatherData},
		config.GetConfig().GetString("kafka.consumerGroupID"))
}
