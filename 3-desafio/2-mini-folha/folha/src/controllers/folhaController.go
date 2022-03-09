package controllers

import (
	"modulo-folha/src/controllers/dto"
	"modulo-folha/src/repositories"
	"net/http"
)

func ListarFolhas(w http.ResponseWriter, r *http.Request) {
	list, err := repositories.BuscarTodasFolhas()
	if err != nil {
		dto.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	dto.JsonResponse(w, http.StatusOK, list)
}
