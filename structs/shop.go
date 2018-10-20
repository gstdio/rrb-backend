package structs

type Shop struct {
	Id       int `json:"id"`
	RegionId int `json:"region_id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Desc     string `json:"desc"`
	Url      string `json:"url"`
}
