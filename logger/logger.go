package logger

import (
	"fmt"
	"os"
)

func Info(format string, values ...interface{}) {
	fmt.Printf("INFO\t"+format, values...)
}

func Error(err ...interface{}) {
	fmt.Println("ERROR\t", err)
}

func Fatal(err error) {
	fmt.Println("FATAL\t", err)
	os.Exit(1)
}
