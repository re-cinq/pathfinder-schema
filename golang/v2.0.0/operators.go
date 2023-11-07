package schema

import (
	"encoding/json"
	"errors"
)

type PCROperator string

// Error parsing the Status
var ErrPCROperatorParse = errors.New("unsupported PCROperator")

var operators = map[string]PCROperator{
	"PEF":               PEF,
	"EPD International": EPD,
	"Other":             Other,
}

const (
	// for EU / PEF Methodology PCRs
	// https://wayback.archive-it.org/12090/20230313054137/https://ec.europa.eu/environment/archives/eussd/pdf/footprint/PEF%20methodology%20final%20draft.pdf
	PEF PCROperator = "PEF"

	// for PCRs authored or published by EPD International
	// https://www.environdec.com/home
	EPD PCROperator = "EPD International"

	// for a PCR not published by the operators mentioned above
	Other PCROperator = "Other"
)

func (u PCROperator) String() string {
	return string(u)
}

func (u *PCROperator) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	if status, ok := operators[value]; !ok {
		return ErrPCROperatorParse
	} else {
		*u = status
	}

	return nil
}
