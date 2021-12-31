package routing

import (
  "net/http"
  "github.com/gorilla/mux"
  apiHandlers "backend/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route


var routes = Routes{
  	Route{
  		"TodoShow",
  		"GET",
  		"/quotes/{sourceName}",
  		apiHandlers.GetRandomQuote,
  	},
}


func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
