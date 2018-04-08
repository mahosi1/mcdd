package mcdf

import (
	"testing"
)

func TestSimpleHeaderWrite(t *testing.T) {
	buf := new(FileBuffer)
	header := NewHeader()
	header.Write(buf)
	expected := 512
	actual := len(buf.Bytes())
	if expected != actual {
		t.Errorf("header write wrote wrong number of bytes, expected %v, got %v", expected, actual)
	}
}

func TestHeaderWriteRead(t *testing.T) {
	buf := new(FileBuffer)
	h1 := NewHeader()
	h1.Write(buf)
	buf.Index = 0
	h2 := new(Header)
	h2.Read(buf)

	if h1.Signature != h2.Signature {
		t.Errorf("incorrect signature, wrote %x, read %x", h1.Signature, h2.Signature)
	}
	if h1.Clsid != h2.Clsid {
		t.Errorf("incorrect Clsid, wrote %x, read %x", h1.Clsid, h2.Clsid)
	}
	if h1.MinorVersion != h2.MinorVersion {
		t.Errorf("incorrect MinorVersion, wrote %x, read %x", h1.MinorVersion, h2.MinorVersion)
	}
	if h1.MajorVersion != h2.MajorVersion {
		t.Errorf("incorrect MajorVersion, wrote %x, read %x", h1.MajorVersion, h2.MajorVersion)
	}
	if h1.ByteOrder != h2.ByteOrder {
		t.Errorf("incorrect ByteOrder, wrote %x, read %x", h1.ByteOrder, h2.ByteOrder)
	}
	if h1.SectorShift != h2.SectorShift {
		t.Errorf("incorrect SectorShift, wrote %x, read %x", h1.SectorShift, h2.SectorShift)
	}
	if h1.MiniSectorShift != h2.MiniSectorShift {
		t.Errorf("incorrect MiniSectorShift, wrote %x, read %x", h1.MiniSectorShift, h2.MiniSectorShift)
	}
	if h1.Reserved != h2.Reserved {
		t.Errorf("incorrect Reserved, wrote %x, read %x", h1.Reserved, h2.Reserved)
	}
	if h1.Unused != h2.Unused {
		t.Errorf("incorrect Unused, wrote %x, read %x", h1.Unused, h2.Unused)
	}
	if h1.DirectorySectorCount != h2.DirectorySectorCount {
		t.Errorf("incorrect DirectorySectorCount, wrote %x, read %x", h1.DirectorySectorCount, h2.DirectorySectorCount)
	}
	if h1.FATSectorsNumber != h2.FATSectorsNumber {
		t.Errorf("incorrect FATSectorsNumber, wrote %x, read %x", h1.FATSectorsNumber, h2.FATSectorsNumber)
	}
	if h1.FirstDirectorySectorID != h2.FirstDirectorySectorID {
		t.Errorf("incorrect FirstDirectorySectorID, wrote %x, read %x", h1.FirstDirectorySectorID, h2.FirstDirectorySectorID)
	}
	if h1.Unused2 != h2.Unused2 {
		t.Errorf("incorrect Unused2, wrote %x, read %x", h1.Unused2, h2.Unused2)
	}
	if h1.MinSizeStandardStream != h2.MinSizeStandardStream {
		t.Errorf("incorrect MinSizeStandardStream, wrote %x, read %x", h1.MinSizeStandardStream, h2.MinSizeStandardStream)
	}
	if h1.FirstMiniFATSectorID != h2.FirstMiniFATSectorID {
		t.Errorf("incorrect FirstMiniFATSectorID, wrote %x, read %x", h1.FirstMiniFATSectorID, h2.FirstMiniFATSectorID)
	}
	if h1.MiniFATSectorsNumber != h2.MiniFATSectorsNumber {
		t.Errorf("incorrect MiniFATSectorsNumber, wrote %x, read %x", h1.MiniFATSectorsNumber, h2.MiniFATSectorsNumber)
	}
	if h1.FirstDIFATSectorID != h2.FirstDIFATSectorID {
		t.Errorf("incorrect FirstDIFATSectorID, wrote %x, read %x", h1.FirstDIFATSectorID, h2.FirstDIFATSectorID)
	}
	if h1.DIFATSectorsNumber != h2.DIFATSectorsNumber {
		t.Errorf("incorrect DIFATSectorsNumber, wrote %x, read %x", h1.DIFATSectorsNumber, h2.DIFATSectorsNumber)
	}
}
