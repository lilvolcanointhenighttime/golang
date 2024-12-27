package modules

import (
	"context"
)


func (db *DB) createTableCustomers() error {
	ctx := context.Background()

	conn, err := db.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	err = conn.QueryRow(ctx,
		`CREATE TABLE customers (id SERIAL PRIMARY KEY, name TEXT NOT NULL, name TEXT, birthdate TIMESTAMP)`,
	).Scan()

	return err
}