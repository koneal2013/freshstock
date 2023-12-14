package model

type Produce struct {
	Code      string  `json:"code" binding:"required,alphanum,len=19"`
	Name      string  `json:"name" binding:"required,alphanum"`
	UnitPrice float64 `json:"unit_price" binding:"required,gte=0"`
}
