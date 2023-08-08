package main

import (
	"flag"

	"der-ems/apps"
	"der-ems/config"
	"der-ems/infra"
	"der-ems/models"
	"der-ems/repository"
)

func main() {
	name := "local-cc-worker"
	dir := flag.String("d", "../../../config", "")
	env := flag.String("e", "ut", "")
	flag.Parse()
	config.Init(*dir, *env)
	cfg := config.GetConfig()
	models.Init(cfg)
	db := models.GetDB()
	defer models.Close()
	repo := repository.NewRepository(db)

	mdLocalCCWorker := apps.NewLocalCCWorker(infra.GetGracefulShutdownCtx(), cfg, repo, name)
	mdLocalCCWorker.MainLoop()
}
