package item

import (
	"github.com/qystishere/mtatool/gta/parser/item"
	gtaItem "github.com/qystishere/mtatool/gta/parser/item"
)

type PickupType string

type Pickup struct {
	Base
	gtaItem.NotImplemented
	Typ     PickupType `json:"type"`
	Amount  string     `json:"amount"`
	Respawn string     `json:"respawn"`
}

func (p Pickup) Type() item.Type {
	return item.TypePickup
}
