package internal

import (
	"fmt"
	"io"
	"net/http"
)

func WebFetcher(url string) ([]byte,error){
	res,err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ",err)
	}
	defer res.Body.Close()

	body,err := io.ReadAll(res.Body)
	if err !=nil {
		fmt.Println("Error : ", err)
	}

	return body,err
}