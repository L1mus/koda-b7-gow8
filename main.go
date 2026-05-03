package main

import (
	"fmt"
	"koda-b7-gow8/internal"
	"sync"
)

func main() {

	// take Response
	chFetchData := make(chan *internal.Character)
	exit := make(chan bool)
	var wg sync.WaitGroup

	var urls = []string{"https://rickandmortyapi.com/api/character/1","https://rickandmortyapi.com/api/character/2","https://rickandmortyapi.com/api/character/3","https://rickandmortyapi.com/api/character/4","https://rickandmortyapi.com/api/character/5"}

	for _, url := range urls {
		wg.Add(1)
        go func(u string) {
            defer wg.Done()
            internal.WebFetcher(u, chFetchData)
        }(url)
	}
	
	go func() {
        wg.Wait()         
        close(chFetchData) 
    }()

	go internal.DisplayResultFetchData(chFetchData,exit)
	<-exit
	
	
	// User Management

	management := internal.NewUserManagement()
	
	// adduser case
	management.AddUser(1, "Ali Mustadji")
	management.AddUser(2, "Ilham")
	management.AddUser(3, "Hanif")
	
	//duplicate Case
	management.AddUser(2, "Naufal")

	//getuser case
	user, err := management.GetUser(1)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		fmt.Printf("User found!!!\nId=%d\nName=%s\n \n", user.Id, user.Name)
	}

	// getuser edge case
	user1, err := management.GetUser(10)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		fmt.Printf("User found: Id=%d, Name=%s\n \n", user1.Id, user1.Name)
	}

	//ShapeGeometry

		shapes := []internal.ShapeGeometry{
		internal.Circle{Radius: 12.3},
		internal.Circle{Radius: 5.5},
		internal.Rectangle{Height: 50, Width: 25},
		internal.Rectangle{Height: 36, Width: 72},
	}

	for _, shape := range shapes {
		switch s := shape.(type) {
		case internal.Circle:
			fmt.Printf("Circle radius= %.2f\nArea: %.4f\n \n", s.Radius, s.Area())
		case internal.Rectangle:
			fmt.Printf("Rectangle Height= %.2f, Width= %.2f\nArea: %.4f\n \n", s.Height, s.Width, s.Area())
		}
	}

}