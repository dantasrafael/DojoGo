package repositories

import (
	"encoding/json"
	"fmt"
	"modulo-folha/src/models"

	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/db"
)

func Salvar(message string) error {
	fmt.Println("Salvando o colaborador: " + message)

	var model models.Colaborador
	if err := json.Unmarshal([]byte(message), &model); err != nil {
		return err
	}

	db, err := db.Connect()
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO colaborador (id, nome, salario) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(model.Id, model.Nome, model.Salario); err != nil {
		return err
	}

	return nil
}

func Excluir(message string) error {
	fmt.Println("Excluindo o colaborador: " + message)

	var model models.Colaborador
	if err := json.Unmarshal([]byte(message), &model); err != nil {
		return err
	}

	db, err := db.Connect()
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("DELETE FROM colaborador WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(model.Id); err != nil {
		return err
	}

	return nil
}
