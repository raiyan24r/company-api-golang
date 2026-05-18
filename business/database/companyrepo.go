package database

import (
	"database/sql"
	_ "embed"
)

type CompanyRepo struct {
	db *sql.DB
}

type Company struct {
	ID                string
	Name              string
	Description       string
	AmountOfEmployees int
	Registered        bool
	Type              string
}

var (
	//go:embed queries/company/insert.sql
	insertCompanyQuery string
	//go:embed queries/company/select_all.sql
	selectAllCompaniesQuery string
	//go:embed queries/company/select_by_id.sql
	selectCompanyByIDQuery string
	//go:embed queries/company/update.sql
	updateCompanyQuery string
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

func (repo *CompanyRepo) GetAllCompanies() ([]Company, error) {
	rows, err := repo.db.Query(selectAllCompaniesQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companies []Company
	for rows.Next() {
		var company Company
		err := rows.Scan(&company.ID, &company.Name, &company.Description, &company.AmountOfEmployees, &company.Registered, &company.Type)
		if err != nil {
			return nil, err
		}
		companies = append(companies, company)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return companies, nil
}

func (repo *CompanyRepo) GetCompanyByID(id string) (*Company, error) {
	row := repo.db.QueryRow(selectCompanyByIDQuery, id)

	var company Company
	err := row.Scan(&company.ID, &company.Name, &company.Description, &company.AmountOfEmployees, &company.Registered, &company.Type)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &company, nil
}

func (repo *CompanyRepo) UpdateCompany(id string, company Company) error {
	_, err := repo.db.Exec(updateCompanyQuery, company.Name, company.Description, company.AmountOfEmployees, company.Registered, company.Type, id)
	if err != nil {
		return err
	}

	return nil
}