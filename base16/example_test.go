// +build !integration

package base16

import (
	"fmt"
	"testing"
)

func ExampleBase16(t *testing.T) {
	var scheme Scheme
	var err error

	scheme, err = NewScheme("test", "nobody")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	fmt.Printf("scheme author: %s", scheme.Author())
}
