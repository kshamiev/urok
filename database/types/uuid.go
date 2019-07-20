package types

import (
	"database/sql/driver"
	"errors"

	"github.com/gofrs/uuid"
)

// String is a nullable string. It supports SQL and JSON serialization.
type UUID struct {
	uuid.UUID
}

func NewUUID() UUID {
	u, _ := uuid.NewV4()
	return UUID{u}
}

// Randomize for sqlboiler
func (uu *UUID) Randomize(nextInt func() int64, fieldType string, shouldBeNull bool) {
	uu.UUID, _ = uuid.NewV4()
}

// Value implements the driver.Valuer interface.
func (uu UUID) Value() (driver.Value, error) {
	v := uuid.UUID{}
	if uu.UUID == v {
		return nil, errors.New("UUID is null")
	}
	return uu.Bytes(), nil
}


