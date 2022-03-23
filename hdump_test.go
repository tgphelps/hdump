package hdump

import (
	"fmt"
	"strings"
	"testing"
)

func TestHdumpBytes(t *testing.T) {
	const result = "00000000 09 31 32 33 34 35 36 37 38 39 61 62 63 64 65 66   .123456789abcdef\n" +
		"00000010 48 65 6c 6c 6f                                    Hello\n"

	var b1 strings.Builder
	dest := NewHdumper(&b1)

	dest.DumpBytes(21, []byte("\t123456789abcdefHello"))

	actual := b1.String()
	if actual != result {
		t.Errorf("DumpBytes returned:\n%s", actual)
		fmt.Printf("expected:\n")
		fmt.Printf(result)
		fmt.Printf("\ngot:\n")
		fmt.Print(actual)
	}
}

func TestHdumpInts(t *testing.T) {
	const result = "00000000 00000001 00000002 00000003 00000004 00000005 00000006 00000007 00000008 \n" +
		"00000008 000000ff 00010000 01000000 ffffffff \n"

	var buff [12]int32 = [12]int32{1, 2, 3, 4, 5, 6, 7, 8, 255, 65536, 1 << 24, -1}
	var b1 strings.Builder
	dest := NewHdumper(&b1)

	dest.DumpInt32s(12, buff[:])

	actual := b1.String()
	if actual != result {
		t.Errorf("DumpInt32s returned:\n%s", actual)
		fmt.Printf("expected:\n")
		fmt.Printf(result)
		fmt.Printf("\ngot:\n")
		fmt.Print(actual)
	}
}
func TestDummy(t *testing.T) {
	fmt.Printf("dummy test\n")
}
