package client

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

type Toml struct {
	Sections struct {
		Title     string
		Atributes []string
	}
}

func ReadTOMLFile(filepath string) (Toml, error)
