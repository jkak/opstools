package dusub2

import "fmt"

// PathStatus for du status
type PathStatus struct {
	Name    string  `json:"name"`
	Unit    string  `json:"unit"`
	Size    float64 `json:"size"`
	PathNum uint64  `json:"pathnum"`
	FileNum uint64  `json:"filenum"`
}

func (p PathStatus) String() string {
	retStr := ""
	retStr += fmt.Sprintf("{\"name\":\"%v\",\"size\":%.2f,", p.Name, p.Size)
	retStr += fmt.Sprintf("\"unit\":\"%v\",\"pathnum\":%v,\"filenum\":%v}", p.Unit, p.PathNum, p.FileNum)
	return retStr
}

func (p *PathStatus) updateUnit() {
	switch {
	case p.Size < KB:
		p.Unit = "B"
	case p.Size < MB:
		p.Unit = "KB"
		p.Size /= KB
	case p.Size < GB:
		p.Unit = "MB"
		p.Size /= MB
	case p.Size < TB:
		p.Unit = "GB"
		p.Size /= GB
	default:
		p.Unit = "TB"
		p.Size /= TB
	}
}
