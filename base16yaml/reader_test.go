package base16yaml

import (
	"fmt"
	"reflect"
	"testing"
)

type ReaderMockError struct {
}

func (rm *ReaderMockError) ReadFile(fname string) ([]byte, error) {
	return nil, fmt.Errorf("read file error")
}

type ReaderMock struct {
}

func (rm *ReaderMock) ReadFile(fname string) ([]byte, error) {
	return []byte(base16TestData[fname]), nil
}

func TestBase16YamlLoad(t *testing.T) {
	mock := &ReaderMock{}
	base16Scheme, _ := Load("default-dark.yaml", mock)

	tests := map[string]struct {
		method string
		args   []reflect.Value
		kind   reflect.Kind
		want   interface{}
	}{
		"Author":               {method: "Author", args: nil, kind: reflect.String, want: "Chris Kempson (http://chriskempson.com)"},
		"Scheme":               {method: "Scheme", args: nil, kind: reflect.String, want: "Default Dark"},
		"CountColors":          {method: "CountColors", args: nil, kind: reflect.Int, want: 16},
		"ExtendedModeOn":       {method: "ExtendedModeOn", args: nil, kind: reflect.Bool, want: false},
		"GetColor(\"base00\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base00")}, kind: reflect.Int, want: 1579032},
		"GetColor(\"base01\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base01")}, kind: reflect.Int, want: 2631720},
		"GetColor(\"base02\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base02")}, kind: reflect.Int, want: 3684408},
		"GetColor(\"base03\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base03")}, kind: reflect.Int, want: 5789784},
		"GetColor(\"base04\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base04")}, kind: reflect.Int, want: 12105912},
		"GetColor(\"base05\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base05")}, kind: reflect.Int, want: 14211288},
		"GetColor(\"base06\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base06")}, kind: reflect.Int, want: 15263976},
		"GetColor(\"base07\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base07")}, kind: reflect.Int, want: 16316664},
		"GetColor(\"base08\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base08")}, kind: reflect.Int, want: 11224642},
		"GetColor(\"base09\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base09")}, kind: reflect.Int, want: 14456406},
		"GetColor(\"base0A\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base0A")}, kind: reflect.Int, want: 16239240},
		"GetColor(\"base0B\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base0B")}, kind: reflect.Int, want: 10597740},
		"GetColor(\"base0C\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base0C")}, kind: reflect.Int, want: 8831417},
		"GetColor(\"base0D\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base0D")}, kind: reflect.Int, want: 8171458},
		"GetColor(\"base0E\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base0E")}, kind: reflect.Int, want: 12225455},
		"GetColor(\"base0F\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base0F")}, kind: reflect.Int, want: 10578246},
	}
	for name, tc := range tests {
		gotValue := getMethodValue(tc.method, tc.args, tc.kind, base16Scheme)
		if !reflect.DeepEqual(tc.want, gotValue) {
			t.Fatalf("%s>: expected: %v, kind: %v, got: %v", name, tc.want, tc.kind, gotValue)
		}
	}
}

func TestBase16YamlLoadExtended(t *testing.T) {
	mock := &ReaderMock{}
	base16Scheme, _ := Load("default-dark-extended.yaml", mock)

	tests := map[string]struct {
		method string
		args   []reflect.Value
		kind   reflect.Kind
		want   interface{}
	}{
		"Author":               {method: "Author", args: nil, kind: reflect.String, want: "Chris Kempson (http://chriskempson.com)"},
		"Scheme":               {method: "Scheme", args: nil, kind: reflect.String, want: "Default Dark (Extended)"},
		"CountColors":          {method: "CountColors", args: nil, kind: reflect.Int, want: 20},
		"ExtendedModeOn":       {method: "ExtendedModeOn", args: nil, kind: reflect.Bool, want: true},
		"GetColor(\"base00\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base00")}, kind: reflect.Int, want: 1579032},
		"GetColor(\"base01\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base01")}, kind: reflect.Int, want: 2631720},
		"GetColor(\"base02\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base02")}, kind: reflect.Int, want: 3684408},
		"GetColor(\"base03\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base03")}, kind: reflect.Int, want: 5789784},
		"GetColor(\"base04\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base04")}, kind: reflect.Int, want: 12105912},
		"GetColor(\"base05\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base05")}, kind: reflect.Int, want: 14211288},
		"GetColor(\"base06\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base06")}, kind: reflect.Int, want: 15263976},
		"GetColor(\"base07\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base07")}, kind: reflect.Int, want: 16316664},
		"GetColor(\"base08\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base08")}, kind: reflect.Int, want: 11224642},
		"GetColor(\"base09\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base09")}, kind: reflect.Int, want: 14456406},
		"GetColor(\"base0A\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base0A")}, kind: reflect.Int, want: 16239240},
		"GetColor(\"base0B\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base0B")}, kind: reflect.Int, want: 10597740},
		"GetColor(\"base0C\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base0C")}, kind: reflect.Int, want: 8831417},
		"GetColor(\"base0D\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base0D")}, kind: reflect.Int, want: 8171458},
		"GetColor(\"base0E\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base0E")}, kind: reflect.Int, want: 12225455},
		"GetColor(\"base0F\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base0F")}, kind: reflect.Int, want: 10578246},
		"GetColor(\"base10\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base10")}, kind: reflect.Int, want: 16711680},
		"GetColor(\"base11\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base11")}, kind: reflect.Int, want: 65280},
		"GetColor(\"base12\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base12")}, kind: reflect.Int, want: 255},
		"GetColor(\"base13\")": {method: "GetColor", args: []reflect.Value{reflect.ValueOf("base13")}, kind: reflect.Int, want: 65535},
	}
	for name, tc := range tests {
		gotValue := getMethodValue(tc.method, tc.args, tc.kind, base16Scheme)
		if !reflect.DeepEqual(tc.want, gotValue) {
			t.Fatalf("%s>: expected: %v, kind: %v, got: %v", name, tc.want, tc.kind, gotValue)
		}
	}
}

func TestBase16YamlLoadError(t *testing.T) {
	mock := &ReaderMock{}
	_, err := Load("default-dark-extended-invalid.yaml", mock)
	if err == nil {
		t.Fatalf("expected error not nil")
	}
	_, err = Load("default-dark-missing-colors.yaml", mock)
	if err == nil {
		t.Fatalf("expected error not nil")
	}
	_, err = Load("invalid-yaml.yaml", mock)
	if err == nil {
		t.Fatalf("expected error not nil")
	}

	mockErr := &ReaderMockError{}
	_, err = Load("default-dark.yaml", mockErr)
	if err == nil {
		t.Fatalf("expected error not nil")
	}

	_, err = Load("default-dark.yaml", mock)
	if err != nil {
		t.Fatalf("expected no error, got err: %v ", err)
	}
}
