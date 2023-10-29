package libsql

import (
	"database/sql"
	_ "github.com/libsql/libsql-client-go/libsql"
)

func NewLibSqlConnection(url string) (*sql.DB, error) {
	return sql.Open("libsql", url)
}
