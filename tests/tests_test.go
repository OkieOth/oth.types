package tests

import (
	encJson "encoding/json"
	"testing"

	types "oth.types/generated/pkg/model0101"
	json_helper "oth.types/pkg/json_helper"
	optional_types "oth.types/pkg/optional_types"
)

func TestOptionalString(t *testing.T) {
	var s optional_types.Optional[string]
	if s.IsSet {
		t.Errorf("uninitialized value is set after creation")
	}
}

func TestBundleLoad(t *testing.T) {
	var x types.Model0101
	err := json_helper.LoadOneObjFromFile(&x, "../configs/models/example_data/0101/Model0101_single.json")
	if err != nil {
		t.Errorf("error while loading bundle file")
		return
	}
	bytes, err := x.MarshalJSON()
	if err != nil {
		t.Errorf("error while json serialize solution")
		return
	}
	var y types.Model0101

	err2 := encJson.Unmarshal(bytes, &y)
	if err2 != nil {
		t.Errorf("error while json unmarshal solution again")
		return
	}

	if !x.Equals(y) {
		t.Errorf(":-(")
		return
	}
}
