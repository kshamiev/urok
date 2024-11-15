package test

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"

	"github.com/kshamiev/urok/database/bbolt/stbb"
)

const indexUser = "users"

type User struct {
	ID          uint64            `db:"id" pg:"id,pk" boil:"id" json:"id"`
	Name        string            `db:"login" pg:"login,use_zero" boil:"login" json:"login"`
	Description null.String       `db:"description" pg:"description,use_zero" boil:"description" json:"description,omitempty" swaggertype:"string"`
	Price       decimal.Decimal   `db:"price" pg:"price,use_zero" boil:"price" json:"price" swaggertype:"number" example:"0.01"`
	SummaOne    float32           `db:"summa_one" pg:"summa_one,use_zero" boil:"summa_one" json:"summa_one" swaggertype:"number" example:"0.01"`
	SummaTwo    float64           `db:"summa_two" pg:"summa_two,use_zero" boil:"summa_two" json:"summa_two" swaggertype:"number" example:"0.01"`
	ShardingID  uuid.UUID         `db:"sharding_id" pg:"sharding_id,use_zero" boil:"sharding_id" json:"sharding_id" example:"8ca3c9c3-cf1a-47fe-8723-3f957538ce42"`
	IsOnline    bool              `db:"is_online" pg:"is_online,use_zero" boil:"is_online" json:"is_online"`
	Duration    time.Duration     `db:"duration" pg:"duration,use_zero" boil:"duration" json:"duration" swaggertype:"number" example:"0"`
	Alias       types.StringArray `db:"alias" pg:"alias,use_zero" boil:"alias" json:"alias,omitempty"`
	CreatedAt   time.Time         `db:"created_at" pg:"created_at,use_zero" boil:"created_at" json:"created_at" example:"2006-01-02T15:04:05Z"`
	UpdatedAt   time.Time         `db:"updated_at" pg:"updated_at,use_zero" boil:"updated_at" json:"updated_at" example:"2006-01-02T15:04:05Z"`
	DeletedAt   null.Time         `db:"deleted_at" pg:"deleted_at,use_zero" boil:"deleted_at" json:"deleted_at,omitempty" example:"2006-01-02T15:04:05Z"`
}

func (self User) GetIndex() string {
	return indexUser
}

func (self User) GetBID() []byte {
	return stbb.Itob(self.ID)
}

func (self *User) SetBID(id []byte) {
	self.ID = stbb.Btoi(id)
}

//

type UserSlice []*User

func (self UserSlice) GetIndex() string {
	return indexUser
}

func (self *UserSlice) ParseObject(i int, value []byte) {
	o := &User{}
	_ = json.Unmarshal(value, o)
	if i > -1 {
		(*self)[i] = o
	} else {
		*self = append(*self, o)
	}
}

func (self UserSlice) GetIds() [][]byte {
	res := make([][]byte, len(self))
	for i := range self {
		res[i] = self[i].GetBID()
	}
	return res
}
