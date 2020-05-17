package item

type Type string
type Section string

const (
	TypeUnknown  Type = ""
	TypeObject   Type = "object"
	TypeVehicle  Type = "vehicle"
	TypePed      Type = "ped"
	TypePickup   Type = "pickup"
	TypeMarker   Type = "marker"
	TypeCullZone Type = "cull"

	SectionUnknown Section = ""
	SectionINST    Section = "inst"
	SectionCULL    Section = "cull"
	SectionGRGE    Section = "grge"
	SectionENEX    Section = "enex"
	SectionPICK    Section = "pick"
	SectionJUMP    Section = "jump"
	SectionTCYC    Section = "tcyc"
	SectionAUZO    Section = "auzo"
	SectionMULT    Section = "mult"
	SectionCARS    Section = "cars"
	SectionOCCL    Section = "occl"
)

type Item interface {
	Type() Type
	Section() Section
	Compile() string
}
