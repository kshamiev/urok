package test

import (
	"strconv"
	"testing"

	"github.com/kshamiev/urok/database/bbolt/stbb"
)

// 3 Третий этап тестирования (удаление)
func TestDelete(t *testing.T) {
	inst := newInstance(t)
	deleteObject(t, inst)
	getObjectNotExists(t, inst, 135)
}

func deleteObject(t *testing.T, inst *stbb.Instance) {
	var err error
	var i int
	var objR *Role
	var objU *User
	var objO *Order
	for i = 100; i < 200; i++ {
		objO = &Order{ID: uint64(i)}
		err = inst.Delete(objO)
		if err != nil {
			t.Fatal(err)
		}
	}
	for i = 100; i < 200; i++ {
		objR = &Role{ID: uint64(i)}
		err = inst.Delete(objR)
		if err != nil {
			t.Fatal(err)
		}
	}
	for i = 100; i < 200; i++ {
		objU = &User{ID: uint64(i)}
		err = inst.Delete(objU)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func getObjectNotExists(t *testing.T, inst *stbb.Instance, id uint64) {
	objU := &User{ID: id}
	err := inst.Load(objU)
	if err == nil || objU.Name == "User - "+strconv.FormatUint(id, 10) {
		t.Fatal("объект не должен существовать")
	}
}
