package typs

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/kshamiev/urok/database/manticore/manti"
)

type Film struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category_id int64     `json:"category_id"`
	ReleaseYear int       `json:"release_year"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type Document struct {
	ID          int64           `json:"id"`
	Title       string          `json:"title"`
	CategoryID  int64           `json:"category_id"`
	ReleaseYear uint            `json:"release_year"`
	Price       decimal.Decimal `json:"price"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	DeletedAt   time.Time       `json:"deleted_at"`
	IsFlag      bool            `json:"is_flag"`
	Description string          `json:"description"`
}

type Documents struct {
	Total int        `json:"total"`
	Data  []Document `json:"data"`
}

func NewDocuments(limit int) *Documents {
	return &Documents{
		Total: 0,
		Data:  make([]Document, 0, limit),
	}
}

func (doc *Documents) Parse(row map[string]interface{}) {
	doc.Data = append(doc.Data, Document{
		ID:          manti.ConvertBigint(row["id"]),
		Title:       manti.ConvertString(row["title"]),
		CategoryID:  manti.ConvertBigint(row["category_id"]),
		ReleaseYear: manti.ConvertUint(row["release_year"]),
		Price:       manti.ConvertFloat(row["price"]),
		CreatedAt:   manti.ConvertTime(row["created_at"]),
		UpdatedAt:   manti.ConvertTime(row["updated_at"]),
		DeletedAt:   manti.ConvertTime(row["deleted_at"]),
		IsFlag:      manti.ConvertBool(row["is_flag"]),
		Description: manti.ConvertString(row["description"]),
	})
}

func (doc *Documents) SetCount(cnt int) {
	doc.Total = cnt
}
