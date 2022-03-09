package repositories

import (
	"modulo-folha/src/models"

	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/db"
)

func CalcularFolha() error {
	db, err := db.Connect()
	if err != nil {
		return err
	}

	rows, err := db.Query("SELECT id, nome, salario FROM colaborador")
	if err != nil {
		return err
	}
	defer rows.Close()

	list := []models.Colaborador{}
	for rows.Next() {
		var record models.Colaborador
		if err := rows.Scan(&record.Id, &record.Nome, &record.Salario); err != nil {
			return err
		}
		list = append(list, record)
	}

	stmt, err := db.Prepare("INSERT INTO folha (colaborador_id, colaborador_nome, colaborador_salario, inss, irpf, total) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return err
	}

	for _, col := range list {
		folha := models.Folha{}
		folha.ColaboradorId = col.Id
		folha.ColaboradorNome = col.Nome
		folha.ColaboradorSalario = col.Salario
		folha.Inss = col.Salario * 0.11
		folha.Irpf = col.Salario * 0.05
		folha.Total = col.Salario - (folha.Inss + folha.Irpf)

		if _, err := stmt.Exec(folha.ColaboradorId, folha.ColaboradorNome, folha.ColaboradorSalario, folha.Inss, folha.Irpf, folha.Total); err != nil {
			return err
		}
	}

	return nil
}

func BuscarTodasFolhas() ([]models.Folha, error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT id, colaborador_id, colaborador_nome, colaborador_salario, inss, irpf, total FROM folha")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := []models.Folha{}
	for rows.Next() {
		var record models.Folha
		if err := rows.Scan(&record.Id, &record.ColaboradorId, &record.ColaboradorNome, &record.ColaboradorSalario, &record.Inss, &record.Irpf, &record.Total); err != nil {
			return nil, err
		}
		list = append(list, record)
	}

	return list, nil
}