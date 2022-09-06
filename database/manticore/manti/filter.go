package manti

type Filter struct {
	Search    string
	Attribute []Attribute
	Order     []Order
	Page      int
	Limit     int
}

type Attribute struct {
	Name  string
	Sign  string
	Value string
}

type Order struct {
	Name  string
	Order string
}
