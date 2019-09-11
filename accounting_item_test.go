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
		AccountingItems: hogia.AccountingItems{{
			Account: 1515,
			Amount:  -44420.54,
			// Quanttiy: 0,
			// Dimension1: 0,
			// Project: 0,
			// Specification: "",
			Text: "Avräkning reception",
			// VATCode: nil,
			// Dimensino2: 0,
		}},
		// BFO-rad	1515	-44420,54					Avräkning reception
		// BFO-rad	1509	85432,00					FAKTURASTÄNGNING
		// BFO-rad	1910	5057,00					Kontant
		// BFO-rad	1912	300,00					BWR-poäng
		// BFO-rad	1914	1196,00					FIT BW
		// BFO-rad	1916	300,00					Fri natt (ej sas)
		// BFO-rad	1583	47167,48					Eurocard/Master/VISA /SHB
		// BFO-rad	1582	12239,00					KORTMANUELL & Restaurang
		// BFO-rad	2420	-531,00					Förskott skuld
		// BFO-rad	3009	-8083,00		102			Parkering
		// BFO-rad	3010	-60282,56		102			Logi
		// BFO-rad	3011	-106,25		102			EM-Kaffe
		// BFO-rad	3013	-16550,00		102			Frukost
		// BFO-rad	3015	-9869,00		101			MAT
		// BFO-rad	3016	-510,00		102			Dryck
		// BFO-rad	3060	-7380,00		102			Transfer
		// BFO-rad	3070	-58,00		101			Cider
		// BFO-rad	3071	-216,00		101			Sprit
		// BFO-rad	3072	-3175,00		101			Vin
		// BFO-rad	3073	-3390,00		101			Starköl
		// BFO-rad	3090	-150,00		102			Övrig försäljning. 25%
		// BFO-rad	3540	-100,00					Exp.avgift
		// BFO-rad	3740	-0,13					Öresutjämning
		// BFO-rad	6056	3130,00					Restaurangcheck
		AccountingItemsFooter: hogia.AccountingItemsFooter{},
	}

	expected := `Rubrik	Bokföringsorder
Datumformat	YYYY-MM-DD
BFO	O	From PMS 	2018-09-06		-1
BFO-rad	1515	-44420,54					Avräkning reception		
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
		t.Errorf("Header didn't marshal properly. Got: \n'%s', expected:\n'%s'", b.String(), expected)
	}
}
