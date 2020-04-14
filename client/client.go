package client

import (
	"errors"
	"fmt"

	"github.com/BurntSushi/toml"
)

type client struct {
	Session *sessionConfig
	Router  *routerConfig
	Group   *groupConfig
	Options *optionConfig

	config *clientConfig
}

type clientConfig struct {
	fpath   string
	Threads int    `toml:"threads"`
	Enabled bool   `toml:"enabled"`
	Color   string `toml:"color"` // hexstring
	Name    string `toml:"name"`
}
type routerConfig struct {
	Port     int    `toml:"port,omitempty"`
	Address  string `toml:"address,omitempty"`
	Password []byte `toml:"address,omitempty"`
	Enabled  bool   `toml:"enabled,omitempty"`
	Name     string `toml:"name,omitempty"`
}

type sessionConfig struct {
	Salt       string `toml:"salt,omitempty"`
	PrimeSize  int    `toml:"prime_size,omitempty"`
	Algorithm  string `toml:"algorithm,omitempty"`
	Hash       string `toml:"hash,omitempty"`
	Cipher     string `toml:"cipher,omitempty"`
	KeySize    int    `toml:"key_size,omitempty"`
	Iterations int    `toml:"iterations,omitempty"`
}

type optionConfig struct {
	MaxUploads int    `toml:"max_uploads,omitempty"`
	Path       string `toml:"path,omitempty"`
	// original == "%m/%d/%G %I:%M:%S %p"
	Datestamp    string `toml:"datestamp,omitempty"`
	Mode         int    `toml:"mode,omitempty"`
	ChunkSize    int64  `toml:"chunk_size,omitempty"`
	Timestamp    string `toml:"timestamp,omitempty"`
	Salt         []byte `toml:"salt,omitempty"`
	MaxDownloads int    `toml:"max_downloads,omitempty"`
	Hash         string `toml:"hash,omitempty"`
	BufferSize   int    `toml:"buffer_size,omitempty"`
}

type groupConfig struct {
	Salt       []byte  `toml:"salt,omitempty"`
	Entropy    string  `toml:"entropy,omitempty"`
	Percent    float64 `toml:"percent,omitempty"`
	Cipher     string  `toml:"cipher,omitempty"`
	Enabled    bool    `toml:"enabled,omitempty"`
	Iterations int     `toml:"iterations,omitempty"`
	KeySize    int     `toml:"key_size,omitempty"`
	Hash       []byte  `toml:"hash,omitempty"`
}

type shareConfig struct {
	Path    string `toml:"path,omitempty"`
	Enabled bool   `toml:"enabled"`
}

func (c *client) loadConfig(tomlpath string) error {
	var conf = new(clientConfig)
	m, err := toml.DecodeFile(tomlpath, conf)
	if err != nil {
		return err
	}

	conf.fpath = tomlpath
	c.config = conf
	if len(m.Undecoded()) > 0 {
		return errors.New(fmt.Sprintf("error: unable to decode %d values from toml: %v\n", len(m.Undecoded()), m.Undecoded()))
	}
	return nil

}
