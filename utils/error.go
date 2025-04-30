package utils

import (
	"fmt"
	"runtime"
)

func TraceError(err error) {
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			fmt.Printf("Error: %v\nOccurred at: %s:%d\n", err, file, line)
		}
	}
}
