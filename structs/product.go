package structs

type Product struct {
	Id int `json:"id"`
	Subclass_id int `json:"subclass_id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

