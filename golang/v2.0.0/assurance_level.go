package schema

import (
	"encoding/json"
	"errors"
)

type AssuranceLevel string

// Error parsing the AssuranceLevel
var ErrAssuranceLevelParse = errors.New("unsupported AssuranceLevel")

var assuranceLevels = map[string]AssuranceLevel{
	"limited":    Limited,
	"reasonable": Reasonable,
}

const (
	// for limited assurance
	Limited AssuranceLevel = "limited"

	// for reasonable assurance
	Reasonable AssuranceLevel = "reasonable"
)

func (u AssuranceLevel) String() string {
	return string(u)
}

func (u *AssuranceLevel) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	if status, ok := assuranceLevels[value]; !ok {
		return ErrAssuranceLevelParse
	} else {
		*u = status
	}

	return nil
}
