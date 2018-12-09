package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Health Check",
		"GET",
		"/health",
		HealthCheck,
	},
	Route{
		"getalldata",
		"GET",
		"/data/getalldata",
		getalldata,
	},

	Route{
		"getdatabyid ",
		"GET",
		"/data/getdatabyid/{id}",
		getdatabyid,
	},

	Route{
		"getdatabytitle",
		"GET",
		"/data/getdatabytitle/{searchtitle}",
		getdatabytitle,
	},

	Route{
		"getdatasorted",
		"GET",
		"/data/getdatasorted/{orderby}",
		getdatasorted,
	},
}
