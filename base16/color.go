package base16

import (
	"fmt"
	"strconv"
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
