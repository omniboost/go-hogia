package hogia

type File struct {
	Sections Sections
}

type Sections []Section

type Section interface {
}
