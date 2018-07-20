package dusub2

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

// PathStatus for du status
type PathStatus struct {
	Name    string  `json:"name"`
	Unit    string  `json:"unit"`
	Size    float64 `json:"size"`
	PathNum uint64  `json:"pathnum"`
	FileNum uint64  `json:"filenum"`
}

// KB _
const (
	_          = iota
	KB float64 = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)

func (p PathStatus) String() string {
	retStr := ""
	retStr += fmt.Sprintf("{\"name\":\"%v\",\"size\":\"%.2f\",", p.Name, p.Size)
	retStr += fmt.Sprintf("\"unit\":%v,\"pathnum\":%v,\"filenum\":%v}", p.Unit, p.PathNum, p.FileNum)
	return retStr
}

// Result for Dusub()
var (
	Result    []PathStatus
	ignores   []string
	blockSize float64
	pointer   int64
)

// Dusub2 stat sub directory disk usage like linux du cmd
func Dusub2(path, ign string) {
	fd, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()
	infos, err := fd.Readdir(0)
	if err != nil {
		return
	}

	getBlkSize(path)
	ignores = strings.Split(ign, ",")

	processSub(path, infos)
	for _, r := range Result {
		updateUnit(&r)
		fmt.Println(r)
	}

}

func processSub(path string, infos []os.FileInfo) {
	// init
	Result = make([]PathStatus, 0)
	pointer = -1

OUT_FOR:
	for _, sub := range infos {
		// check to ignore the sub path
		for _, i := range ignores {
			if sub.Name() == i {
				continue OUT_FOR
			}
		}
		r := PathStatus{Name: sub.Name()}
		Result = append(Result, r)
		pointer++

		filepath.Walk(filepath.Join(path, sub.Name()), walkFn)
	}
}

// update fs block size to BlkSize
func getBlkSize(p string) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(p, &fs)
	if err != nil {
		log.Fatal("get block size error with:", err)
	}
	blockSize = float64(fs.Bsize)
}

func walkFn(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		Result[pointer].PathNum++
	} else {
		if info.Size() < int64(blockSize) {
			Result[pointer].Size += blockSize
		} else {
			Result[pointer].Size += float64(info.Size())
		}
		Result[pointer].FileNum++
	}
	return nil
}

func updateUnit(p *PathStatus) {
	switch {
	case p.Size < KB:
		p.Unit = "B"
	case p.Size < 8*MB:
		p.Unit = "KB"
		p.Size /= KB
	case p.Size < 8*GB:
		p.Unit = "MB"
		p.Size /= MB
	case p.Size < 8*TB:
		p.Unit = "GB"
		p.Size /= GB
	default:
		p.Unit = "TB"
		p.Size /= TB
	}
}
