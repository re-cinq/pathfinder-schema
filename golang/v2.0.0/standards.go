package schema

import (
	"encoding/json"
	"errors"
)

type Standard string

// Error parsing the Standard
var ErrStandardParse = errors.New("unsupported Standard")

var standards = map[string]Standard{
	"GHG Protocol Product standard": GHGProtocol,
	"ISO Standard 14067":            ISO14067,
	"ISO Standard 14044":            ISO14044,
}

const (
	// GHG Protocol Product standard
	GHGProtocol Standard = "GHG Protocol Product standard"

	// ISO Standard 14067
	ISO14067 Standard = "ISO Standard 14067"

	// ISO Standard 14044
	ISO14044 Standard = "ISO Standard 14044"
)

func (u Standard) String() string {
	return string(u)
}

func (u *Standard) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	if value, ok := standards[value]; !ok {
		return ErrStandardParse
	} else {
		*u = value
	}

	return nil
}
