package router

import (
	"errors"
	"fmt"

	"github.com/BurntSushi/toml"
)

// 	[[router.room]]
type room struct {
	// enabled = true
	enabled bool `toml:"enabled"`

	// name = "Room #1"
	name string `toml:"name"` // maybe []byte?

	// color = "ff52c175"
	color [4]byte `toml:"color"`
}

func (r *msgRouter) loadConfig(tomlpath string) error {
	var conf = new(routerConfig)
	m, err := toml.DecodeFile(tomlpath, conf)
	if err != nil {
		return err
	}

	conf.fpath = tomlpath
	r.config = conf
	if len(m.Undecoded()) > 0 {
		return errors.New(fmt.Sprintf("error: unable to decode %d values from toml: %v\n", len(m.Undecoded()), m.Undecoded()))
	}
	return nil
}
