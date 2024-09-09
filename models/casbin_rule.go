package models

import (
	"database/sql"
	"encoding/json"
)

// MyNullString opakowuje sql.NullString i dodaje metodę UnmarshalJSON
type MyNullString struct {
	sql.NullString
}

// UnmarshalJSON konwertuje JSON na MyNullString
func (ns *MyNullString) UnmarshalJSON(b []byte) error {
	var str *string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}

	if str == nil {
		ns.Valid = false
		ns.String = ""
		return nil
	}

	ns.String = *str
	ns.Valid = true
	return nil
}

// CasbinRule reprezentuje regułę Casbin z potencjalnymi wartościami NULL
type CasbinRule struct {
	ID    int          `json:"id"`
	Ptype string       `json:"ptype"`
	V0    MyNullString `json:"v0"`
	V1    MyNullString `json:"v1"`
	V2    MyNullString `json:"v2"`
	V3    MyNullString `json:"v3"`
	V4    MyNullString `json:"v4"`
	V5    MyNullString `json:"v5"`
}
