// Attention, this file is generated. Manual changes get lost with the next
// run of the code generation.
// created by yacg (template: golang_types_tests.mako v1.1.0)
package model0201

import (
	"testing"
	encJson "encoding/json"
	json_helper "oth.types/pkg/json_helper"
)

func TestMakeType1(t *testing.T) {
	o1 :=MakeType1()
	o2 :=MakeType1()
	if ! o1.Equals(o2) {
		t.Error("two fresh created objects are not equal, type: Type1")
		return
	}
}

func TestMakeOptionalType1(t *testing.T) {
	o1 :=MakeOptionalType1()
	o2 :=MakeOptionalType1()
	if o1.IsSet || o2.IsSet {
		t.Error("Optional type objects have a value after creation, type: Type1")
		return
	}
	if ! o1.Value.Equals(o2.Value) {
		t.Error("two fresh created objects are not equal, type: Type1")
		return
	}
}

func TestJsonType1(t *testing.T) {
	var x []Type1
	err := json_helper.LoadOneObjFromFile(&x, "../../../configs/models/example_data/0201/Type1.json")
	if err != nil {
		t.Errorf("error while loading bundle file: 0201/Type1.json, error: %v", err)
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


	for i, value := range y {
		value2 := x[i]
		if !value.Equals(value2) {
			b1, _ := encJson.Marshal(value)
			b2, _ := encJson.Marshal(value2)
			t.Errorf("objects are not equal after json marshal/unmarshal, type: Type1\n%s\n%s\n", b1, b2)
			return
		}
	}
}
func TestMakeType2(t *testing.T) {
	o1 :=MakeType2()
	o2 :=MakeType2()
	if ! o1.Equals(o2) {
		t.Error("two fresh created objects are not equal, type: Type2")
		return
	}
}

func TestMakeOptionalType2(t *testing.T) {
	o1 :=MakeOptionalType2()
	o2 :=MakeOptionalType2()
	if o1.IsSet || o2.IsSet {
		t.Error("Optional type objects have a value after creation, type: Type2")
		return
	}
	if ! o1.Value.Equals(o2.Value) {
		t.Error("two fresh created objects are not equal, type: Type2")
		return
	}
}

func TestJsonType2(t *testing.T) {
	var x []Type2
	err := json_helper.LoadOneObjFromFile(&x, "../../../configs/models/example_data/0201/Type2.json")
	if err != nil {
		t.Errorf("error while loading bundle file: 0201/Type2.json, error: %v", err)
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


	for i, value := range y {
		value2 := x[i]
		if !value.Equals(value2) {
			b1, _ := encJson.Marshal(value)
			b2, _ := encJson.Marshal(value2)
			t.Errorf("objects are not equal after json marshal/unmarshal, type: Type2\n%s\n%s\n", b1, b2)
			return
		}
	}
}
func TestMakeType3(t *testing.T) {
	o1 :=MakeType3()
	o2 :=MakeType3()
	if ! o1.Equals(o2) {
		t.Error("two fresh created objects are not equal, type: Type3")
		return
	}
}

func TestMakeOptionalType3(t *testing.T) {
	o1 :=MakeOptionalType3()
	o2 :=MakeOptionalType3()
	if o1.IsSet || o2.IsSet {
		t.Error("Optional type objects have a value after creation, type: Type3")
		return
	}
	if ! o1.Value.Equals(o2.Value) {
		t.Error("two fresh created objects are not equal, type: Type3")
		return
	}
}

func TestJsonType3(t *testing.T) {
	var x []Type3
	err := json_helper.LoadOneObjFromFile(&x, "../../../configs/models/example_data/0201/Type3.json")
	if err != nil {
		t.Errorf("error while loading bundle file: 0201/Type3.json, error: %v", err)
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


	for i, value := range y {
		value2 := x[i]
		if !value.Equals(value2) {
			b1, _ := encJson.Marshal(value)
			b2, _ := encJson.Marshal(value2)
			t.Errorf("objects are not equal after json marshal/unmarshal, type: Type3\n%s\n%s\n", b1, b2)
			return
		}
	}
}
func TestMakeType4(t *testing.T) {
	o1 :=MakeType4()
	o2 :=MakeType4()
	if ! o1.Equals(o2) {
		t.Error("two fresh created objects are not equal, type: Type4")
		return
	}
}

func TestMakeOptionalType4(t *testing.T) {
	o1 :=MakeOptionalType4()
	o2 :=MakeOptionalType4()
	if o1.IsSet || o2.IsSet {
		t.Error("Optional type objects have a value after creation, type: Type4")
		return
	}
	if ! o1.Value.Equals(o2.Value) {
		t.Error("two fresh created objects are not equal, type: Type4")
		return
	}
}

func TestJsonType4(t *testing.T) {
	var x []Type4
	err := json_helper.LoadOneObjFromFile(&x, "../../../configs/models/example_data/0201/Type4.json")
	if err != nil {
		t.Errorf("error while loading bundle file: 0201/Type4.json, error: %v", err)
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


	for i, value := range y {
		value2 := x[i]
		if !value.Equals(value2) {
			b1, _ := encJson.Marshal(value)
			b2, _ := encJson.Marshal(value2)
			t.Errorf("objects are not equal after json marshal/unmarshal, type: Type4\n%s\n%s\n", b1, b2)
			return
		}
	}
}
