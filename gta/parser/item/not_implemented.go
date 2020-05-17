package item

type NotImplemented struct{}

func (s NotImplemented) Type() Type {
	return TypeUnknown
}

func (s NotImplemented) Section() Section {
	return SectionUnknown
}

func (s NotImplemented) Compile() string {
	return "not implemented"
}
