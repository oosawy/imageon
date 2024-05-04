package main

import (
	"fmt"
	"net/http"

	"github.com/oosawy/imageon"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		res, err := imageon.HandleRequest(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, *res)
	})
	go func() {
		err := http.ListenAndServe(":8080", handler)
		if err != nil {
			panic(err)
		}
	}()

	fmt.Println("Server started at http://localhost:8080")

	// Keep the main goroutine running
	select {}
}
