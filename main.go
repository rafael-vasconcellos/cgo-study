package main

import (
	"cgo_study/dlls"
	"cgo_study/message"
	"fmt"
)

func main() {
	message.Ping()
	message.SayHello("Rafael")
	title, err := dlls.GetActiveWindowTitle()
	if err == nil { fmt.Println(title) }
}