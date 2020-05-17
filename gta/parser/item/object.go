package item

import (
	"fmt"
)

type Object struct {
	Base
	ModelID string `json:"model"`
	Model   string `json:"-"` // may be anything
	LOD     string `json:"lod"`
}

func (o Object) Type() Type {
	return TypeObject
}

func (o Object) Section() Section {
	return SectionINST
}

func (o Object) Compile() string {
	if o.Model == "" {
		o.Model = "mtatool"
	}

	if o.RotW == "" {
		o.RotW = "1"
	}

	if o.LOD == "" {
		o.LOD = "-1"
	}

	return fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s",
		o.ModelID, o.Model, o.Interior, o.PosX, o.PosY, o.PosZ, o.RotX, o.RotY, o.RotZ, o.RotW, o.LOD)
}
