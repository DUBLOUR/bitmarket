package fsDataHolder

type FsDataHolder struct {
	DataPath string
}

func (FsDataHolder) Append(string) error {
	return nil
}

func (FsDataHolder) Delete(string) error {
	return nil
}

func (FsDataHolder) FindByKey(string) (string, error) {
	return "", nil
}
