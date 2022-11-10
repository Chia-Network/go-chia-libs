package types

import (
	"time"
)

// Timestamp Helper type to go to/from uint64 timestamps in json but represent as time.Time in go applications
type Timestamp time.Time

// @TODO Unmarshal
// @TODO Marshal
