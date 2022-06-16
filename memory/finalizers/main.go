package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
)

type TestHeap struct {
	ID             int64             `db:"id" pg:"id,pk" boil:"id" json:"id"`
	Login          string            `db:"login" pg:"login,use_zero" boil:"login" json:"login"`
	Description    null.String       `db:"description" pg:"description,use_zero" boil:"description" json:"description,omitempty" swaggertype:"string"`
	Price          decimal.Decimal   `db:"price" pg:"price,use_zero" boil:"price" json:"price" swaggertype:"number" example:"0.01"`
	SummaOne       float32           `db:"summa_one" pg:"summa_one,use_zero" boil:"summa_one" json:"summa_one" swaggertype:"number" example:"0.01"`
	SummaTwo       float64           `db:"summa_two" pg:"summa_two,use_zero" boil:"summa_two" json:"summa_two" swaggertype:"number" example:"0.01"`
	CNT            int               `db:"cnt" pg:"cnt,use_zero" boil:"cnt" json:"cnt"`
	CNT2           int16             `db:"cnt2" pg:"cnt2,use_zero" boil:"cnt2" json:"cnt2"`
	CNT4           int               `db:"cnt4" pg:"cnt4,use_zero" boil:"cnt4" json:"cnt4"`
	CNT8           int64             `db:"cnt8" pg:"cnt8,use_zero" boil:"cnt8" json:"cnt8"`
	ShardingID     uuid.UUID         `db:"sharding_id" pg:"sharding_id,use_zero" boil:"sharding_id" json:"sharding_id" example:"8ca3c9c3-cf1a-47fe-8723-3f957538ce42"`
	IsOnline       bool              `db:"is_online" pg:"is_online,use_zero" boil:"is_online" json:"is_online"`
	JSONStringNull null.String       `db:"json_string_null" pg:"json_string_null,use_zero" boil:"json_string_null" json:"json_string_null,omitempty" swaggertype:"string"`
	JSONString     string            `db:"json_string" pg:"json_string,use_zero" boil:"json_string" json:"json_string"`
	Duration       time.Duration     `db:"duration" pg:"duration,use_zero" boil:"duration" json:"duration" swaggertype:"number" example:"0"`
	Data           null.Bytes        `db:"data" pg:"data,use_zero" boil:"data" json:"data,omitempty" swaggertype:"string" example:"BYTES"`
	Alias          types.StringArray `db:"alias" pg:"alias,use_zero" boil:"alias" json:"alias,omitempty"`
	CreatedAt      time.Time         `db:"created_at" pg:"created_at,use_zero" boil:"created_at" json:"created_at" example:"2006-01-02T15:04:05Z"`
	UpdatedAt      time.Time         `db:"updated_at" pg:"updated_at,use_zero" boil:"updated_at" json:"updated_at" example:"2006-01-02T15:04:05Z"`
	DeletedAt      null.Time         `db:"deleted_at" pg:"deleted_at,use_zero" boil:"deleted_at" json:"deleted_at,omitempty" example:"2006-01-02T15:04:05Z"`
}

func testHeap() {
	// create pointer
	b := &TestHeap{}
	// add finalizer which just prints
	runtime.SetFinalizer(b, func(b *TestHeap) { fmt.Println("I AM DEAD - HEAP") })
}

type TestStack struct {
	A int
}

func testStack() {
	// create pointer
	a := &TestStack{}
	// add finalizer which just prints
	runtime.SetFinalizer(a, func(a *TestStack) { fmt.Println("I AM DEAD - STACK") })
}

func main() {
	testStack()
	testHeap()
	// run garbage collection
	runtime.GC()
	// sleep to switch to finalizer goroutine
	time.Sleep(1 * time.Millisecond)
}
