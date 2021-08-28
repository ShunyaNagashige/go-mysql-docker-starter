package main

import (
	"github.com/ShunyaNagashige/go-mysql-docker-starter/config"
	"github.com/ShunyaNagashige/go-mysql-docker-starter/model"
	"github.com/ShunyaNagashige/go-mysql-docker-starter/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)

	model.Open()
}
