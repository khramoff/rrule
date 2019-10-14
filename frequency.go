package rrule

import (
	"encoding/json"
	"errors"
	"github.com/spf13/cast"
	"log"
	"strconv"
)

// Frequency defines a set of constants for a base factor for how often recurrences happen.
type Frequency int

// String returns the RFC 5545 string for supported frequencies, and panics otherwise.
func (f Frequency) String() string {
	switch f {
	case Secondly:
		return "SECONDLY"
	case Minutely:
		return "MINUTELY"
	case Hourly:
		return "HOURLY"
	case Daily:
		return "DAILY"
	case Weekly:
		return "WEEKLY"
	case Monthly:
		return "MONTHLY"
	case Yearly:
		return "YEARLY"
	}
	log.Panicf("%d is not a supported frequency constant", f)
	return ""
}

// Frequencies specified in RFC 5545.
const (
	Secondly Frequency = iota
	Minutely
	Hourly
	Daily
	Weekly
	Monthly
	Yearly
)

func (f Frequency) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(f))
}

func (d *Frequency) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case int,int32,float64,float32,int64:
		*d = Frequency(cast.ToInt(value))
		return nil
	case string:
		var err error
		i, err := strconv.Atoi(v.(string))
		if err != nil {
			return err
		}
		*d = Frequency(i)
		return nil
	default:
		return errors.New("invalid frequency")
	}
}
