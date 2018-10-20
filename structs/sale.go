package structs

type Sale struct {
	Id         int `json:"id"`
	ProductId  int `json:"product_id"`
	ShopId     int `json:"shop_id"`
	Price      float64 `json:"price"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	CreateTime int `json:"create_time"`
}
