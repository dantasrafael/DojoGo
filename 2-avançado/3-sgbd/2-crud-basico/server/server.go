package server

import (
	"crud-basico/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ocorreu um erro ao ler o corpo da requisição"))
		return
	}

	var userStruct user
	if errUnm := json.Unmarshal(requestBody, &userStruct); errUnm != nil {
		fmt.Println(errUnm.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ocorreu um erro ao converter um usuário para struct"))
		return
	}

	db, errDb := db.Connect()
	if errDb != nil {
		fmt.Println(errDb.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ocorreu um erro ao conectar com o banco"))
		return
	}
	defer db.Close()

	// stmt, errStt := db.Prepare(`INSERT INTO users(name, email) VALUES($1, $2) RETURNING id`)
	// if errStt != nil {
	// 	fmt.Println(errStt.Error())
	// 	w.Write([]byte("Ocorreu um erro ao criar o statement"))
	// 	return
	// }
	// defer stmt.Close()

	// res, errExec := stmt.Exec(userStruct.Name, userStruct.Email)
	// if errExec != nil {
	// 	fmt.Println(errExec.Error())
	// 	w.Write([]byte("Ocorreu um erro ao executar o statement"))
	// 	return
	// }

	// fmt.Println("Res:", res)
	// fmt.Println(res.LastInsertId()) // Não suportado por este driver
	// fmt.Println(res.RowsAffected())

	var userId int
	errQuery := db.QueryRow(`INSERT INTO users(name, email) VALUES($1, $2) RETURNING id`, userStruct.Name, userStruct.Email).Scan(&userId)
	if errQuery != nil {
		fmt.Println(errQuery.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ocorreu um erro ao executar a inserção"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", userId)))
}

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	db, errDb := db.Connect()
	if errDb != nil {
		fmt.Println(errDb.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ocorreu um erro ao conectar com o banco"))
		return
	}
	defer db.Close()

	rows, errQuery := db.Query(`SELECT * FROM users`)
	if errQuery != nil {
		fmt.Println(errQuery.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ocorreu um erro ao executar a consulta"))
		return
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var u user
		if errScan := rows.Scan(&u.ID, &u.Name, &u.Email); errScan != nil {
			fmt.Println(errScan.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Ocorreu um erro ao scanear o usuário"))
			return
		}
		users = append(users, u)
	}

	w.WriteHeader(http.StatusOK)
	if erroEcd := json.NewEncoder(w).Encode(users); erroEcd != nil {
		fmt.Println(erroEcd.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ocorreu um erro ao converter os usuários para JSON"))
		return
	}
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	paramId, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ocorreu um erro ao converter um id para inteiro"))
		return
	}

	db, errDb := db.Connect()
	if errDb != nil {
		fmt.Println(errDb.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ocorreu um erro ao conectar com o banco"))
		return
	}
	defer db.Close()

	var u user
	errQuery := db.QueryRow(`SELECT * FROM users WHERE id = $1`, paramId).Scan(&u.ID, &u.Name, &u.Email)
	if errQuery != nil {
		fmt.Println(errQuery.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	if erroEcd := json.NewEncoder(w).Encode(u); erroEcd != nil {
		fmt.Println(erroEcd.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ocorreu um erro ao converter o usuário para JSON"))
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	paramId, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ocorreu um erro ao converter um id para inteiro"))
		return
	}

	requestBody, errReadBody := ioutil.ReadAll(r.Body)
	if errReadBody != nil {
		fmt.Println(errReadBody.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ocorreu um erro ao ler o corpo da requisição"))
		return
	}

	var userStruct user
	if errUnm := json.Unmarshal(requestBody, &userStruct); errUnm != nil {
		fmt.Println(errUnm.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ocorreu um erro ao converter um usuário para struct"))
		return
	}

	db, errDb := db.Connect()
	if errDb != nil {
		fmt.Println(errDb.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ocorreu um erro ao conectar com o banco"))
		return
	}
	defer db.Close()

	stmt, errStt := db.Prepare(`UPDATE users SET name = $1, email = $2 WHERE id = $3`)
	if errStt != nil {
		fmt.Println(errStt.Error())
		w.Write([]byte("Ocorreu um erro ao criar o statement"))
		return
	}
	defer stmt.Close()

	if _, errExec := stmt.Exec(userStruct.Name, userStruct.Email, paramId); errExec != nil {
		fmt.Println(errExec.Error())
		w.Write([]byte("Ocorreu um erro ao executar o statement"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	paramId, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ocorreu um erro ao converter um id para inteiro"))
		return
	}

	db, errDb := db.Connect()
	if errDb != nil {
		fmt.Println(errDb.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ocorreu um erro ao conectar com o banco"))
		return
	}
	defer db.Close()

	stmt, errStt := db.Prepare(`DELETE FROM users WHERE id = $1`)
	if errStt != nil {
		fmt.Println(errStt.Error())
		w.Write([]byte("Ocorreu um erro ao criar o statement"))
		return
	}
	defer stmt.Close()

	if _, errExec := stmt.Exec(paramId); errExec != nil {
		fmt.Println(errExec.Error())
		w.Write([]byte("Ocorreu um erro ao executar o statement"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
