package hogia

import "fmt"

type InvoicesSection struct {
	Header         Header
	InvoicesHeader InvoicesHeader
	Contacts       Contacts
	Invoices       Invoices
}

func (s InvoicesSection) MarshalCSV() ([][]string, error) {
	ss := [][]string{}

	tmp, err := s.Header.MarshalCSV()
	ss = append(ss, tmp...)
	if err != nil {
		return ss, err
	}

	tmp, err = s.InvoicesHeader.MarshalCSV()
	ss = append(ss, tmp...)
	if err != nil {
		return ss, err
	}

	tmp, err = s.Contacts.MarshalCSV()
	ss = append(ss, tmp...)
	if err != nil {
		return ss, err
	}

	tmp, err = s.Invoices.MarshalCSV()
	ss = append(ss, tmp...)
	if err != nil {
		return ss, err
	}

	return ss, nil
}

type InvoicesHeader struct {
	// “KUNDRESK”, fixed text
	// Mandatory
	PostType    string
	Text        string
	VoucherDate Date
}

func (h InvoicesHeader) MarshalCSV() ([][]string, error) {
	if h.Text == "" {
		return [][]string{}, nil
	}

	return [][]string{
		[]string{
			"KUNDRESK",
			h.Text,
			h.VoucherDate.String(),
		},
	}, nil
}

type Invoices []Invoice

func (ii Invoices) MarshalCSV() ([][]string, error) {
	ss := [][]string{}

	for _, i := range ii {
		tmp, err := i.MarshalCSV()
		if err != nil {
			return ss, err
		}
		ss = append(ss, tmp...)
	}

	return ss, nil
}

type Invoice struct {
	Header InvoiceHeader
	Items  InvoiceItems
	Footer InvoiceFooter
}

func (i Invoice) MarshalCSV() ([][]string, error) {
	ss := [][]string{}

	tmp, err := i.Header.MarshalCSV()
	ss = append(ss, tmp...)
	if err != nil {
		return ss, err
	}

	tmp, err = i.Items.MarshalCSV()
	ss = append(ss, tmp...)
	if err != nil {
		return ss, err
	}

	tmp, err = i.Footer.MarshalCSV()
	ss = append(ss, tmp...)
	if err != nil {
		return ss, err
	}

	return ss, nil
}

type InvoiceHeader struct {
	// “Kundreskontra”, fixed text
	// Mandatory
	PostType string
	// Must be unique, meaning that the Invoice# should not already exist in Hogia.
	// Mandatory
	InvoiceNumber string
	// Mandatory
	// Length: 20
	CustomerNumber string
	// Mandatory
	InvoiceDate Date
	// Mandatory
	ExpiryDate Date
	// Invoice# to credit
	OriginalInvoice string
	Comment         string
	// On claim account
	// 0-9999999
	Dimension1 Dimension
	// On claim account
	Project string
	// 0 = Debet
	// 1 = credit
	InvoiceType InvoiceType
	// Incl VAT, 2 decimals
	// Mandatory
	Amount Amount
	// Mandatory
	VAT Amount
	// Length: 30
	PaymentTerms string
	// Length: 4
	VoucherSeries string
	// Length: 30
	VoucherText string
	// True = -1
	SettleInvoice *IntOpt
	ThirdParty    *IntBool
	EUInvoice     *IntBool
	// Length: 2
	Reserved1 Reserved
	// If left out picked up from customer
	// Length: 3
	Currency string
	// If left out fetched from the list of currencies
	ExchangeRate FloatOpt
	Reserved2    string
	// On claim account
	// 0-999999
	Dimension2 Dimension
	// -1 if the first line is the account entry for the claim account
	SkipFirstAccountLine IntOpt
	// Mandatory if invoice in foreign currency
	InvoiceAmountInCurrency Amount
	VoucherDate             Date
	// True = -1
	AutoGiro IntOpt
}

func (h InvoiceHeader) MarshalCSV() ([][]string, error) {
	return [][]string{
		[]string{
			"Kundreskontra",
			h.InvoiceNumber,
			h.CustomerNumber,
			h.InvoiceDate.String(),
			h.ExpiryDate.String(),
			h.OriginalInvoice,
			h.Comment,
			h.Dimension1.String(),
			h.Project,
			fmt.Sprint(h.InvoiceType),
			h.Amount.Round2(),
			h.VAT.Round4(),
			h.PaymentTerms,
			h.VoucherSeries,
			// fmt.Printf(h.SettleInvoice),
			// h.ThirdParty.String(),
			// h.EUInvoice.String(),
			// h.Reserved,
			// h.Currency,
			// h.ExchangeRate.String(),
			// h.Reserved2,
			// h.Dimension2.String(),
			// h.SkipFirstAccountLine.String(),
			// h.InvoiceAmountInCurrency.String(),
			// h.VoucherDate.String(),
			// h.AutoGiro.String(),
		},
	}, nil
}

type InvoiceItems []InvoiceItem

func (ii InvoiceItems) MarshalCSV() ([][]string, error) {
	ss := [][]string{}

	for _, i := range ii {
		tmp, err := i.MarshalCSV()
		if err != nil {
			return ss, err
		}
		ss = append(ss, tmp...)
	}

	return ss, nil
}

type InvoiceItem struct {
	// “Kontering”, fixed text
	// Mandatory
	PostType string
	// If the rest is left out only A/R is updated and not the accounting.
	// If internal accounts are used, then they must balance.
	// If the accounting isn’t wanted in the A/R then specify ‘*’
	// Mandatory
	Account string
	// Amount with 2 decimals
	// Mandatory
	Amount   Amount
	Quantity Quantity
	// 0–9999999
	Dimension1 Dimension
	Project    string
	// Length: 10
	Specification string
	Text          string
	// If excluded information is picked up from the account definition
	VATCode string
	// 0–9999999
	Dimension2 Dimension
}

func (i InvoiceItem) MarshalCSV() ([][]string, error) {
	return [][]string{
		[]string{
			"Kontering",
			i.Account,
			i.Amount.String(),
			fmt.Sprint(i.Quantity),
			i.Dimension1.String(),
			fmt.Sprint(i.Project),
			i.Specification,
			i.Text,
			i.VATCode,
			i.Dimension2.String(),
		},
	}, nil
}

type InvoiceFooter struct{}

func (f *InvoiceFooter) MarshalCSV() ([][]string, error) {
	return [][]string{
		[]string{"Kundreskontra-Slut"},
	}, nil
}

type InvoiceType int

var (
	InvoiceTypeDebit  InvoiceType = 0
	InvoiceTypeCredit InvoiceType = 1
)
