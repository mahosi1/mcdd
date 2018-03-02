package main

import (
	"encoding/binary"
	"io"
	"os"
)

type Header struct {
	headerSignature []byte
	majorVersion    uint16
	sectorShift     uint16
	difat           []int
	clsid           []byte
	minorVersion    uint16
	byteOrder       uint
	miniSectorShift uint
	unUsed2         uint
}

func (h *Header) SetClsid(b []byte) {
	h.clsid = b
}

func NewHeader() *Header {
	h := &Header{}
	h.headerSignature = []byte{0xD0, 0xCF, 0x11, 0xE0, 0xA1, 0xB1, 0x1A, 0xE1}
	h.majorVersion = 3
	h.minorVersion = 0x003E
	h.byteOrder = 0xFFFE
	h.clsid = make([]byte, 16)
	h.sectorShift = 0x0009
	h.miniSectorShift = 6
	h.unUsed2 = 0 // need to check this
	h.difat = make([]int, 109)
	for i := range h.difat {
		h.difat[i] = 0xFFFFFFFF
	}
	return h
}

func (h *Header) Write(w io.Writer) {
	w.Write(h.headerSignature)
	w.Write(h.clsid)
	binary.Write(w, binary.LittleEndian, h.minorVersion)
	binary.Write(w, binary.LittleEndian, h.majorVersion)
	binary.Write(w, binary.LittleEndian, h.byteOrder)
	binary.Write(w, binary.LittleEndian, h.sectorShift)
	binary.Write(w, binary.LittleEndian, h.miniSectorShift)
	binary.Write(w, binary.LittleEndian, h.unUsed2)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	h := NewHeader()
	f, err := os.Create("./data")
	check(err)
	defer f.Close()
	h.Write(f)
}
