package main

import (
	"net/http"

	"golang.org/x/net/websocket"
	"github.com/hendrikhoffmann/smartfblib"
)

func main() {
	http.Handle("/switchon", websocket.Handler(func(ws *websocket.Conn) {
		ws.Write([]byte(smartfblib.SwitchOn()))
	}))
	http.ListenAndServe(":3000", nil)
}