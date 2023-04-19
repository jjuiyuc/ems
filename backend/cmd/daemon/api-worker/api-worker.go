package main

import (
	"flag"

	"der-ems/config"
	"der-ems/models"
	"der-ems/repository"
	"der-ems/routers"
	"der-ems/services"
)

func main() {
	dir := flag.String("d", "../../../config", "")
	env := flag.String("e", "template", "")
	flag.Parse()
	config.Init(*dir, *env)
	cfg := config.GetConfig()
	models.Init(cfg)
	db := models.GetDB()
	defer models.Close()

	repo := repository.NewRepository(db)
	services := services.NewServices(cfg, repo)

	routers.NewAPIWorker(*dir, cfg, services)
}
