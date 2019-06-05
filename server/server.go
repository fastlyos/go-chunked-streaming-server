package server

import (
	"log"
	"net/http"
)

func StartHttpServer() error {
	return http.ListenAndServe(":9094", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer w.(http.Flusher).Flush()
		log.Printf("%s %s", r.Method, r.URL.String())
		switch r.Method {
		case http.MethodGet:
			GetHandler(w, r)
		case http.MethodHead:
			HeadHandler(w, r)
		case http.MethodPost:
			PostHandler(w, r)
		case http.MethodPut:
			PutHandler(w, r)
		case http.MethodDelete:
			DeleteHandler(w, r)
		case http.MethodOptions:
			OptionsHandler(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}))
}