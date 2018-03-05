package mcdf

type CompoundFile struct {
	header      *Header
	rootStorage *CfStorage
}

func NewCompoundFile() *CompoundFile {
	c := &CompoundFile{}
	c.header = NewHeader()
	de := &DirectoryEntry{}
	c.rootStorage = NewCfStorage(c, de)

	return c
}
