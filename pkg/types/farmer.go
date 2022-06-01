package types

// EventFarmerSubmittedPartial is the event data for `submitted_partial` from the farmer
type EventFarmerSubmittedPartial struct {
	LauncherID                   string `json:"launcher_id"`
	PoolURL                      string `json:"pool_url"`
	CurrentDifficulty            uint64 `json:"current_difficulty"`
	PointsAcknowledgedSinceStart uint64 `json:"points_acknowledged_since_start"`
}
