package main

import (
	"net/http"
 //   "encoding/json"
    "log"
    "github.com/gorilla/mux"
 	"github.com/hendrikhoffmann/smartfblib"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/switchon", smartfblib.SwitchOn).Methods("GET")
	router.HandleFunc("/switchoff", smartfblib.SwitchOff).Methods("GET")
    router.HandleFunc("/switchchannel/{channel}", smartfblib.SwitchChannel).Methods("GET")
    router.HandleFunc("/setvolume/{volume}", smartfblib.SetVolume).Methods("GET")
    log.Fatal(http.ListenAndServe(":3000", router))

}

//import	"golang.org/x/net/websocket"
//Websocket Communication Postponed
//	http.Handle("/switchon", websocket.Handler(func(ws *websocket.Conn) {
//		ws.Write([]byte(smartfblib.SwitchOn()))
//	}))
//	http.ListenAndServe(":3000", nil)
