package compressor

import (
	"bytes"
	"compress/gzip"
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
			testDecompressDataReader := Decompress(bytes.NewReader(compressedData))
			reader, err := gzip.NewReader(bytes.NewReader(compressedData))

			Expect(err).NotTo(HaveOccurred())

			uncompressedData, e := ioutil.ReadAll(reader)
			decompressData, _ := ioutil.ReadAll(testDecompressDataReader)

			Expect(e).NotTo(HaveOccurred())
			Expect(string(uncompressedData)).To(Equal("TestString"))
			Expect(string(decompressData)).To(Equal("TestString"))
		})
	})
	Context("Failure", func() {
		It("should fail on decompress", func() {
			defer func() {
				if r := recover(); r == nil {
					Fail("The code did not panic")
				}
			}()
			reader := Decompress(bytes.NewReader([]byte("test")))
			Expect(reader).To(BeNil())
		})
	})
})
