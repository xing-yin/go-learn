package main

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func marshal() {
	data, err := json.Marshal(movies) // 紧凑的形式
	//data, err := json.MarshalIndent(movies, "", "") // 整齐缩进的形式
	if err != nil {
		log.Fatalf("JSON marshaling failed:%s", err)
	}
	fmt.Printf("%s\n", data)
}

func unmarshal() {
	//var titles []struct{ Title string }
	//jsonStr := "[{\"Title\":\"Casablanca\",\"released\":1942,\"Actors\":[\"Humphrey Bogart\",\"Ingrid Bergman\"]}]"
	//data := []byte(jsonStr)
	//if err := json.Unmarshal(data, &titles); err != nil {
	//	log.Fatalf("JSON unmashaling failed:%s", err)
	//}
	//fmt.Println(titles)

	var movies2 []Movie
	jsonStr := "[{\"Title\":\"Casablanca\",\"released\":1942,\"Actors\":[\"Humphrey Bogart\",\"Ingrid Bergman\"]}]"
	data := []byte(jsonStr)
	if err := json.Unmarshal(data, &movies2); err != nil {
		log.Fatalf("JSON unmashaling failed:%s", err)
	}
	fmt.Println(movies2)
}

func TestMarshal(t *testing.T) {
	//marshal()

	unmarshal()
}
