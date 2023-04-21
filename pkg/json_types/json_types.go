package json_types

import (
	"strings"
	"time"
)

type JsonTimestamp time.Time

type JsonDate time.Time

type JsonTime time.Time

const FORMAT_STR_DATE = "2006-01-02"
const FORMAT_STR_TIMESTAMP = "2006-01-02T15:04:05"
const FORMAT_STR_TIME = "15:04:05"

func (d JsonDate) MarshalJSON() ([]byte, error) {
	if time.Time(d).IsZero() {
		var ret []byte
		return ret, nil
	} else {
		return []byte(`"` + time.Time(d).Format(FORMAT_STR_DATE) + `"`), nil
	}
}

func (d *JsonDate) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse(FORMAT_STR_DATE, value) //parse time
	if err != nil {
		return err
	}
	*d = JsonDate(t)
	return nil
}

func (d JsonTime) MarshalJSON() ([]byte, error) {
	if time.Time(d).IsZero() {
		var ret []byte
		return ret, nil
	} else {
		return []byte(`"` + time.Time(d).Format(FORMAT_STR_DATE) + `"`), nil
	}
}

func (d *JsonTime) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse(FORMAT_STR_TIME, value) //parse time
	if err != nil {
		return err
	}
	*d = JsonTime(t)
	return nil
}

func (d JsonTimestamp) MarshalJSON() ([]byte, error) {
	if time.Time(d).IsZero() {
		var ret []byte
		return ret, nil
	} else {
		return []byte(`"` + time.Time(d).Format(FORMAT_STR_TIMESTAMP) + `"`), nil
	}
}

func (d *JsonTimestamp) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse(FORMAT_STR_TIMESTAMP, value) //parse time
	if err != nil {
		return err
	}
	*d = JsonTimestamp(t)
	return nil
}
