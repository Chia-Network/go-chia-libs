package config

import (
	"gopkg.in/yaml.v3"
)

// MarshalYAML marshals the NetworkOverrides value to yaml handling anchors where necessary
func (u *NetworkOverrides) MarshalYAML() (interface{}, error) {
	if u.YamlAnchor != nil {
		node := &yaml.Node{
			Kind:  yaml.AliasNode,
			Alias: u.YamlAnchor,
			Value: "network_overrides",
		}
		return node, nil
	}

	// Marshal the struct to a yaml.Node
	var node yaml.Node
	if err := node.Encode(*u); err != nil {
		return nil, err
	}
	node.Anchor = "network_overrides"

	// Store the node as the anchor for future iterations
	u.YamlAnchor = &node

	// Return the node to be marshalled
	return &node, nil
}
