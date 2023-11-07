package schema

import (
	"encoding/json"
	"errors"
)

type Status string

// Error parsing the Status
var ErrStatusParse = errors.New("unsupported Status")

var statuses = map[string]Status{
	"Active":     Active,
	"Deprecated": Deprecated,
}

const (
	// Active product
	Active Status = "Active"

	// Deprecated product
	Deprecated Status = "Deprecated"
)

func (u Status) String() string {
	return string(u)
}

func (u *Status) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	if status, ok := statuses[value]; !ok {
		return ErrStatusParse
	} else {
		*u = status
	}

	return nil
}
