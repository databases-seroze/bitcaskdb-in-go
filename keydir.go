package bitcaskdbingo

import (
	"encoding/gob"
	"os"
)

// KeyDir represents an in-memory hash map for fast look ups of the key
// it stores mappings between key to Meta data. 
type KeyDir map[string]Meta 

// Meta represents file location at which you can find the value
type Meta struct {
	Timestamp  int 
	RecordSize int 
	RecordPos  int 
	FileID 	   int 
}

// Encode encodes the map to a gob file 
// This is typically used to generate a hints file
func (k *KeyDir) Encode(filePath string) error {
	// Create a file for storing gob data 
	file, err := os.Create(filePath)
	if err != nil{
		return err 
	}
	defer file.Close() 

	encoder := gob.NewEncoder(file)

	err = encoder.Encode(k)
	if err!=nil {
		return err 
	}

	return nil 
}

func (k *KeyDir) Decode(filePath string) error {

	file,err := os.Open(filePath)

	if err!=nil{
		return err 
	}

	defer file.Close() 

	decoder := gob.NewDecoder(file)

	err = decoder.Decode(k)
	if err!=nil{
		return err 
	}

	return nil 
}



