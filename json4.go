package main

import (
	"encoding/json"
	"net/http"
)

type Datos struct {
	Addres, Ssl_grade, Country, Owner string
}

type Info struct {
	Servers_changed    bool
	Ssl_grade          string
	Previous_ssl_grade string
	Logo               string
	Title              string
	Is_down            bool
}

type Employee struct {
	Servers  []Datos
	Servers2 []Info
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
			//}
			//json.NewEncoder(w).Encode(data)

			//file, _ := json.MarshalIndent(data, "", " ")
			///aqui inicia el otro
			//data = Employee{

			Servers2: []Info{
				Info{
					Servers_changed:    true,
					Ssl_grade:          "B",
					Previous_ssl_grade: "A+",
					Logo:               "https://server.com/icon.png",
					Title:              "Title of the page",
					Is_down:            false,
				},
			},
		}
		json.NewEncoder(w).Encode(data)
	})

	//	file, _ := json.MarshalIndent(data2, "", " ")

	//	_ = ioutil.WriteFile("test.json", file, 0644)
	//json.NewEncoder(w).Encode(data)
	//json.NewEncoder(w).Encode(data2)
	http.ListenAndServe(":8000", nil)

}
