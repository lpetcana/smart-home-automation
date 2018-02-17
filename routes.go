package main

import (
	"smart-home-automation/controllers"
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunctions []negroni.HandlerFunc

}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		[]negroni.HandlerFunc{controllers.HelloController},
	},
	/*Route{
		"TodoIndex",
		"GET",
		"/todos",
		[]negroni.HandlerFunc{TodoIndex},
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		[]negroni.HandlerFunc{TodoShow},
	},*/
}

func SetAuthenticationRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/token-auth", controllers.Login).Methods("POST")
	/*	router.Handle("/refresh-token-auth",
			negroni.New(
				negroni.HandlerFunc(authentication.RequireTokenAuthentication),
				negroni.HandlerFunc(controllers.RefreshToken),
			)).Methods("GET")*/
	router.Handle("/logout",
		negroni.New(
			//negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.Logout),
		)).Methods("GET")

	return router
}

func SetUnathenticatedRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/token-info/", controllers.CheckTokenInfo).Methods("GET")

	router.Handle("/test/settings",
		negroni.New(
			negroni.HandlerFunc(controllers.HelloController),
		)).Methods("GET")

	router.Handle("/",
		negroni.New(
			negroni.HandlerFunc(controllers.Index),
		)).Methods("GET")

	router.Handle("/todos",
		negroni.New(
			negroni.HandlerFunc(TodoIndex),
		)).Methods("GET")

	router.Handle("/todos/{todoId}",
		negroni.New(
			negroni.HandlerFunc(TodoShow),
		)).Methods("GET")

	router.Handle("/api/weather",
		negroni.New(
			negroni.HandlerFunc(controllers.Weather),
		)).Methods("GET")

	return router
}
