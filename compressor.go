package compressor

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"io/ioutil"
)

type decoder interface {
	DecodeString(string) ([]byte, error)
}

type encoder interface {
	EncodeToString([]byte) string
}

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
func Decompress(reader io.Reader) (io.Reader, error) {
	reader, err := gzip.NewReader(reader)
	if err != nil {
		return nil, err
	}

	return reader, nil
}

// DecodeAndDecompress decodes the given data string with the passed decoder e.g base64 decoder and decompresses the returned bytes
func DecodeAndDecompress(dataString string, d decoder, target interface{}) error {
	bts, err := d.DecodeString(dataString)
	if err == nil {
		r, err := Decompress(bytes.NewReader(bts))
		if err == nil {
			decodedBytes, err := ioutil.ReadAll(r)
			if err == nil {
				err = json.Unmarshal(decodedBytes, target)
			}
		}
	}
	return err
}

// CompressAndEncodeObjectToString compresses and encodes the given object and returns
// a encoded string containing a gzip byte array
func CompressAndEncodeObjectToString(object interface{}, e encoder) string {
	data, _ := json.Marshal(object)
	return e.EncodeToString(Compress(data))
}
