package types

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/volatiletech/sqlboiler/randomize"
)

// Time is a unix timestamp supported time.Time. It supports SQL and JSON serialization.
type Time struct {
	Time time.Time
	Unix int64
}

// NewTime creates a new Time.
func NewTime(t time.Time) Time {
	return Time{
		Time: t,
		Unix: t.Unix(),
	}
}

// TimeFrom creates a new Time
func TimeFrom(t time.Time) Time {
	return NewTime(t)
}

// TimeFromPtr creates a new Time
func TimeFromPtr(t *time.Time) Time {
	if t == nil {
		return NewTime(time.Time{})
	}
	return NewTime(*t)
}

// MarshalJSON implements json.Marshaler.
func (t Time) MarshalJSON() ([]byte, error) {
	return t.Time.MarshalJSON()
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *Time) UnmarshalJSON(data []byte) error {
	if err := t.Time.UnmarshalJSON(data); err != nil {
		return err
	}

	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (t Time) MarshalText() ([]byte, error) {
	return t.Time.MarshalText()
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (t *Time) UnmarshalText(text []byte) error {
	if text == nil || len(text) == 0 {
		return nil
	}
	if err := t.Time.UnmarshalText(text); err != nil {
		return err
	}
	return nil
}

// Ptr returns a pointer to this Time's value, or a nil pointer if this Time is null.
func (t Time) Ptr() *time.Time {
	return &t.Time
}

// Scan implements the Scanner interface.
func (t *Time) Scan(value interface{}) error {
	var err error
	switch x := value.(type) {
	case time.Time:
		t.Time = x
		t.Unix = x.Unix()
	case nil:
		return nil
	default:
		err = fmt.Errorf("null: cannot scan type %T into time.Time: %v", value, value)
	}
	return err
}

// Value implements the driver Valuer interface.
func (t Time) Value() (driver.Value, error) {
	return t.Time, nil
}

// Randomize for sqlboiler
func (t *Time) Randomize(nextInt func() int64, fieldType string, shouldBeNull bool) {
	if shouldBeNull {
		t.Time = time.Time{}
		t.Unix = t.Time.Unix()
	} else {
		t.Time = randomize.Date(nextInt)
		t.Unix = t.Time.Unix()
	}
}
