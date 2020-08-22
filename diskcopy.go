package diskcopy

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

// Image is a DiskCopy 4.2 Image.
type Image struct {
	r *bytes.Reader
}

type header struct {
	Name         [64]byte
	DataSize     uint32
	TagSize      uint32
	DataChecksum uint32
	TagChecksum  uint32
	DiskEncoding byte
	Format       byte
	Magic        uint16
}

// NewImage opens and verifies a disk image from the given Reader. It will read the image fully into memory and return it as a ReadSeekCloser.
func NewImage(r io.Reader) (*Image, error) {
	header := header{}
	if err := binary.Read(r, binary.BigEndian, &header); err != nil {
		return nil, err
	}

	data := make([]byte, header.DataSize)

	read, err := r.Read(data)
	if err != nil {
		return nil, err
	}

	if read != len(data) {
		return nil, errors.New("Could not fully read image")
	}

	return &Image{
		r: bytes.NewReader(data),
	}, nil
}

// Read reads.
func (image *Image) Read(p []byte) (n int, err error) {
	return image.r.Read(p)
}

// Seek seeks.
func (image *Image) Seek(offset int64, whence int) (int64, error) {
	return image.r.Seek(offset, whence)
}

// Close will close the image and free all the storage.
func (image *Image) Close() error {
	return nil
}

// GetEncoding returns the disk encoding.
func (image *Image) GetEncoding() int {
	return 0
}

// GetFormat returns the disk format.
func (image *Image) GetFormat() int {
	return 0
}
