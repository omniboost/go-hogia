package hogia

type AccountingItemsFile struct {
	Header                Header
	AccountingItemsHeader AccountingItemsHeader
	AccountingItems
	AccountingItemsFooter AccountingItemsFooter
}

func (f *AccountingItemsFile) MarshalCSV() ([][]string, error) {
	ss := [][]string{}

	tmp, err := f.Header.MarshalCSV()
	ss = append(ss, tmp...)
	if err != nil {
		return ss, err
	}

	tmp, err = f.AccountingItemsHeader.MarshalCSV()
	ss = append(ss, tmp...)
	if err != nil {
		return ss, err
	}

	tmp, err = f.AccountingItems.MarshalCSV()
	ss = append(ss, tmp...)
	if err != nil {
		return ss, err
	}

	tmp, err = f.AccountingItemsFooter.MarshalCSV()
	ss = append(ss, tmp...)
	if err != nil {
		return ss, err
	}

	return ss, nil
}

type AccountingItems []AccountingItem

func (ii AccountingItems) MarshalCSV() ([][]string, error) {
	return [][]string{}, nil
}

type AccountingItemsHeader struct {
	// “BFO”, fixed text
	// Mandatory
	PostType string
	// If excluded, it will automatically be set to A
	// Length: 4
	VoucherSeries string
	// Length: 4
	Text        string
	VoucherDate Date
	// Length: 1
	Reserved interface{}
	// -1 means that the system checks if the accounts are defined with VAT and
	// books accordingly. VAT accounts should not be included in the Voucher
	// If 0 or if left out, no VAT calculation will be made. The amounts should
	// then be sent as excl VAT. The VAT accounts must then be included in the
	// Voucher.
	// Mandatory
	// Length: 1
	InclVAT IntBool
}

func (h AccountingItemsHeader) MarshalCSV() ([][]string, error) {
	return [][]string{
		[]string{
			"BFO",
			h.VoucherSeries,
			h.Text,
			h.VoucherDate.String(),
			"",
			h.InclVAT.String(),
		},
	}, nil
}

type AccountingItemsFooter string

func (h *AccountingItemsFooter) MarshalCSV() ([][]string, error) {
	return [][]string{
		[]string{"BFO-slut"},
	}, nil
}

type AccountingItem struct {
	// "BFO-rad", fixed text
	// Mandatory
	Posttyp string
	// Accountnumbers allowed are in the range 1000-8999. Debit and Credit must
	// balance – if not, import will abort.  If internal accounts are used they
	// must balance (accounts 9000-9999).
	// Mandatory
	Account int
	// Mandatory
	Amount   float64
	Quantity float64
	// 0–999999
	Dimension1 int
	Project    int
	// Length: 10
	Specification string
	// Length: 30
	Text string
	// If excluded, the VAT code set on the account in Hogia will be used.
	VATCode interface{}
	// 0-999999
	Dimension2 int
}
