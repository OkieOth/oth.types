// Attention, this file is generated. Manual changes get lost with the next
// run of the code generation.
// created by yacg (template: golang_types_tests.mako v1.1.0)
package model0202

import (
	encJson "encoding/json"
	"testing"
    "reflect"
	json_helper "oth.types/pkg/json_helper"
)

func TestJsonType1(t *testing.T) {
	var x []Type1
	err := json_helper.LoadOneObjFromFile(&x, "../../../configs/models/example_data/0202/Type1.json")
	if err != nil {
		t.Errorf("error while loading bundle file: 0202/Type1.json, error: %v", err)
		return
	}
	bytes, err := encJson.Marshal(x)
	if err != nil {
		t.Errorf("error while json serialize type: Type1, error: %v", err)
		return
	}
	var y []Type1

	err2 := encJson.Unmarshal(bytes, &y)
	if err2 != nil {
		t.Errorf("error while json unmarshal type: Type1, error: %v", err)
		return
	}

	if !reflect.DeepEqual(x, y) {
		t.Errorf("objects are not equal after json marshal/unmarshal, type: Type1")
	}
}

func TestJsonType2(t *testing.T) {
	var x []Type2
	err := json_helper.LoadOneObjFromFile(&x, "../../../configs/models/example_data/0202/Type2.json")
	if err != nil {
		t.Errorf("error while loading bundle file: 0202/Type2.json, error: %v", err)
		return
	}
	bytes, err := encJson.Marshal(x)
	if err != nil {
		t.Errorf("error while json serialize type: Type2, error: %v", err)
		return
	}
	var y []Type2

	err2 := encJson.Unmarshal(bytes, &y)
	if err2 != nil {
		t.Errorf("error while json unmarshal type: Type2, error: %v", err)
		return
	}

	if !reflect.DeepEqual(x, y) {
		t.Errorf("objects are not equal after json marshal/unmarshal, type: Type2")
	}
}

func TestJsonType3(t *testing.T) {
	var x []Type3
	err := json_helper.LoadOneObjFromFile(&x, "../../../configs/models/example_data/0202/Type3.json")
	if err != nil {
		t.Errorf("error while loading bundle file: 0202/Type3.json, error: %v", err)
		return
	}
	bytes, err := encJson.Marshal(x)
	if err != nil {
		t.Errorf("error while json serialize type: Type3, error: %v", err)
		return
	}
	var y []Type3

	err2 := encJson.Unmarshal(bytes, &y)
	if err2 != nil {
		t.Errorf("error while json unmarshal type: Type3, error: %v", err)
		return
	}

	if !reflect.DeepEqual(x, y) {
		t.Errorf("objects are not equal after json marshal/unmarshal, type: Type3")
	}
}

func TestJsonType4(t *testing.T) {
	var x []Type4
	err := json_helper.LoadOneObjFromFile(&x, "../../../configs/models/example_data/0202/Type4.json")
	if err != nil {
		t.Errorf("error while loading bundle file: 0202/Type4.json, error: %v", err)
		return
	}
	bytes, err := encJson.Marshal(x)
	if err != nil {
		t.Errorf("error while json serialize type: Type4, error: %v", err)
		return
	}
	var y []Type4

	err2 := encJson.Unmarshal(bytes, &y)
	if err2 != nil {
		t.Errorf("error while json unmarshal type: Type4, error: %v", err)
		return
	}

	if !reflect.DeepEqual(x, y) {
		t.Errorf("objects are not equal after json marshal/unmarshal, type: Type4")
	}
}

