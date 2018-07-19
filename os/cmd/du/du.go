package du

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"syscall"
)

// PathStatus for du status
type PathStatus struct {
	Name    string   `json:"name"`
	Size    ByteSize `json:"size"`
	PathNum uint64   `json:"pathnum"`
	FileNum uint64   `json:"filenum"`
}

func (p PathStatus) String() string {
	retStr := ""
	retStr += fmt.Sprintf("{\"name\":\"%v\",\"size\":\"%s\",", p.Name, p.Size)
	retStr += fmt.Sprintf("\"pathnum\":%v,\"filenum\":%v}", p.PathNum, p.FileNum)
	return retStr
}

// Result of Du()
var Result PathStatus

// BlockSize for fs block size
var BlockSize uint64

// update fs block size to BlkSize
func getBlkSize(p string) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(p, &fs)
	if err != nil {
		log.Fatal("get block size error with:", err)
	}
	BlockSize = uint64(fs.Bsize)
}

// Du stat disk usage like linux du cmd
func Du(paths []string) {
	for _, path := range paths {
		// check whether path is a normal dir or file.
		fd, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		fd.Close()

		getBlkSize(paths[0])
		Result = PathStatus{}
		Result.Name = path
		filepath.Walk(path, walkFn)
		fmt.Printf("%+v\n", Result)
	}
}

func walkFn(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		Result.PathNum++
	} else {
		if info.Size() < int64(BlockSize) {
			Result.Size += ByteSize(BlockSize)
		} else {
			Result.Size += ByteSize(info.Size())
		}
		Result.FileNum++
	}
	return nil
}
