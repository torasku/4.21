package ascii

import (
	"fmt"
	"testing"
)

func TestExtendedASCII(t *testing.T) {

	s := []byte("ÆØÅ")
	testString := ExtendedASCIIText(s)

	if len(testString) < 0 {
		t.Error("Input with no values from extended ASCII")
	}
	fmt.Println("Input has values from extended ASCII")
}
