package ipdb

import "encoding/json"

type IPInfo struct {
	CountryName	string	`json:"country_name"`
	RegionName string 	`json:"region_name"`
	CityName string 	`json:"city_name"`
	OwnerDomain string 	`json:"owner_domain"`
	IspDomain string 	`json:"isp_domain"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	Timezone string 	`json:"timezone"`
	UtcOffset string 	`json:"utc_offset"`
	ChinaAdminCode string	`json:"china_admin_code"`
	IddCode string 			`json:"idd_code"`
	CountryCode string 		`json:"country_code"`
	ContinentCode string 	`json:"continent_code"`
	IDC string 				`json:"idc"`
	BaseStation string 		`json:"base_station"`
	CountryCode3 string 	`json:"country_code3"`
	EuropeanUnion string 	`json:"european_union"`
	CurrencyCode string 	`json:"currency_code"`
	CurrencyName string 	`json:"currency_name"`
	Anycast string 			`json:"anycast"`
}

func (info IPInfo) ToJson() []byte {
	all, err := json.Marshal(info)
	if err == nil {
		return all
	}

	return nil
}