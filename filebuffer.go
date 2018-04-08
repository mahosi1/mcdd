package mcdf

import (
	"bytes"
	"errors"
	"io"
)

// FileBuffer is an in-memory io.Reader, io.Writer, io.Seeker used for testing file operations
type FileBuffer struct {
	Buffer bytes.Buffer
	Index  int64
}

// Bytes returns the byte slice from the in-memory file
func (b *FileBuffer) Bytes() []byte {
	return b.Buffer.Bytes()
}

// Read reads bytes from the FileBuffer into the given byte slice
func (b *FileBuffer) Read(p []byte) (int, error) {
	n, err := bytes.NewBuffer(b.Buffer.Bytes()[b.Index:]).Read(p)

	if err == nil {
		if b.Index+int64(len(p)) < int64(b.Buffer.Len()) {
			b.Index += int64(len(p))
		} else {
			b.Index = int64(b.Buffer.Len())
		}
	}
	return n, err
}

// Write writes bytes to the in-memory file
func (b *FileBuffer) Write(p []byte) (int, error) {
	n, err := b.Buffer.Write(p)

	if err == nil {
		b.Index = int64(b.Buffer.Len())
	}

	return n, err
}

// Seek moves the in-memory file pointer
func (b *FileBuffer) Seek(offset int64, whence int) (int64, error) {
	var err error
	var Index int64

	switch whence {
	case io.SeekStart:
		if offset >= int64(b.Buffer.Len()) || offset < 0 {
			err = errors.New("invalid offset")
		} else {
			b.Index = offset
			Index = offset
		}
	case io.SeekEnd:
		newIndex := int64(b.Buffer.Len()) - offset
		if newIndex > int64(b.Buffer.Len()) || newIndex < 0 {
			err = errors.New("invalid offset")
		} else {
			b.Index = newIndex
			Index = newIndex
		}
	case io.SeekCurrent:
		newIndex := int64(b.Index) + offset
		if newIndex >= int64(b.Buffer.Len()) || newIndex < 0 {
			err = errors.New("invalid offset")
		} else {
			b.Index = newIndex
			Index = newIndex
		}
	default:
		err = errors.New("unsupported seek method")
	}
	return Index, err
}
