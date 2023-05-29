package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RawJokeData struct {
	Categories  []string `json:"categories"`
	CreatedAt   string   `json:"created_at"`
	IconURL     string   `json:"icon_url"`
	ID          string   `json:"id"`
	UpdatedAt   string   `json:"updated_at"`
	URL         string   `json:"url"`
	Value       string   `json:"value"`
}

func PrintJoke() {
	data := RawJokeData{}
	url := "https://api.chucknorris.io/jokes/random"
	resp, err := http.Get(url)

	if err != nil {
		panic("Error fetching from api")
	}
	
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	
	er := json.Unmarshal(body, &data)
	if er != nil {
		panic("Error reading response body")
	}
	fmt.Println(data.Categories)
	fmt.Println(data.Value)
}