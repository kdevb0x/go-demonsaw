package router

import (
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestHash(t *testing.T) {
	// create and write random bits to a file //
	src := rand.New(rand.NewSource(time.Now().UnixNano()))
	var b = make([]byte, src.Int())
	_, err := src.Read(b)
	if err != nil {
		t.Fatal(err)
	}
	// tmp file to hash
	err = ioutil.WriteFile(os.TempDir()+"goDemonsawTestFile", b, os.ModeTemporary)
	//-------------//

	var r = new(transferRouter)
	f, err := r.chunkFile(os.TempDir()+"goDemonsawTestFile", 1024)
	if err != nil {
		t.Fatalf("failed to chunk file; error: %w\n", err)
	}
	hashstr := f.hash()
	if len(hashstr.Error()) > 0 {
		t.Fail()
	}

	t.Logf("hashstring: %s\n", hashstr.String())
}
