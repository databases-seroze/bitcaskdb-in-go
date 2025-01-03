package bitcaskdbingo

import (
	"bytes"
	"encoding/binary"
	"hash/crc32"
	"time"
)

/*
 *
 * Record is a binary representation of how each record is persisted in the disk.
 * Header represents how the record is stored and some metadata with it.
 * we also store checksum, timestamp and expiry of the record each with 4 bytes i.e uint32 = 32 bits
 *
 *
 * Representation of the record stored on disk.
 * ----------------------------------------------------------------------------------------------------------------------------
 * | crc(4 bytes) | time (4 bytes) | expiry (4 bytes) | key_size (4 bytes) | val_size (4 bytes) | key | val |
 * ----------------------------------------------------------------------------------------------------------------------------
 */
const (
	MaxKeySize   = 1<<32 - 1
	MaxValueSize = 1<<32 - 1
)

type Record struct {
	Header Header
	Key    string
	Value  []byte
}

type Header struct {
	Checksum  uint32
	Timestamp uint32
	Expiry    uint32
	KeySize   uint32
	ValSize   uint32
}

// encodes the header object to bytes.Buffer
func (h *Header) encode(buf *bytes.Buffer) error {
	return binary.Write(buf, binary.LittleEndian, h)
}

// decodes the record byte slice into header object
func (h *Header) decode(record []byte) error {
	return binary.Read(bytes.NewReader(record), binary.LittleEndian, h)
}

// checks if the record has expired
func (r *Record) isExpired() bool {
	if r.Header.Expiry == 0 {
		return false
	}
	return time.Now().Unix() > int64(r.Header.Expiry)
}

// returns if the record's value is not corrupted
func (r *Record) isValidCheckSum() bool {
	// do we only track checksum of value ??
	return crc32.ChecksumIEEE(r.Value) == r.Header.Checksum
}
