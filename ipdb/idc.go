package ipdb

import (
	"encoding/json"
	"reflect"
)

type IDCInfo struct {
	CountryName	string	`json:"country_name"`
	RegionName string 	`json:"region_name"`
	CityName string 	`json:"city_name"`
	OwnerDomain string 	`json:"owner_domain"`
	IspDomain string 	`json:"isp_domain"`
	IDC string 			`json:"idc"`
}

func (info IDCInfo) ToJson() []byte {
	all, err := json.Marshal(info)
	if err == nil {
		return all
	}

	return nil
}

type IDC struct {
	reader *Reader
}

func NewIDC(name string) (*IDC, error) {

	r, e := New(name, &IDCInfo{})
	if e != nil {
		return nil, e
	}

	return &IDC{
		reader: r,
	}, nil
}

func (db *IDC) Find(addr, language string) ([]string, error) {
	return db.reader.find1(addr, language)
}

func (db *IDC) FindMap(addr, language string) (map[string]string, error) {

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

func (db *IDC) FindInfo(addr, language string) (*IDCInfo, error) {

	data, err := db.reader.FindMap(addr, language)
	if err != nil {
		return nil, err
	}

	info := &IDCInfo{}

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