package main

type DataSource interface {
	WriteData(data []byte)
	ReadData() []byte
}

type FileDataSource struct{}

func (f *FileDataSource) WriteData(data []byte) {
	// ...
}
func (f *FileDataSource) ReadData() []byte {
	return []byte{}
}

type DataSourceDecorator struct {
	wrapee DataSource
}

func (f *DataSourceDecorator) WriteData(data []byte) {
	f.wrapee.WriteData(data)
}
func (f *DataSourceDecorator) ReadData() []byte {
	return f.wrapee.ReadData()
}

type EncryptionDecorator struct {
	wrapee DataSource
}

func (f *EncryptionDecorator) WriteData(data []byte) {
	f.wrapee.WriteData(data)
}
func (f *EncryptionDecorator) ReadData() []byte {
	return f.wrapee.ReadData()
}

func main() {
}
