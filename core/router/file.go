package router

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// FileMeta carries metadata for a file transfer.
type FileMeta struct {
	FileID string
	Name   string
	Size   int64
}

// SerializeFileMeta serializes FileMeta into bytes: id_len|id|name_len|name|size(int64)
func SerializeFileMeta(m *FileMeta) []byte {
	idb := []byte(m.FileID)
	nameb := []byte(m.Name)
	buf := &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, uint16(len(idb)))
	buf.Write(idb)
	binary.Write(buf, binary.BigEndian, uint16(len(nameb)))
	buf.Write(nameb)
	binary.Write(buf, binary.BigEndian, m.Size)
	return buf.Bytes()
}

// ParseFileMeta parses bytes produced by SerializeFileMeta.
func ParseFileMeta(b []byte) (*FileMeta, error) {
	r := bytes.NewReader(b)
	var l uint16
	if err := binary.Read(r, binary.BigEndian, &l); err != nil {
		return nil, err
	}
	id := make([]byte, l)
	if _, err := r.Read(id); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &l); err != nil {
		return nil, err
	}
	name := make([]byte, l)
	if _, err := r.Read(name); err != nil {
		return nil, err
	}
	var size int64
	if err := binary.Read(r, binary.BigEndian, &size); err != nil {
		return nil, err
	}
	return &FileMeta{FileID: string(id), Name: string(name), Size: size}, nil
}

// ValidateFileMeta does basic checks on fields.
func ValidateFileMeta(m *FileMeta) error {
	if m.FileID == "" {
		return errors.New("empty file id")
	}
	if m.Name == "" {
		return errors.New("empty file name")
	}
	if m.Size < 0 {
		return errors.New("invalid file size")
	}
	return nil
}
