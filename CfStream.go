package mcdf

type CfStream struct {
	compoundFile *CompoundFile
	de           *DirectoryEntry
}

func NewCfStream(compoundFile *CompoundFile, de *DirectoryEntry) *CfStream {
	cfs := &CfStream{}
	cfs.compoundFile = compoundFile
	cfs.de = de
	return cfs
}

func (c *CfStream) SetData(b []byte) {

}
