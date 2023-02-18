package test

import (
	"testing"
)

func TestOrders(t *testing.T) {
	inst := newInstance(t)
	obj := &Order{ID: 23}
	objs := OrderSlice{}
	objs = append(objs, &Order{})

	err := inst.LoadRelation(obj, &objs)
	if err != nil {
		t.Fatal(err)
	}

	err = inst.Select(&objs)
	if err != nil {
		t.Fatal(err)
	}
	// err = inst.DeleteRelation(obj, objs1...)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	t.Log(objs)

}
