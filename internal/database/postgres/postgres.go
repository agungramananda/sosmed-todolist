package postgres

import (
	"fmt"

	"github.com/agungramananda/sosmed-todolist/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

func New(logger *zerolog.Logger, conf *config.DBConfig) (*sqlx.DB, error) {
	datasource := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		conf.Username, conf.Password, conf.Host, conf.Port, conf.Database,
	)

	db, err := sqlx.Connect("pgx", datasource)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to connect to DB")
		return nil, err
	}

	logger.Info().Msg("DB initialized successfully")

	dbMigrate, err := migrate.New("file://internal/database/postgres/migrations", datasource)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to connect to migration engine")
		return nil, err
	}
	defer dbMigrate.Close()

	if err = dbMigrate.Up(); err != nil && err != migrate.ErrNoChange {
		logger.Error().Err(err).Msg("Failed to perform migrations")
		return nil, err
	} else if err == migrate.ErrNoChange {
		logger.Info().Msg("No new migrations to apply")
	}

	rev, isDirty, err := dbMigrate.Version()
	if err != nil && err != migrate.ErrNilVersion {
		logger.Error().Err(err).Msg("Failed to fetch migration version")
		return nil, err
	}

	if isDirty {
		logger.Warn().Msg("Database migration is in a dirty state")
	}

	if err == migrate.ErrNilVersion {
		logger.Info().Msg("DB Migration Version: None")
	} else {
		logger.Info().Msgf("DB Migration Version: %d", rev)
	}

	return db, nil
}
