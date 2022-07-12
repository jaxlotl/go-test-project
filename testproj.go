package main

import (
	"fmt"

	lbd "github.com/jaxlotl/go-libdogecoin"
)

func main() {
	lbd.W_context_start()
	fmt.Println("Hello from libdogecoin!")
	lbd.W_context_stop()
}
