package main

import (
	"flag"

	"der-ems/apps"
	"der-ems/config"
	"der-ems/infra"
	"der-ems/models"
	"der-ems/repository"
	"der-ems/services"
)

func main() {
	name := "weather-worker"
	dir := flag.String("d", "../../../config", "")
	env := flag.String("e", "ut", "")
	flag.Parse()
	config.Init(*dir, *env)
	cfg := config.GetConfig()
	models.Init(cfg)
	db := models.GetDB()
	defer models.Close()
	repo := repository.NewRepository(db)
	weather := services.NewWeatherService(repo)

	mdWeatherWorker := apps.NewWeatherWorker(infra.GetGracefulShutdownCtx(), cfg, repo, weather, name)
	mdWeatherWorker.MainLoop()
}
