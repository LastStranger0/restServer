package sqlite

import (
	"log"
	"restserver/internal/structs/companystruct"
)

const getAllCompanysQuery = `SELECT * FROM Companys`

func (p postgresSQL) GetAllCompanys() []companystruct.Company {
	rows, err := p.dbConnection.Query(getAllCompanysQuery)
	if err != nil {
		log.Println("GetAllCompanys: ", err)
		return []companystruct.Company{}
	}
	defer rows.Close()

	companys := make([]companystruct.Company, 0, 10)
	for rows.Next() {
		company := companystruct.Company{}
		err = rows.Scan(&company.Id, &company.Name, &company.Origin, &company.Telephone)
		if err != nil {
			log.Println("GetAllCompanys: parse err", err)
			continue
		}
		companys = append(companys, company)
	}

	return companys
}

const addCompanyExec = `
	insert 
	into Companys 
	(name, origin, telephone)
	values (?, ?, ?)`

func (p postgresSQL) AddCompany(company companystruct.Company) bool {
	successful := true
	_, err := p.dbConnection.Exec(addCompanyExec,
		company.Name, company.Origin, company.Telephone)
	if err != nil {
		log.Println("AddCompany", err)
		successful = false
	}

	return successful
}

const deleteCompanyExec = `
	delete 
	from Companys 
	where id = ?`

func (p postgresSQL) DeleteCompany(id int) bool {
	successful := true
	_, err := p.dbConnection.Exec(deleteCompanyExec, id)
	if err != nil {
		log.Println("DeleteCompany", err)
		successful = false
	}

	return successful
}

const createTableCompanysExec = `
	CREATE TABLE companys (
	id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
	name varchar(30) not null,
	origin varchar(30) not null,
	telephone int not null
)`

func (p postgresSQL) CreateTableCompanys() {
	_, err := p.dbConnection.Exec(createTableCompanysExec)
	if err != nil {
		log.Println("CreateTableCompanys", err)
	}
}
