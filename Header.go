package mcdf

import (
	"encoding/binary"
	"io"
)

const (
	signature            uint64 = 0xE11AB1A1E011CFD0
	miniStreamSectorSize uint32 = 64
	miniStreamCutoffSize uint32 = 4096
	dirEntrySize         uint32 = 128        //128 bytes
	maxRegSect           uint32 = 0xFFFFFFFA // Maximum regular sector number
	difatSect            uint32 = 0xFFFFFFFC //Specifies a DIFAT sector in the FAT
	fatSect              uint32 = 0xFFFFFFFD // Specifies a FAT sector in the FAT
	endOfChain           uint32 = 0xFFFFFFFE // End of linked chain of sectors
	freeSect             uint32 = 0xFFFFFFFF // Speficies unallocated sector in the FAT, Mini FAT or DIFAT
	maxRegStreamID       uint32 = 0xFFFFFFFA // maximum regular stream ID
	noStream             uint32 = 0xFFFFFFFF // empty pointer
)

type Header struct {
	Signature              uint64
	Clsid                  [16]byte
	MinorVersion           uint16
	MajorVersion           uint16
	ByteOrder              uint16
	SectorShift            uint16
	MiniSectorShift        uint16
	Reserved               uint16
	Unused                 uint32
	DirectorySectorCount   uint32
	FATSectorsNumber       uint32
	FirstDirectorySectorID uint32
	Unused2                uint32
	MinSizeStandardStream  uint32
	FirstMiniFATSectorID   uint32
	MiniFATSectorsNumber   int32
	FirstDIFATSectorID     uint32
	DIFATSectorsNumber     uint32
	InitialDifats          [109]uint32
}

func NewHeader() *Header {
	h := &Header{}
	h.Signature = signature
	h.MajorVersion = 3
	h.MinorVersion = 0x003E
	h.ByteOrder = 0xFFFE
	h.SectorShift = 0x0009
	h.MiniSectorShift = 6
	h.FirstDirectorySectorID = 0xFFFFFFFE
	h.MinSizeStandardStream = miniStreamCutoffSize
	h.FirstMiniFATSectorID = 0xFFFFFFFE
	h.FirstDIFATSectorID = 0xFFFFFFFE
	for i := range h.InitialDifats {
		h.InitialDifats[i] = 0xFFFFFFFF
	}
	return h
}

func (h *Header) Write(w io.Writer) {
	binary.Write(w, binary.LittleEndian, h)
	if h.MajorVersion == 4 {
		zeroHead := make([]byte, 3584)
		w.Write(zeroHead)
	}
}

func (h *Header) Read(r io.Reader) {
	binary.Read(r, binary.LittleEndian, h)
	if h.MajorVersion == 4 {
		zeroHead := make([]byte, 3584)
		r.Read(zeroHead)
	}
}
