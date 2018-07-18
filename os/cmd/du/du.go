package du

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// PathStatus for du status
type PathStatus struct {
	Name    string `json:"name"`
	Size    uint64 `json:"size"`
	PathNum uint64 `json:"pathnum"`
	FileNum uint64 `json:"filenum"`
}

func (p PathStatus) String() string {
	buf, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	return string(buf)
}

// Result of Du()
var Result PathStatus

// Du stat disk usage like linux du cmd
func Du(paths []string) {
	for _, path := range paths {
		Result.Name = path
		filepath.Walk(path, walkFn)
	}
	fmt.Print(Result)
}

func walkFn(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		Result.PathNum++
	} else {
		Result.Size += uint64(info.Size())
		Result.FileNum++
	}
	return nil
}
