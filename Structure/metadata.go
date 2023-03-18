package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

//In Go, the encoding/json and encoding/xml packages provide support for serialization and deserialization of JSON and XML data, respectively. 

func main() {
	user := User{
		ID:       1,
		Name:     "Priya",
		Password: "Priya",
	}
	
	//Serialization - Serialization is the process of converting an object or data structure into a format that can be transmitted over a network or stored in a file
	//The goal of serialization is to make it possible to transfer or store data in a format that is platform-independent 
	//and can be used by different programming languages
	
	userJson, err := json.Marshal(user)
	if err != nil {
		fmt.Println("There was an error when we tried to marshal the data")
	}
	fmt.Println(string(userJson))

	//Deserialization - Deserialization is the process of converting the serialized data back into an object or data structure.This involves recreating the original object.
	
	err = json.Unmarshal(userJson, &deserializedUser)
	if err != nil {
		fmt.Println("There was an error when we tried to unmarshal the data")
	}
	fmt.Println(deserializedUser)

}
