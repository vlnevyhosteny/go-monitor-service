package database

import "database/sql"

const (
	pragma = `PRAGMA foreign_keys = ON;`

	endpoint = `
		CREATE TABLE IF NOT EXISTS endpoint (
			id STRING PRIMARY KEY,
			name STRING NOT NULL UNIQUE,
			url STRING NOT NULL,
			schedule STRING 
		);	
	`

	result = `
		CREATE TABLE IF NOT EXISTS result (
			id STRING PRIMARY KEY,
			statusCode INTEGER NOT NULL,
			body STRING,
			endpointId STRING NOT NULL,
			FOREIGN KEY (endpointId) REFERENCES endpoint(id)
		);
	`
)

func Init(db *sql.DB) error {
	_, err := db.Exec(pragma)
	if err != nil {
		return err
	}

	_, err = db.Exec(endpoint)
	if err != nil {
		return err
	}

	_, err = db.Exec(result)
	if err != nil {
		return err
	}

	return nil
}
