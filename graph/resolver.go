package graph

import (
	"github.com/Sanki0/api-university/graph/connection"
	"github.com/jinzhu/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *gorm.DB
}

func (r *Resolver) InitializePool() {
	r.DB = connection.FetchConnection()
}
