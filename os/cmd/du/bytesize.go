package du

import (
	"fmt"
)

// refer from: https://github.com/cydev/du/blob/master/du.go
// but define my own String() func

// ByteSize is formatting wrapper for bytes as human-readable values.
type ByteSize float64

// ByteSize values from kilobyte to yottabyte.
const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func (b ByteSize) String() string {
	switch {
	case b <= 10*MB:
		return fmt.Sprintf("%.2fKB", b/KB)
	case b <= 10*GB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b <= 10*TB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b <= 10*PB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b <= 10*EB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b <= 10*ZB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b <= 10*YB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	}
	return fmt.Sprintf("%.2fYB", b/YB)
}
