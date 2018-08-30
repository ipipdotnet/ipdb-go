package ipdb

import (
	"fmt"
	"testing"
)

var db *Reader

func init() {
	var err error
	db, err = New("C:/WORK/tiantexin/bb/v6/mydata6vipday4.ipdb")
	fmt.Println(err)

	fmt.Println(db.Find("fe80::", "CN"))
}

func BenchmarkIPDatabase_Query(b *testing.B) {
	for i := 0; i < b.N; i++ {
		db.Find("2001:da8:201::", "CN")
	}
}

func BenchmarkIPDatabase_QueryMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		db.FindMap("2001:da8:201::", "CN")
	}
}