// Attention, this file is generated. Manual changes get lost with the next
// run of the code generation.
// created by yacg (template: golang_types.mako v1.1.0)
package model0101


import (
    uuid "github.com/google/uuid"
    "time"
    optional "oth.types/pkg/optional_types"
	encJson "encoding/json"
)


/* An example schema that defines complex types containing basic types, dates,
    timestamps, and UUIDs. Includes examples of complex types with array and
    dictionary properties of complex and basic types.
*/
type Model0101 struct {

    // A string property
    String_prop string

    // A numeric property
    Number_prop float64

    // An integer property
    Integer_prop int32

    // A boolean property
    Boolean_prop bool

    // A date property
    Date_prop time.Time

    // A timestamp property
    Timestamp_prop time.Time

    // A UUID property
    Uuid_prop uuid.UUID

    // An array property
    Array_prop []Model0101Array_prop

    // A dictionary property
    Dict_prop map[string]string
}

// Creates a Model0101 object
func MakeModel0101() Model0101 {
    var ret Model0101
    // TODO: initialize default values
    return ret
}

type OptionalModel0101 struct {
	Value Model0101
	IsSet bool
}

// Creates a OptionalModel0101 object
func MakeOptionalModel0101() OptionalModel0101 {
    var ret OptionalModel0101
    // TODO: initialize default values
    return ret
}

func (m OptionalModel0101) Set(v Model0101) {
	m.Value = v
	m.IsSet = true
}

func (m OptionalModel0101) UnSet() {
	m.IsSet = false
}

func (v Model0101) MarshalJSON() ([]byte, error) {
    var _Array_prop *[]Model0101Array_prop
	if len(v.Array_prop) > 0 {
		_Array_prop = &v.Array_prop
	}
    var _Dict_prop *map[string]string
	if len(v.Dict_prop) > 0 {
		_Dict_prop = &v.Dict_prop
	}

	return encJson.Marshal(&struct {
        String_prop string `json:"string_prop"`
        Number_prop float64 `json:"number_prop"`
        Integer_prop int32 `json:"integer_prop"`
        Boolean_prop bool `json:"boolean_prop"`
        Date_prop time.Time `json:"date_prop"`
        Timestamp_prop time.Time `json:"timestamp_prop"`
        Uuid_prop uuid.UUID `json:"uuid_prop"`
        Array_prop *[]Model0101Array_prop `json:"array_prop,omitempty"`
        Dict_prop *map[string]string `json:"dict_prop,omitempty"`
	}{
        String_prop: v.String_prop,
        Number_prop: v.Number_prop,
        Integer_prop: v.Integer_prop,
        Boolean_prop: v.Boolean_prop,
        Date_prop: v.Date_prop,
        Timestamp_prop: v.Timestamp_prop,
        Uuid_prop: v.Uuid_prop,
        Array_prop: _Array_prop,
        Dict_prop: _Dict_prop,
	})
}

func (v OptionalModel0101) MarshalJSON() ([]byte, error) {
	if v.IsSet {
		return encJson.Marshal(v.Value)
	} else {
		return []byte("null"), nil
	}
}



type Model0101Array_prop struct {

    // A string property in an array
    Array_string_prop optional.Optional[string]

    // A numeric property in an array
    Array_number_prop optional.Optional[float64]
}

// Creates a Model0101Array_prop object
func MakeModel0101Array_prop() Model0101Array_prop {
    var ret Model0101Array_prop
    // TODO: initialize default values
    return ret
}

type OptionalModel0101Array_prop struct {
	Value Model0101Array_prop
	IsSet bool
}

// Creates a OptionalModel0101Array_prop object
func MakeOptionalModel0101Array_prop() OptionalModel0101Array_prop {
    var ret OptionalModel0101Array_prop
    // TODO: initialize default values
    return ret
}

func (m OptionalModel0101Array_prop) Set(v Model0101Array_prop) {
	m.Value = v
	m.IsSet = true
}

func (m OptionalModel0101Array_prop) UnSet() {
	m.IsSet = false
}

func (v Model0101Array_prop) MarshalJSON() ([]byte, error) {
    var _Array_string_prop string
	if v.Array_string_prop.IsSet {
		_Array_string_prop = v.Array_string_prop.Value
	}
    var _Array_number_prop float64
	if v.Array_number_prop.IsSet {
		_Array_number_prop = v.Array_number_prop.Value
	}

	return encJson.Marshal(&struct {
        Array_string_prop string `json:"array_string_prop,omitempty"`
        Array_number_prop float64 `json:"array_number_prop,omitempty"`
	}{
        Array_string_prop: _Array_string_prop,
        Array_number_prop: _Array_number_prop,
	})
}

func (v OptionalModel0101Array_prop) MarshalJSON() ([]byte, error) {
	if v.IsSet {
		return encJson.Marshal(v.Value)
	} else {
		return []byte("null"), nil
	}
}




