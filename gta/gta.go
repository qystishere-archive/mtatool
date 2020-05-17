package gta

import (
	"fmt"
)

type GTA struct {
	Path string
	Data *Data
}

func Load(path string) (*GTA, error) {
	relativePath := func(name string) string {
		return path + "/" + name
	}

	data, err := LoadData(relativePath("data/"))
	if err != nil {
		return nil, fmt.Errorf("can't load data: %w", err)
	}

	return &GTA{
		Path: path,
		Data: data,
	}, nil
}
