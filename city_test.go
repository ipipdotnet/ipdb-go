package ipdb

import (
	"io/ioutil"
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
}

func TestNewCityInMemory(t *testing.T) {
	data, err := ioutil.ReadFile("city.free.ipdb")
	if err != nil {
		t.Log(err)
	}

	db, err := NewCityInMemory(data)
	if err != nil {
		t.Log(err)
	}

	t.Log(db.BuildTime())

	res, err := db.Find("118.28.1.1", "CN")
	if err != nil {
		t.Log(err)
	}

	t.Log(res)
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
