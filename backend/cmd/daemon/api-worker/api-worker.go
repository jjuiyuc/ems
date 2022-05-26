package main

import (
	"der-ems/config"
	"der-ems/models"
	"der-ems/routers"
	"flag"
)

func main() {
	dir := flag.String("d", "../../../config", "")
	env := flag.String("e", "template", "")
	flag.Parse()
	config.Init(*dir, *env)
	models.Init()
	defer models.Close()

	routers.NewAPIWorker(config.GetConfig())
}
