package server

import (
	"crud-gorm/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID    uint32 `json:"id" gorm:"primaryKey"`
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

	var userStruct User
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

	db.Create(&userStruct)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", userStruct.ID)))
}

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	db, errDb := db.Connect()
	if errDb != nil {
		fmt.Println(errDb.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ocorreu um erro ao conectar com o banco"))
		return
	}

	var users []User
	db.Find(&users)

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

	var u User
	db.First(&u, paramId)
	if u.ID == 0 {
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

	var userStruct User
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

	var dbUser User
	db.First(&dbUser, paramId)
	if dbUser.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	db.Model(dbUser).Updates(userStruct)
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

	db.Delete(&User{}, paramId)
	w.WriteHeader(http.StatusNoContent)
}
