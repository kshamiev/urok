package test

import "strconv"

type testStruct struct {
	X int
	Y string
}

// func (t *testStruct) ToJSON() ([]byte, error) {
// 	return json.Marshal(t)
// }

func (t *testStruct) ToJSON() ([]byte, error) {
	return []byte(`{"X": ` + strconv.Itoa(t.X) + `, "Y": "` + t.Y + `"}`), nil
}
