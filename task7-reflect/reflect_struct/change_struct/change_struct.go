package changestruct

import (
	"fmt"
	"reflect"
)

// ChangeIntoStruct the function which gets interface (structure) and map
// where you define the key[field in struct] and value -> change into struct
// for this value, it works with slices and sub-structure, returns error
// can work with String and int64 types inside of Struct (or Sub-Struct)
func ChangeIntoStruct(in interface{}, mapInit map[string]interface{}) error {
	if in == nil {
		return fmt.Errorf("interface is empty")
	}
	val := reflect.ValueOf(in)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() == reflect.Struct {
		for i := 0; i < val.NumField(); i += 1 {
			typeField := val.Type().Field(i)
			if typeField.Type.Kind() == reflect.Slice { // if we need to parse Slice
				for j := 0; j < val.Field(i).Len(); j += 1 {
					itemStruct := val.Field(i).Index(j)
					ChangeIntoStruct(itemStruct.Addr().Interface(), mapInit)
					continue
				}
			}
			if typeField.Type.Kind() == reflect.Struct { // if we need to parse Struct
				ChangeIntoStruct(val.Field(i).Addr().Interface(), mapInit)
				continue
			}
			mapValue, ok := mapInit[typeField.Name]
			if !ok {
				continue
			}
			switch typeField.Type.Kind() {
			case reflect.Int64:
				val.Field(i).SetInt(mapValue.(int64))
			case reflect.String:
				val.Field(i).SetString(mapValue.(string))
			default:
				return fmt.Errorf("type mistamtch to assign the value")
			}
		}
	}
	return nil
}
