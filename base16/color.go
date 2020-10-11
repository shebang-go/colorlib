package base16

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Uses code parts of the following go module
// --------------------------------------------------------------------------
// Copyright 2015 The TCell Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
// You may obtain a copy of the license at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// --------------------------------------------------------------------------

// Color represents a base16 color as a 32 bit value. The first 24 bits are used
// to encode the color as an RGB value. The last 8 bits are reserved (see also
// gdamore/tcell's color encoding)
type Color int32

const (
	// NoColor is used to indicate a non color value.
	NoColor Color = -1
)

// NewColor returns a new color value by parsing the 24 bit color value in W3C
// #rrggbb format as used in base16 scheme files.
func NewColor(rrggbb string) Color {
	if len(rrggbb) == 6 {
		if v, e := strconv.ParseInt(rrggbb, 16, 32); e == nil {
			return Color(int32(v))
		}
	}
	return NoColor
}

// ToHexString returns a 6 characters long hex string of the color value
func (c Color) ToHexString() string {
	return fmt.Sprintf("%02x%02x%02x", (c>>16)&0xff, (c>>8)&0xff, c&0xff)
}

// ColorNameIndex returns the index (zero based) of the color name
func ColorNameIndex(colorname string) int {
	if strings.HasPrefix(colorname, "base") && len(colorname) == 6 {
		if v, e := strconv.ParseInt(strings.TrimLeft(colorname, "base"), 16, 8); e == nil {
			return int(v)
		}
	}
	return -1
}

// ColorIndexName returns the color name of the index (zero based)
func ColorIndexName(index int) string {
	return fmt.Sprintf("base%02x", index)
}

// ColorNames generates a slice of strings with base16 color names.
func ColorNames(count int) []string {
	keys := make([]string, 0, count)
	for i := 0; i < count; i++ {
		keys = append(keys, ColorIndexName(i))
	}
	sort.Strings(keys)
	return keys
}

// ValidColorName returns true is the color name is a valid base16 color name.
// The second argument is a flag when using the base16 extended mode.
func ValidColorName(colorname string, extended ...bool) bool {
	colorNameRe := `(?i)base[0][0-9a-f]`
	if len(extended) == 1 && extended[0] {
		colorNameRe = `(?i)base[01][0-9a-f]`
	}
	re := regexp.MustCompile(colorNameRe)
	return re.MatchString(colorname)
}
