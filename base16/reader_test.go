// +build !integration

package base16

import (
	"fmt"
	"testing"
)

var base16TestData = map[string]string{
	"default-dark.yaml": `
scheme: "Default Dark"
author: "Chris Kempson (http://chriskempson.com)"
base00: "181818"
base01: "282828"
base02: "383838"
base03: "585858"
base04: "b8b8b8"
base05: "d8d8d8"
base06: "e8e8e8"
base07: "f8f8f8"
base08: "ab4642"
base09: "dc9656"
base0A: "f7ca88"
base0B: "a1b56c"
base0C: "86c1b9"
base0D: "7cafc2"
base0E: "ba8baf"
base0F: "a16946"
`,
	"default-dark-extended.yaml": `
scheme: "Default Dark (Extended)"
author: "Chris Kempson (http://chriskempson.com)"
base00: "181818"
base01: "282828"
base02: "383838"
base03: "585858"
base04: "b8b8b8"
base05: "d8d8d8"
base06: "e8e8e8"
base07: "f8f8f8"
base08: "ab4642"
base09: "dc9656"
base0A: "f7ca88"
base0B: "a1b56c"
base0C: "86c1b9"
base0D: "7cafc2"
base0E: "ba8baf"
base0F: "a16946"
base10: "ff0000"
base11: "00ff00"
base12: "0000ff"
base13: "00ffff"
`,
	"default-dark-missing-colors.yaml": `
scheme: "Default Dark (Extended)"
author: "Chris Kempson (http://chriskempson.com)"
base00: "181818"
base01: "282828"
base02: "383838"
base03: "585858"
base04: "b8b8b8"
base05: "d8d8d8"
base06: "e8e8e8"
base07: "f8f8f8"
base08: "ab4642"
base09: "dc9656"
base0A: "f7ca88"
base0B: "a1b56c"
base0C: "86c1b9"
`,
	"default-dark-extended-invalid.yaml": `
scheme: "Default Dark (Extended)"
author: "Chris Kempson (http://chriskempson.com)"
base00: "181818"
base01: "282828"
base02: "383838"
base03: "585858"
base04: "b8b8b8"
base05: "d8d8d8"
base06: "e8e8e8"
base07: "f8f8f8"
base08: "ab4642"
base09: "dc9656"
base0A: "f7ca88"
base0B: "a1b56c"
base0C: "86c1b9"
base0D: "7cafc2"
base0E: "ba8baf"
base0F: "a16946"
base10: "ff0000"
base11: "00ff00"
base12: "0000ff"
base13: "00ffff"
base14: "00ffff"
base15: "00ffff"
base16: "00ffff"
base17: "00ffff"
base18: "00ffff"
base19: "00ffff"
base1a: "00ffff"
base1b: "00ffff"
base1c: "00ffff"
base1d: "00ffff"
base1e: "00ffff"
base1f: "00ffff"
base20: "00ffff"
`,
	"invalid-yaml.yaml": `
this wil fail
`,
}

func ReadFileMock(fname string) ([]byte, error) {
	return []byte(base16TestData[fname]), nil
}

func ReadFileErrorMock(fname string) ([]byte, error) {
	return nil, fmt.Errorf("read file error")
}

func TestLoadBase16Scheme(t *testing.T) {
	var gotString string
	var expectString string
	var gotInt int
	var expectInt int

	base16Scheme, err := LoadBase16Scheme("default-dark.yaml", ReadFileMock)
	if err != nil {
		t.Errorf("expected no error err=%v", err)
	}

	gotString = base16Scheme.Author()
	expectString = "Chris Kempson (http://chriskempson.com)"
	if gotString != expectString {
		t.Errorf("expected value=%s, got=%s", expectString, gotString)
	}

	gotString = base16Scheme.Scheme()
	expectString = "Default Dark"
	if gotString != expectString {
		t.Errorf("expected value=%s, got=%s", expectString, gotString)
	}

	gotInt = base16Scheme.CountColors()
	expectInt = 16
	if gotInt != expectInt {
		t.Errorf("expected value=%d, got=%d", expectInt, gotInt)
	}

	if base16Scheme.ExtendedModeOn() {
		t.Errorf("expected ExtendedModeOn() to be false got=%t", base16Scheme.ExtendedModeOn())
	}

}

func TestLoadBase16SchemeExtended(t *testing.T) {
	var gotString string
	var expectString string
	var gotInt int
	var expectInt int

	base16Scheme, err := LoadBase16Scheme("default-dark-extended.yaml", ReadFileMock)

	if err != nil {
		t.Errorf("expected no error err=%v", err)
	}

	gotString = base16Scheme.Author()
	expectString = "Chris Kempson (http://chriskempson.com)"
	if gotString != expectString {
		t.Errorf("expected value=%s, got=%s", expectString, gotString)
	}

	gotString = base16Scheme.Scheme()
	expectString = "Default Dark (Extended)"
	if gotString != expectString {
		t.Errorf("expected value=%s, got=%s", expectString, gotString)
	}

	gotInt = base16Scheme.CountColors()
	expectInt = 20
	if gotInt != expectInt {
		t.Errorf("expected value=%d, got=%d", expectInt, gotInt)
	}

	if !base16Scheme.ExtendedModeOn() {
		t.Errorf("expected ExtendedModeOn() to be true got=%t", base16Scheme.ExtendedModeOn())
	}
}

func TestLoadBase16SchemeError(t *testing.T) {

	_, err := LoadBase16Scheme("default-dark-extended-invalid.yaml", ReadFileMock)

	if err == nil {
		t.Errorf("expected error not nil")
	}

	_, err = LoadBase16Scheme("default-dark.yaml", ReadFileErrorMock)
	if err == nil {
		t.Errorf("expected error not nil")
	}

	_, err = LoadBase16Scheme("invalid-yaml.yaml", ReadFileMock)
	if err == nil {
		t.Errorf("expected error not nil")
	}
}

func TestGetColorAfterLoadingScheme(t *testing.T) {
	var gotColor Color
	var expectColor Color

	base16Scheme, _ := LoadBase16Scheme("default-dark.yaml", ReadFileMock)

	gotColor = base16Scheme.GetColor("base0a")
	expectColor = 16239240
	if gotColor != expectColor {
		t.Errorf("expected value=%d, got=%d", expectColor, gotColor)
	}
	gotColor = base16Scheme.GetColor("base0a")
	expectColor = 16239240
	if gotColor != expectColor {
		t.Errorf("expected value=%d, got=%d", expectColor, gotColor)
	}

}
