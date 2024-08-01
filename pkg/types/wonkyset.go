package types

// WonkySet is just an alias for map[string]string that will unmarshal correctly from the empty forms {} and []
// In chia-blockchain, the default for some sets was incorrectly [] in the initial config
// so this ensures compatability with both ways the empty values could show up
type WonkySet map[string]string

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (ws *WonkySet) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// Attempt to unmarshal into a slice of strings.
	var slice []string
	if err := unmarshal(&slice); err == nil {
		*ws = make(map[string]string)
		return nil
	}

	// Attempt to unmarshal into a map.
	var m map[string]string
	if err := unmarshal(&m); err != nil {
		return err
	}

	*ws = m
	return nil
}
