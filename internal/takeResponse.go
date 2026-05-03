package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Character struct {
	ID       int        `json:"id"`
	Name     string     `json:"name"`
	Status   string     `json:"status"`
	Species  string     `json:"species"`
	Type     string     `json:"type"`
	Gender   string     `json:"gender"`
	Origin   Location   `json:"origin"`
	Location Location   `json:"location"`
	Image    string     `json:"image"`
	Episode  []string   `json:"episode"`
	URL      string     `json:"url"`
	Created  string     `json:"created"` 
}


func WebFetcher(url string , chFetchData chan *Character) error{
	res,err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ",err)
	}
	defer res.Body.Close()

	// body,err := io.ReadAll(res.Body)
	// if err !=nil {
	// 	fmt.Println("Error : ", err)
	// }

	var character Character
	// err = json.Unmarshal(body,&character)
	// if err != nil{
	// 	println("Error ", err)
	// }
	
	err = json.NewDecoder(res.Body).Decode(&character)
	if err != nil{
		println("Error ", err)
	}

	chFetchData <- &character
	return nil
}

func DisplayResultFetchData(chFetchData chan *Character,exit chan bool) {
	defer func(){
		exit <- true
	}()
	for v := range chFetchData {
		fmt.Printf("%+v\n", v)
	}
}