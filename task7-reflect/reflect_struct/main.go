package main

import (
	"fmt"
	"log"
	"reflect"
)

type person struct {
	name   string
	skills []Skills
}
type Skills struct {
	name string
	exp  uint32
}

func main() {
	pers := &person{
		name: "john",
		skills: []Skills{
			{"c++", 2},
			{"python", 3},
		},
	}
	mapTest := make(map[string]interface{}, 1)
	mapTest["python"] = 5
	result, err := changeIntoStruct(pers, mapTest)
	if err != nil {
		log.Fatal("Can't change struct by map")
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
	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("the interface is not structure")
	}
	for i := 0; i < val.NumField(); i += 1 {
		typeField := val.Type().Field(i)
		if typeField.Type.Kind() == reflect.Slice {
			log.Println("We have the slice to parse next")
			sli, ok := val.Field(i).Interface().([]Skills)
			if !ok {
				return nil, fmt.Errorf("can't get slice of stuct Skills")
			}
			for j := 0; j < len(sli); i += 1 {
				for valuesMap, keyMap := range mapInit {
					if sli[j].name == valuesMap {
						realMapValue, ok := keyMap.(uint32)
						if !ok {
							return nil, fmt.Errorf("Can't extract real value from map")
						}
						sli[j].exp = realMapValue
					}
				}
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
