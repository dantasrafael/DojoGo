package controller

import (
	"encoding/json"
	"net/http"
)

type InvoiceController struct {
}

func NewInvoiceController() *InvoiceController {
	return &InvoiceController{}
}

func (c InvoiceController) Find(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode("{\"msg\":\"Hello\"}")
}
