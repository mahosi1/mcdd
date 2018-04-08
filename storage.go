package mcdf

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/sakeven/RbTree"
)

var ErrIllegalCharacters = errors.New("illegal characters in entry name")
var ErrEntryNameTooLong = errors.New("entry name must be shorter than 31 characters")

type DirectoryEntry struct {
	dirRepository []DirectoryEntry
	stgType       uint8
	storageClsid  string
	creationDate  []byte
	modifyDate    []byte
	entryName     string
	nameLength    uint16
	Sid           int32
	Child         int32
}

func (de *DirectoryEntry) LessThan(b interface{}) bool {
	de2, _ := b.(*DirectoryEntry)
	if de.nameLength > de2.nameLength {
		return false
	}
	if de.nameLength < de2.nameLength {
		return true
	}
	thisName := string(de.entryName)
	otherName := string(de2.entryName)
	for z := 0; z < len(de.entryName); z++ {
		thisChar := strings.ToUpper(string(thisName[z]))[0]
		otherChar := strings.ToUpper(string(otherName[z]))[0]
		if thisChar > otherChar {
			return false
		} else if thisChar <= otherChar {
			return true
		}
	}

	return false
}

func NewDirectoryEntry(name string, stageType uint8, directoryEntries []DirectoryEntry) *DirectoryEntry {

	var de DirectoryEntry
	de.dirRepository = directoryEntries
	de.stgType = stageType
	de.Sid = -1
	if stageType == 2 {
		de.storageClsid = "00000000000000000000000000000000"
		de.creationDate = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
		de.modifyDate = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	} else if stageType == 2 {
		val := uint64(time.Now().Nanosecond())
		buf := make([]byte, 8)
		binary.LittleEndian.PutUint64(buf, val)
		de.creationDate = buf
	} else if stageType == 5 {
		de.creationDate = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
		de.modifyDate = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	}
	de.SetEntryName(name)
	return &de

}

func (de *DirectoryEntry) SetEntryName(name string) error {
	if strings.Contains(name, "\\") || strings.Contains(name, "/") || strings.Contains(name, ":") || strings.Contains(name, "!") {
		return ErrIllegalCharacters
	}
	if len(name) > 31 {
		return ErrEntryNameTooLong
	}

	temp := []byte(name)
	newName := make([]byte, 64)
	for index, val := range temp {
		newName[index] = val
	}
	newName[len(temp)] = 0x00
	newName[len(temp)+1] = 0x00
	de.nameLength = uint16(len(temp) + 2)
	de.entryName = name
	return nil
}

func TryNew(streamName string, stageType uint8, directoryEntries []DirectoryEntry) *DirectoryEntry {
	var de DirectoryEntry
	for index, val := range directoryEntries {
		if val.stgType == 0 {
			directoryEntries[index] = de
			de.Sid = int32(index)
			return &de
		}
	}
	directoryEntries = append(directoryEntries, de)
	de.Sid = int32(len(directoryEntries) - 1)
	return &de
}

type CfStorage struct {
	compoundFile   *CompoundFile
	directoryEntry *DirectoryEntry
	children       *rbtree.Tree
}

func NewCfStorage(compoundFile *CompoundFile, directoryEntry *DirectoryEntry) *CfStorage {
	cf := &CfStorage{}
	cf.children = rbtree.NewTree()
	cf.compoundFile = compoundFile
	cf.directoryEntry = directoryEntry
	return cf
}

func newUUID() string {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		panic("error making uid")
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}

func (cf *CfStorage) AddStream(streamName string) *CfStream {
	dirEntry := TryNew(streamName, 2, cf.compoundFile.DirectoryEntries)

	cf.children.Insert(dirEntry, dirEntry)

	// value := cf.children.GetRoot()
	// cf.directoryEntry.Child = value.(*DirectoryEntry).Sid

	cfStream := NewCfStream(cf.compoundFile, dirEntry)

	return cfStream
}
