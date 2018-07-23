package dusub2

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

// KB _
const (
	_          = iota
	KB float64 = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)

// Result for Dusub()
var (
	Result    PathStatus
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

	for _, sub := range infos {
		processSub(path, sub)
	}
}

func processSub(path string, sub os.FileInfo) {
	// check to ignore the sub path
	for _, i := range ignores {
		if sub.Name() == i {
			return
		}
	}
	Result.Name = sub.Name()
	filepath.Walk(filepath.Join(path, sub.Name()), walkFn)
	Result.updateUnit()
	fmt.Println(Result)
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
		Result.PathNum++
	} else {
		if info.Size() < int64(blockSize) {
			Result.Size += blockSize
		} else {
			Result.Size += float64(info.Size())
		}
		Result.FileNum++
	}
	return nil
}
