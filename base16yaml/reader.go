package base16yaml

import (
	// "fmt"
	"github.com/shebang-go/colorlib/base16"
	"io/ioutil"
	// "strings"
)

// Reader defines an interface for reading a file.
type Reader interface {
	// ReadFile reads a file and returns a byte slice on success or an error on
	// failure.
	ReadFile(fname string) ([]byte, error)
}

// FileReader implements the Reader interface.
type FileReader struct {
}

// ReadFile proxies to ioutil.ReadFile
func (fr *FileReader) ReadFile(fname string) ([]byte, error) { return ioutil.ReadFile(fname) }

// Load loads a base16 yaml file given by fname and returns a Scheme interface
// on success or an error on failure.
func Load(fname string, readerArg ...Reader) (base16.Scheme, error) {
	var reader Reader = &FileReader{}
	var err error
	var data []byte
	var base16Yaml *Base16Yaml

	if len(readerArg) == 1 {
		reader = readerArg[0]
	}

	data, err = reader.ReadFile(fname)
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

func fromBase16Yaml(base16Yaml *Base16Yaml) (base16.Scheme, error) {
	extendedMode := false
	if len(base16Yaml.colorNames) > 16 {
		extendedMode = true
	}

	scheme, err := base16.NewScheme(base16Yaml.Data["author"], base16Yaml.Data["scheme"], len(base16Yaml.colorNames))
	if err != nil {
		return nil, err
	}

	for k, v := range base16Yaml.Data {
		if base16.ValidColorName(k, extendedMode) {
			scheme.SetColor(k, base16.NewColor(v))
		} else if k == "author" {
			scheme.SetAuthor(v)
		} else if k == "scheme" {
			scheme.SetScheme(v)
		}
	}
	return scheme, nil
}
