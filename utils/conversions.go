package utils

import (
	"fmt"
	"regexp"
	"strconv"
)

// MatcherLookup returns the map value of the named captures.
func MatcherLookup(match []string, re *regexp.Regexp) map[string]string {
	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 {
			result[name] = match[i]
		}
	}
	return result
}

// MustAtoi converts the given string to an integer.
// It panics if the conversion fails.
func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Errorf("MustAtoi: %s", err))
	}
	return i
}

// MustBeString converts the given value to a string.
// It panics if the conversion fails.
func MustBeString(value interface{}) string {
	v, ok := value.(string)
	if !ok {
		panic(fmt.Errorf("MustString: underlying type of input interface must be a string."))
	}
	return v
}

// MustBeMapStringInterface converts the given value to a map[string]interface{}.
// It panics if the conversion fails.
func MustBeMapStringInterface(value interface{}) map[string]interface{} {
	v, ok := value.(map[string]interface{})
	if !ok {
		panic(fmt.Errorf("MustBeMapStringInterface: underlying type of input interface must be a map[string]interface{}."))
	}
	return v
}

// MustBeMapStringInterface converts the given value to a []map[string]interface{}.
// It panics if the conversion fails.
func MustBeSliceOfMapStringInterface(value interface{}) []map[string]interface{} {
	elts, ok := value.([]interface{})
	if !ok {
		panic(fmt.Errorf("MustBeSliceOfMapStringInterface: underlying type of input interface must be a []map[string]interface{}."))
	}

	sl := make([]map[string]interface{}, len(elts))
	for _, elt := range elts {
		v, ok := elt.(map[string]interface{})
		if !ok {
			panic(fmt.Errorf("MustBeSliceOfMapStringInterface: underlying type of input interface must be a []map[string]interface{}."))
		}
		sl = append(sl, v)
	}
	return sl
}
