package graph

import "github.com/lucas-code42/graphql-api/mongoDatabase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Account *mongoDatabase.Account
}
