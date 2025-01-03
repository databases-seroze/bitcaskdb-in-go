package bitcaskdbingo

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

const (
	ACTIVE_DATAFILE = "bitcaskdb_%d.go"
)

type DataFile struct {
	sync.RWMutex

	writer *os.File
	reader *os.File
	id     int
	offset int
}


// New initialises a db store for storing/reading an active db file. 
// At a given time only one file can be active 
func New(dir string, index int) (*DataFile, error) {

	path := filepath.Join(dir, fmt.Sprintf(ACTIVE_DATAFILE, index))
	writer, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return nil, fmt.Errorf("error opening file for writing to db: %w", err)
	}

	reader, err := os.Open(path)
	if err!=nil{
		return nil, fmt.Errorf("error opening file for reading db: %w", err)
	}

	stat, err := writer.Stat()
	if err != nil {
		return nil, fmt.Errorf("error fetching file stats: %v", err)
	}

	df := &DataFile{
		writer: writer, 
		reader: reader, 
		id: index, 
		offset: int(stat.Size()),
	}

	return df, nil 
}

// returns the id of the data file 
func (d *DataFile) ID() int{
	return d.id
}

// returns the size of the db file in bytes 
func (d *DataFile) Size() (int64, error) {

	stat, err := d.writer.Stat() 
	if err!=nil{
		return -1, fmt.Errorf("error fetching file stats: %v", err)
	}

	return stat.Size(), nil 
}

// sync flushes the in-memory buffers to the disk 
func (d *DataFile) Sync() error {
	return d.writer.Sync() 
}

func (d* DataFile) Read(pos int, size int) ([]byte, error) {

	start := int64(pos-size) 

	// Initialize a buffer for reading the data 
	record := make([]byte, size)

	// Read the file with the given offset 
	n, err := d.reader.ReadAt(record, start)
	if err!=nil {
		return nil, err
	}

	if n != int(size) {
		return nil, fmt.Errorf("error fetching record, invalid size")
	}

	return record, nil
}

// writes teh byte array to the dataFile 
func (d* DataFile) Write(data []byte) (int, error) {
	if _, err := d.writer.Write(data); err != nil {
		return -1, err
	}

	// save the current offset 
	offset := d.offset 

	// increase the offset 
	d.offset += len(data)

	return offset, nil 
}

// Close the Data file. i.e closes any file descriptors associated with the underlying db file 
func (d* DataFile) Close() error {
	if err := d.writer.Close(); err!=nil {
		return err 
	}

	if err := d.reader.Close(); err!=nil {
		return err 
	}

	return nil 
}
