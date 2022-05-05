package main

import (
	"flag"

	"der-ems/config"
	"der-ems/kafka"
)

func main() {
	dir := flag.String("d", "./config", "")
	env := flag.String("e", "template", "")
	flag.Parse()

	config.Init(*dir, *env)

	kafka.ConsumerWorker(
		[]string{config.GetConfig().GetString("kafka.topic.receiveWeatherData")},
		config.GetConfig().GetString("kafka.consumerGroupID"))
}
