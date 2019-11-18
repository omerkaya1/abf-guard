package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/omerkaya1/abf-guard/internal/domain/config"
	"github.com/omerkaya1/abf-guard/internal/domain/errors"

	// We absolutely need this import and this comment.
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

// Add method is used to add an IP address to a specified list (black or white)
func (ps *PsqlStorage) Add(ctx context.Context, ip string, blacklist bool) error {
	// Check IP
	if ip == "" {
		return errors.ErrEmptyIP
	}
	// Check whether an IP is present in the DB
	if ok, err := ps.checkIPIsPresent(ctx, blacklist, ip); err != nil {
		return err
	} else if ok {
		return errors.ErrAlreadyStored
	}
	// Prepare a query
	query := ""
	if blacklist {
		query = "insert into blacklist (ip) values (:ip)"
	} else {
		query = "insert into whitelist (ip) values (:ip)"
	}
	// Make a DB request
	_, err := ps.db.ExecContext(ctx, query, ip)
	if err != nil {
		return err
	}
	return nil
}

// Delete method is used to delete an IP address from a specified list (black or white)
func (ps *PsqlStorage) Delete(ctx context.Context, ip string, blacklist bool) error {
	// Check IP
	if ip == "" {
		return errors.ErrEmptyIP
	}
	// Check whether an IP is present in the DB
	if ip, err := ps.checkIPIsPresent(ctx, blacklist, ip); err != nil {
		return err
	} else if !ip {
		return errors.ErrDoesNotExist
	}
	// Prepare a query
	query := ""
	if blacklist {
		query = "delete from blacklist where ip=$1"
	} else {
		query = "delete from whitelist where ip=$1"
	}
	// Make a DB request
	_, err := ps.db.ExecContext(ctx, query, ip)
	if err != nil {
		return err
	}
	return nil
}

// GetIPList returns an IP list requested by the callee (black or white)
func (ps *PsqlStorage) GetIPList(ctx context.Context, blacklist bool) ([]string, error) {
	// Prepare a query
	query := ""
	if blacklist {
		query = "select * from blacklist"
	} else {
		query = "select * from whitelist"
	}
	// Make a DB request
	rows, err := ps.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// Prepare the return value
	ips := make([]string, 10)
	for rows.Next() {
		select {
		case <-ctx.Done():
			return ips, ctx.Err()
		default:
			ip := ""
			if err := rows.Scan(&ip); err != nil {
				return ips, err
			}
			ips = append(ips, ip)
		}
	}
	// Always check the rows' error
	if err := rows.Err(); err != nil {
		return ips, err
	}
	return ips, nil
}

func (ps *PsqlStorage) checkIPIsPresent(ctx context.Context, blacklist bool, ip string) (bool, error) {
	query := ""
	if blacklist {
		query = "select * from blacklist where ip=$1"
	} else {
		query = "select * from whitelist where ip=$1"
	}
	err := ps.db.GetContext(ctx, &ip, query, ip)
	if err != nil && err == sql.ErrNoRows {
		return false, nil
	}
	return true, err
}
