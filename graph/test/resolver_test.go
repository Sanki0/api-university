package test

import (
	"testing"

	"github.com/Sanki0/api-university/graph/resolver"
)

func TestResolverInitialize(t *testing.T) {
	var r resolver.Resolver
	r.InitializePool()
	if r.DB == nil {
		t.Error("Resolver.InitializePool() failed")
	}
}
