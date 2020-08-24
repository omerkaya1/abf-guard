package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/omerkaya1/abf-guard/internal/config"

	// We absolutely need this import and this comment.
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

// PsqlStorage object holds everything related to the DB interactions
type PsqlStorage struct {
	db *sqlx.DB
}

var (
	// ErrAlreadyStored reports IP duplication errors
	ErrAlreadyStored = errors.New("provided IP is already stored")
	// ErrDoesNotExist reports unknown IP querying errors
	ErrDoesNotExist = errors.New("provided IP does not exist in the DB")
	// ErrIsInTheBlacklist reports querying blacklisted IP from the whitelist
	ErrIsInTheBlacklist = errors.New("the ip is in the blacklist")
)

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
	// Check whether an IP is present in the DB
	if ok, err := ps.checkIPIsPresent(ctx, blacklist, ip); err != nil {
		return err
	} else if ok {
		return ErrAlreadyStored
	}
	// Make a DB request
	_, err := ps.db.ExecContext(ctx, "INSERT INTO ip_list(id, ip, bl) VALUES(default, $1, $2)", ip, blacklist)
	return err
}

// Delete method is used to delete an IP address from a specified list (black or white)
func (ps *PsqlStorage) Delete(ctx context.Context, ip string, blacklist bool) (int64, error) {
	// Check whether an IP is present in the DB
	if ok, err := ps.checkIPIsPresent(ctx, blacklist, ip); err != nil {
		return 0, err
	} else if !ok {
		return 0, ErrDoesNotExist
	}
	// Make a DB request
	r, err := ps.db.ExecContext(ctx, "DELETE FROM ip_list WHERE ip=$1", ip)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// GetIPList returns an IP list requested by the callee (black or white)
func (ps *PsqlStorage) GetIPList(ctx context.Context, blacklist bool) ([]string, error) {
	// Make a DB request
	rows, err := ps.db.QueryxContext(ctx, "SELECT * FROM ip_list WHERE bl=$1", blacklist)
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
		return ErrDoesNotExist
	} else if ok {
		return ErrIsInTheBlacklist
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
