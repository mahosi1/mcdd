package mcdf

import (
	"encoding/binary"
	"io"
)

type Header struct {
	headerSignature        []byte
	majorVersion           uint16
	sectorShift            uint16
	difat                  []uint32
	Clsid                  []byte
	minorVersion           uint16
	byteOrder              uint16
	miniSectorShift        uint32
	unUsed                 []byte
	DirectorySectorNumbers int32
	FATSectorsNumber       int32
	FirstDirectorySectorID uint32
	unUsed2                uint32
	MinSizeStandardStream  uint32
	FirstMiniFATSectorID   uint32
	MiniFATSectorsNumber   int32
	FirstDIFATSectorID     uint32
	DIFATSectorsNumber     uint32
}

func NewHeader() *Header {
	h := &Header{}
	h.headerSignature = []byte{0xD0, 0xCF, 0x11, 0xE0, 0xA1, 0xB1, 0x1A, 0xE1}
	h.majorVersion = 3
	h.minorVersion = 0x003E
	h.byteOrder = 0xFFFE
	h.Clsid = make([]byte, 16)
	h.sectorShift = 0x0009
	h.miniSectorShift = 6
	h.unUsed = make([]byte, 6)
	h.FirstDirectorySectorID = 0xFFFFFFFE
	h.MinSizeStandardStream = 4096
	h.FirstMiniFATSectorID = 0xFFFFFFFE
	h.FirstDIFATSectorID = 0xFFFFFFFE
	h.difat = make([]uint32, 109)
	for i := range h.difat {
		h.difat[i] = 0xFFFFFFFF
	}
	return h
}

func (h *Header) Write(w io.Writer) {
	binary.Write(w, binary.LittleEndian, h.headerSignature)
	binary.Write(w, binary.LittleEndian, h.Clsid)
	binary.Write(w, binary.LittleEndian, h.minorVersion)
	binary.Write(w, binary.LittleEndian, h.majorVersion)
	binary.Write(w, binary.LittleEndian, h.byteOrder)
	binary.Write(w, binary.LittleEndian, h.sectorShift)
	binary.Write(w, binary.LittleEndian, h.miniSectorShift)
	binary.Write(w, binary.LittleEndian, h.unUsed)
	binary.Write(w, binary.LittleEndian, h.DirectorySectorNumbers)
	binary.Write(w, binary.LittleEndian, h.FATSectorsNumber)
	binary.Write(w, binary.LittleEndian, h.FirstDirectorySectorID)
	binary.Write(w, binary.LittleEndian, h.unUsed2)
	binary.Write(w, binary.LittleEndian, h.MinSizeStandardStream)
	binary.Write(w, binary.LittleEndian, h.FirstMiniFATSectorID)
	binary.Write(w, binary.LittleEndian, h.MiniFATSectorsNumber)
	binary.Write(w, binary.LittleEndian, h.FirstDIFATSectorID)
	binary.Write(w, binary.LittleEndian, h.DIFATSectorsNumber)

	for _, i := range h.difat {
		binary.Write(w, binary.LittleEndian, i)
	}

	if h.majorVersion == 4 {
		zeroHead := make([]byte, 3584)
		w.Write(zeroHead)
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
