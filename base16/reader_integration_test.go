// +build integration

package base16

import (
	"io/ioutil"
	"path"
	"runtime"
	"testing"
)

func TestIntegrationLoadBase16Scheme(t *testing.T) {
	var gotString string
	var expectString string
	var gotInt int
	var expectInt int

	_, filename, _, _ := runtime.Caller(0)
	testFile := path.Join(path.Dir(filename), "./test/data/default-dark.yaml")

	base16Scheme, err := LoadBase16Scheme(testFile)

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

func TestIntegrationLoadBase16SchemeExtended(t *testing.T) {
	var gotString string
	var expectString string
	var gotInt int
	var expectInt int

	_, filename, _, _ := runtime.Caller(0)
	testFile := path.Join(path.Dir(filename), "./test/data/default-dark-extended.yaml")

	base16Scheme, err := LoadBase16Scheme(testFile, ioutil.ReadFile)

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

func TestIntegrationLoadBase16SchemeError(t *testing.T) {

	_, filename, _, _ := runtime.Caller(0)
	testFile := path.Join(path.Dir(filename), "./test/data/default-dark-extended-invalid.yaml")

	_, err := LoadBase16Scheme(testFile, ioutil.ReadFile)

	if err == nil {
		t.Errorf("expected error not nil")
	}
}

func TestIntegrationGetColor(t *testing.T) {
	var gotColor Color
	var expectColor Color

	_, filename, _, _ := runtime.Caller(0)
	testFile := path.Join(path.Dir(filename), "./test/data/default-dark.yaml")

	base16Scheme, _ := LoadBase16Scheme(testFile, ioutil.ReadFile)

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
