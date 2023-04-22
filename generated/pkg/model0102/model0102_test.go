// Attention, this file is generated. Manual changes get lost with the next
// run of the code generation.
// created by yacg (template: golang_types_tests.mako v1.1.0)
package model0102

import (
	"testing"
	encJson "encoding/json"
	json_helper "oth.types/pkg/json_helper"
)

func TestMakeModel0102(t *testing.T) {
	o1 :=MakeModel0102()
	o2 :=MakeModel0102()
	if ! o1.Equals(o2) {
		t.Error("two fresh created objects are not equal, type: Model0102")
		return
	}
}


func TestJsonModel0102(t *testing.T) {
	var x []Model0102
	err := json_helper.LoadOneObjFromFile(&x, "../../../configs/models/example_data/0102/Model0102.json")
	if err != nil {
		t.Errorf("error while loading bundle file: 0102/Model0102.json, error: %v", err)
		return
	}
	bytes, err := encJson.Marshal(x)
	if err != nil {
		t.Errorf("error while json serialize type: Model0102, error: %v", err)
		return
	}
	var y []Model0102

	err2 := encJson.Unmarshal(bytes, &y)
	if err2 != nil {
		t.Errorf("error while json unmarshal type: Model0102, error: %v", err)
		return
	}


	for i, value := range y {
		value2 := x[i]
		if !value.Equals(value2) {
			b1, _ := encJson.Marshal(value)
			b2, _ := encJson.Marshal(value2)
			t.Errorf("objects are not equal after json marshal/unmarshal, type: Model0102\n%s\n%s\n", b1, b2)
			return
		}
	}
}
func TestMakeModel0102Array_prop(t *testing.T) {
	o1 :=MakeModel0102Array_prop()
	o2 :=MakeModel0102Array_prop()
	if ! o1.Equals(o2) {
		t.Error("two fresh created objects are not equal, type: Model0102Array_prop")
		return
	}
}

func TestMakeOptionalModel0102Array_prop(t *testing.T) {
	o1 :=MakeOptionalModel0102Array_prop()
	o2 :=MakeOptionalModel0102Array_prop()
	if o1.IsSet || o2.IsSet {
		t.Error("Optional type objects have a value after creation, type: Model0102Array_prop")
		return
	}
	if ! o1.Value.Equals(o2.Value) {
		t.Error("two fresh created objects are not equal, type: Model0102Array_prop")
		return
	}
}

func TestJsonModel0102Array_prop(t *testing.T) {
	var x []Model0102Array_prop
	err := json_helper.LoadOneObjFromFile(&x, "../../../configs/models/example_data/0102/Model0102Array_prop.json")
	if err != nil {
		t.Errorf("error while loading bundle file: 0102/Model0102Array_prop.json, error: %v", err)
		return
	}
	bytes, err := encJson.Marshal(x)
	if err != nil {
		t.Errorf("error while json serialize type: Model0102Array_prop, error: %v", err)
		return
	}
	var y []Model0102Array_prop

	err2 := encJson.Unmarshal(bytes, &y)
	if err2 != nil {
		t.Errorf("error while json unmarshal type: Model0102Array_prop, error: %v", err)
		return
	}


	for i, value := range y {
		value2 := x[i]
		if !value.Equals(value2) {
			b1, _ := encJson.Marshal(value)
			b2, _ := encJson.Marshal(value2)
			t.Errorf("objects are not equal after json marshal/unmarshal, type: Model0102Array_prop\n%s\n%s\n", b1, b2)
			return
		}
	}
}
