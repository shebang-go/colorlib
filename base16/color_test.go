// +build !integration

package base16

import (
	"reflect"
	"testing"
)

func TestNewColor(t *testing.T) {
	var gotColor Color
	var expectColor Color

	// black
	gotColor = NewColor("000000")
	expectColor = 0
	if gotColor != expectColor {
		t.Errorf("expected value=%d, got=%d", expectColor, gotColor)
	}

	// white
	gotColor = NewColor("ffffff")
	expectColor = 16777215
	if gotColor != expectColor {
		t.Errorf("expected value=%d, got=%d", expectColor, gotColor)
	}
	// red
	gotColor = NewColor("ff0000")
	expectColor = 16711680
	if gotColor != expectColor {
		t.Errorf("expected value=%d, got=%d", expectColor, gotColor)
	}

	// green
	gotColor = NewColor("00ff00")
	expectColor = 65280
	if gotColor != expectColor {
		t.Errorf("expected value=%d, got=%d", expectColor, gotColor)
	}

	// blue
	gotColor = NewColor("0000ff")
	expectColor = 255
	if gotColor != expectColor {
		t.Errorf("expected value=%d, got=%d", expectColor, gotColor)
	}
}

func TestNewColorErrorHandling(t *testing.T) {
	var gotColor Color
	var expectColor Color

	// unsupported returns no color
	gotColor = NewColor("000000ff")
	expectColor = NoColor
	if gotColor != expectColor {
		t.Errorf("expected value=%d, got=%d", expectColor, gotColor)
	}
}

func TestToHexString(t *testing.T) {
	var color Color
	var expectString string
	var gotString string

	// black
	color = NewColor("000000")
	expectString = "000000"
	gotString = color.ToHexString()

	if gotString != expectString {
		t.Errorf("expected value=%s, got=%s", expectString, gotString)
	}
	// white
	color = NewColor("ffffff")
	expectString = "ffffff"
	gotString = color.ToHexString()

	if gotString != expectString {
		t.Errorf("expected value=%s, got=%s", expectString, gotString)
	}
	// #f7ca88
	color = NewColor("f7ca88")
	expectString = "f7ca88"
	gotString = color.ToHexString()

	if gotString != expectString {
		t.Errorf("expected value=%s, got=%s", expectString, gotString)
	}
}

func TestColorNameIndex(t *testing.T) {
	var gotInt int
	var expectInt int

	gotInt = ColorNameIndex("base00")
	expectInt = 0
	if gotInt != expectInt {
		t.Errorf("expected value=%d, got=%d", expectInt, gotInt)
	}

	gotInt = ColorNameIndex("base0a")
	expectInt = 10
	if gotInt != expectInt {
		t.Errorf("expected value=%d, got=%d", expectInt, gotInt)
	}

	gotInt = ColorNameIndex("base0f")
	expectInt = 15
	if gotInt != expectInt {
		t.Errorf("expected value=%d, got=%d", expectInt, gotInt)
	}
}

func TestColorIndexName(t *testing.T) {
	var gotString string
	var expectString string

	gotString = ColorIndexName(0)
	expectString = "base00"
	if gotString != expectString {
		t.Errorf("expected value=%s, got=%s", expectString, gotString)
	}

	gotString = ColorIndexName(10)
	expectString = "base0a"
	if gotString != expectString {
		t.Errorf("expected value=%s, got=%s", expectString, gotString)
	}

	gotString = ColorIndexName(15)
	expectString = "base0f"
	if gotString != expectString {
		t.Errorf("expected value=%s, got=%s", expectString, gotString)
	}

}

func TestGetColorNames(t *testing.T) {

	got := []string(ColorNames(16))
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
		"base0a",
		"base0b",
		"base0c",
		"base0d",
		"base0e",
		"base0f",
	}
	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expected value=%v, got=%v", expect, got)
	}

}

func TestValidColorName(t *testing.T) {

	colorname := "base00"
	if !ValidColorName(colorname) {
		t.Errorf("expected colorname=%s to be valid (true)", colorname)
	}
	colorname = "base0a"
	if !ValidColorName(colorname) {
		t.Errorf("expected colorname=%s to be valid (true)", colorname)
	}
	colorname = "base0f"
	if !ValidColorName(colorname) {
		t.Errorf("expected colorname=%s to be valid (true)", colorname)
	}
	colorname = "base10"
	if ValidColorName(colorname) {
		t.Errorf("expected colorname=%s to be invalid (true)", colorname)
	}
	colorname = "base10"
	if !ValidColorName(colorname, true) {
		t.Errorf("expected colorname=%s to be valid (true)", colorname)
	}
}
