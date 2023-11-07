package schema

import (
	"encoding/json"
	"errors"
)

type AccountingMethodology string

// Error parsing the AccountingMethodology
var ErrAccountingMethodologyParse = errors.New("unsupported AccountingMethodology")

var accountings = map[string]AccountingMethodology{
	"PEF":     PEFAccounting,
	"ISO":     ISOAccounting,
	"GHGP":    GHGPAccounting,
	"Quantis": QuantisAccounting,
}

const (
	// for the EU Product Environmental Footprint Guide
	PEFAccounting AccountingMethodology = "PEF"

	// For the ISO 14067 standard
	ISOAccounting AccountingMethodology = "ISO"

	// For the Greenhouse Gas Protocol (GHGP) Land sector and Removals Guidance
	GHGPAccounting AccountingMethodology = "GHGP"

	// For the Quantis Accounting for Natural Climate Solutions Guidance
	QuantisAccounting AccountingMethodology = "Quantis"
)

func (u AccountingMethodology) String() string {
	return string(u)
}

func (u *AccountingMethodology) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	if status, ok := accountings[value]; !ok {
		return ErrAccountingMethodologyParse
	} else {
		*u = status
	}

	return nil
}
