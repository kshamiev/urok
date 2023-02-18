package test

import "encoding/json"

type OrderSlice []*Order

func (self OrderSlice) GetIndex() string {
	return "orders"
}

func (self *OrderSlice) ParseByte(_, value []byte) {
	o := &Order{}
	_ = json.Unmarshal(value, o)
	*self = append(*self, o)
}
