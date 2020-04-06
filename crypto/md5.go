// Crypto package contains cryptographic routines and helpers utilized by godemonsaw/client
// ** Although I try my best, at this time this implementation is NOT assumed to be secure **

package crypto

import (
	m5 "crypto/md5"
	"github.com/pkg/errors"
	"hash"
	"io"
	"log"
	"os"
)

type Hash uint8

type HashDigest []byte

const (
	NONE Hash = iota
	MD5
	SHA1
	SHA224
	SHA265
	SHA384
	SHA512
	RC2
)

var hashtype = []string{
	NONE:   "",
	MD5:    "md5",
	SHA1:   "sha1",
	SHA224: "sha224",
	SHA265: "sha256",
	SHA384: "sha384",
	SHA512: "sha512",
	RC2:    "rc2",
}

type FileHasher interface {
	HashFile(file *os.File) HashDigest
}

type Md5 struct {
	Md5Hash hash.Hash
	salt    []byte
}

func NewMd5(salt ...string) *Md5 {
	var saltbuff = make([]byte, len(salt[0], len(salt[:]*len(salt))))
	i := len(salt)
	switch i {
	case i > 1:
		for i, s := range salt {
			append(saltbuff, []byte(s))
		}
		return &Md5{m5.New(), saltbuff}

	case 1:
		return &Md5{m5.New(), []byte(salt[0])}

	default:
		return &Md5{m5.New(), nil}
	}
}

func (s *Md5) HashFile(file ...*os.File) (digest []HashDigest, err error) {
	switch {
	case len(file) < 1:
		warn := errors.New("Warning: No file arguments given in call to HashFile!")
		digest = nil
		err = warn
		break
	case len(file) >= 1:
		for i, f := range file {
			if _, err = io.Copy(s.Md5Hash, f); err != nil {
				log.Printf("Error hashing file: %s", err)
				return nil, err
			}
			digest = append(digest, s.Md5Hash.Sum(nil))
			s.Md5Hash.Reset()
			err = nil
			break
		}
	}
	return
}
