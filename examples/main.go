
package main

import "fmt"

import (
	"bytes"
	"encoding/binary"
)

// Header represents the fixed width fields present at the start of every record.
type Header struct {
    Checksum  uint32
    Timestamp uint32
    Expiry    uint32
    KeySize   uint32
    ValSize   uint32
}

// Encode takes a byte buffer, encodes the value of header and writes to the buffer.
func (h *Header) encode(buf *bytes.Buffer) error {
    return binary.Write(buf, binary.LittleEndian, h)
}

// Decode takes a record object decodes the binary value the buffer.
func (h *Header) decode(record []byte) error {
    // why are we using LittleEndian
    return binary.Read(bytes.NewReader(record), binary.LittleEndian, h)
}

type Record struct {
    Header Header
    Key    string //what's the diff b/w string and []byte why use one over the other 
    Value  []byte
}

type Barrel struct {
    keydir KeyDir 
    df     
}

func main() {
    // fmt.Println("Hello, World inside examples!")

    // lets start a db 

}
