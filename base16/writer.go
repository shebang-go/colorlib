package base16

import (
	// "fmt"
	"io/ioutil"
	"os"
)

// FileWriterFunc declares the function signature to write a file
type FileWriterFunc func(filename string, data []byte, perm os.FileMode) error

// SaveBase16Scheme saves the  base16 scheme to the file fname. writer can be used to
// pass a file writer function as dependecy injection.
func SaveBase16Scheme(fname string, scheme *Scheme, perm os.FileMode, writer ...FileWriterFunc) error {
	base16Yaml := toBase16Yaml(scheme)

	fileWriter := ioutil.WriteFile
	if len(writer) == 1 {
		fileWriter = writer[0]
	}

	data, err := MarshalBase16Yaml(base16Yaml)
	if err != nil {
		return err
	}

	err = fileWriter(fname, data, perm)
	if err != nil {
		return err
	}
	return nil
}

func toBase16Yaml(scheme *Scheme) *Base16Yaml {

	base16Yaml := Base16Yaml{
		Data: make(map[string]string, scheme.CountColors()+2),
	}

	for key, targetKey := range scheme.fileKeys.colorNameKeys {
		base16Yaml.Data[targetKey] = scheme.colors[key].ToHexString()
	}
	base16Yaml.Data[scheme.fileKeys.otherKeys["author"]] = scheme.Author()
	base16Yaml.Data[scheme.fileKeys.otherKeys["scheme"]] = scheme.Scheme()
	return &base16Yaml
}
