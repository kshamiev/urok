// Комментарий к пакету

package mockone

import "context"

//go:generate mockery  --dir . --name IOAuthProvider --output bublik
type IOAuthProvider interface {
	GetConfig() Employee
	GetAuthURL() string
	Exchange(ctx context.Context, code string) (*Employee, error)
	GetAccessToken(token Employee) string
	RefreshToken(ctx context.Context, t string) (*Employee, error)
}

// Employee Какой-то тип
type Employee struct {
	Name   string // fdsfsd
	Salary int    // fsdf
	Sales  int    // fsd
	Bonus  int    // fdsfsd
}
