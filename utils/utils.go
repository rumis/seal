package utils

import (
	"errors"
	"regexp"

	"github.com/rumis/mapstructure"
)

// Struct2Map 将结构体转为Map
func Struct2Map(input interface{}) (map[string]interface{}, error) {
	r := make(map[string]interface{})
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:   &r,
		Metadata: nil,
		TagName:  "seal",
	})
	if err != nil {
		return r, err
	}
	err = decoder.Decode(input)
	return r, err
}

// Struct2MapSlice 将结构体数组转为Map数组
func Struct2MapSlice(input interface{}) ([]map[string]interface{}, error) {
	r := make([]map[string]interface{}, 0)
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:   &r,
		Metadata: nil,
		TagName:  "seal",
	})
	if err != nil {
		return r, err
	}
	err = decoder.Decode(input)
	return r, err
}

// map转结构体
// params out must be a pointer
func Map2Struct(input interface{}, out interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:     out,
		Metadata:   nil,
		TagName:    "seal",
		DecodeHook: mapstructure.StringToTimeHookFunc("2006-01-02 15:04:05"),
	})
	if err != nil {
		return err
	}
	err = decoder.Decode(input)
	return err
}

var plRegex = regexp.MustCompile(`\{:\w+\}`)

// ReplacePlaceHolders  replace the params placeholders
func ReplacePlaceHolders(s string, pl string, params map[string]interface{}) (string, []interface{}, error) {
	var lastErr error
	args := make([]interface{}, 0, 8)
	s = plRegex.ReplaceAllStringFunc(s, func(m string) string {
		key := m[2 : len(m)-1]
		if arg, ok := params[key]; ok {
			args = append(args, arg)
		} else {
			lastErr = errors.New("parameter not found: " + key)
		}
		return pl
	})
	return s, args, lastErr
}

// the regexp for columns and tables.
var selectRegex = regexp.MustCompile(`(?i:\s+as\s+|\s+)([\w\-_\.]+)$`)

// SelectRegex the regexp for columns and tables
func SelectRegex() *regexp.Regexp {
	return selectRegex
}

// AliasName get the table or column alias name from user expression
func AliasName(s string) string {
	matches := selectRegex.FindStringSubmatch(s)
	if len(matches) == 0 {
		// no alias
		return s
	}
	return matches[1]
}
