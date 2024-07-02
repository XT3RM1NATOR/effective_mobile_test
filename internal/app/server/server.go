package server

import (
	"github.com/XT3RM1NAT0R/time-tracker/config"
	"github.com/XT3RM1NAT0R/time-tracker/internal/delivery"
	"github.com/jmoiron/sqlx"

	_ "github.com/XT3RM1NAT0R/time-tracker/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
	"os"
)

// RunHTTPServer
// @title Time Tracker
// @version 1.0
// @description This is the backend server for the test assignment.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name .
// @license.url .
// @host petstore.swagger.io
// @externalDocs.description  OpenAPI 2.0
// @BasePath /
func RunHTTPServer(cfg *config.Config, db *sqlx.DB) {
	e := echo.New()

	logger := logrus.New()
	logger.Out = os.Stdout

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} [${status}] ${method} ${uri} (${latency_human})\n",
		Output: logger.Out,
	}))
	e.Use(middleware.CORS())

	delivery.RegisterRoutes(e, cfg, db)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	//http://localhost:4200/swagger/index.html что бы увидеть доку

	if err := e.Start(cfg.Server.Port); err != nil {
		panic(err)
	}
}
