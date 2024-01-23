package repository

type locationAreaBrief struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaList struct {
	Count    int                 `json:"count"`
	Next     *string             `json:"next"`
	Previous *string             `json:"previous"`
	Results  []locationAreaBrief `json:"results"`
}

func NewLocationAreaList() *LocationAreaList {
	return &LocationAreaList{}
}
