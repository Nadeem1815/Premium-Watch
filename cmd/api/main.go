package main

import (
	"log"

	_ "github.com/nadeem1815/premium-watch/cmd/api/docs"
	"github.com/nadeem1815/premium-watch/pkg/config"
	"github.com/nadeem1815/premium-watch/pkg/di"
)

// @title Ecommerce REST API
// @version 1.0
// @description Ecommerce REST API built using Go Lang, PSQL, REST API following Clean Architecture. Hosted with Ngnix, AWS EC2 and RDS

// @contact.name Nadeem Fahad
// @contact.url https://github.com/Nadeem1815
// @contact.email nadeemf408@gmail.com

// @license.name MIT
////@host permium-watch.shop
// @host localhost:3000
// @license.url https://opensource.org/licenses/MIT

// @BasePath /
// //@schemes https
// @query.collection.format multi
func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config", configErr)

	}

	server, diErr := di.InitializerAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server:", diErr)
	} else {
		server.Start()
	}

}
