package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var i = 1
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if i >= 2 {
			log.Fatal("Nope")
		}
		fmt.Fprintln(w, i)
		i++
	})
	log.Fatal(http.ListenAndServe(listenAddress(), nil))
}

func listenAddress() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	fmt.Println("Listening on port", port)
	return ":" + port

}
