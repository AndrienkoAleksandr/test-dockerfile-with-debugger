package main

import (
"log"
"net/http"
)

func main() {
	log.Println("launching server...")

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(`Hello world`)); err != nil {
			log.Println("Fail to send response ", err)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
