package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Datos struct {
	Addres, Ssl_grade, Country, Owner string
}

type Info struct {
	Servers_changed    string
	Ssl_grade          string
	Previous_ssl_grade string
	Logo               string
	Title              string
	Is_down            string
}

type Employee struct {
	Servers []Datos
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := Employee{

			Servers: []Datos{
				Datos{
					Addres:    "server1",
					Ssl_grade: "B",
					Country:   "US",
					Owner:     "Amazon.com, Inc.",
				},
				Datos{
					Addres:    "server2",
					Ssl_grade: "A+",
					Country:   "US",
					Owner:     "Amazon.com, Inc.",
				},
				Datos{
					Addres:    "server3",
					Ssl_grade: "A",
					Country:   "US",
					Owner:     "Amazon.com, Inc.",
				},
			},
		}

		file, _ := json.MarshalIndent(data, "", " ")

		_ = ioutil.WriteFile("test.json", file, 0644)
		json.NewEncoder(w).Encode(data)

	})
	http.ListenAndServe(":8000", nil)
}
