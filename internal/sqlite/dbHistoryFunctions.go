package sqlite

import (
	"log"
)

const saveChangeHistoryExec = `insert into change_history (id_user, id_field) values (?,?)`

func (p postgresSQL) SaveChangeHistory(idUser, idProduct string) error {
	_, err := p.dbConnection.Exec(saveChangeHistoryExec, idUser, idProduct)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

const createTableHistoryExec = `
	CREATE TABLE IF NOT EXISTS change_history (
	id	INTEGER NOT NULL UNIQUE,
	id_user	TEXT NOT NULL,
	id_field	TEXT NOT NULL,
	PRIMARY KEY(id))`

func (p postgresSQL) CreateTableHistory() {
	_, err := p.dbConnection.Exec(createTableHistoryExec)
	if err != nil {
		log.Println("CreateTableCompanys", err)
	}
}
