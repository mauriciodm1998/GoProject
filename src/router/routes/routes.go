package routes

import (
	"API/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI          string
	Method       string
	Func         func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

func Config(r *mux.Router) *mux.Router {
	routes := UserRoutes
	routes = append(routes, loginRoute)

	for _, route := range routes {

		if route.AuthRequired {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authorization(route.Func)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Func)).Methods(route.Method)
		}
	}

	return r
}
