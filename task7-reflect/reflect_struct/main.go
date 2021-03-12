package main

import (
	"fmt"
	"log"
	"reflect"
)

type person struct {
	name   string
	skills []Skill
}
type Skill struct {
	name string
	exp  uint32
}

func main() {
	pers := &person{
		name: "john",
		skills: []Skill{
			{"c++", 2},
			{"python", 3},
			{"golang", 5},
		},
	}
	mapTest := make(map[string]interface{}, 1)
	mapTest["python"] = 5
	_, err := changeIntoStruct(pers, mapTest)
	if err != nil {
		log.Fatal("Can't change struct by map: ", err)
	}

	//fmt.Println(result)
}

func changeIntoStruct(in interface{}, mapInit map[string]interface{}) (interface{}, error) {
	if in == nil {
		return nil, fmt.Errorf("interface is empty")
	}
	val := reflect.ValueOf(in)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("the interface is not structure")
	}
	for i := 0; i < val.NumField(); i += 1 {
		typeField := val.Type().Field(i)
		if typeField.Type.Kind() == reflect.Slice {
			log.Println("We have the slice to parse next")
			for j := 0; j < val.Field(i).Len(); j += 1 {
				itemStruct := val.Field(i).Index(j)
				if itemStruct.Kind() != reflect.Struct {
					return nil, fmt.Errorf("something wrong with Skill struct definition")
				}
				changeIntoStruct(itemStruct.Interface(), mapInit)
				continue
			}
		}
		if typeField.Type.Kind() == reflect.Struct {
			log.Println("We have just one more struct inside")
			changeIntoStruct(val.Field(i).Interface(), mapInit)
			continue
		}

	}
	return val.Interface(), nil
}
