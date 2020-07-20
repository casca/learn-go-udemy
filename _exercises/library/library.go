package library

import (
	"fmt"
	"runtime"
)

func PrintVer() {
	fmt.Println(runtime.Version())
}
