package repositories

import (
	"database/sql"
	"fmt"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/db"
	"log"
	"folha/src/domain/entities"
)

func EmployeeFindAll() ([]*entities.Employee, error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, errQuery := db.Query(`SELECT e.id, e.name, e.salary FROM employee e`)
	if errQuery != nil {
		return nil, err
	}
	defer rows.Close()

	list := []entities.Employee{}
	for rows.Next() {
		var record entities.Employee
		if errScan := rows.Scan(&record.ID, &record.Name, &record.Salary); errScan != nil {
			log.Println(errScan.Error())
			return nil, err
		}
		list = append(list, record)
	}

	return list, nil
}

func EmployeeFindById(id *uint64) (*entities.Employee, error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var record entities.Employee
	if err = db.QueryRow(`SELECT e.id, e.name, e.salary FROM employee e WHERE e.id = $1`, id).Scan(&record.ID, &record.Name, &record.Salary); err != nil {
		if err == sql.ErrNoRows {
			return &record, nil
		}
		return nil, err
	}

	return &record, nil
}

func EmployeeSave(model *entities.Employee) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.QueryRow(`INSERT INTO employee(name, salary) VALUES($1, $2) RETURNING id`, model.Name, model.Salary).Scan(&model.ID)
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	return nil
}
