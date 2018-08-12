package sqltype

import (
	"database/sql/driver"
	"fmt"
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
	switch v := value.(type) {
	default:
		return fmt.Errorf("unable to parse NullDuration value for type %T", v)
	case time.Duration:
		nd.Duration = v
		nd.Valid = true
	case *time.Duration:
		if v != nil {
			nd.Duration = *v
			nd.Valid = true
		}
	case []uint8:
		if len(v) > 0 {
			d, err := time.ParseDuration(string(v))
			if err == nil {
				nd.Duration = d
				nd.Valid = true
			}
		}
	case string:
		d, err := time.ParseDuration(v)
		if err == nil {
			nd.Duration = d
			nd.Valid = true
		}
	case *string:
		if v != nil {
			d, err := time.ParseDuration(*v)
			if err == nil {
				nd.Duration = d
				nd.Valid = true
			}
		}
	}
	return nil
}

// Value implements the driver Valuer interface.
func (nd NullDuration) Value() (driver.Value, error) {
	if !nd.Valid {
		return nil, nil
	}
	return nd.Duration.String(), nil
}
