package base16

import (
	"fmt"
	"strings"
)

const (
	// ExtendedModeMaxColors specifies how many colors can be used in base16
	// extended mode.
	ExtendedModeMaxColors = 32
	Base16DefaultColors   = 16
)

// Scheme defines the interface for a base16 scheme.
type Scheme interface {
	// Author returns the author of a scheme
	Author() string

	// Scheme returns the scheme identifier (name)
	Scheme() string

	// CountColors returns the number of colors
	CountColors() int

	// GetColor returns the color specified by colorname
	GetColor(colorname string) Color

	// GetColorNames returns a sorted string slice of all color names
	GetColorNames() []string

	// SetColor sets the color c for color name
	SetColor(colorname string, c Color)

	// ExtendedModeOn returns the extended mode flag
	ExtendedModeOn() bool
}

// SchemeData is the internal representation of a base16 colors scheme. All color
// names are converted to lower case characters in order to avoid confusion when
// accessing color names.
type SchemeData struct {
	// auther contains the author of the scheme
	author string

	// scheme holds the scheme identifier (or name)
	scheme string

	// colors holds all base16 colors in a string map. When accessing colors
	// keys are automatically converted to lower case characters.
	colors map[string]Color

	// fileKeys maps lower case key names (memory) to the case of the original key
	// names in the yaml file so that the original case is preserved when writing
	// files.
	// fileKeys Base16FileKeys

	// // sortedColorNames contains all color names sorted alphabetically.
	sortedColorNames []string

	// extendedMode is a flag which will be set when more than 16 colors are
	// defined.
	extendedMode bool
}

// NewScheme creates a new scheme. Use schemeName and author to define the basic
// data fields for a scheme. The 3rd argument can be used to extend the scheme
// in order to use more than 16 colors (experimental and non-standard).
func NewScheme(schemeName string, author string, countColorsOverride ...int) (Scheme, error) {
	countColors := Base16DefaultColors
	extendedMode := false
	if len(countColorsOverride) == 1 {
		if countColorsOverride[0] > ExtendedModeMaxColors || countColorsOverride[0] < Base16DefaultColors {
			return nil, fmt.Errorf(
				"scheme must have at least %d colors and at most %d colors",
				Base16DefaultColors,
				ExtendedModeMaxColors,
			)
		}
		if countColorsOverride[0] > Base16DefaultColors {
			extendedMode = true
			countColors = countColorsOverride[0]
		}
	}

	scheme := SchemeData{
		scheme:           schemeName,
		author:           author,
		extendedMode:     extendedMode,
		sortedColorNames: ColorNames(countColors),
		colors:           make(map[string]Color, countColors),
		// fileKeys: Base16FileKeys{
		// 	colorNameKeys: make(map[string]string, countColors),
		// 	otherKeys:     make(map[string]string, 2),
		// },
	}

	for _, k := range scheme.sortedColorNames {
		scheme.colors[k] = NoColor
	}
	return &scheme, nil
}

// Author returns the author of the scheme
func (scheme *SchemeData) Author() string {
	return scheme.author
}

// Scheme returns the scheme identifier (name)
func (scheme *SchemeData) Scheme() string {
	return scheme.scheme
}

// CountColors returns the number of colors
func (scheme *SchemeData) CountColors() int {
	return len(scheme.colors)
}

// GetColor returns the color specified by colorname
func (scheme *SchemeData) GetColor(colorname string) Color {
	return scheme.colors[strings.ToLower(colorname)]
}

// GetColorNames returns a sorted string slice of all color names
func (scheme *SchemeData) GetColorNames() []string {
	return scheme.sortedColorNames
}

// SetColor sets the color c for color name
func (scheme *SchemeData) SetColor(colorname string, c Color) {
	scheme.colors[strings.ToLower(colorname)] = c
}

// ExtendedModeOn returns the extended mode flag
func (scheme *SchemeData) ExtendedModeOn() bool {
	return scheme.extendedMode
}
