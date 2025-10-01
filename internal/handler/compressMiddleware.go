package handler

import (
	"compress/flate"
	"net/http"
	"strings"

	"compress/gzip"
	"compress/lzw"
	"fmt"
	"io"
)

type CustomReader struct {
	SourceReader   io.ReadCloser
	CompressReader io.Reader
}

func NewCustomReader(source io.ReadCloser, encoding string) (*CustomReader, error) {
	compressReader, _ := getDecompressReader(encoding, source)
	return &CustomReader{
		SourceReader:   source,
		CompressReader: compressReader,
	}, nil
}

func (r *CustomReader) Read(p []byte) (n int, err error) {
	return r.CompressReader.Read(p)
}

func (r *CustomReader) Close() error {
	err := r.SourceReader.Close()
	if err != nil {
		return err
	}
	if rc, ok := r.CompressReader.(io.Closer); ok {
		return rc.Close()
	}

	return nil
}

func getDecompressReader(encoding string, reader io.Reader) (io.Reader, error) {
	switch encoding {
	case "gzip":
		return gzip.NewReader(reader)
	case "compress":
		return lzw.NewReader(reader, lzw.LSB, 8), nil
	case "deflate":
		return flate.NewReader(reader), nil
	}

	return nil, fmt.Errorf("unknown decompress encoding: %s", encoding)
}

func CustomDecompress(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encoding := strings.ToLower(r.Header.Get("Content-Encoding"))
		if encoding != "" {
			dr, err := NewCustomReader(r.Body, encoding)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)

				return
			}
			defer dr.Close()

			r.Body = dr
		}

		// no decompression
		next.ServeHTTP(w, r)
	})
}
