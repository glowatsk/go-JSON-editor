package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	f, err := ioutil.ReadFile("/home/glowatsk/go/src/github.com/glowatsk/go-JSON-editor/hard.json")
	if err != nil {
		panic(err)
	}

	var i interface{}
	err = json.Unmarshal(f, &i)

	fetchValue(i)

}

func fetchValue(value interface{}) {
	switch value.(type) {
	case string:
		fmt.Printf("%v is an interface \n ", value)
	case bool:
		fmt.Printf("%v is bool \n ", value)
	case float64:
		fmt.Printf("%v is float64 \n ", value)
	case []interface{}:
		fmt.Printf("%v is a slice of interface \n ", value)
		for _, v := range value.([]interface{}) { // use type assertion to loop over []interface{}
			fetchValue(v)
		}
	case map[string]interface{}:
		fmt.Printf("%v is a map \n ", value)
		for _, v := range value.(map[string]interface{}) { // use type assertion to loop over map[string]interface{}
			fetchValue(v)
		}
	default:
		fmt.Printf("%v is unknown \n ", value)
	}
}
