package main

import (
	"golang.org/x/net/websocket"
	"io"
	"log"
	"net/http"
)

func handleWS(ws *websocket.Conn) {
	log.SetOutput(io.MultiWriter(log.Writer(), ws))
	var ch = make(chan string)
	<-ch
}

func init() {
	http.Handle("/ws", websocket.Server{
		Handler: handleWS,
	})
	go func() {
		err := http.ListenAndServe("localhost:18964", nil)
		if err != nil {
			log.Println("could not set up websocket server: ", err)
		}
	}()
}
