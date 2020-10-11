package base16yaml

import (
	// "fmt"
	"github.com/shebang-go/colorlib/base16"
	"io/ioutil"
	"os"
)

// Writer defines an interface for writing a file.
type Writer interface {
	// WriteFile writes the file given by fname with bytes given by data and
	// permissions given by perm. Returns error on failure.
	WriteFile(filename string, data []byte, perm os.FileMode) error
}

// FileWriter implements the Writer interface.
type FileWriter struct {
}

// WriteFile proxies to ioutil.WriteFile
func (fr *FileWriter) WriteFile(fname string, data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(fname, data, perm)
}

// Save saves the  base16 scheme to the file fname. writer can be used to
// pass a file writer interface for dependecy injection.
func Save(fname string, scheme base16.Scheme, perm os.FileMode, writerArg ...Writer) error {
	var fileWriter Writer = &FileWriter{}
	base16Yaml := toBase16Yaml(scheme)

	if len(writerArg) == 1 {
		fileWriter = writerArg[0]
	}

	data, err := MarshalBase16Yaml(base16Yaml)
	if err != nil {
		return err
	}

	err = fileWriter.WriteFile(fname, data, perm)
	if err != nil {
		return err
	}
	return nil
}

func toBase16Yaml(scheme base16.Scheme) *Base16Yaml {

	base16Yaml := Base16Yaml{
		Data: make(map[string]string, scheme.CountColors()+2),
	}

	for _, key := range scheme.GetColorNames() {
		base16Yaml.Data[key] = scheme.GetColor(key).ToHexString()
	}
	base16Yaml.Data["author"] = scheme.Author()
	base16Yaml.Data["scheme"] = scheme.Scheme()
	return &base16Yaml
}
