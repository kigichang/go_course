package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	type Movie struct {
		Title  string
		Year   int  `json:"released"`
		Color  bool `json:"color,omitempty"`
		Actors []string
	}

	var movies = []Movie{
		{
			Title:  "Casablanca",
			Year:   1942,
			Color:  false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"},
		},
		{
			Title:  "Cool Hand Luke",
			Year:   1967,
			Color:  true,
			Actors: []string{"Paul Newman"},
		},
		{
			Title:  "Bullitt",
			Year:   1968,
			Color:  true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"},
		},
	}

	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	fmt.Println("movies: ", string(data))

	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"

	var movie1 Movie
	err = json.Unmarshal([]byte(
		`{
			"Title": "Casablanca",
			"released":1942,
			"Actors":["Humphrey Bogart","Ingrid Bergman"]}
		`), &movie1)

	if err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}

	fmt.Println("movie1", movie1)

}
