package router

import (
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/crypto/blake2b"
)

// chunkedFile contains all of the bytes representing a file, split into equally
// sized chunks.
type chunkedFile struct {
	filename string
	// total number of chunks that make up the file
	chunkCount int
	// the size of each chunk
	chunksize int
	// really [chunkCount][]byte
	chunks [][]byte

	checksum []byte // easier to get this durring chunking
}

func (t *transferRouter) chunckFile(filepath string, chunksize int) (*chunkedFile, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return chunk(f, f.Name(), chunksize)

}

func chunk(src io.Reader, filename string, chunksize int) (*chunkedFile, error) {
	var b []byte
	_, err := io.ReadFull(src, b)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var f = new(chunkedFile)
	f.filename = filename
	f.chunks = make([][]byte, len(b))
	// end is start + chunksize
	var start int
	for i := range f.chunks {
		f.chunks[i] = make([]byte, chunksize)
		// TODO: check the math here this doesnt feel right [kda 4/15/20]
		n := copy(f.chunks[i], b[start:start+chunksize])
		start += chunksize
		b = b[n:]
	}
	return f, nil
}

func (cf *chunkedFile) hash() hashstring {
	var e hashstring
	// tmp buff
	var bs = make([]byte, cf.chunkCount*cf.chunksize)
	for _, n := range cf.chunks {
		// copy(bs[cf.chunksize * i:cf.chunksize * i + cf.chunksize], n)
		bs = append(bs, n...)
	}
	hasher, err := blake2b.New512(nil)
	if err != nil {
		e.err = fmt.Errorf("error hashing with %s: %w\n", "blake2b", err)
		log.Println(e.Error())
		return e
	}
	_, err = hasher.Write(bs)
	if err != nil {
		e.err = fmt.Errorf("error hashing with %s: %w\n", "blake2b", err)
		log.Println(e.Error())
		return e
	}
	e.hash = hasher.Sum(nil)
	fmt.Println(e.hash)
	return e
}

// hashstring is a string of a hash that implements the error interface, in
// order to be used as an optional type.
type hashstring struct {
	h   []byte
	err error
}

func (h hashstring) Error() string {
	if h.err != nil {
		return h.err.Error()
	}
	return ""
}

func (h hashstring) String() string {
	return hex.EncodeToString(h.h)
}
