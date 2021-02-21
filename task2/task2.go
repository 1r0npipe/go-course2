package main

import (
	"encoding/json"
	"fmt"
  )
  
  // MyStruct my structure
  type MyStruct struct {}
  
  // MarshalJSON is a function which return byte from JSON format or error
  func (f MyStruct) MarshalJSON() ([]byte, error) {
	return []byte(`{"a": 0}`), nil
  }
  
  func main() {
	j, _ := json.Marshal(MyStruct{})
	fmt.Println(string(j))
  }
  