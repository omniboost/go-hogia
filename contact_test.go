package hogia_test

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"testing"
	"time"

	"github.com/omniboost/go-hogia"
)

func TestContactsSection(t *testing.T) {
	var b bytes.Buffer

	s := hogia.ContactsSection{
		Header: hogia.Header{
			Rubrik:     "Kontakt",
			DateFormat: "YYYY-MM-DD",
		},
		ContactsHeader: hogia.ContactsHeader{
			PostType:    "Kontakt",
			Text:        "from PMS ",
			VoucherDate: hogia.Date{time.Date(2018, 9, 6, 0, 0, 0, 0, time.UTC)},
		},
		Contacts: hogia.Contacts{
			hogia.Contact{
				Kontakt: hogia.Kontakt{
					Updates:       11,
					ContactNumber: "4055715",
					Name:          "Magnus Andersson AB",
					Address1:      "Hakenäs 101",
					ZipCode:       "44428",
					Address5:      "Stenungsund",
					Address6:      "SE",
					Customer:      true,
					Supplier:      false,
					Phone:         "30366290",
				},
				Kund: hogia.Kund{
					Updates:       11,
					ContactNumber: "4055715",
					PaymentDays:   30,
					ARAccount:     1510,
					ReminderType:  0,
					EUCustomer:    -1,
				},
			},
		},
	}

	expected := `Rubrik	Kontakt
Datumformat	YYYY-MM-DD
Kontakt	from PMS 20180906	
Kontakt	11	4055715	Magnus Andersson AB           	Hakenäs 101                   	                              	                              	                              	44428     	Stenungsund                   	SE                            															-1	0	30366290            	                    			
kund	11	4055715             											                              	                              	30	1510	                              		0		-1											                              	
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
