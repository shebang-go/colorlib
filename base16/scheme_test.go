// +build !integration

package base16

import (
	"reflect"
	"testing"
)

func TestNewScheme(t *testing.T) {
	var expectedString string
	var gotString string
	var expectedInt int
	var gotInt int
	var scheme *Scheme
	var err error

	scheme, err = NewScheme("test", "nobody")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	expectedString = "test"
	gotString = scheme.Scheme()

	if expectedString != gotString {
		t.Errorf("expected value=%s, got=%s", expectedString, gotString)

	}

	expectedString = "nobody"
	gotString = scheme.Author()

	if expectedString != gotString {
		t.Errorf("expected value=%s, got=%s", expectedString, gotString)
	}

	if scheme.ExtendedModeOn() {
		t.Errorf("expected ExtendedModeOn()=false, got=true")
	}

	expectedInt = 16
	gotInt = scheme.CountColors()

	if expectedInt != gotInt {
		t.Errorf("expected value=%d, got=%d", expectedInt, gotInt)
	}

	gotColorNames := scheme.GetColorNames()
	expectedColorNames := []string{
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
		"base0a",
		"base0b",
		"base0c",
		"base0d",
		"base0e",
		"base0f",
	}
	if !reflect.DeepEqual(expectedColorNames, gotColorNames) {
		t.Errorf("expected value=%v, got=%v", expectedColorNames, gotColorNames)
	}

}

func TestNewSchemeExtended(t *testing.T) {
	// var expectedString string
	// var gotString string
	var expectedInt int
	var gotInt int
	var scheme *Scheme
	var err error

	scheme, err = NewScheme("test", "nobody", 20)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if !scheme.ExtendedModeOn() {
		t.Errorf("expected ExtendedModeOn()=true, got=false")
	}

	expectedInt = 20
	gotInt = scheme.CountColors()

	if expectedInt != gotInt {
		t.Errorf("expected value=%d, got=%d", expectedInt, gotInt)
	}

	gotColorNames := scheme.GetColorNames()
	expectedColorNames := []string{
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
		"base0a",
		"base0b",
		"base0c",
		"base0d",
		"base0e",
		"base0f",
		"base10",
		"base11",
		"base12",
		"base13",
	}
	if !reflect.DeepEqual(expectedColorNames, gotColorNames) {
		t.Errorf("expected value=%v, got=%v", expectedColorNames, gotColorNames)
	}

	expectedColor := NoColor
	gotColor := scheme.GetColor("base00")
	if expectedColor != gotColor {
		t.Errorf("expected value=%d, got=%d", expectedColor, gotColor)
	}

	expectedColor = NoColor
	gotColor = scheme.GetColor("base13")
	if expectedColor != gotColor {
		t.Errorf("expected value=%d, got=%d", expectedColor, gotColor)
	}

}
func TestNewSchemeErrorHandling(t *testing.T) {

	_, err := NewScheme("test", "nobody", 15)
	if err == nil {
		t.Errorf("expected error not nil")
	}

	_, err = NewScheme("test", "nobody", 33)
	if err == nil {
		t.Errorf("expected error not nil")
	}

}
func TestGetColor(t *testing.T) {
	scheme, _ := NewScheme("test", "nobody")

	expectedColor := NoColor
	gotColor := scheme.GetColor("base00")
	if expectedColor != gotColor {
		t.Errorf("expected value=%d, got=%d", expectedColor, gotColor)
	}

	expectedColor = NoColor
	gotColor = scheme.GetColor("base0f")
	if expectedColor != gotColor {
		t.Errorf("expected value=%d, got=%d", expectedColor, gotColor)
	}

	scheme, _ = NewScheme("test", "nobody", 20)

	expectedColor = NoColor
	gotColor = scheme.GetColor("base00")
	if expectedColor != gotColor {
		t.Errorf("expected value=%d, got=%d", expectedColor, gotColor)
	}

	expectedColor = NoColor
	gotColor = scheme.GetColor("base13")
	if expectedColor != gotColor {
		t.Errorf("expected value=%d, got=%d", expectedColor, gotColor)
	}

}

func TestSetColor(t *testing.T) {
	scheme, _ := NewScheme("test", "nobody")

	scheme.SetColor("base00", NewColor("ff0000"))
	expectColor := NewColor("ff0000")
	gotColor := scheme.GetColor("base00")
	if gotColor != expectColor {
		t.Errorf("expected value=%d, got=%d", expectColor, gotColor)
	}
}
