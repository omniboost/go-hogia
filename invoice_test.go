package hogia_test

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"testing"
	"time"

	"github.com/omniboost/go-hogia"
)

func TestInvoicesSection(t *testing.T) {
	var b bytes.Buffer

	s := hogia.InvoicesSection{
		Header: hogia.Header{
			Rubrik:     "Kundreskontra",
			DateFormat: "YYYY-MM-DD",
		},
		InvoicesHeader: hogia.InvoicesHeader{
			PostType:    "KUNDRESK",
			Text:        "From PMS",
			VoucherDate: hogia.Date{time.Date(2018, 9, 6, 0, 0, 0, 0, time.UTC)},
		},
		Invoices: hogia.Invoices{
			hogia.Invoice{
				Header: hogia.InvoiceHeader{
					InvoiceNumber:  "00065216",
					CustomerNumber: "4055715",
					InvoiceDate:    hogia.Date{time.Date(2018, 9, 6, 0, 0, 0, 0, time.UTC)},
					ExpiryDate:     hogia.Date{time.Date(2018, 10, 6, 0, 0, 0, 0, time.UTC)},
					InvoiceType:    hogia.InvoiceTypeDebit,
					Amount:         hogia.Amount(24410.00),
					VAT:            hogia.Amount(2535.32),
				},
				Items: hogia.InvoiceItems{
					{
						Account: "*",
					},
				},
				Footer: hogia.InvoiceFooter{},
			},
			hogia.Invoice{
				Header: hogia.InvoiceHeader{
					InvoiceNumber:  "00065217",
					CustomerNumber: "4055715",
					InvoiceDate:    hogia.Date{time.Date(2018, 9, 6, 0, 0, 0, 0, time.UTC)},
					ExpiryDate:     hogia.Date{time.Date(2018, 10, 6, 0, 0, 0, 0, time.UTC)},
					InvoiceType:    hogia.InvoiceTypeCredit,
					Amount:         hogia.Amount(-24410.00),
					VAT:            hogia.Amount(-2535.32),
				},
				Items: hogia.InvoiceItems{
					{
						Account: "*",
					},
				},
				Footer: hogia.InvoiceFooter{},
			},
			hogia.Invoice{
				Header: hogia.InvoiceHeader{
					InvoiceNumber:  "00065218",
					CustomerNumber: "4055715",
					InvoiceDate:    hogia.Date{time.Date(2018, 9, 6, 0, 0, 0, 0, time.UTC)},
					ExpiryDate:     hogia.Date{time.Date(2018, 10, 6, 0, 0, 0, 0, time.UTC)},
					InvoiceType:    hogia.InvoiceTypeDebit,
					Amount:         hogia.Amount(6050.00),
					VAT:            hogia.Amount(616.48),
				},
				Items: hogia.InvoiceItems{
					{
						Account: "*",
					},
				},
				Footer: hogia.InvoiceFooter{},
			},
		},
	}

	expected := `Rubrik	Kundreskontra
Datumformat	YYYY-MM-DD
KUNDRESK	From PMS	2018-09-06
Kundreskontra	00065216	4055715	2018-09-06	2018-10-06					0	24410,00	2535,3200		
Kontering	*								
Kundreskontra-Slut
Kundreskontra	00065217	4055715	2018-09-06	2018-10-06					1	-24410,00	-2535,3200		
Kontering	*								
Kundreskontra-Slut
Kundreskontra	00065218	4055715	2018-09-06	2018-10-06					0	6050,00	616,4800		
Kontering	*								
Kundreskontra-Slut
`

	ss, err := s.MarshalCSV()
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
