package main

import (
	"fmt"
	"github.com/1r0npipe/go-course2/task7-reflect/reflect_struct/change_struct"
	"log"
)

type person struct {
	FirstName string
	Skills    []Skill
}
type Skill struct {
	Name string
	Exp  int64
}

func main() {
	// Define an example of person struct to replace fields in
	pers := &person{
		FirstName: "john",
		Skills: []Skill{
			{"c++", int64(2)},
			{"python", int64(3)},
			{"golang", int64(5)},
		},
	}
	mapTest := make(map[string]interface{}, 1)
	mapTest["Exp"] = int64(7)
	mapTest["Name"] = "test"
	mapTest["FirstName"] = "Robby"
	// output befor change
	fmt.Printf("%+v\n", pers)
	err := changestruct.ChangeIntoStruct(pers, mapTest)
	if err != nil {
		log.Fatal("Can't change struct by map: ", err)
	}
	// output after change
	fmt.Printf("%+v", pers)
}
