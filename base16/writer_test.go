// +build !integration

package base16

import (
	"os"
	"strings"
	"testing"
)

func FileWriterMock(filename string, data []byte, perm os.FileMode) error {

	return nil
}

func TestSaveBase16Scheme(t *testing.T) {
	base16Scheme, _ := LoadBase16Scheme("default-dark.yaml", ReadFileMock)

	SaveBase16Scheme("test", base16Scheme, 0700, func(filename string, data []byte, perm os.FileMode) error {

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
		return nil
	})
}
