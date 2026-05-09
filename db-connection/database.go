package dbconnection

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jimmymuthoni/async_api/config"
)


func NewPostgresConn(conf *config.Config) (*sql.DB, error){
	dburl := conf.DatabaseURL()
	db, err := sql.Open("posgres", dburl)
	if err != nil {
		return nil, fmt.Errorf("Failed to open database connction: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("Failed to ping database connection: %w", err)
	}
	return db, nil
}
