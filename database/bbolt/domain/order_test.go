package domain

import (
	"testing"
)

func TestOrder(t *testing.T) {

	obj := Order{
		ID:   34,
		Name: "Popcorn",
	}
	t.Log(obj)

}
