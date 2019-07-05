// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package dep

import (
	"tinyURL/app/graphql"
	"tinyURL/fw"
)

// Injectors from wire.go:

func InitializeApp() fw.App {
	graphQl := graphql.NewTinyUrlGraphQl()
	server := fw.NewGraphGophers(graphQl)
	app := fw.NewApp(server)
	return app
}