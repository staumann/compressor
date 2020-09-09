package compressor

import (
	"bytes"
	"compress/gzip"
	"io"
)

// Compress compresses the given byte array and returns a compressed byte array (using gzip)
func Compress(data []byte) []byte {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)

	_, err := gz.Write(data)
	if err != nil {
		return nil
	}

	if err = gz.Flush(); err != nil {
		return nil
	}

	if err = gz.Close(); err != nil {
		return nil
	}

	return b.Bytes()
}

// Decompress decompresses the given io.Reader and returns a decompressed io.Reader (using gzip)
func Decompress(reader io.Reader) io.Reader {

	reader, err := gzip.NewReader(reader)
	if err != nil {
		panic(err)
	}

	return reader
}
