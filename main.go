package main

import (
	"context"
	"fmt"
)

func main() {

	a := []int{1}
	b := a[0:1]

	a = append(a, 2)
	a[0] = 10

	fmt.Println(a, b)
	// [10 2] [0]

	u := &User{Name: "popcorn"}
	ctx := context.Background()
	ctx = context.WithValue(ctx, "test", u)
	uu, err := TokenFromContext(ctx)
	fmt.Println(uu, err)
}

func TokenFromContext(ctx context.Context) (*User, error) {
	token, ok := ctx.Value("test").(*User)
	if !ok {
		return nil, fmt.Errorf("user not present in context")
	}
	return token, nil
}

type User struct {
	Name string
}
