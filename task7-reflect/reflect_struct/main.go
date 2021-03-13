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
type person2 struct {
	name string
	age  int
}
type Skill struct {
	name string
	exp  uint32
}

func main() {
	//pers := &person{
	//	name: "john",
	//	skills: []Skill{
	//		{"c++", 2},
	//		{"python", 3},
	//		{"golang", 5},
	//	},
	//}
	pers2 := &person2{"bob", 34}
	mapTest := make(map[string]interface{}, 1)
	mapTest["python"] = 5
	result, err := changeIntoStruct(*pers2, mapTest)
	if err != nil {
		log.Fatal("Can't change struct by map: ", err)
	}
	fmt.Println(result)
}

func changeIntoStruct(in interface{}, mapInit map[string]interface{}) (interface{}, error) {
	if in == nil {
		return nil, fmt.Errorf("interface is empty")
	}
	val := reflect.ValueOf(in)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() == reflect.Struct {
		fmt.Println("checking struct...")
		for i := 0; i < val.NumField(); i += 1 {
			typeField := val.Type().Field(i)
			if typeField.Type.Kind() == reflect.Slice {
				fmt.Println("We have the slice to parse next")
				for j := 0; j < val.Field(i).Len(); j += 1 {
					itemStruct := val.Field(i).Index(j)
					fmt.Println("wow", itemStruct)
					changeIntoStruct(itemStruct.Addr().Interface(), mapInit)
					continue
				}
			}
			if typeField.Type.Kind() == reflect.Struct {
				fmt.Println("We have just one more struct inside")
				changeIntoStruct(val.Field(i).Interface(), mapInit)
				continue
			}
			for key, value := range mapInit {
				if val.Field(i).Kind() == reflect.String {
					f := val.Field(i)
					if f == key {
						if val.Field(i + 1).IsValid() {
							val.Field(i + 1).Set(value)
						}
					}
				}
			}
		}
	}

	return val.Interface(), nil
}
