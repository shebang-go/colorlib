package main

import (
	"fmt"
	"github.com/shebang-go/colorlib/base16"
)

func main() {

	scheme, _ := base16.NewScheme("demo", "wili")
	fmt.Printf("scheme: %v", scheme)
}
