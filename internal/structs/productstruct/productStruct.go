package productstruct

type Product struct {
	Model   string `json:"model"`
	Company string `json:"company"`
	Price   int    `json:"price"`
	Id      int    `json:"id"`
}
