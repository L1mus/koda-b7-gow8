package internal

import (
	"errors"
	"fmt"
)


func ProcessNumber(numbers []int) (int,error) {
	if numbers == nil {
		err := errors.New("Not data provided")
		return 0,err
	}
	if len(numbers) == 0 {
		panic("empty list provided")
	}
	
	for idx,v := range numbers {
		fmt.Printf("Angka ke-%d: %d", idx,v)
	}
	return 0 , nil
}