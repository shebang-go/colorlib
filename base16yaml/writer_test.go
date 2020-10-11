package base16yaml

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

// WriterMock implements the Writer interface with an error side effect.
type WriterMockError struct {
}

// WriteFile calls callback with the WriteFile func signature for testing results
func (wm *WriterMockError) WriteFile(fname string, data []byte, perm os.FileMode) error {
	return fmt.Errorf("write file error")
}

// WriterMock implements the Writer interface.
type WriterMock struct {
	callback func(fname string, data []byte, perm os.FileMode) error
}

// WriteFile calls callback with the WriteFile func signature for testing results
func (wm *WriterMock) WriteFile(fname string, data []byte, perm os.FileMode) error {
	return wm.callback(fname, data, perm)
}
func TestBase16YamlSaveError(t *testing.T) {
	mockReader := &ReaderMock{}
	base16Scheme, _ := Load("default-dark.yaml", mockReader)

	mockWriter := &WriterMockError{}
	err := Save("default-dark.yaml", base16Scheme, 0700, mockWriter)

	if err == nil {
		t.Fatalf("expected error")
	}

}

func TestSaveBase16Scheme(t *testing.T) {

	var writeTestFunc = func(filename string, data []byte, perm os.FileMode) error {

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
	}
	mockReader := &ReaderMock{}
	base16Scheme, _ := Load("default-dark.yaml", mockReader)
	mockWriter := &WriterMock{
		callback: writeTestFunc,
	}
	err := Save("default-dark.yaml", base16Scheme, 0700, mockWriter)

	if err != nil {
		t.Fatalf("expected no error, got err: %v ", err)
	}
}
