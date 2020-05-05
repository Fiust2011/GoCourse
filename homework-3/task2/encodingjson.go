package main

import (
    "encoding/json"
    "fmt"
)

type Automobile struct{
	Carbody string
	Year int
	Bodyvolume string

}

var jsonStr = '{"carbody": "car", "year": 2020, "bodyvolume": "200"}'

func main() {
	
	data := []byte(jsonStr)

	a := &Automobile{}
	err := json.Unmarshal(data, a)
	if err != nil {
		panic(err)
	}
	fmt.Printf("struct:\n\t%#v\n\n", a)
	
	a.bodyvolume = "200"
	result, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	fmt.Printf("json string:\n\t%s\n", string(result))
}