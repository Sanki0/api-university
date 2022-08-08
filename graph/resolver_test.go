package graph

import "testing"



func TestResolverInitialize(t *testing.T) {
	var r Resolver
	r.InitializePool()
	if r.DB == nil {
		t.Error("Resolver.InitializePool() failed")
	}
}