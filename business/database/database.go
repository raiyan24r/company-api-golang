package database

import "database/sql"

type Database struct {
	*CompanyRepo	
}

func New(mysqlDb *sql.DB) *Database {
	return &Database{
		&CompanyRepo{mysqlDb},
	}
}