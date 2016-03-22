package sdk

import "fmt"

// Printer is the base interface
type Printer struct {
}

// Hello prints Hello [string] to the output
func (p *Printer) Hello(s string) {
	fmt.Printf("Hello %s", s)
}

// New creates a new instance of the Printer type
func New() *Printer {
	return &Printer{}
}
