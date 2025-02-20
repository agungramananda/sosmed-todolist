package domain

import (
	"github.com/agungramananda/sosmed-todolist/internal/common/custom_validator"
	"github.com/agungramananda/sosmed-todolist/internal/common/exceptions"
	"github.com/agungramananda/sosmed-todolist/internal/common/middleware"
	"github.com/agungramananda/sosmed-todolist/internal/domain/brands"
	"github.com/agungramananda/sosmed-todolist/internal/domain/platforms"
	"github.com/agungramananda/sosmed-todolist/internal/domain/tasks"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	ecmiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func InitDomain(db *sqlx.DB, e *echo.Echo, logger *zerolog.Logger, validator *custom_validator.Validator){
	e.GET("/api/swagger/*", echoSwagger.WrapHandler)
	root := e.Group("/api/v1",
		ecmiddleware.RequestIDWithConfig(ecmiddleware.RequestIDConfig{Generator: uuid.NewString}),
		ecmiddleware.CORS(),
		middleware.RequestLogger(logger),
	)

	e.Validator = validator
	e.HTTPErrorHandler = exceptions.CustomHTTPErrorHandler(*logger)

	//brands
	brandsRepo := brands.NewRepository(db)
	brandsSvc := brands.NewService(brandsRepo)
	brands.NewController(brandsSvc).Route(root)

	//platforms
	platformsRepo := platforms.NewRepository(db)
	platformsSvc := platforms.NewService(platformsRepo)
	platforms.NewController(platformsSvc).Route(root)

	//tasks
	tasksRepo := tasks.NewRepository(db)
	tasksSvc := tasks.NewService(tasksRepo)
	tasks.NewController(tasksSvc).Route(root)
}