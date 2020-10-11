package base16yaml

import (
	// "fmt"
	"github.com/shebang-go/colorlib/base16"
	"io/ioutil"
	"strings"
)

// FileReaderFunc declares the function signature to read a file
type FileReaderFunc func(string) ([]byte, error)

// LoadBase16Scheme loads a base16 scheme from file fname. reader can be used to
// pass a file reader function as dependecy injection.
func LoadBase16Scheme(fname string, reader ...FileReaderFunc) (*Scheme, error) {
	var err error
	var data []byte
	var base16Yaml *Base16Yaml
	var fileReader FileReaderFunc = ioutil.ReadFile

	if len(reader) == 1 {
		fileReader = reader[0]
	}

	data, err = fileReader(fname)
	if err != nil {
		return nil, err
	}

	base16Yaml, err = UnmarshalBase16Yaml(data)
	if err != nil {
		return nil, err
	}

	base16Scheme, err := fromBase16Yaml(base16Yaml)
	if err != nil {
		return nil, err
	}

	return base16Scheme, nil

}

func fromBase16Yaml(base16Yaml *Base16Yaml) (*Scheme, error) {
	extendedMode := false
	if len(base16Yaml.colorNames) > 16 {
		extendedMode = true
	}

	scheme := Scheme{
		author: base16Yaml.Data["author"],
		scheme: base16Yaml.Data["scheme"],
		colors: make(map[string]Color, len(base16Yaml.colorNames)),
		fileKeys: Base16FileKeys{
			colorNameKeys: make(map[string]string, len(base16Yaml.colorNames)),
			otherKeys:     make(map[string]string, len(base16Yaml.Data)-len(base16Yaml.colorNames)),
		},
		extendedMode: extendedMode,
	}

	colorName := ""
	for k, v := range base16Yaml.Data {
		if ValidColorName(k, extendedMode) {
			colorName = strings.ToLower(k)
			scheme.fileKeys.colorNameKeys[colorName] = k
			scheme.colors[colorName] = NewColor(v)
		} else {
			scheme.fileKeys.otherKeys[strings.ToLower(k)] = k
		}
	}
	return &scheme, nil
}
