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
func NewPsqlStorage(cfg *config.DBConf) (*PsqlStorage, error) {
	dsn := fmt.Sprintf("host=%s port=%s password=%s user=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Password, cfg.User, cfg.Name, cfg.SSL)
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(3)
	db.SetMaxIdleConns(2)
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
	query := "INSERT INTO ip_list(id, ip, bl) VALUES(default, $1, $2)"
	// Make a DB request
	_, err := ps.db.ExecContext(ctx, query, ip, blacklist)
	return err
}

// Delete method is used to delete an IP address from a specified list (black or white)
func (ps *PsqlStorage) Delete(ctx context.Context, ip string, blacklist bool) error {
	// Check IP
	if ip == "" {
		return errors.ErrEmptyIP
	}
	// Check whether an IP is present in the DB
	if ok, err := ps.checkIPIsPresent(ctx, blacklist, ip); err != nil {
		return err
	} else if !ok {
		return errors.ErrDoesNotExist
	}
	// Prepare a query
	query := "DELETE FROM ip_list WHERE ip=$1"
	// Make a DB request
	_, err := ps.db.ExecContext(ctx, query, ip)
	return err
}

// GetIPList returns an IP list requested by the callee (black or white)
func (ps *PsqlStorage) GetIPList(ctx context.Context, blacklist bool) ([]string, error) {
	// Prepare a query
	query := "SELECT * FROM ip_list WHERE bl=$1"
	// Make a DB request
	rows, err := ps.db.QueryxContext(ctx, query, blacklist)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// Prepare the return value
	ips := make([]string, 0)
	for rows.Next() {
		select {
		case <-ctx.Done():
			return ips, ctx.Err()
		default:
			id, ip, bl := "", "", false
			if err := rows.Scan(&id, &ip, &bl); err != nil {
				return ips, err
			}
			ips = append(ips, ip)
		}
	}
	return ips, rows.Err()
}

// GreenLightPass is a method that checks whether an ip is not in the black list
func (ps *PsqlStorage) GreenLightPass(ctx context.Context, ip string) error {
	if ok, err := ps.checkIPIsList(ctx, ip); err != nil && err == sql.ErrNoRows {
		// Does not exist, - needs a bucket
		return errors.ErrDoesNotExist
	} else if ok {
		//
		return errors.ErrIsInTheBlacklist
	} else {
		return nil
	}
}

func (ps *PsqlStorage) checkIPIsPresent(ctx context.Context, blacklist bool, ip string) (bool, error) {
	query, result := "SELECT ip FROM ip_list WHERE ip=$1", ""
	err := ps.db.GetContext(ctx, &result, query, ip)
	if err != nil && err == sql.ErrNoRows {
		return false, nil
	}
	return result != "", err
}

func (ps *PsqlStorage) checkIPIsList(ctx context.Context, ip string) (bool, error) {
	query, result := "SELECT bl FROM ip_list WHERE ip=$1", new(bool)
	return *result, ps.db.GetContext(ctx, result, query, ip)
}
