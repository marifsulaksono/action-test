package main

import (
	"fmt"
	"net/http"
)

// Fungsi handle Index menggunakan function
// func handlerIndex(w http.ResponseWriter, r *http.Request) {
// 	var message = "Welcome"
// 	w.Write([]byte(message))
// }

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello World"
	w.Write([]byte(message))
}

func main() {
	// ========== Routing, membuat multiplexer sendiri dengan memanfaatkan struct http.servemux
	handlerIndex := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hallo Dunia!"))
	}

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	// routing menggunakan anonymous function
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello again!"))
	})

	var address = "localhost:9000"
	fmt.Printf("Server started at %s\n", address)
	// server := new(http.Server)
	// server.Addr = address
	err := http.ListenAndServe(address, nil)
	// err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
