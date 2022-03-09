package models

type Colaborador struct {
	Id uint64 `json:"id"`
	Nome string `json:"nome"`
	Salario float64 `json:"salario"`
}