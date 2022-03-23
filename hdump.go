// Package hdump dumps data in the traditional hex/ASCII format.
package hdump

// version 0.1.0

import (
	"fmt"
	"io"
	"log"
	"unsafe"
)

// Hdumper contains all context needed to dump data.
type Hdumper struct {
	dest    io.Writer // Where the dump gets printed
	offset  int       // Next data offset to print
	hexOnly bool      // true => Show only hex. No ASCII.
}

// NewHdumper return a new Hdumper which will print to 'wr'.
// offset defaults to 0, and hexOnly to false
func NewHdumper(wr io.Writer) *Hdumper {
	var hd Hdumper
	hd.dest = wr
	return &hd
}

// SetOffset sets the offset to be printed next time.
func (h *Hdumper) SetOffset(num int) {
	h.offset = num
}

// SetHexOnly sets or clears hexOnly
func (h *Hdumper) SetHexOnly(b bool) {
	h.hexOnly = b
}

// DumpBytes dumps 'size' bytes from the slice 'buff'.
func (h *Hdumper) DumpBytes(size int, buff []byte) {
	offset := 0
	for size > 0 {
		this := min(size, 16)
		h.dump16(buff[offset:offset+this], this)
		offset += this
		size -= this
		if size < 0 {
			log.Fatal("size went negative")
		}
	}
}

// DumpInt32s dumps 'size' int32s from the slice 'buff'.
func (h *Hdumper) DumpInt32s(size int, buff []int32) {
	offset := 0
	for size > 0 {
		this := min(size, 8)
		h.dump8ints(buff[offset:offset+this], this)
		offset += this
		size -= this
		if size < 0 {
			log.Fatal("size went negative")
		}
	}
}

// dump16 prints the next line of up to 16 bytes dumped.
func (h *Hdumper) dump16(buff []byte, n int) {
	fmt.Fprintf(h.dest, "%08x ", h.offset)
	h.offset += 16
	buff = buff[:n]
	for _, b := range buff {
		fmt.Fprintf(h.dest, "%02x ", b)
	}
	if !h.hexOnly {
		if n < 16 {
			for i := n; i < 16; i++ {
				fmt.Fprint(h.dest, "   ")
			}
		}
		fmt.Fprint(h.dest, "  ")
		for _, b := range buff {
			fmt.Fprint(h.dest, asc(b))
		}
		fmt.Fprintf(h.dest, "\n")
	}
}

// dump8ints prints the next line of up to 8 ints dumped.
func (h *Hdumper) dump8ints(buff []int32, n int) {
	fmt.Fprintf(h.dest, "%08x ", h.offset)
	h.offset += 8
	buff = buff[:n]
	for _, i := range buff {
		j := *((*uint32)(unsafe.Pointer(&i)))
		fmt.Fprintf(h.dest, "%08x ", j)
	}
	fmt.Fprintf(h.dest, "\n")
}

// asc returns the printable version of a byte.
func asc(b byte) string {
	if b < 32 || b > 126 {
		return "."
	} else {
		return string(b)
	}
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
