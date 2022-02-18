package controllers

import (
	"encoding/json"
	"io/ioutil"
	"modulo-escolar/src/core/utils"
	"modulo-escolar/src/domain/entities"
	usecases "modulo-escolar/src/domain/usecases/enrollment"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func FindAllEnrollments(w http.ResponseWriter, r *http.Request) {
	studentName := strings.ToLower(r.URL.Query().Get("studentName"))
	courseName := strings.ToLower(r.URL.Query().Get("courseName"))

	list, err := usecases.GetAllEnrollmentsUsecase(&studentName, &courseName)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.JsonResponse(w, http.StatusOK, list)
}

func CreateEnrollment(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	var model entities.Enrollment
	if err = json.Unmarshal(body, &model); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = model.Prepare(); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = usecases.CreateEnrollmentUsecase(&model); err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.JsonResponse(w, http.StatusCreated, model)
}

func DeleteEnrollment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	paramId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = usecases.DeleteEnrollmentUsecase(&paramId); err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
