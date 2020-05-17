package item

import (
	"github.com/qystishere/mtatool/gta/parser/item"
	gtaItem "github.com/qystishere/mtatool/gta/parser/item"
)

type Vehicle struct {
	Base
	gtaItem.NotImplemented
	Sirens   string `json:"sirens"`
	PaintJob string `json:"paintjob"`
	Model    string `json:"model"`
	Plate    string `json:"plate"`
}

func (v Vehicle) Type() item.Type {
	return item.TypeVehicle
}
