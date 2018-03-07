package mcdf

type CompoundFile struct {
	header           *Header
	rootStorage      *CfStorage
	DirectoryEntries []DirectoryEntry
}

func NewCompoundFile() *CompoundFile {
	c := &CompoundFile{}
	c.header = NewHeader()
	de := &DirectoryEntry{}
	c.rootStorage = NewCfStorage(c, de)
	c.DirectoryEntries = make([]DirectoryEntry, 0)
	return c
}
