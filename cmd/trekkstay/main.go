package main

import (
	"flag"
	"os"
	"strconv"
	"trekkstay/api"
	"trekkstay/pkgs/log"
	"trekkstay/pkgs/transport/http/server"
)

var (
	configPath = flag.String("conf", "./env/dev.env", "application config path")
	migration  = flag.Bool("migrate", false, "migrate database")
)

func init() {
	flag.Parse()

	if err := os.Setenv("CONFIG_PATH", *configPath); err != nil {
		panic(err)
	}

	if err := os.Setenv("CONFIG_ALLOW_MIGRATION",
		strconv.FormatBool(*migration)); err != nil {
		panic(err)
	}
}

// @title							Trekkstay API - V1
// @version        					v1.0
// @description     				API system for Trekkstay - Hotel Booking System
// @termsOfService  				https://swagger.io/

// @contact.name   					Trekkstay Team
// @contact.url    					https://www.trekkstay.com
// @contact.email  					support@trekkstay.com

// @license.name  					Apache 2.0
// @license.url   					https://www.apache.org/licenses/LICENSE-2.0.html

// @host      						localhost:8888
// @BasePath  						/
// @securitydefinitions.apikey  	JWT
// @in                          	header
// @name                        	Authorization
func main() {
	log.JsonLogger.Info("Starting server...")

	server.MustRun(api.NewServer())
}
