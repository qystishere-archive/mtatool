package item

import (
	"github.com/qystishere/mtatool/gta/parser/item"
	gtaItem "github.com/qystishere/mtatool/gta/parser/item"
)

type Object struct {
	Base
	gtaItem.Object
	Breakable   string `json:"breakable"`
	Collisions  string `json:"collisions"`
	Model       string `json:"model"`
	Doublesided string `json:"doublesided"`
	Scale       string `json:"scale"`
}

func (o Object) Type() item.Type {
	return item.TypeObject
}
