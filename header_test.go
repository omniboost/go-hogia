package hogia_test

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"testing"

	"github.com/omniboost/go-hogia"
)

func TestHeader(t *testing.T) {
	var b bytes.Buffer

	header := &hogia.Header{
		Rubrik:     "Invoices from XXXXX 2018-11-30",
		DateFormat: "YYYY-MM-DD",
	}

	expected := `Rubrik	Invoices from XXXXX 2018-11-30
Datumformat	YYYY-MM-DD
`

	ss, err := header.MarshalCSV()
	if err != nil {
		t.Error()
	}

	w := csv.NewWriter(bufio.NewWriter(&b))
	w.Comma = '\t'
	err = w.WriteAll(ss)
	if err != nil {
		t.Error(err)
	}

	if b.String() != expected {
		t.Errorf("Header didn't marshal properly. Got %s, expected %s", b.String(), expected)
	}
}
