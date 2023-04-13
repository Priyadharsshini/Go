package main

import "fmt"

type DataSources interface{
	getData() ([]byte, error)
}

type API struct {
	// some fields specific to this
}

type File struct {
	// some fields specific to this
}

type DataBase struct {
	// some fields specific to this 
}

func(a API) getData()([]byte, error){
	// Implementation specific to this 
}

func(f File) getData()([]byte, error){
	// Implementation specific to this
}

func(d DataBase) getData()([]byte, error){
	// Implentation specific to this
}

func processData(ds DataSources){
	// Any common preprocessing required for them all 
	value , err := ds.getData()
	if err != nil {
		//do something
	}
}

func main(){
	f := File{}
	d := DataBase{}
	a := API{}
	processData(f)
}
