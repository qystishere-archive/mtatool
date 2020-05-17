package parser

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/qystishere/mtatool/gta/parser/resource"
	"github.com/qystishere/mtatool/shared"
)

type GTADat struct {
	Path      string
	Resources []resource.Resourcer

	FileInfo os.FileInfo
}

func LoadGTADat(path string) (*GTADat, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	bb, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var (
		resources = make([]resource.Resourcer, 0)
	)
	for i, line := range strings.Split(string(bb), shared.CRLF) {
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		data := strings.Split(line, " ")
		if len(data) != 2 {
			return nil, fmt.Errorf("can't parse line (%d): %s", i, line)
		}

		switch resource.Type(data[0]) {
		case resource.TypeIMG:
			resources = append(resources, resource.NewSimple(resource.TypeIMG, data[1]))
		case resource.TypeSplash:
			resources = append(resources, resource.NewSimple(resource.TypeSplash, data[1]))
		case resource.TypeIDE:
			resources = append(resources, resource.NewSimple(resource.TypeIDE, data[1]))
		case resource.TypeIPL:
			path := path[:15] + strings.ToLower(data[1])
			ipl, err := resource.LoadIPL(path)
			if err != nil {
				return nil, fmt.Errorf("can't load ipl: %w", err)
			}

			resources = append(resources, ipl)
		}
	}

	return &GTADat{
		Path:      path,
		Resources: resources,

		FileInfo: stat,
	}, nil
}

func (dat *GTADat) Compile() string {
	var b strings.Builder
	b.WriteString(shared.FilePrefix)
	var previousType resource.Type
	for _, r := range dat.Resources {
		if r.Type() != previousType {
			b.WriteString(shared.CRLF + shared.CRLF + fmt.Sprintf("# - %s", r.Type()))
			previousType = r.Type()
		}
		b.WriteString(shared.CRLF + fmt.Sprintf("%s %s", r.Type(), r.Path()))
	}
	return b.String()
}

func (dat *GTADat) Save() error {
	return ioutil.WriteFile(dat.Path, []byte(dat.Compile()), 0644)
}