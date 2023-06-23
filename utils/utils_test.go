package utils

import "testing"

func TestStructToMap(t *testing.T) {
	type Query struct {
		Name string `json:"courseName"`
		No   string `json:"lessonNo"`
		Code string `json:"courseCode"`
	}
	m := StructToMap(Query{})
	t.Log(m)
}
