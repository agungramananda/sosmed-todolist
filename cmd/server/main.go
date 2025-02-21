package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/agungramananda/sosmed-todolist/config"
	"github.com/agungramananda/sosmed-todolist/docs"
	"github.com/agungramananda/sosmed-todolist/internal/common/custom_validator"
	"github.com/agungramananda/sosmed-todolist/internal/common/logger"
	"github.com/agungramananda/sosmed-todolist/internal/database/postgres"
	"github.com/agungramananda/sosmed-todolist/internal/domain"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// @title Sosmed Todolist API
// @version 1.0
// @description Simple API for to-do-list management posts on social media
// @termsOfService http://swagger.io/terms/

// @BasePath /api/v1

func main() {
	config := config.New()

	docs.SwaggerInfo.Host = config.SwaggerHost
	logger := logger.New()
	validator := custom_validator.NewCustomValidator(validator.New(validator.WithRequiredStructEnabled()))

	db, err := postgres.New(logger, config.DbConf)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to initialized db")
	}

	e := echo.New()
	domain.InitDomain(db,e,logger, validator)
	
	e.HideBanner = true
	e.HidePort = true

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func (){
		logger.Info().Msgf("starting service, listening at %s:%s", config.ServiceHost, config.ServicePort)
		if err := e.Start(fmt.Sprintf("%s:%s", config.ServiceHost, config.ServicePort)); err != http.ErrServerClosed {
			logger.Error().Err(err).Msg("failed to start service")
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		logger.Fatal().Err(err).Msg(err.Error())
	}
}