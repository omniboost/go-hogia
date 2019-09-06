package hogia_test

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"testing"
	"time"

	"github.com/omniboost/go-hogia"
)

func TestAccountingItemsFile(t *testing.T) {
	var b bytes.Buffer

	f := hogia.AccountingItemsFile{
		Header: hogia.Header{
			Rubrik:     "Bokföringsorder",
			DateFormat: "YYYY-MM-DD",
		},
		AccountingItemsHeader: hogia.AccountingItemsHeader{
			PostType:      "BFO",
			VoucherSeries: "O",
			Text:          "From PMS ",
			VoucherDate:   hogia.Date{time.Date(2018, 9, 6, 0, 0, 0, 0, time.UTC)},
			Reserved:      nil,
			InclVAT:       false,
		},
		AccountingItemsFooter: hogia.AccountingItemsFooter{},
	}

	expected := `Rubrik	Bokföringsorder
Datumformat	YYYY-MM-DD
BFO	O	From PMS 	2018-09-06		-1
BFO-slut
`

	ss, err := f.MarshalCSV()
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
		t.Errorf("Header didn't marshal properly. Got: \n%s, expected:\n%s", b.String(), expected)
	}
}
