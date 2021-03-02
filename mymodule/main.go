package prtinthelloname

import (
	"fmt"
	//_ "github.com/valyala/fasthttp"
	_ "github.com/gorilla/websocket"
)

// PrintJesonHello just prints hello to anyone, 
// for simple test importing in golang
func PrintJesonHello(name string) {
	fmt.Printf("Hello dear %s", name)
}