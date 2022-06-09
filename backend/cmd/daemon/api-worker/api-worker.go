package main

import (
	"der-ems/config"
	"der-ems/models"
	"der-ems/repository"
	"der-ems/routers"
	"der-ems/services"
	"flag"
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

	repo := repository.NewUserRepository(db)
	authService := services.NewAuthService(repo)
	emailService := services.NewEmailService(cfg)
	userService := services.NewUserService(repo)

	routers.NewAPIWorker(cfg, authService, emailService, userService)
}
