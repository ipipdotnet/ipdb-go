package ipdb

import "testing"

func TestNewRisk(t *testing.T) {
	r, e := NewRisk("c:/work/ipdb/v6risk.ipdb")
	if e != nil {
		t.Fatal(e)
	}
	info, e := r.FindInfo("2001:240:2a3b:700::")
	if e == nil {
		t.Log(info)
	}

	_, e = r.FindInfo("2001:240:2a3b:7100::")
	if e != nil {
		t.Log(e)
	}
}
