// Attention, this file is generated. Manual changes get lost with the next
// run of the code generation.
// created by yacg (template: golang_types_tests.mako v1.1.0)
package model0101

import (
	encJson "encoding/json"
	"testing"
	json_helper "oth.types/pkg/json_helper"
)

func TestJsonModel0101(t *testing.T) {
	var x []Model0101
	err := json_helper.LoadOneObjFromFile(&x, "../../../configs/models/example_data/0101/Model0101.json")
	if err != nil {
		t.Errorf("error while loading bundle file: 0101/Model0101.json, error: %v", err)
		return
	}
	bytes, err := encJson.Marshal(x)
	if err != nil {
		t.Errorf("error while json serialize type: Model0101, error: %v", err)
		return
	}
	var y []Model0101

	err2 := encJson.Unmarshal(bytes, &y)
	if err2 != nil {
		t.Errorf("error while json unmarshal type: Model0101, error: %v", err)
		return
	}


	for i, value := range y {
		value2 := x[i]
		if !value.Equals(value2) {
			b1, _ := encJson.Marshal(value)
			b2, _ := encJson.Marshal(value2)
			t.Errorf("objects are not equal after json marshal/unmarshal, type: Model0101\n%s\n%s\n", b1, b2)
			return
		}
	}
}

func TestJsonModel0101Array_prop(t *testing.T) {
	var x []Model0101Array_prop
	err := json_helper.LoadOneObjFromFile(&x, "../../../configs/models/example_data/0101/Model0101Array_prop.json")
	if err != nil {
		t.Errorf("error while loading bundle file: 0101/Model0101Array_prop.json, error: %v", err)
		return
	}
	bytes, err := encJson.Marshal(x)
	if err != nil {
		t.Errorf("error while json serialize type: Model0101Array_prop, error: %v", err)
		return
	}
	var y []Model0101Array_prop

	err2 := encJson.Unmarshal(bytes, &y)
	if err2 != nil {
		t.Errorf("error while json unmarshal type: Model0101Array_prop, error: %v", err)
		return
	}


	for i, value := range y {
		value2 := x[i]
		if !value.Equals(value2) {
			b1, _ := encJson.Marshal(value)
			b2, _ := encJson.Marshal(value2)
			t.Errorf("objects are not equal after json marshal/unmarshal, type: Model0101Array_prop\n%s\n%s\n", b1, b2)
			return
		}
	}
}

