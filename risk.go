package ipdb

import "strconv"

type RiskInfo struct {
	Score       int
	Behavior    string
	CountryCode string
}

type Risk struct {
	reader *reader
}

func NewRisk(fn string) (*Risk, error) {
	r, e := newReader(fn, &RiskInfo{})
	if e != nil {
		return nil, e
	}
	return &Risk{reader: r}, nil
}

func (r *Risk) FindInfo(addr string) (*RiskInfo, error) {
	info := &RiskInfo{}

	m, e := r.reader.FindMap(addr, "CN")
	if e != nil {
		return info, e
	}

	if v, ok := m["score"]; ok {
		info.Score, _ = strconv.Atoi(v)
	}
	if v, ok := m["behavior"]; ok {
		info.Behavior = v
	}
	if v, ok := m["country_code"]; ok {
		info.CountryCode = v
	}

	return info, nil
}
