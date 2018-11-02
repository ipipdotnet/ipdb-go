package ipdb

import (
	"encoding/json"
	"reflect"
)

type CityInfo struct {
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

func (info CityInfo) ToJson() []byte {
	all, err := json.Marshal(info)
	if err == nil {
		return all
	}

	return nil
}

type City struct {
	reader *Reader
}

func NewCity(name string) (*City, error) {

	r, e := New(name, &CityInfo{})
	if e != nil {
		return nil, e
	}

	return &City{
		reader: r,
	}, nil
}

func (db *City) Find(addr, language string) ([]string, error) {
	return db.reader.find1(addr, language)
}

func (db *City) FindMap(addr, language string) (map[string]string, error) {

	data, err := db.reader.find1(addr, language)
	if err != nil {
		return nil, err
	}
	info := make(map[string]string, len(db.reader.meta.Fields))
	for k, v := range data {
		info[db.reader.meta.Fields[k]] = v
	}

	return info, nil
}

func (db *City) FindInfo(addr, language string) (*CityInfo, error) {

	data, err := db.reader.FindMap(addr, language)
	if err != nil {
		return nil, err
	}

	info := &CityInfo{}

	for k, v := range data {
		sv := reflect.ValueOf(info).Elem()
		sfv := sv.FieldByName(db.reader.refType[k])

		if !sfv.IsValid() {
			continue
		}
		if !sfv.CanSet() {
			continue
		}

		sft := sfv.Type()
		fv := reflect.ValueOf(v)
		if sft == fv.Type() {
			sfv.Set(fv)
		}
	}

	return info, nil
}