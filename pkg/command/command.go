package command

import (
	"os"
	"strconv"
	"strings"
)

// ParseArgs Returns a map of arguments.
//
// Anything in the format of "--key=value" will be added to the map as key=value. If there's no = or/and value,
// the value will be empty in the map.
//
// Anything not suffixed with "--" will be added to the map in ascending order e.g. "0", "1" etc. For example,
// the args "foo bar" would be 0=foo, 1=bar.
func ParseArgs() map[string]string {
	result := map[string]string{}

	// Find pairs of --key value, and use remaining items as args
	args := os.Args

	argIndex := 0
	for _, arg := range args {
		if strings.HasPrefix(arg, "--") && len(arg) > 2 {
			equalsIndex := strings.Index(arg, "=")
			key := arg[2:]
			value := ""
			if equalsIndex > 0 {
				key = arg[2:equalsIndex]
				if len(arg) > equalsIndex+2 {
					value = arg[equalsIndex+1:]
				}
			}
			result[key] = value
		} else {
			key := strconv.Itoa(argIndex)
			result[key] = arg
			argIndex++
		}
	}

	return result
}
