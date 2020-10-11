// +build integration

package base16yaml

import (
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"
)

func TestIntegrationBase16YamlSave(t *testing.T) {
	schemeSaveFile, err := ioutil.TempFile("", "base16scheme")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(schemeSaveFile.Name())

	_, filename, _, _ := runtime.Caller(0)
	testFile := path.Join(path.Dir(filename), "./testdata/data/default-dark.yaml")

	base16Scheme, _ := Load(testFile)

	Save(schemeSaveFile.Name(), base16Scheme, 0700)
	data, err := ioutil.ReadFile(schemeSaveFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	gotData := strings.Split(string(data), "\n")

	expectedFields := map[string]string{
		"scheme": "\"Default Dark\"",
		"author": "\"Chris Kempson (http://chriskempson.com)\"",
		"base00": "\"181818\"",
		"base01": "\"282828\"",
		"base02": "\"383838\"",
		"base03": "\"585858\"",
		"base04": "\"b8b8b8\"",
		"base05": "\"d8d8d8\"",
		"base06": "\"e8e8e8\"",
		"base07": "\"f8f8f8\"",
		"base08": "\"ab4642\"",
		"base09": "\"dc9656\"",
		"base0A": "\"f7ca88\"",
		"base0B": "\"a1b56c\"",
		"base0C": "\"86c1b9\"",
		"base0D": "\"7cafc2\"",
		"base0E": "\"ba8baf\"",
		"base0F": "\"a16946\"",
	}

	for _, line := range gotData {
		fields := strings.Split(line, ": ")

		if len(fields) == 2 {
			if expectedValue, ok := expectedFields[fields[0]]; ok {
				if strings.Compare(expectedValue, fields[1]) != 0 {

					t.Errorf("expected value=%s, got value=%s", expectedValue, fields[1])
				}
			} else {
				t.Errorf("expected fieldName=%s to be present in expected map", fields[0])

			}
		}
	}
}
