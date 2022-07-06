package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World!")
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Printf("data=%s\n", data)
		fmt.Fprintf(w, "Hello %s\n!!", data)

	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye World!")
	})

	http.ListenAndServe(":9090", nil)
}
