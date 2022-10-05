// Комментарий к пакету

package mockery

import "context"

//go:generate mockery  --dir . --name IOAuthProvider --output mockerytest
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
