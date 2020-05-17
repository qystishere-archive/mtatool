package mta

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/qystishere/mtatool/mta/parser"
)

type Mod struct {
	Path string
	Maps []*parser.Map
}

func LoadMod(path string) (*Mod, error) {
	relativePath := func(name string) string {
		return path + "/" + name
	}

	maps := make([]*parser.Map, 0)
	err := filepath.Walk(relativePath("resources"),
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if strings.Contains(path, "editor_") {
				return nil
			}

			if strings.HasSuffix(info.Name(), ".map") {
				m, err := parser.LoadMap(path)
				if err != nil {
					return fmt.Errorf("can't load map: %w", err)
				}

				maps = append(maps, m)
			}

			return nil
		})
	if err != nil {
		return nil, fmt.Errorf("can't read resources: %w", err)
	}

	return &Mod{
		Path: path,
		Maps: maps,
	}, nil
}
