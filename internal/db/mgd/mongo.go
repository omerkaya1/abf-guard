package mgd

import (
	"context"
	"fmt"
	"github.com/omerkaya1/abf-guard/internal/domain/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

// MongoStorage object holds everything related to the DB interactions
type MongoStorage struct {
	db *mongo.Client
}

// NewMongoStorage returns new PsqlStorage object to the callee
func NewMongoStorage(cfg config.DBConf) (*MongoStorage, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", cfg.Host, cfg.Port)))
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	return &MongoStorage{db: client}, nil
}

// Authenticate .
func (ps *MongoStorage) Authenticate(context.Context, string, string, string) (bool, error) {
	return true, nil
}

// Flash .
func (ps *MongoStorage) Flash(context.Context, string, string, string) (bool, error) {
	return true, nil
}

// Add .
func (ps *MongoStorage) Add(context.Context, string, string, string) (bool, error) {
	return true, nil
}

// Delete .
func (ps *MongoStorage) Delete(context.Context, string, string, string) (bool, error) {
	return true, nil
}
