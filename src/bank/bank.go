package bank

import (
	"API/src/config"
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mssql", config.ConnectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
