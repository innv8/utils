package utils

import (
	"database/sql"
	"time"
)

// DBConnect
// Connects to DB
// Takes driver, uri
// Returns client, error
func DBConnect(driver, uri string) (client *sql.DB, err error) {
	LogINFO("[db] connecting to %s db at %s", driver, uri)
	client, err = sql.Open(driver, uri)
	if err != nil {
		LogError("[db] unable to connect to db because %v", err)
		return
	}

	if err = client.Ping(); err != nil {
		LogError("[db] unable to ping db because %v", err)
		return
	}

	client.SetMaxOpenConns(100)
	client.SetMaxIdleConns(64)
	client.SetConnMaxIdleTime(time.Second * 3600)
	client.SetConnMaxLifetime(time.Hour)
	LogINFO("[db] connected to db successfully")
	return client, nil
}

// LogDBStats
// Log db Stats
// Takes db client
// Retusn nothing
func LogDBStats(client *sql.DB) {
	stats := client.Stats()
	LogINFO("[db] Stats :: Max=%d\tOpen=%d\tInUse=%d\tIdle=%d\tWaitCount=%d\tWaitDuration=%.4fs",
		stats.MaxOpenConnections,
		stats.OpenConnections,
		stats.InUse,
		stats.Idle,
		stats.WaitCount,
		stats.WaitDuration.Seconds(),
	)
}
