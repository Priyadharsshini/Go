package main

import (
	"encoding/json"
	"fmt"
)

/*
The field names must remain in caps so that it is exported and only exported values will be marshalled and unmarshalled in JSON.
Marshal function only considers the exported fields for the encoding.
The second value to Unmarshal can either be a pointer or an interface
*/

type Address struct {
	Street string
	City   string
}

func main() {
	var add Address
	
	// Any additional fields not declared in the struct will not be unmarshalled 
	
	data := []byte(`
		{
			"Street" : "teststreet",
			"City" : "testcity",
			"Pincode" : 600
		}`)
	err := json.Unmarshal(data, &add)
	if err == nil {
		fmt.Println(add)
	}
}
