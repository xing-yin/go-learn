package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(PlayerServer)
	err := http.ListenAndServe(":5000", handler)
	if err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}

func PlayerServer(w http.ResponseWriter, req *http.Request) {
	player := req.URL.Path[len("/players/"):]

	fmt.Fprintf(w, GetPlayerScore(player))
}

func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}

	if name == "Floyd" {
		return "10"
	}

	return ""
}
