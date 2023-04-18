// Attention, this file is generated. Manual changes get lost with the next
// run of the code generation.
// created by yacg (template: golang_types.mako v1.1.0)
package model0201


import (
	encJson "encoding/json"
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

type OptionalType1 struct {
	Value Type1
	IsSet bool
}

// Creates a OptionalType1 object
func MakeOptionalType1() OptionalType1 {
    var ret OptionalType1
    // TODO: initialize default values
    return ret
}

func (m OptionalType1) Set(v Type1) {
	m.Value = v
	m.IsSet = true
}

func (m OptionalType1) UnSet() {
	m.IsSet = false
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

func (v OptionalType1) MarshalJSON() ([]byte, error) {
	if v.IsSet {
		return encJson.Marshal(v.Value)
	} else {
		return []byte("null"), nil
	}
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

type OptionalType2 struct {
	Value Type2
	IsSet bool
}

// Creates a OptionalType2 object
func MakeOptionalType2() OptionalType2 {
    var ret OptionalType2
    // TODO: initialize default values
    return ret
}

func (m OptionalType2) Set(v Type2) {
	m.Value = v
	m.IsSet = true
}

func (m OptionalType2) UnSet() {
	m.IsSet = false
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

func (v OptionalType2) MarshalJSON() ([]byte, error) {
	if v.IsSet {
		return encJson.Marshal(v.Value)
	} else {
		return []byte("null"), nil
	}
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

type OptionalType3 struct {
	Value Type3
	IsSet bool
}

// Creates a OptionalType3 object
func MakeOptionalType3() OptionalType3 {
    var ret OptionalType3
    // TODO: initialize default values
    return ret
}

func (m OptionalType3) Set(v Type3) {
	m.Value = v
	m.IsSet = true
}

func (m OptionalType3) UnSet() {
	m.IsSet = false
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

func (v OptionalType3) MarshalJSON() ([]byte, error) {
	if v.IsSet {
		return encJson.Marshal(v.Value)
	} else {
		return []byte("null"), nil
	}
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

type OptionalType4 struct {
	Value Type4
	IsSet bool
}

// Creates a OptionalType4 object
func MakeOptionalType4() OptionalType4 {
    var ret OptionalType4
    // TODO: initialize default values
    return ret
}

func (m OptionalType4) Set(v Type4) {
	m.Value = v
	m.IsSet = true
}

func (m OptionalType4) UnSet() {
	m.IsSet = false
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

func (v OptionalType4) MarshalJSON() ([]byte, error) {
	if v.IsSet {
		return encJson.Marshal(v.Value)
	} else {
		return []byte("null"), nil
	}
}



