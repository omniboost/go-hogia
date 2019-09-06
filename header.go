package hogia

type Header struct {
	Rubrik     string
	DateFormat string
}

func (h *Header) MarshalCSV() ([][]string, error) {
	return [][]string{
		[]string{"Rubrik", h.Rubrik},
		[]string{"Datumformat", h.DateFormat},
	}, nil
}
