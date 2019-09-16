package hogia

import "fmt"

type ContactsSection struct {
	Header         Header
	ContactsHeader ContactsHeader
	Contacts       Contacts
}

func (s ContactsSection) MarshalCSV() ([][]string, error) {
	ss := [][]string{}

	tmp, err := s.Header.MarshalCSV()
	ss = append(ss, tmp...)
	if err != nil {
		return ss, err
	}

	tmp, err = s.ContactsHeader.MarshalCSV()
	ss = append(ss, tmp...)
	if err != nil {
		return ss, err
	}

	tmp, err = s.Contacts.MarshalCSV()
	ss = append(ss, tmp...)
	if err != nil {
		return ss, err
	}

	return ss, nil
}

type ContactsHeader struct {
	// “Kontakt”, fixed text
	// Mandatory
	PostType    string
	Text        string
	VoucherDate Date
}

func (h ContactsHeader) MarshalCSV() ([][]string, error) {
	if h.Text == "" {
		return [][]string{}
	}

	return [][]string{
		[]string{
			"Kontakt",
			h.Text,
			h.VoucherDate.String(),
		},
	}, nil
}

type Contacts []Contact

func (cc Contacts) MarshalCSV() ([][]string, error) {
	ss := [][]string{}

	for _, c := range cc {
		tmp, err := c.MarshalCSV()
		if err != nil {
			return ss, err
		}
		ss = append(ss, tmp...)
	}

	return ss, nil
}

type Contact struct {
	Kontakt Kontakt
	Kund    Kund
}

func (c Contact) MarshalCSV() ([][]string, error) {
	ss := [][]string{}

	tmp, err := c.Kontakt.MarshalCSV()
	ss = append(ss, tmp...)
	if err != nil {
		return ss, err
	}

	tmp, err = c.Kund.MarshalCSV()
	ss = append(ss, tmp...)
	if err != nil {
		return ss, err
	}

	return ss, nil
}

type Kontakt struct {
	// "Kontakt", fixed test
	// Mandatory
	PostType string
	// “10”, update existing and create new
	// Mandatory
	Updates int
	// Mandatory
	// Length: 20
	ContactNumber string
	// Mandatory
	// Length: 50
	Name string
	// Length: 30
	Address1 string
	// Length: 30
	Address2 string
	// Length: 30
	Address3 string
	// Length: 30
	Address4 string
	// Length: 5
	ZipCode string
	// Length: 30
	Address5 string
	// Length: 30
	Address6   string
	Reserved12 Reserved
	Reserved13 Reserved
	Reserved14 Reserved
	Reserved15 Reserved
	Reserved16 Reserved
	Reserved17 Reserved
	Reserved18 Reserved
	// If left blank the name will be used for sorting
	// Length: 50
	SortName   string
	Reserved20 Reserved
	Reserved21 Reserved
	Reserved22 Reserved
	Comment    string
	Reserved24 Reserved
	Reserved25 Reserved
	// -1 if customer, other 0
	Customer IntBool
	// -1 if supplier, other 0
	Supplier IntBool
	// Length: 20
	Phone string
	// Length: 20
	Fax string
	// Length: 30
	Reference string
	// Length: 40
	EANLocalCode string
	// Length: 60
	Email string
	// Length: 50
	IBAN string
}

func (k Kontakt) MarshalCSV() ([][]string, error) {
	return [][]string{
		[]string{
			"Kontakt",
			fmt.Sprint(k.Updates),
			k.ContactNumber,
			k.Name,
			k.Address1,
			k.Address2,
			k.Address3,
			k.Address4,
			k.ZipCode,
			k.Address5,
			k.Address6,
			fmt.Sprint(k.Reserved12),
			fmt.Sprint(k.Reserved13),
			fmt.Sprint(k.Reserved14),
			fmt.Sprint(k.Reserved15),
			fmt.Sprint(k.Reserved16),
			fmt.Sprint(k.Reserved17),
			fmt.Sprint(k.Reserved18),
			k.SortName,
			fmt.Sprint(k.Reserved20),
			fmt.Sprint(k.Reserved21),
			fmt.Sprint(k.Reserved22),
			k.Comment,
			fmt.Sprint(k.Reserved24),
			fmt.Sprint(k.Reserved25),
			k.Customer.String(),
			k.Supplier.String(),
			k.Phone,
			k.Fax,
			k.Reference,
			k.EANLocalCode,
			k.Email,
			k.IBAN,
		},
	}, nil
}

type Kund struct {
	// "Kund", fixed test
	// Mandatory
	PostType string
	// "10", update existing and create new
	// Mandatory
	Updates int
	// Mandatory
	// Length: 20
	ContactNumber string
	Claim         Amount
	CreditLimit   Amount
	Reserved6     Reserved
	Reserved7     Reserved
	Reserved8     Reserved
	Reserved9     Reserved
	Reserved10    Reserved
	Reserved11    Reserved
	Reserved12    Reserved
	Reserved13    Reserved
	// Length: 30
	DeliveryTerms string
	// Length: 30
	PaymentTerms string
	PaymentDays  IntOpt
	// Mandatory
	ARAccount int
	// Default reference for the customer
	// Length: 30
	OurReference string
	// -1 if the customer shouldn't be invoiced
	NoInvoicing IntOpt
	// 0 - demand
	// 1 - Current account
	// 2 - Nothing
	ReminderType IntOpt
	// -1 if export customer
	ExportCustomer      IntOpt
	InterestCalculation IntOpt
	InvoiceFee          IntOpt
	// -1 to mark as selected, other 0
	Note          IntOpt
	EUSupplier    IntOpt
	Comment       string
	CustomerGroup IntOpt
	// Pricelist 1 – 10, no pricelist = 1
	PricelistArticles IntOpt
	// Pricelist 1 – 10, no pricelist = 1
	PricelistTime IntOpt
	// 0-9999999
	Dimension1       Dimension
	InterestFreeDays IntOpt
	// Within EU = -1
	EUCustomer IntOpt
	// Length: 30
	EGVATNo      string
	Reserved34   Reserved
	RemindersFee IntOpt
	Reserved36   Reserved
	CashCustomer IntOpt
	Reserved38   Reserved
	// Length: 30
	Form string
	// Country code
	// Length: 3
	Country string
	// SE - Swedish, NO - Norwegian etc
	// Length: 2
	Language string
	// Length: 3
	Currency    string
	SalesPerson IntOpt
	Reserved44  Reserved
	NoteText    string
	// Length: 30
	OrganizationNumber    string
	Reserved47            Reserved
	LowestInterestInvoice Amount
	Reserved49            Reserved
	Reserved50            Reserved
	Reserved51            Reserved
	Reserved52            Reserved
	Reserved53            Reserved
	Reserved54            Reserved
	Reserved55            Reserved
	Reserved56            Reserved
	Reserved57            Reserved
	Reserved58            Reserved
	Reserved59            Reserved
	Reserved60            Reserved
	Reserved61            Reserved
	Reserved62            Reserved
	Reserved63            Reserved
	Reserved64            Reserved
	// 0=No, 1=Post giro, 2=Bank giro, 3=BG account. When set GiroNr (66) must be set
	Autogiro IntOpt
}

func (k Kund) MarshalCSV() ([][]string, error) {
	return [][]string{
		[]string{
			"Kund",
			fmt.Sprint(k.Updates),
			k.ContactNumber,
			k.Claim.String(),
			k.CreditLimit.String(),
			k.Reserved6.String(),
			k.Reserved7.String(),
			k.Reserved8.String(),
			k.Reserved9.String(),
			k.Reserved10.String(),
			k.Reserved11.String(),
			k.Reserved12.String(),
			k.Reserved13.String(),
			k.DeliveryTerms,
			k.PaymentTerms,
			k.PaymentDays.String(),
			fmt.Sprint(k.ARAccount),
			k.OurReference,
			k.NoInvoicing.String(),
			k.ReminderType.String(),
			k.ExportCustomer.String(),
			k.InterestCalculation.String(),
			k.InvoiceFee.String(),
			k.Note.String(),
			k.EUSupplier.String(),
			k.Comment,
			k.CustomerGroup.String(),
			k.PricelistArticles.String(),
			k.PricelistTime.String(),
			k.Dimension1.String(),
			k.InterestFreeDays.String(),
			k.EUCustomer.String(),
			k.EGVATNo,
			k.Reserved34.String(),
			k.RemindersFee.String(),
			k.Reserved36.String(),
			k.CashCustomer.String(),
			k.Reserved38.String(),
			k.Form,
			k.Country,
			k.Language,
			k.Currency,
			k.SalesPerson.String(),
			k.Reserved44.String(),
			k.NoteText,
			k.OrganizationNumber,
			k.Reserved47.String(),
			k.LowestInterestInvoice.String(),
			k.Reserved49.String(),
			k.Reserved50.String(),
			k.Reserved51.String(),
			k.Reserved52.String(),
			k.Reserved53.String(),
			k.Reserved54.String(),
			k.Reserved55.String(),
			k.Reserved56.String(),
			k.Reserved57.String(),
			k.Reserved58.String(),
			k.Reserved59.String(),
			k.Reserved60.String(),
			k.Reserved61.String(),
			k.Reserved62.String(),
			k.Reserved63.String(),
			k.Reserved64.String(),
			k.Autogiro.String(),
		},
	}, nil
}
