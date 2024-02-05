package model

import (
	"database/sql"
	"database/sql/driver"
	"time"
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

func NewNullTime(t time.Time) NullTime {
	return NullTime{Time: t, Valid: true}
}

func NewNullTimeNull() NullTime {
	var nt NullTime
	return nt
}
