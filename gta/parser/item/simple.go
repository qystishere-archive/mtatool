package item

type Raw struct {
	t  Type
	se Section
	s  string
}

func NewRaw(section Section, source string) *Raw {
	return &Raw{
		se: section,
		s:  source,
	}
}

func (s *Raw) Type() Type {
	return TypeUnknown
}

func (s *Raw) Section() Section {
	return s.se
}

func (s *Raw) Compile() string {
	return s.s
}
