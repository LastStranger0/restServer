package handler

import (
	"net/http"
)

// GetCompanyAndProducts ...
func GetCompanyAndProducts(w http.ResponseWriter, r *http.Request) {}

// структуры должны находится в другом файле
type companyCrossProductStruct struct {
	Company string
	Model   string
}
