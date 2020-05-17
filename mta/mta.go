package mta

import (
	"fmt"
)

type MTA struct {
	Path string
	Mod  *Mod
}

func Load(path string) (*MTA, error) {
	relativePath := func(name string) string {
		return path + "/" + name
	}

	mod, err := LoadMod(relativePath("server/mods/deathmatch"))
	if err != nil {
		return nil, fmt.Errorf("can't read mod: %w", err)
	}

	return &MTA{
		Path: path,
		Mod:  mod,
	}, nil
}
