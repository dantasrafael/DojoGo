package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"modulo-colaboradores/controller/dto"
	model2 "modulo-colaboradores/model"
	"modulo-colaboradores/producers"
	"modulo-colaboradores/repository"
	"net/http"
	"strconv"
)

func BuscarTodosColaboradores(w http.ResponseWriter, r *http.Request) {
	list, err := repository.BuscarTodosColaboradores()
	if err != nil {
		dto.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	dto.JsonResponse(w, http.StatusOK, list)
}

func SalvarColaborador(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		dto.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	model := model2.Colaborador{}
	err = json.Unmarshal(body, &model)
	if err != nil {
		dto.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = repository.SalvarColaborador(&model)
	if err != nil {
		dto.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	producers.CriacaoColaborador(&model)

	dto.JsonResponse(w, http.StatusCreated, model)
}

func DeletarColaborador(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		dto.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = repository.DeletarColaborador(id)
	if err != nil {
		dto.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	producers.DelecaoColaborador(&model2.Colaborador{Id: id})

	w.WriteHeader(http.StatusNoContent)
}
