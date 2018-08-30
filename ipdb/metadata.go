package ipdb

type MetaData struct {
	Build     int64 			`json:"build"`
	IPVersion uint16 			`json:"ip_version"`
	Languages map[string]int 	`json:"languages"`
	NodeCount int 				`json:"node_count"`
	TotalSize int				`json:"total_size"`
	Fields 	  []string 			`json:"fields"`
}