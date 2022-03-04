package repository

import (
	db2 "github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/db"
	"modulo-colaboradores/model"
)

const (
	insertColaborador       = "insert into colaboradores (nome, funcao, salario) values ($1, $2, $3) returning id"
	buscaTodosColaboradores = "select id, nome, funcao, salario from colaboradores"
	deletarColaborador      = "delete from colaboradores where id = $1"
)

func BuscarTodosColaboradores() ([]model.Colaborador, error) {
	db, err := db2.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(buscaTodosColaboradores)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := []model.Colaborador{}

	for rows.Next() {
		var record model.Colaborador
		if err = rows.Scan(&record.Id, &record.Nome, &record.Funcao, &record.Salario); err != nil {
			return nil, err
		}
		list = append(list, record)
	}

	return list, nil
}

func SalvarColaborador(model *model.Colaborador) error {
	db, err := db2.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.QueryRow(insertColaborador, model.Nome, model.Funcao, model.Salario).Scan(&model.Id)
	if err != nil {
		return err
	}

	return nil
}

func DeletarColaborador(id uint64) error {
	db, err := db2.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	st, err := db.Prepare(deletarColaborador)
	if err != nil {
		return err
	}
	defer st.Close()

	_, err = st.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
