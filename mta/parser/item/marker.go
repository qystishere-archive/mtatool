package item

import (
	"github.com/qystishere/mtatool/gta/parser/item"
	gtaItem "github.com/qystishere/mtatool/gta/parser/item"
)

type MarkerType string

type Marker struct {
	Base
	gtaItem.NotImplemented
	Typ   MarkerType `json:"type"`
	Color string     `json:"color"`
	Size  string     `json:"size"`
}

func (m Marker) Type() item.Type {
	return item.TypeMarker
}
