package main

import (
	"fmt"
	"log"
	"reflect"
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
	pers := person{
		FirstName: "john",
		Skills: []Skill{
			{"c++", int64(2)},
			{"python", int64(3)},
			{"golang", int64(5)},
		},
	}
	//pers2 := &person2{"bob", 34}
	mapTest := make(map[string]interface{}, 1)
	mapTest["Name"] = "test"
	fmt.Printf("%+v\n", pers)
	_, err := changeIntoStruct(pers, mapTest)
	if err != nil {
		log.Fatal("Can't change struct by map: ", err)
	}
	fmt.Printf("%+v", pers)
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
		for i := 0; i < val.NumField(); i += 1 {
			typeField := val.Type().Field(i)
			fmt.Println(typeField.Name)
			if typeField.Type.Kind() == reflect.Slice {
				//fmt.Println("We have the slice to parse next")
				for j := 0; j < val.Field(i).Len(); j += 1 {
					itemStruct := val.Field(i).Index(j)
					changeIntoStruct(itemStruct.Interface(), mapInit)
					continue
				}
			}
			if typeField.Type.Kind() == reflect.Struct {
				//fmt.Println("We have just one more struct inside")
				changeIntoStruct(val.Field(i).Interface(), mapInit)
				continue
			}
			mapValue, ok := mapInit[typeField.Name]
			if !ok {
				continue
			}
			fmt.Printf("%+v - %+v\n", mapValue, typeField.Name)
			switch val.Field(i).Type().Kind() {
			case reflect.Int:
				val.Field(i).SetInt(mapValue.(int64))
			case reflect.Float64:
				val.Field(i).SetFloat(mapValue.(float64))
			case reflect.String:
				val.Field(i).SetString(mapValue.(string))
			case reflect.Bool:
				val.Field(i).SetBool(mapValue.(bool))
			default:
				return nil, fmt.Errorf("type mistamtch to assign the value")
			}
		}
	}
	return val.Interface(), nil
}
