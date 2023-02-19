package test

import (
	"strconv"
	"testing"

	"github.com/shopspring/decimal"

	"github.com/kshamiev/urok/database/bbolt/stbb"
)

const (
	customID1 = "fikus"
	customID2 = "zegota"
	customID3 = "bugor"
)

// 1 Первый этап тестирования (создание)
func TestSave(t *testing.T) {
	inst := newInstance(t)
	saveNewObject(t, inst)
	saveNewObjectCustomIDs(t, inst)
	getObjectExists(t, inst, 135)
}

func saveNewObject(t *testing.T, inst *stbb.Instance) {
	var err error
	var i int
	var objR *Role
	var objU *User
	var objO *Order
	for i = 100; i < 200; i++ {
		objO = &Order{
			ID:    uint64(i),
			Name:  "Order - " + strconv.FormatInt(int64(i), 10),
			Price: decimal.NewFromInt(int64(i)),
		}
		err = inst.Save(objO)
		if err != nil {
			t.Fatal(err)
		}
	}
	for i = 100; i < 200; i++ {
		objR = &Role{
			ID:    uint64(i),
			Name:  "Role - " + strconv.FormatInt(int64(i), 10),
			Price: decimal.NewFromInt(int64(i)),
		}
		err = inst.Save(objR)
		if err != nil {
			t.Fatal(err)
		}
	}
	for i = 100; i < 200; i++ {
		objU = &User{
			ID:    uint64(i),
			Name:  "User - " + strconv.FormatInt(int64(i), 10),
			Price: decimal.NewFromInt(int64(i)),
		}
		err = inst.Save(objU)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func saveNewObjectCustomIDs(t *testing.T, inst *stbb.Instance) {
	var err error
	var i int
	var objO *Order
	for i = 100; i < 120; i++ {
		objO = &Order{
			ID:    uint64(i),
			Name:  "Order - " + strconv.FormatInt(int64(i), 10),
			Price: decimal.NewFromInt(int64(i)),
		}
		err = inst.SaveByID(objO, customID1+strconv.FormatInt(int64(i), 10))
		if err != nil {
			t.Fatal(err)
		}
	}
	for i = 120; i < 140; i++ {
		objO = &Order{
			ID:    uint64(i),
			Name:  "Order - " + strconv.FormatInt(int64(i), 10),
			Price: decimal.NewFromInt(int64(i)),
		}
		err = inst.SaveByID(objO, customID2+strconv.FormatInt(int64(i), 10))
		if err != nil {
			t.Fatal(err)
		}
	}
	for i = 140; i < 160; i++ {
		objO = &Order{
			ID:    uint64(i),
			Name:  "Order - " + strconv.FormatInt(int64(i), 10),
			Price: decimal.NewFromInt(int64(i)),
		}
		err = inst.SaveByID(objO, customID3+strconv.FormatInt(int64(i), 10))
		if err != nil {
			t.Fatal(err)
		}
	}
}

func getObjectExists(t *testing.T, inst *stbb.Instance, id uint64) {
	objU := &User{ID: id}
	err := inst.Load(objU)
	if err != nil {
		t.Fatal(err)
	}
	if objU.Name == "User - "+strconv.FormatUint(id, 10) {
		t.Log("USER VALID")
	} else {
		t.Fatal("USER INVALID")
	}
}
