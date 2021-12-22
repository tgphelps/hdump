package hdump

import (
	"fmt"
	"strings"
	"testing"
)

func TestHdump(t *testing.T) {
	const result = "00000000 09 31 32 33 34 35 36 37 38 39 61 62 63 64 65 66   .123456789abcdef\n" +
		"00000010 48 65 6c 6c 6f                                    Hello\n"

	var b1 strings.Builder
	dest := NewHdumper(&b1)

	dest.DumpBytes(21, []byte("\t123456789abcdefHello"))

	actual := b1.String()
	if actual != result {
		t.Errorf("DumpBytes returned:\n%s", actual)
		// fmt.Printf("expected:\n");
		// fmt.Printf(result);
		// fmt.Printf("\ngot:\n");
		// fmt.Print(actual);
	}
}

func TestDummy(t *testing.T) {
	fmt.Printf("dummy test\n")
}
