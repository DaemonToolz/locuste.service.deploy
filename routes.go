package main

import (
	"net/http"
)

// Route Chemin Web
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes Ensemble de chemins HTTP
type Routes []Route

// A mettre dans un JSON (et charger via Swagger ?)
var routes = Routes{
	Route{
		"Upload file",
		"POST",
		"/upload/{version}",
		Upload,
	},
	Route{
		"Start install procedure",
		"GET",
		"/install/{version}",
		Install,
	},

	Route{
		"Get available versions",
		"GET",
		"/versions",
		GetAvailableVersions,
	},

	Route{
		"Get available versions",
		"DELETE",
		"/delete/{version}",
		DeleteVersion,
	},
}
