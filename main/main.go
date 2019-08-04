package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func main() {
	parseJSON()
	serveHTTP()
}

func serveHTTP() {
	tmpl := template.Must(template.ParseFiles("/home/glowatsk/go/src/github.com/glowatsk/go-JSON-editor/tpl.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := "You are getting data"
		tmpl.Execute(w, data)
		fmt.Println(r)
	})

	http.ListenAndServe(":8080", nil)
}

func parseJSON() {
	f, err := ioutil.ReadFile("/home/glowatsk/go/src/github.com/glowatsk/go-JSON-editor/hard.json")
	if err != nil {
		panic(err)
	}

	var i interface{}
	err = json.Unmarshal(f, &i)

	fetchValue(i)
}

func fetchValue(value interface{}) {
	fmt.Printf("%v \n", value)
	switch value.(type) {
	case string:
		fmt.Printf("%v is an interface \n ", value)
	case bool:
		fmt.Printf("%v is bool \n ", value)
	case float64:
		fmt.Printf("%v is float64 \n ", value)
	case []interface{}:
		// fmt.Printf("%v is a slice of interface \n ", value)
		for _, v := range value.([]interface{}) { // use type assertion to loop over []interface{}
			fetchValue(v)
		}
	case map[string]interface{}:
		// fmt.Printf("%v is a map \n ", value)
		for k, v := range value.(map[string]interface{}) { // use type assertion to loop over map[string]interface{}
			fmt.Printf("%v \n", k)
			fetchValue(v)
		}
	default:
		fmt.Printf("%v is unknown \n ", value)
	}
}
