//go:build cgo && !pq

package core

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/pocketbase/dbx"
)

func connectDB(dbPath string) (*dbx.DB, error) {
	db, err := dbx.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	if err := initPragmas(db); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
