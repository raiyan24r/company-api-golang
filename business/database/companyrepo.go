package database

import (
	"database/sql"
	_ "embed"
)

type CompanyRepo struct {
	db *sql.DB
}

type Company struct {
	Name            string
	Description     string
	AmountOfEmployees int
	Registered      bool
	Type            string
}

var (
	//go.embed queries/company/insert.sql
	insertCompanyQuery string
)

func (repo *CompanyRepo) CreateCompany(company Company) (int64, error) {
	result, err := repo.db.Exec(insertCompanyQuery, company.Name, company.Description, company.AmountOfEmployees, company.Registered, company.Type)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}