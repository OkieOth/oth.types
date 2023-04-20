package jsonhelper

import (
	"testing"
)

func TestLoadOneObjFromFile(t *testing.T) {
	var s string
	err := LoadOneObjFromFile(&s, "../../tests/resources/simple.json")
	if err != nil {
		t.Errorf("error while loading bundle file")
		return
	}
}
