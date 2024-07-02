package main

import (
	"github.com/XT3RM1NAT0R/time-tracker/config"
	"github.com/XT3RM1NAT0R/time-tracker/internal/app/db"
	"github.com/XT3RM1NAT0R/time-tracker/internal/app/server"
)

func main() {
	cfg := config.Load()
	postgres := db.ConnectToDB(cfg)

	server.RunHTTPServer(cfg, postgres)
}
