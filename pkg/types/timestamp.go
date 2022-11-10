package types

import (
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
	intval, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}
	t.Time = time.Unix(int64(intval), 0)
	return nil
}
