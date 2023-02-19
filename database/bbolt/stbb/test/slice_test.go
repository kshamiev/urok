package test

import (
	"testing"

	"github.com/kshamiev/urok/database/bbolt/stbb"
)

func TestOrders(t *testing.T) {
	inst := newInstance(t)
	obj := &Order{ID: 23}
	// objs := OrderSlice{}
	var objs OrderSlice
	objs = append(objs, &Order{})

	objLiist := []stbb.Modeler{}
	for i := range objs {
		objLiist = append(objLiist, objs[i])
	}

	err := inst.SaveRelation(obj, objs.GetIndex(), objs.GetIds())
	if err != nil {
		t.Fatal(err)
	}

	if objs == nil {
		t.Log("OK")
	}

	err = inst.Select(&objs, true)
	if err != nil {
		t.Fatal(err)
	}
	// err = inst.DeleteRelation(obj, objs1...)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	if objs == nil {
		t.Log("OK")
	}
	t.Log(objs)

}
