package logutils

import "fmt"

func Print(msg string) {
	fmt.Println(msg)
}

//TODO: Log to a file
func NiceLog(context string) {
	fmt.Println(context)
}
