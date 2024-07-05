package personGen

import "net/http"

func AddRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /info", InfoHandler)
}
