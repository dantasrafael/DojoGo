package entities

import (
	"time"
)

type Payroll struct {
	ID        	uint64     	`json:"id"`
	Employee  	Employee  	`json:"employee"`
	Reference 	time.Time 	`json:"reference"`
	SalaryBase  float32 		`json:"salary_base"`
	Discount		float32			`json:"discount"`
	Total				float32			`json:"total"`
	Created_at	time.Time 	`json:"created_at"`
}
