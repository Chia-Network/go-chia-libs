package util

import (
	"fmt"
	"log"

	"github.com/chia-network/go-chia-libs/pkg/types"
)

// FormatBytes takes bytes as input and outputs a human friendly version
func FormatBytes(bytes types.Uint128) string {
	labels := []string{"MiB", "GiB", "TiB", "PiB", "EiB", "ZiB", "YiB"}
	base := uint64(1024)

	value := bytes.Div64(base)
	log.Printf("%s %s\n", value.String(), "KiB")
	for _, label := range labels {
		if value.FitsInUint64() {
			valueUint64 := float64(value.Uint64()) / float64(base)
			if valueUint64 < float64(base) {
				return fmt.Sprintf("%.3f %s", valueUint64, label)
			}

			// We always start using the Uint128 every iteration, so _now_ we can do this math so it's ready for next time
			value = value.Div64(base)
		} else {
			value = value.Div64(base)
			if value.Cmp64(base) == -1 {
				return fmt.Sprintf("%s %s", value.String(), label)
			}
		}
	}

	return fmt.Sprintf("%s %s", value.String(), labels[len(labels)-1])
}
