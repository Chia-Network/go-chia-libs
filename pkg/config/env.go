package config

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/chia-network/go-chia-libs/pkg/types"
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

		err := c.SetFieldByPath(pAndV.path, pAndV.value)
		if err != nil {
			log.Fatalln(err.Error())
		}
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

// SetFieldByPath iterates through each item in path to find the corresponding `yaml` tag in the struct
// Once found, we move to the next item in path and look for that key within the first element
// If any element is not found, an error will be returned
func (c *ChiaConfig) SetFieldByPath(path []string, value any) error {
	v := reflect.ValueOf(c).Elem()
	return setFieldByPath(v, path, value)
}

func setFieldByPath(v reflect.Value, path []string, value any) error {
	if len(path) == 0 {
		return fmt.Errorf("invalid path")
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		yamlTagRaw := field.Tag.Get("yaml")
		yamlTag := strings.Split(yamlTagRaw, ",")[0]

		if yamlTagRaw == ",inline" && field.Anonymous {
			// Check the inline struct
			if err := setFieldByPath(v.Field(i), path, value); err != nil {
				return err
			}
		} else if yamlTag == path[0] {
			// We found a match for the current level of "paths"
			// If we only have 1 element left in paths, then we can set the value
			// Otherwise, we can recursively call setFieldByPath again, with the remaining elements of path
			fieldValue := v.Field(i)
			if len(path) > 1 {
				if fieldValue.Kind() == reflect.Map {
					mapKey := reflect.ValueOf(path[1])
					if !mapKey.Type().ConvertibleTo(fieldValue.Type().Key()) {
						return fmt.Errorf("invalid map key type %s", mapKey.Type())
					}
					mapValue := fieldValue.MapIndex(mapKey)
					if mapValue.IsValid() {
						if !mapValue.CanSet() {
							// Create a new writable map and copy over the existing data
							newMapValue := reflect.New(fieldValue.Type().Elem()).Elem()
							newMapValue.Set(mapValue)
							mapValue = newMapValue
						}
						err := setFieldByPath(mapValue, path[2:], value)
						if err != nil {
							return err
						}
						fieldValue.SetMapIndex(mapKey, mapValue)
						return nil
					}
				} else {
					return setFieldByPath(fieldValue, path[1:], value)
				}
			}

			if !fieldValue.CanSet() {
				return fmt.Errorf("cannot set field %s", path[0])
			}

			// Special Cases
			if fieldValue.Type() == reflect.TypeOf(types.Uint128{}) {
				strValue, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected string for Uint128 field, got %T", value)
				}
				bigIntValue := new(big.Int)
				_, ok = bigIntValue.SetString(strValue, 10)
				if !ok {
					return fmt.Errorf("invalid string for big.Int: %s", strValue)
				}
				fieldValue.Set(reflect.ValueOf(types.Uint128FromBig(bigIntValue)))
				return nil
			}

			val := reflect.ValueOf(value)

			if fieldValue.Type() != val.Type() {
				if val.Type().ConvertibleTo(fieldValue.Type()) {
					val = val.Convert(fieldValue.Type())
				} else {
					convertedVal, err := convertValue(value, fieldValue.Type())
					if err != nil {
						return err
					}
					val = reflect.ValueOf(convertedVal)
				}
			}

			fieldValue.Set(val)

			return nil
		}
	}

	return nil
}

func convertValue(value interface{}, targetType reflect.Type) (interface{}, error) {
	switch targetType.Kind() {
	case reflect.Uint8:
		v, err := strconv.ParseUint(fmt.Sprintf("%v", value), 10, 8)
		if err != nil {
			return nil, err
		}
		return uint8(v), nil
	case reflect.Uint16:
		v, err := strconv.ParseUint(fmt.Sprintf("%v", value), 10, 16)
		if err != nil {
			return nil, err
		}
		return uint16(v), nil
	case reflect.Uint32:
		v, err := strconv.ParseUint(fmt.Sprintf("%v", value), 10, 32)
		if err != nil {
			return nil, err
		}
		return uint32(v), nil
	case reflect.Uint64:
		v, err := strconv.ParseUint(fmt.Sprintf("%v", value), 10, 64)
		if err != nil {
			return nil, err
		}
		return v, nil
	case reflect.Bool:
		v, err := strconv.ParseBool(fmt.Sprintf("%v", value))
		if err != nil {
			return nil, err
		}
		return v, nil
	default:
		return nil, fmt.Errorf("unsupported conversion to %s", targetType.Kind())
	}
}
