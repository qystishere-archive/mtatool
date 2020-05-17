package item

import (
	"github.com/qystishere/mtatool/gta/parser/item"
	gtaItem "github.com/qystishere/mtatool/gta/parser/item"
)

type Ped struct {
	Base
	gtaItem.NotImplemented
	Model string `json:"model"`
}

func (p Ped) Type() item.Type {
	return item.TypePed
}