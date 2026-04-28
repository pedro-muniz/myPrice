package repositories

import "database/sql"

type IDAO interface {
	Write(query string, args ...any) (sql.Result, error)
	Read(query string, args ...any) (*sql.Rows, error)
}
