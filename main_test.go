package sdk

import "testing"

func TestPrinter(t *testing.T) {
	printer := New()
	printer.Hello("World")
}
