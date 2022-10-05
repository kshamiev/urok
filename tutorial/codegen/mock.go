package codegen

import "context"

//go:generate mockery  --dir . --name IOAuthProvider --output bublik
type IOAuthProvider interface {
	GetConfig() Employee
	GetAuthURL() string
	Exchange(ctx context.Context, code string) (*Employee, error)
	GetAccessToken(token Employee) string
	RefreshToken(ctx context.Context, t string) (*Employee, error)
}
