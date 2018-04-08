package mcdf

type CompoundFile struct {
	header           *Header
	RootStorage      *CfStorage
	DirectoryEntries []DirectoryEntry
}

func NewCompoundFile() *CompoundFile {
	c := &CompoundFile{}
	c.header = NewHeader()
	de := &DirectoryEntry{}
	c.RootStorage = NewCfStorage(c, de)
	c.DirectoryEntries = make([]DirectoryEntry, 0)
	return c
}
