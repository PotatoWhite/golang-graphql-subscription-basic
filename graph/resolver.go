package graph

import (
	"github.com/potatowhite/golang-graphql-basic/graph/model"
	"sync"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	messageSent map[string]chan *model.Message
	mu          sync.Mutex
}
