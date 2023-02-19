package test

import (
	"fmt"
	"testing"

	"github.com/kshamiev/urok/database/bbolt/stbb"
)

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

}

func loadRelation(t *testing.T, inst *stbb.Instance) {

}

func deleteRelation(t *testing.T, inst *stbb.Instance) {

}
