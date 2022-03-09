package models

type Folha struct {
	Id uint64 `json:"id"`
	ColaboradorId uint64 `json:"colaborador_id"`
	ColaboradorNome string `json:"colaborador_nome"`
	ColaboradorSalario float64 `json:"colaborador_salario"`
	Inss float64 `json:"inss"`
	Irpf float64 `json:"irpf"`
	Total float64 `json:"total"`
}