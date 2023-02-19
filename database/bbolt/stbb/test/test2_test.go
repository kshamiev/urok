package test

import (
	"fmt"
	"testing"

	"github.com/kshamiev/urok/database/bbolt/stbb"
)

// 2 Второй этап тестирования (манипуляции с данными)
func TestSlice(t *testing.T) {
	inst := newInstance(t)
	selectPrefix(t, inst)
	selectRange(t, inst)
}

func selectPrefix(t *testing.T, inst *stbb.Instance) {
	t.Log("selectPrefix")
	var objSlice OrderSlice
	err := inst.SelectPrefix(&objSlice, customID2)
	if err != nil {
		t.Fatal(err)
	}
	for i := range objSlice {
		t.Log(objSlice[i])
	}
	fmt.Println(len(objSlice))
}

func selectRange(t *testing.T, inst *stbb.Instance) {
	t.Log("selectRange")
	var objSlice UserSlice
	err := inst.SelectRange(&objSlice, string(stbb.Itob(125)), string(stbb.Itob(135)))
	if err != nil {
		t.Fatal(err)
	}
	for i := range objSlice {
		t.Log(objSlice[i])
	}
	fmt.Println(len(objSlice))
}

func saveRelation(t *testing.T, inst *stbb.Instance) {
	t.Log("saveRelation")
	var err error
	objU := &User{ID: 111}

	// привязываем две роли
	objR := &Role{ID: 143}
	err = inst.SaveRelation(objU, &RoleSlice{objR})
	if err != nil {
		t.Fatal(err)
	}
	objR = &Role{ID: 139}
	err = inst.SaveRelation(objU, &RoleSlice{objR})
	if err != nil {
		t.Fatal(err)
	}

	// привязываем два заказа
	objO := &Order{ID: 183}
	err = inst.SaveRelation(objU, &OrderSlice{objO})
	if err != nil {
		t.Fatal(err)
	}
	objO = &Order{ID: 167}
	err = inst.SaveRelation(objU, &OrderSlice{objO})
	if err != nil {
		t.Fatal(err)
	}
}

func loadRelation(t *testing.T, inst *stbb.Instance) {
	t.Log("loadRelation")
	var err error
	objU := &User{ID: 111}

	// получаем привязанные роли
	var roles RoleSlice
	err = inst.LoadRelation(objU, &roles)
	if err != nil {
		t.Fatal(err)
	}
	for i := range roles {
		t.Log(roles[i])
	}

	// получаем привязанные заказы
	var orders OrderSlice
	err = inst.LoadRelation(objU, &orders)
	if err != nil {
		t.Fatal(err)
	}
	for i := range orders {
		t.Log(orders[i])
	}

	// Также проверяем связь от потомков. Или с обратной стороны

	// Роль
	var users1 UserSlice
	objR := &Role{ID: 143}
	err = inst.LoadRelation(objR, &users1)
	if err != nil {
		t.Fatal(err)
	}
	for i := range users1 {
		t.Log(users1[i])
	}

	// Заказ
	var users2 UserSlice
	objO := &Order{ID: 183}
	err = inst.LoadRelation(objO, &users2)
	if err != nil {
		t.Fatal(err)
	}
	for i := range users2 {
		t.Log(users2[i])
	}
}

func deleteRelation(t *testing.T, inst *stbb.Instance) {
	t.Log("deleteRelation")
	var err error
	objU := &User{ID: 111}

	var roles RoleSlice
	err = inst.DeleteRelation(objU, &roles)
	if err != nil {
		t.Fatal(err)
	}

	var orders OrderSlice
	err = inst.DeleteRelation(objU, &orders)
	if err != nil {
		t.Fatal(err)
	}
}
