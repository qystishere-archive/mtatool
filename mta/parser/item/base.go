package item

import (
	"github.com/qystishere/mtatool/gta/parser/item"
)

type Base struct {
	ID        string `json:"id"`
	Alpha     string `json:"alpha"`
	Dimension string `json:"dimension"`
}

func (bi *Base) ItemType() item.Type {
	return item.TypeUnknown
}
