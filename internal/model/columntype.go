package model

import (
	"database/sql"
	"database/sql/driver"
)

type (
	NullTime sql.NullTime
)

func (nt NullTime) MarshalJSON() ([]byte, error) {
	if nt.Valid {
		return nt.Time.MarshalJSON()
	}
	return []byte("null"), nil
}

func (nt NullTime) Value() (driver.Value, error) {
	return sql.NullTime(nt).Value()
}

func (nt *NullTime) Scan(value interface{}) error {
	return (*sql.NullTime)(nt).Scan(value)
}
