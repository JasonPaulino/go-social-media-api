package v1

import (
	"log"
	"net/http"
)

func Routes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello, World!"))

		if err != nil {
			log.Fatal(err)
		}
	})

	return router
}
