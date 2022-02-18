package controllers

import (
	"encoding/json"
	"io/ioutil"
	"modulo-escolar/src/core/utils"
	"modulo-escolar/src/domain/entities"
	usecases "modulo-escolar/src/domain/usecases/student"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func FindAllStudents(w http.ResponseWriter, r *http.Request) {
	name := strings.ToLower(r.URL.Query().Get("name"))

	list, err := usecases.GetAllStudentsUsecase(&name)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.JsonResponse(w, http.StatusOK, list)
}

func FindStudentById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	paramId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	model, erro := usecases.GetStudentByIdUsecase(&paramId)
	if erro != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	if model.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	utils.JsonResponse(w, http.StatusOK, model)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	var model entities.Student
	if err = json.Unmarshal(body, &model); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = model.Prepare(); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = usecases.CreateStudentUsecase(&model); err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.JsonResponse(w, http.StatusCreated, model)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	paramId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	var model entities.Student
	if err = json.Unmarshal(body, &model); err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = usecases.UpdateStudentUsecase(&paramId, &model); err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	paramId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = usecases.DeleteStudentUsecase(&paramId); err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
