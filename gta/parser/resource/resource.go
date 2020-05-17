package resource

type Type string

const (
	TypeIMG    Type = "IMG"
	TypeSplash Type = "SPLASH"
	TypeIDE    Type = "IDE"
	TypeIPL    Type = "IPL"
)

type Resourcer interface {
	Type() Type
	Path() string
}

// simpleResource
type simpleResource struct {
	t Type
	p string
}

func NewSimple(t Type, p string) *simpleResource {
	return &simpleResource{
		t: t,
		p: p,
	}
}

func (sr *simpleResource) Type() Type {
	return sr.t
}

func (sr *simpleResource) Path() string {
	return sr.p
}
