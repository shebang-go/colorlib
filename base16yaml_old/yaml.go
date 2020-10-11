package base16yaml

import (
	"fmt"
	yaml "gopkg.in/yaml.v3"
	"sort"
	"strings"
	// "unicode"
	// "unicode/utf8"
)

// Base16FileKeys declares a data struct for storing the character case of
// the original key names.
type Base16FileKeys struct {
	colorNameKeys map[string]string
	otherKeys     map[string]string
}

// Base16Yaml declares the data structure for reading and writing a scheme file.
type Base16Yaml struct {
	Data       map[string]string `yaml:",omitempty,inline"`
	colorNames []string
}

// UnmarshalBase16Yaml parses data as yaml and returns a Base16Yaml object on
// success.
func UnmarshalBase16Yaml(data []byte) (*Base16Yaml, error) {

	base16Yaml := Base16Yaml{Data: make(map[string]string)}

	err := yaml.Unmarshal(data, &base16Yaml)
	if err != nil {
		return nil, err
	}

	base16Yaml.colorNames = base16Yaml.getYamlColorNames()
	if len(base16Yaml.colorNames) > ExtendedModeMaxColors {
		return nil, fmt.Errorf("cannot use more than %d colors. Got %d colors", ExtendedModeMaxColors, len(base16Yaml.colorNames))
	}

	if len(base16Yaml.colorNames) < 16 {
		return nil, fmt.Errorf("invalid base16 scheme, expected at leaset 16 color definitions, got=%d", len(base16Yaml.colorNames))
	}

	return &base16Yaml, nil
}
func (y *Base16Yaml) getYamlColorNames() []string {

	colorNames := make([]string, 0, ExtendedModeMaxColors)
	for k := range y.Data {
		if strings.HasPrefix(k, "base") {
			colorNames = append(colorNames, k)
		}
	}
	sort.Strings(colorNames)
	return colorNames
}

// MarshalBase16Yaml parses data as yaml and returns a Base16Yaml object on
// success.
func MarshalBase16Yaml(base16Yaml *Base16Yaml) ([]byte, error) {

	textColors := ""
	textOtherFields := ""
	var targetString *string

	fileKeyNames := make([]string, 0, len(base16Yaml.Data))
	for key := range base16Yaml.Data {
		fileKeyNames = append(fileKeyNames, key)
	}
	sort.Strings(fileKeyNames)

	for _, key := range fileKeyNames {
		if strings.HasPrefix(key, "base") {
			targetString = &textColors
		} else {
			targetString = &textOtherFields
		}
		*targetString = *targetString + key + ": \"" + base16Yaml.Data[key] + "\"\n"
	}
	// cannot get it to work using yaml.Marshal, some color values are not
	// properly quoted
	// data, err := yaml.Marshal(&base16Yaml)
	return []byte(textOtherFields + textColors), nil
}
