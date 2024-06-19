package config

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// FillValuesFromEnvironment reads environment variables starting with `chia.` and edits the config based on the config path
// chia.selected_network=mainnet would set the top level `selected_network: mainnet`
// chia.full_node.port=8444 would set full_node.port to 8444
//
// # Complex data structures can be passed in as JSON strings and they will be parsed out into the datatype specified for the config prior to being inserted
//
// chia.network_overrides.constants.mainnet='{"GENESIS_CHALLENGE":"abc123","GENESIS_PRE_FARM_POOL_PUZZLE_HASH":"xyz789"}'
func (c *ChiaConfig) FillValuesFromEnvironment() {
	valuesToUpdate := getAllChiaVars()
	log.Printf("%+v", valuesToUpdate)
	for key, pAndV := range valuesToUpdate {
		log.Printf("Key is: %v", key)
		log.Printf("Path is: %v", pAndV.path)
		log.Printf("Value is: %s", pAndV.value)
	}
}

type pathAndValue struct {
	path  []string
	value string
}

func getAllChiaVars() map[string]pathAndValue {
	// Most shells don't allow `.` in env names, but docker will and its easier to visualize the `.`, so support both
	// `.` and `__` as valid path segment separators
	// chia.full_node.port
	// chia__full_node__port
	separators := []string{".", "__"}
	envVars := os.Environ()
	finalVars := map[string]pathAndValue{}

	for _, sep := range separators {
		prefix := fmt.Sprintf("chia%s", sep)
		for _, env := range envVars {
			if strings.HasPrefix(env, prefix) {
				pair := strings.SplitN(env, "=", 2)
				if len(pair) == 2 {
					finalVars[pair[0][len(prefix):]] = pathAndValue{
						path:  strings.Split(pair[0], sep)[1:], // This is the path in the config to the value to edit minus the "chia" prefix
						value: pair[1],
					}
				}
			}
		}
	}

	return finalVars
}
