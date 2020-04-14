package ipdb

import (
	"testing"
)

var db *City

func init() {
	db, _ = NewCity("city.free.ipdb")
}

func TestNewCity(t *testing.T) {
	db, err := NewCity("city.free.ipdb")
	if err != nil {
		t.Log(err)
	}

	t.Log(db.BuildTime())

	t.Log(db.Fields())

	loc, err := db.Find("1.1.1.1", "CN")
	if err != nil {
		t.Log(err)
	} else {
		t.Log(loc)
	}

	m, err := db.FindMap("27.190.250.164", "CN")
	if err == nil {
		for k, v := range m {
			t.Log(k, v)
		}
	}

	info1, err := db.FindInfo("1.1.1.1", "CN")
	if err == nil {
		for _, item := range info1.ASNInfo {
			t.Log(item.ASN, item.Registry, item.Country, item.Net, item.Org)
		}
	}

	info, err := db.FindInfo("27.190.250.164", "CN")
	if err != nil {
		t.Log(err)
	} else {
		t.Log(info.Route)
		t.Log(info.ASN)
		t.Log(info.DistrictInfo)
		t.Log(info.ASNInfo)

		for _, af := range info.ASNInfo {
			t.Log(af)
		}
	}
}

func BenchmarkCity_Find(b *testing.B) {

	for i := 0; i < b.N; i++ {
		db.Find("118.28.1.1", "CN")
	}
}

func BenchmarkCity_FindMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		db.FindMap("118.28.1.1", "CN")
	}
}

func BenchmarkCity_FindInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		db.FindInfo("118.28.1.1", "CN")
	}
}
