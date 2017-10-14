package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Yo, I built a thing :)")
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
