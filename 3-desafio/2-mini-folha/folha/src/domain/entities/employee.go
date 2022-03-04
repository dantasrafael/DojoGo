package entities

type Employee struct {
	ID        uint16    `json:"id"`
	Name      string    `json:"name"`
	Salary     float32   `json:"salary"`
}
