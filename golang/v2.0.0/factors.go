package schema

import (
	"encoding/json"
	"errors"
)

type CharacterizationFactor string

// Error parsing the CharacterizationFactor
var ErrCharacterizationFactorParse = errors.New("unsupported CharacterizationFactor")

var factors = map[string]CharacterizationFactor{
	"AR6": AR6,
	"AR5": AR5,
}

const (
	// AR6 for the Sixth Assessment Report of the Intergovernmental Panel on Climate Change (IPCC)
	AR6 CharacterizationFactor = "AR6"

	// for the Fifth Assessment Report of the IPCC.
	AR5 CharacterizationFactor = "AR5"
)

func (u CharacterizationFactor) String() string {
	return string(u)
}

func (u *CharacterizationFactor) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	if value, ok := factors[value]; !ok {
		return ErrCharacterizationFactorParse
	} else {
		*u = value
	}

	return nil
}
