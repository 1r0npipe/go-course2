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
	pers := &person{
		FirstName: "john",
		Skills: []Skill{
			{"c++", int64(2)},
			{"python", int64(3)},
			{"golang", int64(5)},
		},
	}
	//pers2 := &person{"bob", []Skill{}}
	mapTest := make(map[string]interface{}, 1)
	mapTest["Exp"] = int64(7)
	mapTest["Name"] = "test"
	mapTest["FirstName"] = "Robby"
	fmt.Printf("%+v\n", pers)
	err := changeIntoStruct(pers, mapTest)
	if err != nil {
		log.Fatal("Can't change struct by map: ", err)
	}
	fmt.Printf("%+v", pers)
}
