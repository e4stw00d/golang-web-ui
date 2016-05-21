package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"html/template"
	"net/http"
	"os"
)

var opened bool = false

func response(ws *websocket.Conn) {

	if opened {
		return
	}

	opened = true

	for {
		receivedtext := make([]byte, 100)

		n, err := ws.Read(receivedtext)

		if n == 0 {
			fmt.Println("Conn closed!")
			os.Exit(0)
		}

		if err != nil {
			fmt.Printf("Received: %d bytes\n", n)
		}

		s := string(receivedtext[:n])
		fmt.Printf("Received: %d bytes: %s\n", n, s)

		str := "Hello!"
		ws.Write([]byte(str))
		fmt.Printf("Sent: %s\n", str)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("main.html")
		if err != nil {
			println(err.Error())
		}
		t.Execute(w, nil)
	})
	http.Handle("/ws", websocket.Handler(response))
	http.ListenAndServe(":8080", nil)
}
