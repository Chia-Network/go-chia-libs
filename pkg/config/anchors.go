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

	// Create a new node for this instance
	node := &yaml.Node{
		Kind:   yaml.MappingNode,
		Anchor: tag,
	}

	// Encode the struct to the node
	if err := node.Encode(reflect.ValueOf(in).Elem().Interface()); err != nil {
		return nil, err
	}

	// Store the node as the anchor for future iterations
	in.SetAnchorNode(node)

	return node, nil
}
