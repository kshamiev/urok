package stbb

import "encoding/json"

type OrderSlice []*Order

func (self OrderSlice) GetIndex() string {
	return "orders"
}

func (self *OrderSlice) ParseByte(_, value []byte) error {
	o := &Order{}
	err := json.Unmarshal(value, o)
	if err != nil {
		return err
	}
	*self = append(*self, o)
	return nil
}
