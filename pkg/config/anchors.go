package config

import (
	"reflect"

	"gopkg.in/yaml.v3"
)

// MarshalYAML marshals the NetworkOverrides value to yaml handling anchors where necessary
func (nc *NetworkOverrides) MarshalYAML() (interface{}, error) {
	return anchorHelper(nc, "network_overrides")
}

// MarshalYAML marshals the LoggingConfig value to yaml handling anchors where necessary
func (lc *LoggingConfig) MarshalYAML() (interface{}, error) {
	return anchorHelper(lc, "logging")
}

// Anchorable defines the methods a type must implement to support anchors
type Anchorable interface {
	AnchorNode() *yaml.Node
	SetAnchorNode(*yaml.Node)
}

func anchorHelper(in Anchorable, tag string) (*yaml.Node, error) {
	if in.AnchorNode() != nil {
		node := &yaml.Node{
			Kind:  yaml.AliasNode,
			Alias: in.AnchorNode(),
			Value: tag,
		}
		return node, nil
	}

	// Get the underlying value of 'in' for marshalling or else we end up recursively in this function
	value := reflect.ValueOf(in)
	if value.Kind() == reflect.Ptr {
		// Dereference if it's a pointer
		value = value.Elem()
	}

	// Marshal the struct to a yaml.Node
	var node yaml.Node
	if err := node.Encode(value.Interface()); err != nil {
		return nil, err
	}
	node.Anchor = tag

	// Store the node as the anchor for future iterations
	in.SetAnchorNode(&node)

	// Return the node to be marshalled
	return &node, nil
}
