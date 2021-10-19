package logutils

import (
	"fmt"
	"os"
)

func Print(msg string) {
	fmt.Println(msg)
}

//TODO: Log to a file
func NiceLog(context string) {
	fmt.Println(context)
}

func FatalAlert(message string) {
	fmt.Println(message)
	os.Exit(0)
}
