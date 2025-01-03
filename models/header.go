package models

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
    return binary.Read(bytes.NewReader(record), binary.LittleEndian, h)
}