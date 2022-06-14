package main

import (
	"der-ems/apps"
	"der-ems/config"
	"der-ems/infra"
	"der-ems/models"
	"der-ems/repository"
	"flag"
)

func main() {
	name := "weather-worker"
	dir := flag.String("d", "../../../config", "")
	env := flag.String("e", "template", "")
	flag.Parse()
	config.Init(*dir, *env)
	cfg := config.GetConfig()
	models.Init(cfg)
	db := models.GetDB()
	defer models.Close()

	repo := repository.NewRepository(db)

	mdWeatherWorker := apps.NewWeatherWorker(infra.GetGracefulShutdownCtx(), cfg, repo, name)

	mdWeatherWorker.MainLoop()
}
