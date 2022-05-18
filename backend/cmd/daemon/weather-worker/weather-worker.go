package main

import (
	"der-ems/apps"
	"der-ems/config"
	"der-ems/infra"
	"der-ems/models"
	"flag"
)

func main() {
	name := "weather-worker"
	dir := flag.String("d", "../../../config", "")
	env := flag.String("e", "template", "")
	flag.Parse()
	config.Init(*dir, *env)
	models.Init()
	defer models.Close()

	mdWeatherWorker := apps.NewWeatherWorker(infra.GetGracefulShutdownCtx(), config.GetConfig(), name)

	mdWeatherWorker.MainLoop()
}
