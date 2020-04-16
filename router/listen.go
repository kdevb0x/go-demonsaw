package router

import (
	"bytes"
	"crypto"
	"crypto/tls"
	"errors"
	"io"
	"log"
	"net"
)

// socEncConn is a net.Conn encrypted using zero or more layers of social crypto.
type socEncConn struct {
	tls.Conn
	key *crypto.PrivateKey
}

// Read implements io.Reader, decrypting data from the conn.
func (cc *socEncConn) Read(p []byte) (n int, err error) {
	var buff = bytes.NewBuffer(p)
	n, err = cc.Conn.Read(buff.Bytes())
	if err != nil {
		return 0, err
	}
	// TODO: write this function (n == count layers)
	err = decrypt(buff.Bytes(), buff.Bytes()[:buff.Len()])
	if err != nil {
		return 0, err
	}
	p = buff.Bytes()[:]
	return len(p), nil

}

func (cc *socEncConn) cryptoCopyEncrypt(dst io.Writer, src io.Reader) (int, error) {
	var c = make(chan []byte)
	go func(c chan []byte) {
		for {
			var rb = new(bytes.Buffer)
			msgSize, err := rb.ReadFrom(src)
			if err != nil {
				if err != bytes.ErrTooLarge || err != io.EOF {
					log.Printf("error reading from src: %w\n", err)
				}
			}
		}
	}(c)
	_, err := dst.Write(<-c)
	if err != nil {
		return 0, err
	}

}

// cryptoCopy is similar to io.Copy, but performs some cryptographic operation
// during the copy.
type cryptoCopy func(dest iot.Writer, src io.Reader) (written int64, err error)

func (r *router) Accept() (net.Conn, error) {
	if r.Addr() == nil {
		return nil, errors.New("listen address unset")
	}

}
