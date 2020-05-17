package resource

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/qystishere/mtatool/gta/parser/item"
	"github.com/qystishere/mtatool/shared"
)

var (
	ErrNotImplemented = errors.New("not implemented")

	itemsOrder = []string{"zone", "inst", "cull", "grge", "enex", "pick", "jump", "tcyc", "auzo", "mult", "cars", "occl"}
)

type IPL struct {
	path  string
	items map[string][]item.Item

	FileInfo os.FileInfo
}

func CreateIPL(path string) (*IPL, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	return &IPL{
		path:     path,
		items:    make(map[string][]item.Item),
		FileInfo: stat,
	}, nil
}

func LoadIPL(path string) (*IPL, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	bb, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var (
		ipl = &IPL{
			path:     path,
			items:    make(map[string][]item.Item),
			FileInfo: stat,
		}

		lines          = strings.Split(string(bb), shared.CRLF)
		currentSection item.Section
	)

	for _, line := range lines {
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		if line == "end" {
			if currentSection != "" {
				currentSection = ""
			} else {
				return nil, fmt.Errorf("parsing error, unexpected line: %s", line)
			}
			continue
		} else if currentSection == "" {
			currentSection = item.Section(line)
			continue
		}

		if currentSection == item.SectionINST {
			// 16265, des_damlodbit04, 1024, -606.0859375, 1910.101563, 9.140625, 0, 0, 0, 1, -1
			parsed := strings.Split(line, ", ")
			if len(parsed) != 11 {
				return nil, fmt.Errorf("parsing error, unexpected length: %s", line)
			}

			if err = ipl.Add(item.Object{
				Base: item.Base{
					Interior: parsed[2],
					PosX:     parsed[3],
					PosY:     parsed[4],
					PosZ:     parsed[5],
					RotX:     parsed[6],
					RotY:     parsed[7],
					RotZ:     parsed[8],
					RotW:     parsed[9],
				},
				ModelID: parsed[0],
				Model:   parsed[1],
				LOD:     parsed[10],
			}); err != nil {
				return nil, err
			}

			continue
		}

		if err = ipl.Add(item.NewRaw(currentSection, line)); err != nil {
			return nil, err
		}
	}

	return ipl, nil
}

func (i *IPL) Clear() {
	i.items = make(map[string][]item.Item)
}

func (i *IPL) Add(it item.Item) error {
	if it.Section() == item.SectionUnknown {
		return ErrNotImplemented
	}
	section := string(it.Section())
	if _, ok := i.items[section]; !ok {
		i.items[section] = make([]item.Item, 0)
	}
	i.items[section] = append(i.items[section], it)
	return nil
}

func (i *IPL) Compile() string {
	var b strings.Builder
	b.WriteString(shared.FilePrefix + shared.CRLF)
	for _, name := range itemsOrder {
		m, ok := i.items[name]
		if !ok {
			continue
		}

		b.WriteString(shared.CRLF + name)
		for _, i := range m {
			b.WriteString(shared.CRLF + i.Compile())
		}
		b.WriteString(shared.CRLF + "end")
	}
	return b.String()
}

func (i *IPL) Save() error {
	return ioutil.WriteFile(i.path, []byte(i.Compile()), 0644)
}

func (i *IPL) Type() Type {
	return TypeIPL
}

func (i *IPL) Path() string {
	return i.path
}
