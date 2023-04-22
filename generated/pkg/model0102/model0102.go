// Attention, this file is generated. Manual changes get lost with the next
// run of the code generation.
// created by yacg (template: golang_types.mako v1.1.0)
package model0102


import (
    uuid "github.com/google/uuid"
    json_types "oth.types/pkg/json_types"
    optional "oth.types/pkg/optional_types"
	encJson "encoding/json"
    "reflect"
)


/* An example schema that defines complex types containing basic types, dates,
    timestamps, and UUIDs. Includes examples of complex types with array and
    dictionary properties of complex and basic types.
*/
type Model0102 struct {

    // A string property
    String_prop optional.Optional[string]

    // A numeric property
    Number_prop float64

    // An integer property
    Integer_prop optional.Optional[int32]

    // A boolean property
    Boolean_prop optional.Optional[bool]

    // A date property
    Date_prop optional.Optional[json_types.JsonDate]

    // A timestamp property
    Timestamp_prop optional.Optional[json_types.JsonTimestamp]

    // A UUID property
    Uuid_prop optional.Optional[uuid.UUID]

    // An array property
    Array_prop []Model0102Array_prop

    // A dictionary property
    Dict_prop map[string]string
}

// Creates a Model0102 object
func MakeModel0102() Model0102 {
    var ret Model0102
    // TODO: initialize default values
    return ret
}

func (v Model0102) Equals(o Model0102) bool {
    if v.String_prop != o.String_prop {
        return false
    }
    if v.Number_prop != o.Number_prop {
        return false
    }
    if v.Integer_prop != o.Integer_prop {
        return false
    }
    if v.Boolean_prop != o.Boolean_prop {
        return false
    }
    if v.Date_prop != o.Date_prop {
        return false
    }
    if v.Timestamp_prop != o.Timestamp_prop {
        return false
    }
    if v.Uuid_prop != o.Uuid_prop {
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

type OptionalModel0102 struct {
	Value Model0102
	IsSet bool
}

// Creates a OptionalModel0102 object
func MakeOptionalModel0102() OptionalModel0102 {
    var ret OptionalModel0102
    // TODO: initialize default values
    return ret
}

func (m OptionalModel0102) Set(v Model0102) {
	m.Value = v
	m.IsSet = true
}

func (m OptionalModel0102) UnSet() {
	m.IsSet = false
}

func (v Model0102) MarshalJSON() ([]byte, error) {
    var _String_prop string
	if v.String_prop.IsSet {
		_String_prop = v.String_prop.Value
	}
    var _Integer_prop int32
	if v.Integer_prop.IsSet {
		_Integer_prop = v.Integer_prop.Value
	}
    var _Boolean_prop *bool
	if v.Boolean_prop.IsSet {
		_Boolean_prop = &v.Boolean_prop.Value
	}
    var _Date_prop *json_types.JsonDate
	if v.Date_prop.IsSet {
		_Date_prop = &v.Date_prop.Value
	}
    var _Timestamp_prop *json_types.JsonTimestamp
	if v.Timestamp_prop.IsSet {
		_Timestamp_prop = &v.Timestamp_prop.Value
	}
    var _Uuid_prop *uuid.UUID
	if v.Uuid_prop.IsSet {
		_Uuid_prop = &v.Uuid_prop.Value
	}
    var _Array_prop *[]Model0102Array_prop
	if len(v.Array_prop) > 0 {
		_Array_prop = &v.Array_prop
	}
    var _Dict_prop *map[string]string
	if len(v.Dict_prop) > 0 {
		_Dict_prop = &v.Dict_prop
	}

	return encJson.Marshal(&struct {
        String_prop string `json:"string_prop,omitempty"`
        Number_prop float64 `json:"number_prop"`
        Integer_prop int32 `json:"integer_prop,omitempty"`
        Boolean_prop *bool `json:"boolean_prop,omitempty"`
        Date_prop *json_types.JsonDate `json:"date_prop,omitempty"`
        Timestamp_prop *json_types.JsonTimestamp `json:"timestamp_prop,omitempty"`
        Uuid_prop *uuid.UUID `json:"uuid_prop,omitempty"`
        Array_prop *[]Model0102Array_prop `json:"array_prop,omitempty"`
        Dict_prop *map[string]string `json:"dict_prop,omitempty"`
	}{
        String_prop: _String_prop,
        Number_prop: v.Number_prop,
        Integer_prop: _Integer_prop,
        Boolean_prop: _Boolean_prop,
        Date_prop: _Date_prop,
        Timestamp_prop: _Timestamp_prop,
        Uuid_prop: _Uuid_prop,
        Array_prop: _Array_prop,
        Dict_prop: _Dict_prop,
	})
}

func (v OptionalModel0102) MarshalJSON() ([]byte, error) {
	if v.IsSet {
		return encJson.Marshal(v.Value)
	} else {
		return []byte("null"), nil
	}
}



type Model0102Array_prop struct {

    // A string property in an array
    Array_string_prop optional.Optional[string]

    // A numeric property in an array
    Array_number_prop optional.Optional[float64]
}

// Creates a Model0102Array_prop object
func MakeModel0102Array_prop() Model0102Array_prop {
    var ret Model0102Array_prop
    // TODO: initialize default values
    return ret
}

func (v Model0102Array_prop) Equals(o Model0102Array_prop) bool {
    if v.Array_string_prop != o.Array_string_prop {
        return false
    }
    if v.Array_number_prop != o.Array_number_prop {
        return false
    }
	return true
}

type OptionalModel0102Array_prop struct {
	Value Model0102Array_prop
	IsSet bool
}

// Creates a OptionalModel0102Array_prop object
func MakeOptionalModel0102Array_prop() OptionalModel0102Array_prop {
    var ret OptionalModel0102Array_prop
    // TODO: initialize default values
    return ret
}

func (m OptionalModel0102Array_prop) Set(v Model0102Array_prop) {
	m.Value = v
	m.IsSet = true
}

func (m OptionalModel0102Array_prop) UnSet() {
	m.IsSet = false
}

func (v Model0102Array_prop) MarshalJSON() ([]byte, error) {
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

func (v OptionalModel0102Array_prop) MarshalJSON() ([]byte, error) {
	if v.IsSet {
		return encJson.Marshal(v.Value)
	} else {
		return []byte("null"), nil
	}
}




