// Attention, this file is generated. Manual changes get lost with the next
// run of the code generation.
// created by yacg (template: golang_types.mako v1.1.0)
package model0201


import (
	encJson "encoding/json"
    "reflect"
)


type Type1 struct {

    String_prop string

    Number_prop float64

    Array_prop []string

    Dict_prop map[string]float64
}

// Creates a Type1 object
func MakeType1() Type1 {
    var ret Type1
    // TODO: initialize default values
    return ret
}

func (v Type1) Equals(o Type1) bool {
    if v.String_prop != o.String_prop {
        return false
    }
    if v.Number_prop != o.Number_prop {
        return false
    }
    if !reflect.DeepEqual(v.Array_prop, o.Array_prop) {
        return false
    }
    if len(v.Dict_prop) != len(o.Dict_prop) {
        return false
    }
    for key, vValue := range v.Dict_prop {
        oValue, exists := o.Dict_prop[key]
        if (!exists) || (!reflect.DeepEqual(oValue, vValue)) {
            return false
        }
    }
	return true
}


func (v Type1) MarshalJSON() ([]byte, error) {
    var _Array_prop *[]string
	if len(v.Array_prop) > 0 {
		_Array_prop = &v.Array_prop
	}
    var _Dict_prop *map[string]float64
	if len(v.Dict_prop) > 0 {
		_Dict_prop = &v.Dict_prop
	}

	return encJson.Marshal(&struct {
        String_prop string `json:"string_prop"`
        Number_prop float64 `json:"number_prop"`
        Array_prop *[]string `json:"array_prop,omitempty"`
        Dict_prop *map[string]float64 `json:"dict_prop,omitempty"`
	}{
        String_prop: v.String_prop,
        Number_prop: v.Number_prop,
        Array_prop: _Array_prop,
        Dict_prop: _Dict_prop,
	})
}





type Type2 struct {

    String_prop string

    Boolean_prop bool

    Array_prop []int32

    Dict_prop map[string]string
}

// Creates a Type2 object
func MakeType2() Type2 {
    var ret Type2
    // TODO: initialize default values
    return ret
}

func (v Type2) Equals(o Type2) bool {
    if v.String_prop != o.String_prop {
        return false
    }
    if v.Boolean_prop != o.Boolean_prop {
        return false
    }
    if !reflect.DeepEqual(v.Array_prop, o.Array_prop) {
        return false
    }
    if len(v.Dict_prop) != len(o.Dict_prop) {
        return false
    }
    for key, vValue := range v.Dict_prop {
        oValue, exists := o.Dict_prop[key]
        if (!exists) || (!reflect.DeepEqual(oValue, vValue)) {
            return false
        }
    }
	return true
}


func (v Type2) MarshalJSON() ([]byte, error) {
    var _Array_prop *[]int32
	if len(v.Array_prop) > 0 {
		_Array_prop = &v.Array_prop
	}
    var _Dict_prop *map[string]string
	if len(v.Dict_prop) > 0 {
		_Dict_prop = &v.Dict_prop
	}

	return encJson.Marshal(&struct {
        String_prop string `json:"string_prop"`
        Boolean_prop bool `json:"boolean_prop"`
        Array_prop *[]int32 `json:"array_prop,omitempty"`
        Dict_prop *map[string]string `json:"dict_prop,omitempty"`
	}{
        String_prop: v.String_prop,
        Boolean_prop: v.Boolean_prop,
        Array_prop: _Array_prop,
        Dict_prop: _Dict_prop,
	})
}





type Type3 struct {

    Type1_prop Type1

    Type2_prop Type2

    Type2_array_prop []Type2

    Type2_dict_prop map[string]Type2
}

// Creates a Type3 object
func MakeType3() Type3 {
    var ret Type3
    // TODO: initialize default values
    return ret
}

func (v Type3) Equals(o Type3) bool {
    if !reflect.DeepEqual(v.Type1_prop, o.Type1_prop) {
        return false
    }
    if !reflect.DeepEqual(v.Type2_prop, o.Type2_prop) {
        return false
    }
    if !reflect.DeepEqual(v.Type2_array_prop, o.Type2_array_prop) {
        return false
    }
    if len(v.Type2_dict_prop) != len(o.Type2_dict_prop) {
        return false
    }
    for key, vValue := range v.Type2_dict_prop {
        oValue, exists := o.Type2_dict_prop[key]
        if (!exists) || (!reflect.DeepEqual(oValue, vValue)) {
            return false
        }
    }
	return true
}


func (v Type3) MarshalJSON() ([]byte, error) {
    var _Type2_array_prop *[]Type2
	if len(v.Type2_array_prop) > 0 {
		_Type2_array_prop = &v.Type2_array_prop
	}
    var _Type2_dict_prop *map[string]Type2
	if len(v.Type2_dict_prop) > 0 {
		_Type2_dict_prop = &v.Type2_dict_prop
	}

	return encJson.Marshal(&struct {
        Type1_prop Type1 `json:"type1_prop"`
        Type2_prop Type2 `json:"type2_prop"`
        Type2_array_prop *[]Type2 `json:"type2_array_prop,omitempty"`
        Type2_dict_prop *map[string]Type2 `json:"type2_dict_prop,omitempty"`
	}{
        Type1_prop: v.Type1_prop,
        Type2_prop: v.Type2_prop,
        Type2_array_prop: _Type2_array_prop,
        Type2_dict_prop: _Type2_dict_prop,
	})
}





type Type4 struct {

    String_prop string

    Integer_prop int32

    Array_prop []float64

    Dict_prop map[string]int32
}

// Creates a Type4 object
func MakeType4() Type4 {
    var ret Type4
    // TODO: initialize default values
    return ret
}

func (v Type4) Equals(o Type4) bool {
    if v.String_prop != o.String_prop {
        return false
    }
    if v.Integer_prop != o.Integer_prop {
        return false
    }
    if !reflect.DeepEqual(v.Array_prop, o.Array_prop) {
        return false
    }
    if len(v.Dict_prop) != len(o.Dict_prop) {
        return false
    }
    for key, vValue := range v.Dict_prop {
        oValue, exists := o.Dict_prop[key]
        if (!exists) || (!reflect.DeepEqual(oValue, vValue)) {
            return false
        }
    }
	return true
}


func (v Type4) MarshalJSON() ([]byte, error) {
    var _Array_prop *[]float64
	if len(v.Array_prop) > 0 {
		_Array_prop = &v.Array_prop
	}
    var _Dict_prop *map[string]int32
	if len(v.Dict_prop) > 0 {
		_Dict_prop = &v.Dict_prop
	}

	return encJson.Marshal(&struct {
        String_prop string `json:"string_prop"`
        Integer_prop int32 `json:"integer_prop"`
        Array_prop *[]float64 `json:"array_prop,omitempty"`
        Dict_prop *map[string]int32 `json:"dict_prop,omitempty"`
	}{
        String_prop: v.String_prop,
        Integer_prop: v.Integer_prop,
        Array_prop: _Array_prop,
        Dict_prop: _Dict_prop,
	})
}




