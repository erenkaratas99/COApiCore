package models

// models.Address model info
// @Description a nested struct for a field of entities.Customer and entities.Order
type Address struct {
	AddressOf   string `json:"address_of"bson:"address_of"`
	AddressLine string `json:"address_line"bson:"address_line"`
	City        string `json:"city"`
	Country     string `json:"country"`
	CityCode    int    `json:"cityCode"`
}
