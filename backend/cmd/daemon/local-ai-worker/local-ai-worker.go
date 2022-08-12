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
	name := "local-ai-worker"
	dir := flag.String("d", "../../../config", "")
	env := flag.String("e", "template", "")
	flag.Parse()
	config.Init(*dir, *env)
	cfg := config.GetConfig()
	models.Init(cfg)
	db := models.GetDB()
	defer models.Close()
	repo := repository.NewRepository(db)

	mdLocalAIWorker := apps.NewLocalAIWorker(infra.GetGracefulShutdownCtx(), cfg, repo, name)
	mdLocalAIWorker.MainLoop()
}
