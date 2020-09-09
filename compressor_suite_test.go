package compressor

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCompressor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Compressor Suite")
}

var _ = Describe("Compressor", func() {
	Context("Compression", func() {
		It("should compress and decompress data correctly", func() {
			data := []byte("TestString")

			compressedData := Compress(data)
			testDecompressDataReader, err := Decompress(bytes.NewReader(compressedData))
			reader, err := gzip.NewReader(bytes.NewReader(compressedData))

			Expect(err).NotTo(HaveOccurred())

			uncompressedData, e := ioutil.ReadAll(reader)
			decompressData, _ := ioutil.ReadAll(testDecompressDataReader)

			Expect(e).NotTo(HaveOccurred())
			Expect(string(uncompressedData)).To(Equal("TestString"))
			Expect(string(decompressData)).To(Equal("TestString"))
		})
		It("should encode compress the given object", func() {
			encodedString := CompressAndEncodeObjectToString(map[string]string{
				"foo": "bar",
			}, base64.RawURLEncoding)

			bts, _ := base64.RawURLEncoding.DecodeString(encodedString)

			r, _ := Decompress(bytes.NewReader(bts))
			b, _ := ioutil.ReadAll(r)
			returnMap := make(map[string]string)
			Expect(json.Unmarshal(b, &returnMap)).To(BeNil())
			Expect(returnMap["foo"]).To(Equal("bar"))
		})
		It("should decode and decompress the given string", func() {
			encodedString := CompressAndEncodeObjectToString(map[string]string{
				"foo": "bar",
			}, base64.RawURLEncoding)
			returnMap := make(map[string]string)
			Expect(DecodeAndDecompress(encodedString, base64.RawURLEncoding, &returnMap)).To(BeNil())
			Expect(returnMap["foo"]).To(Equal("bar"))
		})
	})
	Context("Failure", func() {
		It("should fail on decompress", func() {
			reader, err := Decompress(bytes.NewReader([]byte("test")))
			Expect(reader).To(BeNil())
			Expect(err).NotTo(BeNil())
		})
	})
})
