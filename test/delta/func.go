package delta

import "strconv"

type testStruct struct {
	X int
	Y string
}

// old.txt
// func (t *testStruct) ToJSON() ([]byte, error) {
// 	return json.Marshal(t)
// }

// new.txt
func (t *testStruct) ToJSON() ([]byte, error) {
	return []byte(`{"X": ` + strconv.Itoa(t.X) + `, "Y": "` + t.Y + `"}`), nil
}
