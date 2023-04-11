/*
One way to convert binary to int in Golang is to use the 
strconv.ParseInt function with the base argument set to 2. 
This will take a binary string as input and return an int64 value as output. 

Another way is to use the encoding/binary package and its ByteOrder interface2. 
You can also write your own function to convert binary to int using bitwise operations
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	number := "1001"
	dec_number, err := strconv.ParseInt(number, 2, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dec_number)
	}
}