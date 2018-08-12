package sqltype

import (
	"database/sql/driver"
	"time"
)

// NullDuration represents a time.Duration that may be null. NullDuration implements the
// sql.Scanner interface so it can be used as a scan destination, similar to
// sql.NullString.
type NullDuration struct {
	Duration time.Duration
	Valid    bool // Valid is true if Duration is not NULL
}

// Scan implements the Scanner interface.
func (nd *NullDuration) Scan(value interface{}) error {
	nd.Duration, nd.Valid = value.(time.Duration)
	return nil
}

// Value implements the driver Valuer interface.
func (nd NullDuration) Value() (driver.Value, error) {
	if !nd.Valid {
		return nil, nil
	}
	return nd.Duration, nil
}
