package logger

import (
	"fmt"
	"os"
)

func Fatal(err error) {
	fmt.Println("FATAL\t", err)
	os.Exit(1)
}
