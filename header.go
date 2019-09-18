package hogia

type Header struct {
	Rubrik     string
	DateFormat string
}

func (h *Header) MarshalCSV() ([][]string, error) {
	s := [][]string{}

	if h.Rubrik != "" {
		s = append(s, []string{"Rubrik", h.Rubrik})
	}

	if h.DateFormat != "" {
		s = append(s, []string{"DatumFormat", h.DateFormat})
	}

	return s, nil
}
