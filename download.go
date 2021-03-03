package ipdb

import "net/url"

//
type Download struct {
	URL *url.URL
}

func (dl *Download) SaveToFile(fn string) error {
	return nil
}

func NewDownload(httpUrl string) (*Download, error) {
	v, e := url.Parse(httpUrl)
	if e != nil {
		return nil, e
	}

	return &Download{
		URL: v,
	}, nil
}