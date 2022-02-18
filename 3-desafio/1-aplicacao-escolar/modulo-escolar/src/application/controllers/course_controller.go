package controllers

import (
	"encoding/json"
	"io/ioutil"
	"modulo-escolar/src/core/utils"
	"modulo-escolar/src/domain/entities"
	usecases "modulo-escolar/src/domain/usecases/course"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func FindAllCourses(w http.ResponseWriter, r *http.Request) {
	name := strings.ToLower(r.URL.Query().Get("name"))

	list, err := usecases.GetAllCoursesUsecase(&name)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.JsonResponse(w, http.StatusOK, list)
}

func FindCourseById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	paramId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	record, erro := usecases.GetCourseByIdUsecase(&paramId)
	if erro != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	if record.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	utils.JsonResponse(w, http.StatusOK, record)
}

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	var model entities.Course
	if err = json.Unmarshal(body, &model); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = model.Prepare(); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = usecases.CreateCourseUsecase(&model); err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.JsonResponse(w, http.StatusCreated, model)
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
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

	var model entities.Course
	if err = json.Unmarshal(body, &model); err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = usecases.UpdateCourseUsecase(&paramId, &model); err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	paramId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = usecases.DeleteCourseUsecase(&paramId); err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
