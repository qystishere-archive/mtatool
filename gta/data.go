package gta

import (
	"fmt"

	"github.com/qystishere/mtatool/gta/parser"
)

type Data struct {
	Path   string
	GTADat *parser.GTADat
}

func LoadData(path string) (*Data, error) {
	relativePath := func(name string) string {
		return path + "/" + name
	}

	gtaDat, err := parser.LoadGTADat(relativePath("gta.dat"))
	if err != nil {
		return nil, fmt.Errorf("can't load gta.dat: %w", err)
	}

	return &Data{
		Path:   path,
		GTADat: gtaDat,
	}, nil
}

func (d *Data) Save() error {
	if err := d.GTADat.Save(); err != nil {
		return fmt.Errorf("can't save gta.dat: %w", err)
	}

	return nil
}
