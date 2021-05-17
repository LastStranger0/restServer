package sqlite

import "database/sql"

type postgresSQL struct {
	dbConnection *sql.DB
}
