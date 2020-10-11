package base16

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

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
