package psql

import (
	"context"
	"fmt"
	"github.com/omerkaya1/abf-guard/internal/domain/config"
	"github.com/omerkaya1/abf-guard/internal/domain/errors"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

// PsqlStorage object holds everything related to the DB interactions
type PsqlStorage struct {
	db *sqlx.DB
}

// NewPsqlStorage returns new PsqlStorage object to the callee
func NewPsqlStorage(cfg config.DBConf) (*PsqlStorage, error) {
	if cfg.Name == "" || cfg.User == "" || cfg.SSL == "" || cfg.Password == "" {
		return nil, errors.ErrBadDBConfiguration
	}
	dsn := fmt.Sprintf("host=%s port=%s password=%s user=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Password, cfg.User, cfg.Name, cfg.SSL)
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, err
	}
	return &PsqlStorage{db: db}, nil
}

// Authenticate .
func (ps *PsqlStorage) Authenticate(context.Context, string, string, string) (bool, error) {
	return true, nil
}

// Flash .
func (ps *PsqlStorage) Flash(context.Context, string, string, string) (bool, error) {
	return true, nil
}

// Add .
func (ps *PsqlStorage) Add(context.Context, string, string, string) (bool, error) {
	return true, nil
}

// Delete .
func (ps *PsqlStorage) Delete(context.Context, string, string, string) (bool, error) {
	return true, nil
}
