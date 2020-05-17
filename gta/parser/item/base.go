package item

type Base struct {
	Interior string `json:"interior"`
	PosX     string `json:"posX"`
	PosY     string `json:"posY"`
	PosZ     string `json:"posZ"`
	RotX     string `json:"rotX"`
	RotY     string `json:"rotY"`
	RotZ     string `json:"rotZ"`
	RotW     string `json:"rotW"`
}

func (bi *Base) ItemType() Type {
	return TypeUnknown
}
