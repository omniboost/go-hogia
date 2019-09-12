package hogia

type File struct {
	Sections Sections
}

type Sections []Section

type Section interface {
	MarshalCSV() ([][]string, error)
}

func (f File) MarshalCSV() ([][]string, error) {
	ss := [][]string{}

	for _, s := range f.Sections {
		tmp, err := s.MarshalCSV()
		if err != nil {
			return ss, err
		}
		ss = append(ss, tmp...)
	}

	return ss, nil
}
