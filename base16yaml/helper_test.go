package base16yaml

import (
	"reflect"
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

func getValue(field string, kind reflect.Kind, v interface{}) interface{} {
	r := reflect.ValueOf(v)
	retval := reflect.Indirect(r).FieldByName(field)
	switch kind {
	case reflect.String:
		return retval.String()
	case reflect.Int:
		return int(retval.Int())
	}
	return retval
}
func getMethodValue(method string, args []reflect.Value, kind reflect.Kind, v interface{}) interface{} {
	meth := reflect.ValueOf(v).MethodByName(method)
	retval := meth.Call(args)[0]
	switch kind {
	case reflect.String:
		return retval.String()
	case reflect.Int:
		return int(retval.Int())
	case reflect.Bool:
		return bool(retval.Bool())
	}
	return retval
}
