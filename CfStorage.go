package mcdf

type DirectoryEntry struct {
}

type CfStorage struct {
	compoundFile   *CompoundFile
	directoryEntry *DirectoryEntry
}

func NewCfStorage(compoundFile *CompoundFile, directoryEntry *DirectoryEntry) *CfStorage {
	cf := &CfStorage{}
	cf.compoundFile = compoundFile
	cf.directoryEntry = directoryEntry
	return cf
}
