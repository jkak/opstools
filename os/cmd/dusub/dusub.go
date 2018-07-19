package dusub

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/jkak/opstools/os/cmd/du"
)

// PathStatus for du status
type PathStatus struct {
	Path    string
	Name    string
	PathNum uint64
	FileNum uint64
	Size    du.ByteSize
}

func (p PathStatus) String() string {
	retStr := ""
	retStr += fmt.Sprintf("{\"name\":\"%v\",\"size\":%s,\"path\":\"%s\"", p.Name, p.Size, p.Path)
	retStr += fmt.Sprintf("\"pathnum\":%v,\"filenum\":%v}", p.PathNum, p.FileNum)
	return retStr
}

// Result for Dusub()
var (
	Result    []PathStatus
	ignores   []string
	blockSize uint64
	pointer   int64
)

// Dusub stat sub directory disk usage like linux du cmd
func Dusub(path, ign string) {
	fd, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()
	infos, err := fd.Readdir(0)
	if err != nil {
		// fmt.Printf("err when Readdir():%v\n", err)
		return
	}
	getBlkSize(path)
	ignores = strings.Split(ign, ",")

	processSub(path, infos)
	for _, r := range Result {
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
		absPath := filepath.Join(path, sub.Name())
		r := PathStatus{Name: sub.Name(), Path: absPath}
		Result = append(Result, r)
		pointer++

		filepath.Walk(absPath, walkFn)
	}
}

// update fs block size to BlkSize
func getBlkSize(p string) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(p, &fs)
	if err != nil {
		log.Fatal("get block size error with:", err)
	}
	blockSize = uint64(fs.Bsize)
}

func walkFn(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		Result[pointer].PathNum++
	} else {
		if info.Size() < int64(blockSize) {
			Result[pointer].Size += du.ByteSize(blockSize)
		} else {
			Result[pointer].Size += du.ByteSize(info.Size())
		}
		Result[pointer].FileNum++
	}
	return nil
}
