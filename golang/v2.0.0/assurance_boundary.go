package schema

import (
	"encoding/json"
	"errors"
)

type AssuranceBoundary string

// Error parsing the AssuranceBoundary
var ErrAssuranceBoundaryParse = errors.New("unsupported AssuranceBoundary")

var boundaries = map[string]AssuranceBoundary{
	"Gate-to-Gate":   GateToGate,
	"Cradle-to-Gate": CradleToGate,
}

const (
	// Gate-to-Gate
	GateToGate AssuranceBoundary = "Gate-to-Gate"

	// Cradle-to-Gate
	CradleToGate AssuranceBoundary = "Cradle-to-Gate"
)

func (u AssuranceBoundary) String() string {
	return string(u)
}

func (u *AssuranceBoundary) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	if status, ok := boundaries[value]; !ok {
		return ErrAssuranceBoundaryParse
	} else {
		*u = status
	}

	return nil
}
