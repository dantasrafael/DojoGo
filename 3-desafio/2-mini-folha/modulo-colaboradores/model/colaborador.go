package model

type Colaborador struct {
	Id      uint64  `json:"id"`
	Nome    string  `json:"nome"`
	Funcao  string  `json:"funcao"`
	Salario float64 `json:"salario"`
}
