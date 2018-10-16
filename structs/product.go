package structs

type Product struct {
	Id int `json:"id"`
	Subclass_id int `json:"subclass_id"`
	Price float64 `json:"price,omitempty"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	Url string `json:"url"`
}

