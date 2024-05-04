package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/oosawy/imageon"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			message := errors.Join(errors.New("failed to read request body"), err).Error()
			http.Error(w, message, http.StatusInternalServerError)
			return
		}

		res, err := imageon.HandleRequest(r.Context(), imageon.RawRequest{
			Path: r.URL.Path,
			Body: string(body),
		})
		if err != nil {
			message := errors.Join(errors.New("internal server error"), err).Error()
			http.Error(w, message, http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "%d: %s", res.StatusCode, res.Body)
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
