package main

import (
	"context"
	"fmt"
)

func main() {

	// u := &User{Name: "popcorn"}
	ctx := context.Background()
	// ctx = context.WithValue(ctx, "test", u)

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
