package ipdb

import "testing"

var db *City

func init() {
	db, _ = NewCity("c:/work/ipdb/city.free.ipdb")
}

func TestNewCity(t *testing.T) {
	db, err := NewCity("c:/work/ipdb/city.free.ipdb")
	if err != nil {
		t.Log(err)
	}

	t.Log(db.BuildTime())
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