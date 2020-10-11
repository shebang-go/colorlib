// +build !integration

package base16yaml

import (
	"reflect"
	"testing"
)

func TestUnmarshalBase16Yaml(t *testing.T) {

	base16Yaml, _ := UnmarshalBase16Yaml([]byte(base16TestData["default-dark.yaml"]))
	var gotString string

	testCases := []struct {
		colorName  string
		colorValue string
	}{
		{"base00", "181818"},
		{"base0F", "a16946"},
	}
	for _, table := range testCases {
		gotString = base16Yaml.Data[table.colorName]
		if gotString != table.colorValue {
			t.Errorf(
				"expected key=%s value=%s, got value=%s",
				table.colorName,
				table.colorValue,
				gotString,
			)
		}
	}
}

func TestUnmarshalBase16YamlError(t *testing.T) {

	_, err := UnmarshalBase16Yaml([]byte(base16TestData["invalid-yaml.yaml"]))
	if err == nil {
		t.Errorf("expected error not nil")
	}
	_, err = UnmarshalBase16Yaml([]byte(base16TestData["default-dark-extended-invalid.yaml"]))
	if err == nil {
		t.Errorf("expected error not nil")
	}
	_, err = UnmarshalBase16Yaml([]byte(base16TestData["default-dark-missing-colors.yaml"]))
	if err == nil {
		t.Errorf("expected error not nil")
	}

}

func TestGetYamlColorNames(t *testing.T) {

	base16Yaml, _ := UnmarshalBase16Yaml([]byte(base16TestData["default-dark.yaml"]))
	got := []string(base16Yaml.getYamlColorNames())
	expect := []string{
		"base00",
		"base01",
		"base02",
		"base03",
		"base04",
		"base05",
		"base06",
		"base07",
		"base08",
		"base09",
		"base0A",
		"base0B",
		"base0C",
		"base0D",
		"base0E",
		"base0F",
	}
	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expected value=%v, got=%v", expect, got)
	}
}
