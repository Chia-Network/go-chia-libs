package types

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

// Timestamp Helper type to go to/from uint64 timestamps in json but represent as time.Time in go applications
type Timestamp struct {
	time.Time
}

// MarshalJSON marshals the time to unix timestamp
func (t Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", t.Time.Unix())), nil
}

// UnmarshalJSON unmarshals from uint64
func (t *Timestamp) UnmarshalJSON(data []byte) error {
	// Check if the data is the JSON "null"
	if string(data) == "null" {
		// Handle the null case
		*t = Timestamp{}
		return nil
	}

	// Split accounts for unix timestamps in float form (1668050986.646834)
	// In these cases, we just parse the seconds and ignore the decimal
	splits := bytes.Split(data, []byte(`.`))
	intval, err := strconv.Atoi(string(splits[0]))
	if err != nil {
		return err
	}
	t.Time = time.Unix(int64(intval), 0)
	return nil
}
